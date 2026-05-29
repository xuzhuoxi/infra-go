package tcpx

import (
	"errors"
	"net"

	"github.com/xuzhuoxi/infra-go/netx"
)

var errProxyNil = errors.New("TcpConn.SRProxy is ni")

// TcpSockConn is a TCP IServerConn implementation.
// TcpSockConn 是 TCP 的 IServerConn 实现。
type TcpSockConn struct {
	TcpConnInfo netx.IConnInfo
	SRProxy     netx.IPackSendReceiver
	Conn        *net.TCPConn
}

// ClientAddress returns the remote client address.
// ClientAddress 返回客户端连接地址。
func (o *TcpSockConn) ClientAddress() string {
	return o.TcpConnInfo.GetRemoteAddress()
}

// SendBytes sends raw bytes on the connection.
// SendBytes 发送二进制数据。
func (o *TcpSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes)
	return err
}

// SendPack sends a framed pack on the connection.
// SendPack 发送二进制消息包。
func (o *TcpSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data)
	return err
}

// CloseConn stops receiving and closes the TCP connection.
// CloseConn 停止接收并关闭 TCP 连接。
func (o *TcpSockConn) CloseConn() error {
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
