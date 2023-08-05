// Package quicx
// Create on 2023/8/5
// @author xuzhuoxi
package quicx

import (
	"errors"
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/infra-go/netx"
)

var errProxyNil = errors.New("QuicConn.SRProxy is ni")

type QuicSockConn struct {
	Address string
	SRProxy netx.IPackSendReceiver
	Session quic.Session
	Stream  quic.Stream
}

func (o *QuicSockConn) ClientAddress() string {
	return o.Address
}

func (o *QuicSockConn) CloseConn() error {
	err1 := o.SRProxy.StopReceiving()
	err2 := o.Stream.Close()
	err3 := o.Session.Close()
	if nil != err1 {
		return err1
	}
	if nil != err2 {
		return err2
	}
	return err3
}

func (o *QuicSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes)
	return err
}

func (o *QuicSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data)
	return err
}
