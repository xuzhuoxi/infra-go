package netx

import (
	"github.com/xuzhuoxi/util-go/errorsx"
	"log"
	"net"
)

func NewUDPDialClient() IUDPClient {
	return &UDPDialClient{SockClientBase: SockClientBase{Name: "UDPDialClient", Network: UDPNetwork}}
}

func NewUDPListenClient() IUDPClient {
	return &UDPListenClient{SockClientBase: SockClientBase{Name: "UDPListenClient", Network: UDPNetwork}}
}

//UDPDialClient
type UDPDialClient struct {
	SockClientBase
}

func (c *UDPDialClient) OpenClient(params SockParams) error {
	funcName := "UDPDialClient.OpenClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if c.opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		c.Network = params.Network
	}
	rAddr, err := GetUDPAddr(c.Network, params.RemoteAddress)
	if nil != err {
		return err
	}
	conn, cErr := net.DialUDP(c.Network, nil, rAddr)
	if nil != cErr {
		return cErr
	}
	c.conn = conn
	c.messageProxy = NewMessageSendReceiver(conn, conn, UdpDialRW, c.Network)
	c.opening = true
	log.Println(funcName + "()")
	return nil
}

func (c *UDPDialClient) CloseClient() error {
	funcName := "UDPDialClient.Close"
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

//UDPListenClient
type UDPListenClient struct {
	SockClientBase
}

func (c *UDPListenClient) OpenClient(params SockParams) error {
	funcName := "UDPListenClient.OpenClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if c.opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		c.Network = params.Network
	}
	lAddr, err := GetUDPAddr(c.Network, params.LocalAddress)
	if nil != err {
		return err
	}
	conn, cErr := net.ListenUDP(c.Network, lAddr)
	if nil != cErr {
		return cErr
	}
	c.conn = conn
	c.messageProxy = NewMessageSendReceiver(conn, conn, UdpListenRW, c.Network)
	c.opening = true
	log.Println(funcName + "()")
	return nil
}

func (c *UDPListenClient) CloseClient() error {
	funcName := "UDPListenClient.Close"
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
