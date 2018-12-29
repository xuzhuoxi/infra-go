package netx

import (
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
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
	connProxy := &ReadWriterProxy{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	c.messageProxy = NewMessageSendReceiver(connProxy, connProxy, UdpDialRW, c.Network)
	c.opening = true
	logx.Infoln(funcName + "()")
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
	logx.Infoln(funcName + "()")
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
	connProxy := &UDPListenReadWriterProxy{ReadWriter: conn}
	c.messageProxy = NewMessageSendReceiver(connProxy, connProxy, UdpListenRW, c.Network)
	c.opening = true
	logx.Infoln(funcName + "()")
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
	logx.Infoln(funcName + "()")
	return nil
}
