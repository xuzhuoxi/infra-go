package netx

import (
	"net"
)

type SockParams struct {
	Network string
	//E.g
	// tcp,udp,quic:	127.0.0.1:9999
	LocalAddress string
	//E.g
	// websocket:	ws://127.0.0.1:9999
	// tcp,udp,quic:	127.0.0.1:9999
	RemoteAddress string

	//E.g: /,/echo
	WSPattern string
	//E.g: http://127.0.0.1/，最后的"/"必须
	WSOrigin string
	//E.g: ""
	WSProtocol string
}

type ISockConn interface {
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

type ISockPackSender interface {
	SendPackTo(data []byte, rAddress ...string) error
}

type ISockByteSender interface {
	SendBytesTo(bytes []byte, rAddress ...string) error
}
