// Package wsx
// Create on 2023/8/5
// @author xuzhuoxi
package wsx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/netx"
	"golang.org/x/net/websocket"
)

var errProxyNil = errors.New("WsConn.SRProxy is ni")

type WsSockConn struct {
	Address string
	SRProxy netx.IPackSendReceiver
	Conn    *websocket.Conn
}

func (o *WsSockConn) ClientAddress() string {
	return o.Address
}

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

func (o *WsSockConn) SendBytes(bytes []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendBytes(bytes)
	return err
}

func (o *WsSockConn) SendPack(data []byte) error {
	if nil == o.SRProxy {
		return errProxyNil
	}
	_, err := o.SRProxy.SendPack(data)
	return err
}
