package netx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"net"
	"sync"
)

const (
	UDPBuffLength = 4096
)

func NewUDPServer() IUDPServer {
	server := &UDPServer{}
	server.Name = "UDPServer"
	server.Network = UDPNetwork
	server.Logger = logx.DefaultLogger()
	server.PackHandler = DefaultPackHandler
	return server
}

type IUDPServer interface {
	ISockServer
}

type UDPServer struct {
	SockServerBase
	conn         *net.UDPConn
	messageProxy IPackSendReceiver
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
	s.Logger.Infoln("[UDPServer] listening on:", params.LocalAddress)
	s.running = true
	s.conn = conn
	connProxy := &UDPConnAdapter{ReadWriter: conn}
	s.messageProxy = NewPackSendReceiver(connProxy, connProxy, s.PackHandler, UdpDataBlockHandler, s.Logger, true)
	s.serverMu.Unlock()
	s.Logger.Infoln(funcName + "()")
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
	s.Logger.Infoln(funcName + "()")
	return nil
}

func (s *UDPServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := UdpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *UDPServer) SendBytesTo(bytes []byte, rAddress ...string) error {
	funcName := "UDPServer.SendPackTo"
	if !s.Running() {
		return ConnNilError(funcName)
	}
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	if len(rAddress) == 0 {
		return NoAddrError(funcName)
	}
	_, err := s.messageProxy.SendBytes(bytes, rAddress...)
	return err
}

func listenUDP(network string, address string) (*net.UDPConn, error) {
	udpAddr, _ := GetUDPAddr(network, address)
	return net.ListenUDP(network, udpAddr)
}
