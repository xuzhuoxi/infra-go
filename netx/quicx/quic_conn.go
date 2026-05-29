package quicx

import (
	"errors"

	"github.com/quic-go/quic-go"
	"github.com/xuzhuoxi/infra-go/netx"
)

var errProxyNil = errors.New("QuicConn.SRProxy is ni")

// QuicSockConn is a QUIC IServerConn implementation.
// QuicSockConn 是 QUIC 的 IServerConn 实现。
type QuicSockConn struct {
	QuicConnInfo netx.IConnInfo
	SRProxy      netx.IPackSendReceiver
	Session      *quic.Conn
	Stream       *quic.Stream
}

// ClientAddress returns the remote client address.
// ClientAddress 返回客户端连接地址。
func (o *QuicSockConn) ClientAddress() string {
	return o.QuicConnInfo.GetRemoteAddress()
}

// SendBytes sends raw bytes on the QUIC stream.
// SendBytes 在 QUIC 流上发送二进制数据。
func (o *QuicSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes)
	return err
}

// SendPack sends a framed pack on the QUIC stream.
// SendPack 在 QUIC 流上发送二进制消息包。
func (o *QuicSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data)
	return err
}

// CloseConn stops receiving and closes the stream and session.
// CloseConn 停止接收并关闭流与会话。
func (o *QuicSockConn) CloseConn() error {
	err1 := o.SRProxy.StopReceiving()
	var err2, err3 error
	if o.Stream != nil {
		err2 = o.Stream.Close()
	}
	if o.Session != nil {
		err3 = o.Session.CloseWithError(0, "server conn closed")
	}
	if nil != err1 {
		return err1
	}
	if nil != err2 {
		return err2
	}
	return err3
}
