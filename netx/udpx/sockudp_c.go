package udpx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
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

func newUDPDialClient() netx.ISockClient {
	return newUdpDC("UDPDialClient", netx.UDPNetwork)
}

func newUDP4DialClient() netx.ISockClient {
	return newUdpDC("UDP4DialClient", netx.UDPNetwork4)
}

func newUDP6DialClient() netx.ISockClient {
	return newUdpDC("UDP6DialClient", netx.UDPNetwork6)
}

func newUDPListenClient() netx.ISockClient {
	return newUdpLC("UDPListenClient", netx.UDPNetwork)
}

func newUDP4ListenClient() netx.ISockClient {
	return newUdpLC("UDP4ListenClient", netx.UDPNetwork4)
}

func newUDP6ListenClient() netx.ISockClient {
	return newUdpLC("UDP6ListenClient", netx.UDPNetwork6)
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

type IUDPClient interface {
	netx.ISockClient
}

//UDPDialClient
type UDPDialClient struct {
	netx.SockClientBase
}

func (c *UDPDialClient) OpenClient(params netx.SockParams) error {
	funcName := "UDPDialClient.OpenClient"
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
	connProxy := &netx.ReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	c.PackProxy = netx.NewPackSendReceiver(connProxy, connProxy, c.PackHandler, UdpDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName + "()")
	return nil
}

func (c *UDPDialClient) CloseClient() error {
	funcName := "UDPDialClient.Close"
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
	c.Logger.Infoln(funcName + "()")
	return nil
}

//UDPListenClient
type UDPListenClient struct {
	netx.SockClientBase
}

func (c *UDPListenClient) OpenClient(params netx.SockParams) error {
	funcName := "UDPListenClient.OpenClient"
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
	c.PackProxy = netx.NewPackSendReceiver(connProxy, connProxy, c.PackHandler, UdpDataBlockHandler, c.Logger, true)
	c.Opening = true
	c.Logger.Infoln(funcName + "()")
	return nil
}

func (c *UDPListenClient) CloseClient() error {
	funcName := "UDPListenClient.Close"
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
	c.Logger.Infoln(funcName + "()")
	return nil
}
