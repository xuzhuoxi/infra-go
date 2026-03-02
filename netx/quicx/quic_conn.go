// Package quicx
// Create on 2023/8/5
// @author xuzhuoxi
package quicx

import (
	"errors"

	"github.com/quic-go/quic-go"
	"github.com/xuzhuoxi/infra-go/netx"
)

var errProxyNil = errors.New("QuicConn.SRProxy is ni")

type QuicSockConn struct {
	QuicConnInfo netx.IConnInfo
	SRProxy      netx.IPackSendReceiver
	Session      *quic.Conn
	Stream       *quic.Stream
}

func (o *QuicSockConn) ClientAddress() string {
	return o.QuicConnInfo.GetRemoteAddress()
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
