package netx

import (
	"errors"
	"net"
	"strings"
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

// ParseSockNetwork 字符串转 SockNetwork
func ParseSockNetwork(str string) SockNetwork {
	if "" == str {
		return Undefined
	}
	lStr := strings.ToLower(str)
	switch lStr {
	case QuicNetwork.String():
		return QuicNetwork
	case TcpNetwork.String():
		return TcpNetwork
	case TcpNetwork4.String():
		return TcpNetwork4
	case TcpNetwork6.String():
		return TcpNetwork6
	case UDPNetwork.String():
		return UDPNetwork
	case UDPNetwork4.String():
		return UDPNetwork4
	case UDPNetwork6.String():
		return UDPNetwork6
	case WSNetwork.String():
		return WSNetwork
	default:
		return Undefined
	}
}

const (
	Undefined SockNetwork = ""

	QuicNetwork SockNetwork = "quic"

	TcpNetwork  SockNetwork = "tcp"
	TcpNetwork4 SockNetwork = "tcp4"
	TcpNetwork6 SockNetwork = "tcp6"

	UDPNetwork  SockNetwork = "udp"
	UDPNetwork4 SockNetwork = "udp4"
	UDPNetwork6 SockNetwork = "udp6"

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

type ISockName interface {
	// SetName
	// 设置标识名称
	SetName(name string)
	// GetName
	// 获取标识名称
	GetName() string
}

// ISockConn
// 这个是裁剪接口，函数签名不能改
// 针对net.Conn、quic.Session、*net.UDPConn、*websocket.Conn
type ISockConn interface {
	// Close
	// 关闭
	Close() error
	// LocalAddr
	// 本地连接地址
	LocalAddr() net.Addr
	// RemoteAddr
	// 远程连接地址
	RemoteAddr() net.Addr
}

type ISockSender interface {
	// SendBytesTo
	// 发送二进制数据, 不会重新打包
	SendBytesTo(bytes []byte, rAddress ...string) error
	// SendPackTo
	// 发送二进制消息包(把数据打包，补充长度信息)
	SendPackTo(data []byte, rAddress ...string) error
}

type ISockSenderSetter interface {
	SetSockSender(sockSender ISockSender)
}

type ISockSenderGetter interface {
	GetSockSender() ISockSender
}

//---------------------------

func RegisterNetwork(network SockNetwork, newServer func() ISockServer, newClient func() ISockClient, newClient2 func() ISockClient) {
	networkServerMap[network] = newServer
	networkClientMap[network] = newClient
	networkClient2Map[network] = newClient2
}
