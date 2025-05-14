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

type IUDPClient interface {
	netx.ISockClient
}

type UDPDialClient struct {
	netx.SockClientBase
}

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

type UDPListenClient struct {
	netx.SockClientBase
}

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
