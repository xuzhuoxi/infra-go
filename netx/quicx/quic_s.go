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
	server := &QUICServer{
		logFuncNameSend: "QUICServer[Quic].SendBytesTo",
	}
	server.Name = "Quic"
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
	logFuncNameSend string

	listener quic.Listener
	mapConn  map[string]netx.IServerConn // [connId:netx.IServerConn]
}

func (s *QUICServer) StartServer(params netx.SockParams) error {
	funcName := fmt.Sprintf("[QUICServer[%s].StartServer]", s.Name)
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
	s.Logger.Infoln(funcName, "listening on:", params.LocalAddress)
	s.listener = listener
	s.mapConn = make(map[string]netx.IServerConn)
	s.Running = true
	s.ServerMu.Unlock()
	s.Logger.Infoln(funcName, "()")
	s.DispatchServerStartedEvent(s)
	for s.Running {
		session, err := listener.Accept()
		if !s.Running || nil != err {
			return err
		}
		go s.handlerSession(session)
	}
	return nil
}

func (s *QUICServer) StopServer() error {
	funcName := fmt.Sprintf("[QUICServer[%s].StopServer]", s.Name)
	s.ServerMu.Lock()
	if !s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.ServerMu.Unlock()
		s.Logger.Infoln(funcName, "()")
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

func (s *QUICServer) CloseConnection(connId string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	if conn, ok := s.mapConn[connId]; ok {
		delete(s.mapConn, connId)
		err = conn.CloseConn()
		return err, nil != err
	}
	return errors.New("QUICServer: No Connection At " + connId), false
}

func (s *QUICServer) FindConnection(connId string) (conn netx.IServerConn, ok bool) {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	conn, ok = s.mapConn[connId]
	return
}

func (s *QUICServer) SendPackTo(pack []byte, connId ...string) error {
	bytes := QuicDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, connId...)
}

func (s *QUICServer) SendBytesTo(data []byte, connId ...string) error {
	funcName := s.logFuncNameSend
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	if !s.Running || nil == s.mapConn {
		return netx.ConnNilError(funcName)
	}
	if 0 == len(connId) {
		return netx.NoAddrError(funcName)
	}
	for _, cId := range connId {
		ts, ok := s.mapConn[cId]
		if ok {
			ts.SendBytes(data)
		}
	}
	return nil
}

func (s *QUICServer) handlerSession(session quic.Session) {
	localAddress := session.LocalAddr().String()
	remoteAddress := session.RemoteAddr().String()
	connInfo := netx.NewConnInfo(localAddress, remoteAddress)
	stream, proxy := s.startConn(connInfo, session)
	if nil == stream || nil == proxy {
		return
	}
	proxy.StartReceiving() // 这里会阻塞
	s.endConn(connInfo, session, stream)
}

func (s *QUICServer) startConn(connInfo netx.IConnInfo, session quic.Session) (stream quic.Stream, proxy netx.IPackSendReceiver) {
	s.ServerMu.Lock()
	stream1, err := session.AcceptStream()
	if nil != err {
		s.ServerMu.Unlock()
		s.Logger.Warnln("[QUICServer.startConn]", err)
		return
	}
	connProxy := &QUICStreamAdapter{Reader: stream1, Writer: stream1, remoteAddress: connInfo.GetRemoteAddress()}
	proxy = netx.NewPackSendReceiver(connInfo, connProxy, connProxy, s.PackHandlerContainer, QuicDataBlockHandler, s.Logger, false)
	s.mapConn[connInfo.GetConnId()] = &QuicSockConn{QuicConnInfo: connInfo, Session: session, Stream: stream1, SRProxy: proxy}
	s.ServerMu.Unlock()

	s.DispatchServerConnOpenEvent(s, connInfo)
	s.Logger.Infoln("[QUICServer.startConn] Quic Connection:", connInfo, "Opened!")
	return stream1, proxy
}

func (s *QUICServer) endConn(connInfo netx.IConnInfo, session quic.Session, stream quic.Stream) {
	s.ServerMu.Lock()
	delete(s.mapConn, connInfo.GetConnId())
	s.ServerMu.Unlock()
	if nil != stream {
		stream.Close()
	}
	if nil != session {
		session.Close()
	}
	s.DispatchServerConnCloseEvent(s, connInfo)
	s.Logger.Infoln("[QUICServer.endConn] Quic Connection:", connInfo, "Closed!")
}

func listenQuic(localAddress string) (quic.Listener, error) {
	return quic.ListenAddr(localAddress, generateTLSConfig(), nil)
}
