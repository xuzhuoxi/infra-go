package netx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
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

func newTCPClient() ISockClient {
	return newTcpC("TCPClient", TcpNetwork)
}

func newTCP4Client() ISockClient {
	return newTcpC("TCP4Client", TcpNetwork4)
}

func newTCP6Client() ISockClient {
	return newTcpC("TCP6Client", TcpNetwork6)
}

func newTcpC(name string, network SockNetwork) ISockClient {
	client := &TCPClient{}
	client.Name = name
	client.Network = network
	client.Logger = logx.DefaultLogger()
	client.PackHandler = NewIPackHandler(nil)
	return client
}

//----------------------------

type ITCPClient interface {
	ISockClient
}

type TCPClient struct {
	SockClientBase
}

func (c *TCPClient) OpenClient(params SockParams) error {
	funcName := "TCPClient.OpenClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if "" != params.Network {
		c.Network = params.Network
	}
	conn, err := net.Dial(c.Network.String(), params.RemoteAddress)
	if nil != err {
		return err
	}
	c.conn = conn
	connProxy := &ReadWriterAdapter{Reader: conn, Writer: conn, remoteAddr: conn.RemoteAddr()}
	c.PackProxy = NewPackSendReceiver(connProxy, connProxy, c.PackHandler, TcpDataBlockHandler, c.Logger, false)
	c.opening = true
	c.Logger.Infoln(funcName + "()")
	return nil
}

func (c *TCPClient) CloseClient() error {
	funcName := "TCPClient.CloseClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if !c.opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	c.opening = false
	if nil != c.conn {
		c.conn.Close()
		c.conn = nil
	}
	c.Logger.Infoln(funcName + "()")
	return nil
}
