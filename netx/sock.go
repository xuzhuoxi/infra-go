package netx

import (
	"net"
)

type SockParams struct {
	Network string
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
