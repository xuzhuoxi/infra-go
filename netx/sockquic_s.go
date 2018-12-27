package netx

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/util-go/errorsx"
	"log"
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
	s.mapSession = make(map[string]quic.Session)
	s.mapProxy = make(map[string]IMessageSendReceiver)
	s.running = true
	s.serverMu.Unlock()
	log.Println(funcName + "()")
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
	for _, value := range s.mapSession {
		value.Close()
	}
	s.mapSession = nil
	s.running = false
	log.Println(funcName + "()")
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
	defer func() {
		s.serverMu.Lock()
		if nil != session {
			session.Close()
		}
		delete(s.mapSession, address)
		delete(s.mapProxy, address)
		s.serverMu.Unlock()
	}()
	s.serverMu.Lock()
	s.mapSession[address] = session
	proxy := NewMessageSendReceiver(session, session, QuicRW, s.Network) //要改
	s.mapProxy[address] = proxy
	proxy.SetSplitHandler(s.splitHandler)
	proxy.SetMessageHandler(s.messageHandler)
	s.serverMu.Unlock()
	log.Println("New Quic Connection:", address)
	proxy.StartReceiving()
}

func listenQuic(address string) (quic.Listener, error) {
	return quic.ListenAddr(address, generateTLSConfig(), nil)
}
