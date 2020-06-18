package udpx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"net"
	"sync"
)

const (
	UDPBuffLength = 4096
)

func NewUDPServer() IUDPServer {
	return newUDPServer().(IUDPServer)
}

func NewUDP4Server() IUDPServer {
	return newUDP4Server().(IUDPServer)
}

func NewUDP6Server() IUDPServer {
	return newUDP6Server().(IUDPServer)
}

func newUDPServer() netx.ISockServer {
	return newUdpS("UDPServer", netx.UDPNetwork)
}

func newUDP4Server() netx.ISockServer {
	return newUdpS("UDP4Server", netx.UDPNetwork4)
}

func newUDP6Server() netx.ISockServer {
	return newUdpS("UDP6Server", netx.UDPNetwork6)
}

func newUdpS(name string, network netx.SockNetwork) netx.ISockServer {
	server := &UDPServer{}
	server.Name = name
	server.Network = network
	server.Logger = logx.DefaultLogger()
	server.PackHandlerContainer = netx.NewIPackHandler(nil)
	return server
}

//---------------------------

type IUDPServer interface {
	netx.ISockServer
	eventx.IEventDispatcher
}

type UDPServer struct {
	eventx.EventDispatcher
	netx.SockServerBase
	lang.ChannelLimitNone

	conn         *net.UDPConn
	messageProxy netx.IPackSendReceiver
	serverMu     sync.RWMutex
}

func (s *UDPServer) StartServer(params netx.SockParams) error {
	funcName := "UDPServer.StartServer"
	s.serverMu.Lock()
	if s.Running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		s.Network = params.Network
	}
	conn, err := listenUDP(s.Network.String(), params.LocalAddress)
	if nil != err {
		defer s.serverMu.Unlock()
		return err
	}
	s.Logger.Infoln("[UDPServer] listening on:", params.LocalAddress)
	s.Running = true
	s.conn = conn
	connProxy := &UDPConnAdapter{ReadWriter: conn}
	s.messageProxy = netx.NewPackSendReceiver(connProxy, connProxy, s.PackHandlerContainer, UdpDataBlockHandler, s.Logger, true)
	s.serverMu.Unlock()
	s.Logger.Infoln(funcName + "()")
	s.DispatchServerStartedEvent(s)
	err2 := s.messageProxy.StartReceiving()
	return err2
}

func (s *UDPServer) StopServer() error {
	funcName := "UDPServer.StopServer"
	s.serverMu.Lock()
	if !s.Running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.serverMu.Unlock()
		s.Logger.Infoln(funcName + "()")
		s.DispatchServerStoppedEvent(s)
	}()
	if nil != s.conn {
		s.conn.Close()
	}
	s.Running = false
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
	if !s.Running || s.messageProxy == nil || s.conn == nil {
		return netx.ConnNilError(funcName)
	}
	if len(rAddress) == 0 {
		return netx.NoAddrError(funcName)
	}
	_, err := s.messageProxy.SendBytes(bytes, rAddress...)
	return err
}

func listenUDP(network string, address string) (*net.UDPConn, error) {
	udpAddr, _ := GetUDPAddr(network, address)
	return net.ListenUDP(network, udpAddr)
}
