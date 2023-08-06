// Package tcpx
// Create on 2023/8/5
// @author xuzhuoxi
package tcpx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/netx"
	"net"
)

var errProxyNil = errors.New("TcpConn.SRProxy is ni")

type TcpSockConn struct {
	Address string
	SRProxy netx.IPackSendReceiver
	Conn    *net.TCPConn
}

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

func (o *TcpSockConn) ClientAddress() string {
	return o.Address
}

func (o *TcpSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes)
	return err
}

func (o *TcpSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data)
	return err
}
