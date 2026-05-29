package wsx

import (
	"errors"

	"github.com/xuzhuoxi/infra-go/netx"
	"golang.org/x/net/websocket"
)

var errProxyNil = errors.New("WsConn.SRProxy is ni")

// WsSockConn is a WebSocket IServerConn implementation.
// WsSockConn 是 WebSocket 的 IServerConn 实现。
type WsSockConn struct {
	WsConnInfo netx.IConnInfo
	SRProxy    netx.IPackSendReceiver
	Conn       *websocket.Conn
}

// ClientAddress returns the remote client address.
// ClientAddress 返回客户端连接地址。
func (o *WsSockConn) ClientAddress() string {
	return o.WsConnInfo.GetRemoteAddress()
}

// SendBytes sends raw bytes on the WebSocket connection.
// SendBytes 在 WebSocket 连接上发送二进制数据。
func (o *WsSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes)
	return err
}

// SendPack sends a framed pack on the WebSocket connection.
// SendPack 在 WebSocket 连接上发送二进制消息包。
func (o *WsSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data)
	return err
}

// CloseConn stops receiving and closes the WebSocket connection.
// CloseConn 停止接收并关闭 WebSocket 连接。
func (o *WsSockConn) CloseConn() error {
	var err1 error
	var err2 error
	if nil != o.SRProxy {
		err1 = o.SRProxy.StopReceiving()
	}
	if nil != o.Conn {
		err2 = o.Conn.Close()
	}
	if nil != err1 {
		return err1
	}
	return err2
}
