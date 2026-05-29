package netx

import (
	"errors"
	"strings"
)

const (
	Undefined SockNetwork = ""

	QuicNetwork SockNetwork = "quic" // "quic" — quic 协议

	TcpNetwork  SockNetwork = "tcp"  // "tcp" — TCP 协议
	Tcp4Network SockNetwork = "tcp4" // "tcp4" — IPv4 的 TCP 协议
	Tcp6Network SockNetwork = "tcp6" // "tcp6" — IPv6 的 TCP 协议

	UDPNetwork  SockNetwork = "udp"  // "udp" — UDP 协议
	UDP4Network SockNetwork = "udp4" // "udp4" — IPv4 的 UDP 协议
	UDP6Network SockNetwork = "udp6" // "udp6" — IPv6 的 UDP 协议

	WSNetwork SockNetwork = "websocket" // "websocket" — WebSocket 协议
)

var (
	networkServerMap = make(map[SockNetwork]func() ISockServer)

	networkClientMap  = make(map[SockNetwork]func() ISockClient)
	networkClient2Map = make(map[SockNetwork]func() ISockClient)

	errNetworkRegister = errors.New("Network is not registered. ")
)

// SockNetwork identifies a registered socket protocol.
// SockNetwork 标识已注册的套接字协议。
type SockNetwork string

// String returns the network name.
// String 返回网络协议名称。
func (n SockNetwork) String() string {
	return string(n)
}

// NewServer creates a new server for this network.
// NewServer 为该网络协议创建服务端实例。
func (n SockNetwork) NewServer() (server ISockServer, err error) {
	if f, ok := networkServerMap[n]; ok {
		return f(), nil
	}
	return nil, errNetworkRegister
}

// NewClient creates a new client for this network.
// NewClient 为该网络协议创建客户端实例。
func (n SockNetwork) NewClient() (server ISockClient, err error) {
	if f, ok := networkClientMap[n]; ok {
		return f(), nil
	}
	return nil, errNetworkRegister
}

// NewClient2 creates an alternate client implementation for this network.
// NewClient2 为该网络协议创建备用客户端实例。
func (n SockNetwork) NewClient2() (server ISockClient, err error) {
	if f, ok := networkClient2Map[n]; ok {
		return f(), nil
	}
	return nil, errNetworkRegister
}

// ParseSockNetwork parses a string into SockNetwork.
// ParseSockNetwork 将字符串解析为 SockNetwork。
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
	case Tcp4Network.String():
		return Tcp4Network
	case Tcp6Network.String():
		return Tcp6Network
	case UDPNetwork.String():
		return UDPNetwork
	case UDP4Network.String():
		return UDP4Network
	case UDP6Network.String():
		return UDP6Network
	case WSNetwork.String():
		return WSNetwork
	default:
		return Undefined
	}
}

// RegisterNetwork registers server and client factories for a network.
// RegisterNetwork 为指定网络协议注册服务端与客户端工厂函数。
func RegisterNetwork(network SockNetwork, newServer func() ISockServer, newClient func() ISockClient, newClient2 func() ISockClient) {
	networkServerMap[network] = newServer
	networkClientMap[network] = newClient
	networkClient2Map[network] = newClient2
}
