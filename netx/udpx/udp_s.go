package udpx

import (
	"fmt"
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
	return newUdpS("UDP4Server", netx.UDP4Network)
}

func newUDP6Server() netx.ISockServer {
	return newUdpS("UDP6Server", netx.UDP6Network)
}

func newUdpS(name string, network netx.SockNetwork) netx.ISockServer {
	server := &UDPServer{
		logFuncNameSend: fmt.Sprintf("UDPServer[%s].SendBytesTo", name),
	}
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
	logFuncNameSend string

	conn         *net.UDPConn
	messageProxy netx.IPackSendReceiver
	serverMu     sync.RWMutex
}

func (s *UDPServer) StartServer(params netx.SockParams) error {
	funcName := fmt.Sprintf("[UDPServer(%s).StartServer]", s.Name)
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
	s.Logger.Infoln(funcName, "listening on:", params.LocalAddress)
	s.Running = true
	s.conn = conn
	connProxy := &UDPConnAdapter{ReadWriter: conn}
	connInfo := netx.NewRemoteMuiltConnInfo(conn.LocalAddr().String(), "")
	s.messageProxy = netx.NewPackSendReceiver(connInfo, connProxy, connProxy, s.PackHandlerContainer, UdpDataBlockHandler, s.Logger, true)
	s.serverMu.Unlock()
	s.Logger.Infoln(funcName, "()")
	s.DispatchServerStartedEvent(s)
	err2 := s.messageProxy.StartReceiving()
	return err2
}

func (s *UDPServer) StopServer() error {
	funcName := fmt.Sprintf("[UDPServer(%s).StopServer]", s.Name)
	s.serverMu.Lock()
	if !s.Running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.serverMu.Unlock()
		s.Logger.Infoln(funcName, "()")
		s.DispatchServerStoppedEvent(s)
	}()
	if nil != s.conn {
		s.conn.Close()
	}
	s.Running = false
	return nil
}

func (s *UDPServer) SetMaxConn(max int) {
	return
}

func (s *UDPServer) Connections() int {
	return 0
}

func (s *UDPServer) CloseConnection(connId string) (err error, ok bool) {
	return nil, false
}

// FindConnection
// connId: RemoteAddress
func (s *UDPServer) FindConnection(connId string) (conn netx.IServerConn, ok bool) {
	// Udp没有实际连接，也不保证NAT端口利用情况下的数据准确性，因此使用远程地址作为连接id
	return &UdpSockConn{ConnId: connId, RemoteAddress: connId, SRProxy: s.messageProxy}, true
}

// SendPackTo
// connId : Use RemoteAddress
func (s *UDPServer) SendPackTo(pack []byte, connId ...string) error {
	bytes := UdpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, connId...)
}

// SendBytesTo
// connId : Use RemoteAddress
func (s *UDPServer) SendBytesTo(bytes []byte, connId ...string) error {
	funcName := s.logFuncNameSend
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	if !s.Running || s.messageProxy == nil || s.conn == nil {
		return netx.ConnNilError(funcName)
	}
	if len(connId) == 0 {
		return netx.NoAddrError(funcName)
	}
	_, err := s.messageProxy.SendBytes(bytes, connId...)
	return err
}

func listenUDP(network string, localAddress string) (*net.UDPConn, error) {
	udpAddr, _ := GetUDPAddr(network, localAddress)
	return net.ListenUDP(network, udpAddr)
}
