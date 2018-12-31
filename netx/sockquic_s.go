package netx

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
)

func NewQuicServer() IQUICServer {
	rs := &QUICServer{}
	rs.Network = QuicNetwork
	rs.splitHandler = DefaultByteSplitHandler
	rs.messageHandler = DefaultMessageHandler
	return rs
}

type QUICServer struct {
	SockServerBase

	listener   quic.Listener
	mapProxy   map[string]IMessageSendReceiver
	mapSession map[string]quic.Session
	mapStream  map[string]quic.Stream
}

func (s *QUICServer) StartServer(params SockParams) error {
	funcName := "QUICServer.StartServer"
	s.serverMu.Lock()
	if s.running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		s.Network = params.Network
	}
	listener, err := listenQuic(params.LocalAddress)
	if err != nil {
		defer s.serverMu.Unlock()
		return err
	}
	s.listener = listener
	s.mapProxy = make(map[string]IMessageSendReceiver)
	s.mapSession = make(map[string]quic.Session)
	s.mapStream = make(map[string]quic.Stream)
	s.running = true
	s.serverMu.Unlock()
	logx.Infoln(funcName + "()")
	for s.running {
		session, err := listener.Accept()
		if !s.running || nil != err {
			return err
		}
		go s.handlerSession(session.RemoteAddr().String(), session)
	}
	return nil
}

func (s *QUICServer) StopServer() error {
	funcName := "QUICServer.StopServer"
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	if !s.running {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if nil != s.listener {
		s.listener.Close()
		s.listener = nil
	}
	for _, sess := range s.mapSession {
		sess.Close()
	}
	s.mapSession = nil
	for _, stream := range s.mapStream {
		stream.Close()
	}
	s.mapStream = nil
	s.running = false
	logx.Infoln(funcName + "()")
	return nil
}

func (s *QUICServer) SendDataTo(data []byte, rAddress ...string) error {
	funcName := "QUICServer.SendDataTo"
	if 0 == len(rAddress) {
		return NoAddrError(funcName)
	}
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	for _, address := range rAddress {
		ts, ok := s.mapProxy[address]
		if ok {
			ts.SendMessage(data)
		}
	}
	return nil
}

func (s *QUICServer) handlerSession(address string, session quic.Session) {
	funcName := "QUICServer.handlerSession"
	defer func() {
		s.serverMu.Lock()
		if nil != session {
			session.Close()
		}
		delete(s.mapProxy, address)
		delete(s.mapSession, address)
		delete(s.mapStream, address)
		s.serverMu.Unlock()
	}()
	s.serverMu.Lock()
	stream, err := session.AcceptStream()
	if nil != err {
		logx.Warnln(funcName, err)
		return
	}
	defer stream.Close()
	s.mapSession[address] = session
	connProxy := &QUICSessionReadWriter{Reader: stream, Writer: stream, RemoteAddr: session.RemoteAddr()}
	proxy := NewMessageSendReceiver(connProxy, connProxy, false)
	s.mapProxy[address] = proxy
	s.mapStream[address] = stream
	proxy.SetSplitHandler(s.splitHandler)
	proxy.SetMessageHandler(s.messageHandler)
	s.serverMu.Unlock()
	logx.Infoln("New Quic Connection:", address)
	proxy.StartReceiving()
}

func listenQuic(address string) (quic.Listener, error) {
	return quic.ListenAddr(address, generateTLSConfig(), nil)
}
