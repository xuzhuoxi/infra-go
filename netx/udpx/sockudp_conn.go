// Package udpx
// Create on 2023/8/5
// @author xuzhuoxi
package udpx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/netx"
)

var errProxyNil = errors.New("UdpConn.SRProxy is ni")

type UdpSockConn struct {
	ConnId        string
	RemoteAddress string
	SRProxy       netx.IPackSendReceiver
}

func (o *UdpSockConn) ClientAddress() string {
	return o.RemoteAddress
}

func (o *UdpSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes, o.ConnId)
	return err
}

func (o *UdpSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data, o.ConnId)
	return err
}

func (o *UdpSockConn) CloseConn() error {
	return nil
}
