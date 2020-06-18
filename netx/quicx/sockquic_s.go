package quicx

import (
	"errors"
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
)

func NewQuicServer() IQUICServer {
	return newQuicServer().(IQUICServer)
}

func newQuicServer() netx.ISockServer {
	server := &QUICServer{}
	server.Name = "QuicServer"
	server.Network = netx.QuicNetwork
	server.Logger = logx.DefaultLogger()
	server.PackHandlerContainer = netx.NewIPackHandler(nil)
	return server
}

//----------------------------

type IQUICServer interface {
	netx.ISockServer
	eventx.IEventDispatcher
}

type QUICServer struct {
	eventx.EventDispatcher
	netx.SockServerBase
	lang.ChannelLimitNone

	listener   quic.Listener
	mapProxy   map[string]netx.IPackSendReceiver
	mapSession map[string]quic.Session
	mapStream  map[string]quic.Stream
}

func (s *QUICServer) StartServer(params netx.SockParams) error {
	funcName := "QUICServer.StartServer"
	s.ServerMu.Lock()
	if s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		s.Network = params.Network
	}
	listener, err := listenQuic(params.LocalAddress)
	if err != nil {
		defer s.ServerMu.Unlock()
		return err
	}
	s.Logger.Infoln("[QUICServer] listening on:", params.LocalAddress)
	s.listener = listener
	s.mapProxy = make(map[string]netx.IPackSendReceiver)
	s.mapSession = make(map[string]quic.Session)
	s.mapStream = make(map[string]quic.Stream)
	s.Running = true
	s.ServerMu.Unlock()
	s.Logger.Infoln(funcName + "()")
	s.DispatchServerStartedEvent(s)
	for s.Running {
		session, err := listener.Accept()
		if !s.Running || nil != err {
			return err
		}
		go s.handlerSession(session.RemoteAddr().String(), session)
	}
	return nil
}

func (s *QUICServer) StopServer() error {
	funcName := "QUICServer.StopServer"
	s.ServerMu.Lock()
	if !s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.ServerMu.Unlock()
		s.Logger.Infoln(funcName + "()")
		s.DispatchServerStoppedEvent(s)
	}()
	if nil != s.listener {
		s.listener.Close()
		s.listener = nil
	}
	for _, proxy := range s.mapProxy {
		proxy.StopReceiving()
	}
	s.mapProxy = nil
	for _, stream := range s.mapStream {
		stream.Close()
	}
	s.mapStream = nil
	for _, sess := range s.mapSession {
		sess.Close()
	}
	s.mapSession = nil
	s.Running = false
	return nil
}

func (s *QUICServer) Connections() int {
	return len(s.mapSession)
}

func (s *QUICServer) CloseConnection(address string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	stream, ok1 := s.mapStream[address]
	session, ok2 := s.mapSession[address]
	if !ok1 && !ok2 {
		return errors.New("QUICServer: No Connection At " + address), false
	}
	delete(s.mapProxy, address)
	delete(s.mapStream, address)
	delete(s.mapSession, address)
	var err1 error
	var err2 error
	if ok1 {
		err1 = stream.Close()
	}
	if ok2 {
		err2 = session.Close()
	}
	if nil != err1 {
		return err1, true
	}
	if nil != err2 {
		return err2, true
	}
	return nil, true
}

func (s *QUICServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := QuicDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *QUICServer) SendBytesTo(data []byte, rAddress ...string) error {
	funcName := "QUICServer.SendBytesTo"
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	if !s.Running || nil == s.mapProxy || nil == s.mapStream || nil == s.mapSession {
		return netx.ConnNilError(funcName)
	}
	if 0 == len(rAddress) {
		return netx.NoAddrError(funcName)
	}
	for _, address := range rAddress {
		ts, ok := s.mapProxy[address]
		if ok {
			ts.SendBytes(data)
		}
	}
	return nil
}

func (s *QUICServer) handlerSession(address string, session quic.Session) {
	funcName := "QUICServer.handlerSession"
	s.ServerMu.Lock()
	var stream quic.Stream
	var err error
	stream, err = session.AcceptStream()
	if nil != err {
		s.ServerMu.Unlock()
		s.Logger.Warnln(funcName, err)
		return
	}
	s.mapSession[address] = session
	s.mapStream[address] = stream
	connProxy := &QUICStreamAdapter{Reader: stream, Writer: stream, RemoteAddr: session.RemoteAddr()}
	proxy := netx.NewPackSendReceiver(connProxy, connProxy, s.PackHandlerContainer, QuicDataBlockHandler, s.Logger, false)
	s.mapProxy[address] = proxy
	s.ServerMu.Unlock()
	s.DispatchServerConnOpenEvent(s, address)
	s.Logger.Infoln("[QUICServer] Quic Connection:", address, "Opened!")

	defer func() {
		s.DispatchServerConnCloseEvent(s, address)
		s.Logger.Infoln("[QUICServer] Quic Connection:", address, "Closed!")
	}()
	defer func() {
		s.ServerMu.Lock()
		delete(s.mapProxy, address)
		delete(s.mapStream, address)
		delete(s.mapSession, address)
		if nil != stream {
			stream.Close()
			stream = nil
		}
		if nil != session {
			session.Close()
			session = nil
		}
		s.ServerMu.Unlock()
	}()
	proxy.StartReceiving() //这里会阻塞
}

func listenQuic(address string) (quic.Listener, error) {
	return quic.ListenAddr(address, generateTLSConfig(), nil)
}
