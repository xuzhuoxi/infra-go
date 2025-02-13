package tcpx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"net"
)

func NewTCPClient() ITCPClient {
	return newTCPClient().(ITCPClient)
}

func NewTCP4Client() ITCPClient {
	return newTCP4Client().(ITCPClient)
}

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

type ITCPClient interface {
	netx.ISockClient
}

type TCPClient struct {
	netx.SockClientBase
}

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
	connProxy := &netx.ReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	c.PackProxy = netx.NewPackSendReceiver(connProxy, connProxy, c.PackHandler, TcpDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

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
