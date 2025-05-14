package netx

import (
	"io"
	"net"
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
	io.Closer
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
	SendBytesTo(bytes []byte, connId ...string) error
	// SendPackTo
	// 发送二进制消息包(把数据打包，补充长度信息)
	SendPackTo(data []byte, connId ...string) error
}

type ISockSenderSetter interface {
	SetSockSender(sockSender ISockSender)
}

type ISockSenderGetter interface {
	GetSockSender() ISockSender
}
