package tcpx

import (
	"fmt"
	"net"

	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
)

// NewTCPClient creates a TCP client.
// NewTCPClient 创建 TCP 客户端。
func NewTCPClient() ITCPClient {
	return newTCPClient().(ITCPClient)
}

// NewTCP4Client creates an IPv4 TCP client.
// NewTCP4Client 创建 IPv4 TCP 客户端。
func NewTCP4Client() ITCPClient {
	return newTCP4Client().(ITCPClient)
}

// NewTCP6Client creates an IPv6 TCP client.
// NewTCP6Client 创建 IPv6 TCP 客户端。
func NewTCP6Client() ITCPClient {
	return newTCP6Client().(ITCPClient)
}

func newTCPClient() netx.ISockClient {
	return newTcpC("TCPClient", netx.TcpNetwork)
}

func newTCP4Client() netx.ISockClient {
	return newTcpC("TCP4Client", netx.Tcp4Network)
}

func newTCP6Client() netx.ISockClient {
	return newTcpC("TCP6Client", netx.Tcp6Network)
}

func newTcpC(name string, network netx.SockNetwork) netx.ISockClient {
	client := &TCPClient{}
	client.Name = name
	client.Network = network
	client.Logger = logx.DefaultLogger()
	client.PackHandler = netx.NewIPackHandler(nil)
	return client
}

//----------------------------

// ITCPClient is a TCP-specific ISockClient.
// ITCPClient 是 TCP 专用的 ISockClient。
type ITCPClient interface {
	netx.ISockClient
}

// TCPClient implements ITCPClient over TCP connections.
// TCPClient 是基于 TCP 连接的 ITCPClient 实现。
type TCPClient struct {
	netx.SockClientBase
}

// OpenClient dials the remote TCP endpoint and starts pack I/O.
// OpenClient 拨号连接远程 TCP 端点并启动封包读写。
func (c *TCPClient) OpenClient(params netx.SockParams) error {
	funcName := fmt.Sprintf("[TCPClient(%s).OpenClient]", c.Name)
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if "" != params.Network {
		c.Network = params.Network
	}
	conn, err := net.Dial(c.Network.String(), params.RemoteAddress)
	if nil != err {
		return err
	}
	c.Conn = conn
	connInfo := netx.NewRemoteOnlyConnInfo(conn.LocalAddr().String(), conn.RemoteAddr().String())
	connProxy := &netx.ConnReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	c.PackProxy = netx.NewPackSendReceiver(connInfo, connProxy, connProxy, c.PackHandler, TcpDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

// CloseClient closes the TCP connection.
// CloseClient 关闭 TCP 连接。
func (c *TCPClient) CloseClient() error {
	funcName := fmt.Sprintf("[TCPClient(%s).CloseClient]", c.Name)
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
