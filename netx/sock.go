package netx

import (
	"errors"
	"net"
)

var errNetworkRegister = errors.New("Network is not registered. ")

type SockNetwork string

func (n SockNetwork) String() string {
	return string(n)
}

func (n SockNetwork) NewServer() (server ISockServer, err error) {
	if f, ok := networkServerMap[n]; ok {
		return f(), nil
	}
	return nil, errNetworkRegister
}

func (n SockNetwork) NewClient() (server ISockClient, err error) {
	if f, ok := networkClientMap[n]; ok {
		return f(), nil
	}
	return nil, errNetworkRegister
}

func (n SockNetwork) NewClient2() (server ISockClient, err error) {
	if f, ok := networkClient2Map[n]; ok {
		return f(), nil
	}
	return nil, errNetworkRegister
}

const (
	QuicNetwork SockNetwork = "quic"

	TcpNetwork  SockNetwork = "tcp"
	TcpNetwork4             = "tcp4"
	TcpNetwork6             = "tcp6"

	UDPNetwork  SockNetwork = "udp"
	UDPNetwork4             = "udp4"
	UDPNetwork6             = "udp6"

	WSNetwork SockNetwork = "websocket"
)

var (
	networkServerMap = make(map[SockNetwork]func() ISockServer)

	networkClientMap  = make(map[SockNetwork]func() ISockClient)
	networkClient2Map = make(map[SockNetwork]func() ISockClient)
)

type SockParams struct {
	Network SockNetwork
	// E.g
	// tcp,udp,quic:	127.0.0.1:9999
	LocalAddress string
	// E.g
	// websocket:	ws://127.0.0.1:9999
	// tcp,udp,quic:	127.0.0.1:9999
	RemoteAddress string

	// E.g: /,/echo
	WSPattern string
	// E.g: http://127.0.0.1/，最后的"/"必须
	WSOrigin string
	// E.g: ""
	WSProtocol string
}

type ISockConn interface {
	// 关闭
	Close() error
	// 本地连接地址
	LocalAddr() net.Addr
	// 远程连接地址
	RemoteAddr() net.Addr
}

type ISockSender interface {
	// 发送二进制消息包(把数据打包，补充长度信息)
	SendPackTo(data []byte, rAddress ...string) error
	// 发送二进制数据
	SendBytesTo(bytes []byte, rAddress ...string) error
}

type ISockName interface {
	// 设置标识名称
	SetName(name string)
	// 获取标识名称
	GetName() string
}

//---------------------------

func RegisterNetwork(network SockNetwork, newServer func() ISockServer, newClient func() ISockClient, newClient2 func() ISockClient) {
	networkServerMap[network] = newServer
	networkClientMap[network] = newClient
	networkClient2Map[network] = newClient2
}

func init() {
	RegisterNetwork(QuicNetwork, newQuicServer, newQUICClient, newQUICClient)

	RegisterNetwork(TcpNetwork, newTCPServer, newTCPClient, newTCPClient)
	RegisterNetwork(TcpNetwork4, newTCP4Server, newTCP4Client, newTCP4Client)
	RegisterNetwork(TcpNetwork6, newTCP6Server, newTCP6Client, newTCP6Client)

	RegisterNetwork(UDPNetwork, newUDPServer, newUDPDialClient, newUDPListenClient)
	RegisterNetwork(UDPNetwork4, newUDP4Server, newUDP4DialClient, newUDP4ListenClient)
	RegisterNetwork(UDPNetwork6, newUDP6Server, newUDP6DialClient, newUDP6ListenClient)

	RegisterNetwork(WSNetwork, newWebSocketServer, newWebSocketClient, newWebSocketClient)
}
