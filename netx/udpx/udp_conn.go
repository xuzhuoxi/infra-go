package udpx

import (
	"errors"

	"github.com/xuzhuoxi/infra-go/netx"
)

var errProxyNil = errors.New("UdpConn.SRProxy is ni")

// UdpSockConn is a virtual UDP IServerConn keyed by remote address.
// UdpSockConn 是以远程地址为键的 UDP 虚拟 IServerConn。
type UdpSockConn struct {
	ConnId        string
	RemoteAddress string
	SRProxy       netx.IPackSendReceiver
}

// ClientAddress returns the remote client address.
// ClientAddress 返回客户端连接地址。
func (o *UdpSockConn) ClientAddress() string {
	return o.RemoteAddress
}

// SendBytes sends raw bytes to the remote address.
// SendBytes 向远程地址发送二进制数据。
func (o *UdpSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes, o.ConnId)
	return err
}

// SendPack sends a framed pack to the remote address.
// SendPack 向远程地址发送二进制消息包。
func (o *UdpSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data, o.ConnId)
	return err
}

// CloseConn is a no-op for UDP virtual connections.
// CloseConn 对 UDP 虚拟连接无操作。
func (o *UdpSockConn) CloseConn() error {
	return nil
}
