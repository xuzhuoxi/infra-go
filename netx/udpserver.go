package netx

import (
	"github.com/xuzhuoxi/go-util/errorsx"
	"log"
	"net"
	"sync"
)

const (
	UDPBuffLength = 4096
)

func NewUDPServer() IUDPServer {
	rs := &UDPServer{}
	rs.Network = UDPNetwork
	rs.splitHandler = DefaultByteSplitHandler
	rs.messageHandler = DefaultMessageHandler
	return rs
}

type UDPServer struct {
	SockServerBase
	conn         *net.UDPConn
	messageProxy IMessageSendReceiver
	serverMu     sync.Mutex
}

func (s *UDPServer) StartServer(params SockParams) error {
	funcName := "UDPServer.StartServer"
	s.serverMu.Lock()
	if s.running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		s.Network = params.Network
	}
	conn, err := listenUDP(s.Network, params.LocalAddress)
	if nil != err {
		defer s.serverMu.Unlock()
		return err
	}
	s.running = true
	s.conn = conn
	s.messageProxy = NewMessageSendReceiver(s.conn, s.conn, UdpListenRW, s.Network)
	s.messageProxy.SetSplitHandler(s.splitHandler)
	s.messageProxy.SetMessageHandler(s.messageHandler)
	s.serverMu.Unlock()
	log.Println(funcName + "()")
	err2 := s.messageProxy.StartReceiving()
	return err2
}

func (s *UDPServer) StopServer() error {
	funcName := "UDPServer.StopServer"
	s.serverMu.Lock()
	if !s.running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.running = false
		s.serverMu.Unlock()
	}()
	if nil != s.conn {
		s.conn.Close()
	}
	log.Println(funcName + "()")
	return nil
}

func (s *UDPServer) SendDataTo(msg []byte, rAddress ...string) error {
	funcName := "UDPServer.SendMessage"
	if !s.Running() {
		return ConnNilError(funcName)
	}
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	if len(rAddress) == 0 {
		return NoAddrError(funcName)
	}
	_, err := s.messageProxy.SendMessage(msg, rAddress...)
	return err
}

func listenUDP(network string, address string) (*net.UDPConn, error) {
	udpAddr, _ := GetUDPAddr(network, address)
	return net.ListenUDP(network, udpAddr)
}
