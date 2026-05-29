package netx

import (
	"io"
	"net"
)

// SockParams holds socket connection parameters for client and server setup.
// SockParams 保存客户端与服务端建立连接所需的参数。
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

// ISockName provides get/set for a socket identity name.
// ISockName 提供套接字标识名称的读写。
type ISockName interface {
	// SetName sets the identity name.
	// SetName 设置标识名称。
	SetName(name string)
	// GetName returns the identity name.
	// GetName 获取标识名称。
	GetName() string
}

// ISockConn is a trimmed connection interface; method signatures must not change.
// ISockConn 是裁剪后的连接接口，函数签名不可更改。
// Targets net.Conn, quic.Conn, *net.UDPConn, and *websocket.Conn.
// 适用于 net.Conn、quic.Conn、*net.UDPConn、*websocket.Conn。
type ISockConn interface {
	io.Closer
	// LocalAddr returns the local network address.
	// LocalAddr 返回本地连接地址。
	LocalAddr() net.Addr
	// RemoteAddr returns the remote network address.
	// RemoteAddr 返回远程连接地址。
	RemoteAddr() net.Addr
}

// ISockSender sends raw bytes or framed packs to connections.
// ISockSender 向连接发送原始字节或已封包的消息。
type ISockSender interface {
	// SendBytesTo sends raw bytes without repacking.
	// SendBytesTo 发送二进制数据，不会重新打包。
	SendBytesTo(bytes []byte, connId ...string) error
	// SendPackTo sends a framed message pack with length prefix.
	// SendPackTo 发送二进制消息包（把数据打包，补充长度信息）。
	SendPackTo(data []byte, connId ...string) error
}

// ISockSenderSetter injects an ISockSender implementation.
// ISockSenderSetter 注入 ISockSender 实现。
type ISockSenderSetter interface {
	SetSockSender(sockSender ISockSender)
}

// ISockSenderGetter returns the bound ISockSender.
// ISockSenderGetter 返回已绑定的 ISockSender。
type ISockSenderGetter interface {
	GetSockSender() ISockSender
}
