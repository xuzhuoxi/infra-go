package udpx

import (
	"net"

	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
)

// NewUDPDialClient creates a UDP client that dials a remote address.
// NewUDPDialClient 创建拨号连接远程地址的 UDP 客户端。
func NewUDPDialClient() IUDPClient {
	return newUDPDialClient().(IUDPClient)
}

// NewUDP4DialClient creates an IPv4 UDP dial client.
// NewUDP4DialClient 创建 IPv4 UDP 拨号客户端。
func NewUDP4DialClient() IUDPClient {
	return newUDP4DialClient().(IUDPClient)
}

// NewUDP6DialClient creates an IPv6 UDP dial client.
// NewUDP6DialClient 创建 IPv6 UDP 拨号客户端。
func NewUDP6DialClient() IUDPClient {
	return newUDP6DialClient().(IUDPClient)
}

// NewUDPListenClient creates a UDP client that listens on a local address.
// NewUDPListenClient 创建在本地地址监听的 UDP 客户端。
func NewUDPListenClient() IUDPClient {
	return newUDPListenClient().(IUDPClient)
}

// NewUDP4ListenClient creates an IPv4 UDP listen client.
// NewUDP4ListenClient 创建 IPv4 UDP 监听客户端。
func NewUDP4ListenClient() IUDPClient {
	return newUDP4ListenClient().(IUDPClient)
}

// NewUDP6ListenClient creates an IPv6 UDP listen client.
// NewUDP6ListenClient 创建 IPv6 UDP 监听客户端。
func NewUDP6ListenClient() IUDPClient {
	return newUDP6ListenClient().(IUDPClient)
}

func newUDPDialClient() netx.ISockClient {
	return newUdpDC("UDPDialClient", netx.UDPNetwork)
}

func newUDP4DialClient() netx.ISockClient {
	return newUdpDC("UDP4DialClient", netx.UDP4Network)
}

func newUDP6DialClient() netx.ISockClient {
	return newUdpDC("UDP6DialClient", netx.UDP6Network)
}

func newUDPListenClient() netx.ISockClient {
	return newUdpLC("UDPListenClient", netx.UDPNetwork)
}

func newUDP4ListenClient() netx.ISockClient {
	return newUdpLC("UDP4ListenClient", netx.UDP4Network)
}

func newUDP6ListenClient() netx.ISockClient {
	return newUdpLC("UDP6ListenClient", netx.UDP6Network)
}

func newUdpDC(name string, network netx.SockNetwork) netx.ISockClient {
	client := &UDPDialClient{}
	client.Name = name
	client.Network = network
	client.Logger = logx.DefaultLogger()
	client.PackHandler = netx.NewIPackHandler(nil)
	return client
}

func newUdpLC(name string, network netx.SockNetwork) netx.ISockClient {
	client := &UDPListenClient{}
	client.Name = name
	client.Network = network
	client.Logger = logx.DefaultLogger()
	client.PackHandler = netx.NewIPackHandler(nil)
	return client
}

//---------------------------

// IUDPClient is a UDP-specific ISockClient.
// IUDPClient 是 UDP 专用的 ISockClient。
type IUDPClient interface {
	netx.ISockClient
}

// UDPDialClient is a UDP client that dials a remote endpoint.
// UDPDialClient 是拨号连接远程端点的 UDP 客户端。
type UDPDialClient struct {
	netx.SockClientBase
}

// OpenClient dials the remote UDP endpoint.
// OpenClient 拨号连接远程 UDP 端点。
func (c *UDPDialClient) OpenClient(params netx.SockParams) error {
	funcName := "[UDPDialClient.OpenClient]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if c.Opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		c.Network = params.Network
	}
	rAddr, err := GetUDPAddr(c.Network.String(), params.RemoteAddress)
	if nil != err {
		return err
	}
	conn, cErr := net.DialUDP(c.Network.String(), nil, rAddr)
	if nil != cErr {
		return cErr
	}
	c.Conn = conn
	connAdapter := &netx.ConnReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()} // 由于这里是客户端，因此是对服务器的一对一连接，不需要使用 UDPConnAdapter
	connInfo := netx.NewRemoteOnlyConnInfo(conn.LocalAddr().String(), conn.RemoteAddr().String())
	c.PackProxy = netx.NewPackSendReceiver(connInfo, connAdapter, connAdapter, c.PackHandler, UdpDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

// CloseClient closes the UDP dial connection.
// CloseClient 关闭 UDP 拨号连接。
func (c *UDPDialClient) CloseClient() error {
	funcName := "[UDPDialClient.Close]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if !c.Opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	c.Opening = false
	if nil != c.Conn {
		c.Conn.Close()
		c.Conn = nil
	}
	c.Logger.Infoln(funcName, "()")
	return nil
}

// UDPListenClient is a UDP client that listens on a local address for datagrams.
// UDPListenClient 是在本地地址监听数据报的 UDP 客户端。
type UDPListenClient struct {
	netx.SockClientBase
}

// OpenClient listens on the local UDP address.
// OpenClient 在本地 UDP 地址上监听。
func (c *UDPListenClient) OpenClient(params netx.SockParams) error {
	funcName := "[UDPListenClient.OpenClient]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if c.Opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		c.Network = params.Network
	}
	lAddr, err := GetUDPAddr(c.Network.String(), params.LocalAddress)
	if nil != err {
		return err
	}
	conn, cErr := net.ListenUDP(c.Network.String(), lAddr)
	if nil != cErr {
		return cErr
	}
	c.Conn = conn
	connProxy := &UDPConnAdapter{ReadWriter: conn}
	connInfo := netx.NewRemoteOnlyConnInfo(conn.LocalAddr().String(), conn.RemoteAddr().String())
	c.PackProxy = netx.NewPackSendReceiver(connInfo, connProxy, connProxy, c.PackHandler, UdpDataBlockHandler, c.Logger, true)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

// CloseClient closes the UDP listen socket.
// CloseClient 关闭 UDP 监听套接字。
func (c *UDPListenClient) CloseClient() error {
	funcName := "[UDPListenClient.Close]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if !c.Opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	c.Opening = false
	if nil != c.Conn {
		c.Conn.Close()
		c.Conn = nil
	}
	c.Logger.Infoln(funcName, "()")
	return nil
}
