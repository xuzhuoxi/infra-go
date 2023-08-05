package quicx

import (
	"errors"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
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

	listener quic.Listener
	mapConn  map[string]netx.IServerConn
}

func (s *QUICServer) StartServer(params netx.SockParams) error {
	funcName := fmt.Sprintf("QUICServer[%s].StartServer", s.Name)
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
	s.mapConn = make(map[string]netx.IServerConn)
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
	funcName := fmt.Sprintf("QUICServer[%s].StopServer", s.Name)
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
	for _, value := range s.mapConn {
		value.CloseConn()
	}
	s.mapConn = nil
	s.Running = false
	return nil
}

func (s *QUICServer) SetMaxConn(max int) {
	return
}

func (s *QUICServer) Connections() int {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return len(s.mapConn)
}

func (s *QUICServer) CloseConnection(address string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	value, ok1 := s.mapConn[address]
	if !ok1 {
		return errors.New("QUICServer: No Connection At " + address), false
	}
	delete(s.mapConn, address)
	if err1 := value.CloseConn(); nil != err1 {
		return err1, false
	}
	return nil, true
}

func (s *QUICServer) FindConnection(address string) (conn netx.IServerConn, ok bool) {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	conn, ok = s.mapConn[address]
	return
}

func (s *QUICServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := QuicDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *QUICServer) SendBytesTo(data []byte, rAddress ...string) error {
	funcName := fmt.Sprintf("QUICServer[%s].SendBytesTo", s.Name)
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	if !s.Running || nil == s.mapConn {
		return netx.ConnNilError(funcName)
	}
	if 0 == len(rAddress) {
		return netx.NoAddrError(funcName)
	}
	for _, address := range rAddress {
		ts, ok := s.mapConn[address]
		if ok {
			ts.SendBytes(data)
		}
	}
	return nil
}

const (
	funcNameHandlerSession = "QUICServer.handlerSession"
)

func (s *QUICServer) handlerSession(address string, session quic.Session) {
	s.ServerMu.Lock()
	var stream quic.Stream
	var err error
	stream, err = session.AcceptStream()
	if nil != err {
		s.ServerMu.Unlock()
		s.Logger.Warnln(funcNameHandlerSession, err)
		return
	}
	connProxy := &QUICStreamAdapter{Reader: stream, Writer: stream, RemoteAddr: session.RemoteAddr()}
	proxy := netx.NewPackSendReceiver(connProxy, connProxy, s.PackHandlerContainer, QuicDataBlockHandler, s.Logger, false)
	s.mapConn[address] = &QuicSockConn{Address: address, Session: session, Stream: stream, SRProxy: proxy}
	s.ServerMu.Unlock()
	s.DispatchServerConnOpenEvent(s, address)
	s.Logger.Infoln("[QUICServer] Quic Connection:", address, "Opened!")

	defer func() {
		s.DispatchServerConnCloseEvent(s, address)
		s.Logger.Infoln("[QUICServer] Quic Connection:", address, "Closed!")
	}()
	defer func() {
		s.ServerMu.Lock()
		delete(s.mapConn, address)
		if nil != stream {
			stream.Close()
		}
		if nil != session {
			session.Close()
		}
		s.ServerMu.Unlock()
	}()
	proxy.StartReceiving() //这里会阻塞
}

func listenQuic(address string) (quic.Listener, error) {
	return quic.ListenAddr(address, generateTLSConfig(), nil)
}
