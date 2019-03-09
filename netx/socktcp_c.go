package netx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"net"
)

func NewTCPClient() ITCPClient {
	client := &TCPClient{}
	client.Name = "TCPClient"
	client.Network = TcpNetwork
	client.Logger = logx.DefaultLogger()
	client.PackHandler = NewIPackHandler(nil)
	return client
}

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
	conn, err := net.Dial(c.Network, params.RemoteAddress)
	if nil != err {
		return err
	}
	c.conn = conn
	connProxy := &ReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
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
