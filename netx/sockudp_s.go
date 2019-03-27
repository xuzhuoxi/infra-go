package netx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
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
	server.PackHandler = NewIPackHandler(nil)
	return server
}

type IUDPServer interface {
	ISockServer
	eventx.IEventDispatcher
}

type UDPServer struct {
	eventx.EventDispatcher
	SockServerBase
	lang.ChannelLimitNone

	conn         *net.UDPConn
	messageProxy IPackSendReceiver
	serverMu     sync.RWMutex
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
	s.dispatchServerStartedEvent(s)
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
		s.serverMu.Unlock()
		s.Logger.Infoln(funcName + "()")
		s.dispatchServerStoppedEvent(s)
	}()
	if nil != s.conn {
		s.conn.Close()
	}
	s.running = false
	return nil
}

func (s *UDPServer) Connections() int {
	return 0
}

func (s *UDPServer) CloseConnection(address string) (err error, ok bool) {
	return nil, false
}

func (s *UDPServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := UdpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *UDPServer) SendBytesTo(bytes []byte, rAddress ...string) error {
	funcName := "UDPServer.SendPackTo"
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	if !s.running || s.messageProxy == nil || s.conn == nil {
		return ConnNilError(funcName)
	}
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
