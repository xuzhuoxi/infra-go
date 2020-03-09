package netx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"net"
)

func NewUDPDialClient() IUDPClient {
	return newUDPDialClient().(IUDPClient)
}

func NewUDP4DialClient() IUDPClient {
	return newUDP4DialClient().(IUDPClient)
}

func NewUDP6DialClient() IUDPClient {
	return newUDP6DialClient().(IUDPClient)
}

func NewUDPListenClient() IUDPClient {
	return newUDPListenClient().(IUDPClient)
}

func NewUDP4ListenClient() IUDPClient {
	return newUDP4ListenClient().(IUDPClient)
}

func NewUDP6ListenClient() IUDPClient {
	return newUDP6ListenClient().(IUDPClient)
}

func newUDPDialClient() ISockClient {
	return newUdpDC("UDPDialClient", UDPNetwork)
}

func newUDP4DialClient() ISockClient {
	return newUdpDC("UDP4DialClient", UDPNetwork4)
}

func newUDP6DialClient() ISockClient {
	return newUdpDC("UDP6DialClient", UDPNetwork6)
}

func newUDPListenClient() ISockClient {
	return newUdpLC("UDPListenClient", UDPNetwork)
}

func newUDP4ListenClient() ISockClient {
	return newUdpLC("UDP4ListenClient", UDPNetwork4)
}

func newUDP6ListenClient() ISockClient {
	return newUdpLC("UDP6ListenClient", UDPNetwork6)
}

func newUdpDC(name string, network SockNetwork) ISockClient {
	client := &UDPDialClient{}
	client.Name = name
	client.Network = network
	client.Logger = logx.DefaultLogger()
	client.PackHandler = NewIPackHandler(nil)
	return client
}

func newUdpLC(name string, network SockNetwork) ISockClient {
	client := &UDPListenClient{}
	client.Name = name
	client.Network = network
	client.Logger = logx.DefaultLogger()
	client.PackHandler = NewIPackHandler(nil)
	return client
}

//---------------------------

type IUDPClient interface {
	ISockClient
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
	rAddr, err := GetUDPAddr(c.Network.String(), params.RemoteAddress)
	if nil != err {
		return err
	}
	conn, cErr := net.DialUDP(c.Network.String(), nil, rAddr)
	if nil != cErr {
		return cErr
	}
	c.conn = conn
	connProxy := &ReadWriterAdapter{Reader: conn, Writer: conn, remoteAddr: conn.RemoteAddr()}
	c.PackProxy = NewPackSendReceiver(connProxy, connProxy, c.PackHandler, UdpDataBlockHandler, c.Logger, false)
	c.opening = true
	c.Logger.Infoln(funcName + "()")
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
	c.Logger.Infoln(funcName + "()")
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
	lAddr, err := GetUDPAddr(c.Network.String(), params.LocalAddress)
	if nil != err {
		return err
	}
	conn, cErr := net.ListenUDP(c.Network.String(), lAddr)
	if nil != cErr {
		return cErr
	}
	c.conn = conn
	connProxy := &UDPConnAdapter{ReadWriter: conn}
	c.PackProxy = NewPackSendReceiver(connProxy, connProxy, c.PackHandler, UdpDataBlockHandler, c.Logger, true)
	c.opening = true
	c.Logger.Infoln(funcName + "()")
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
	c.Logger.Infoln(funcName + "()")
	return nil
}
