package udpx

import (
	"fmt"
	"net"
	"sync"

	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
)

const (
	// UDPBuffLength default UDP read buffer size
	UDPBuffLength = 4096
)

// NewUDPServer creates a UDP server.
// NewUDPServer 创建 UDP 服务端。
func NewUDPServer() IUDPServer {
	return newUDPServer().(IUDPServer)
}

// NewUDP4Server creates an IPv4 UDP server.
// NewUDP4Server 创建 IPv4 UDP 服务端。
func NewUDP4Server() IUDPServer {
	return newUDP4Server().(IUDPServer)
}

// NewUDP6Server creates an IPv6 UDP server.
// NewUDP6Server 创建 IPv6 UDP 服务端。
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

// IUDPServer is a UDP-specific ISockServer with event dispatch.
// IUDPServer 是支持事件分发的 UDP 专用 ISockServer。
type IUDPServer interface {
	netx.ISockServer
	eventx.IEventDispatcher
}

// UDPServer implements IUDPServer on a single UDP socket.
// UDPServer 在单个 UDP 套接字上实现 IUDPServer。
type UDPServer struct {
	eventx.EventDispatcher
	netx.SockServerBase
	lang.ChannelLimitNone
	logFuncNameSend string

	conn         *net.UDPConn
	messageProxy netx.IPackSendReceiver
	serverMu     sync.RWMutex
}

// StartServer listens on the local UDP address and receives datagrams.
// StartServer 在本地 UDP 地址监听并接收数据报。
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

// StopServer closes the UDP socket and stops receiving.
// StopServer 关闭 UDP 套接字并停止接收。
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

// SetMaxConn is a no-op for UDP; connection count is not tracked.
// SetMaxConn 对 UDP 无操作；不跟踪连接数。
func (s *UDPServer) SetMaxConn(max int) {
	return
}

// Connections always returns 0 for UDP.
// Connections 对 UDP 始终返回 0。
func (s *UDPServer) Connections() int {
	return 0
}

// CloseConnection is unsupported for UDP.
// CloseConnection 对 UDP 不支持关闭指定连接。
func (s *UDPServer) CloseConnection(connId string) (err error, ok bool) {
	return nil, false
}

// FindConnection returns a virtual connection for the remote address connId.
// FindConnection 按远程地址 connId 返回虚拟连接。
//   connId: remote address
//   connId: 远程地址
func (s *UDPServer) FindConnection(connId string) (conn netx.IServerConn, ok bool) {
	// Udp没有实际连接，也不保证NAT端口利用情况下的数据准确性，因此使用远程地址作为连接id
	return &UdpSockConn{ConnId: connId, RemoteAddress: connId, SRProxy: s.messageProxy}, true
}

// SendPackTo sends a framed pack to remote addresses identified by connId.
// SendPackTo 向 connId 指定的远程地址发送二进制消息包。
//   connId: remote address
//   connId: 远程地址
func (s *UDPServer) SendPackTo(pack []byte, connId ...string) error {
	bytes := UdpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, connId...)
}

// SendBytesTo sends raw bytes to remote addresses identified by connId.
// SendBytesTo 向 connId 指定的远程地址发送二进制数据。
//   connId: remote address
//   connId: 远程地址
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
