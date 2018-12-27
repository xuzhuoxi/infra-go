package netx

import (
	"github.com/xuzhuoxi/util-go/errorsx"
	"log"
	"net"
)

func NewTCPClient() ITCPClient {
	client := &TCPClient{SockClientBase: SockClientBase{Name: "TCPClient", Network: TcpNetwork}}
	return client
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
	c.messageProxy = NewMessageSendReceiver(conn, conn, TcpRW, c.Network)
	c.opening = true
	log.Println(funcName + "()")
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
	log.Println(funcName + "()")
	return nil
}
