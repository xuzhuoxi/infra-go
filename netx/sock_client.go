//
//Created by xuzhuoxi
//on 2019-02-14.
//@author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

type IClient interface {
	OpenClient(params SockParams) error
	CloseClient() error
	Opening() bool
}

type ISockClient interface {
	IClient
	IPackReceiver
	ISockSender

	LocalAddress() string
}

type SockClientBase struct {
	Name     string
	Network  string
	clientMu sync.RWMutex
	opening  bool

	localAddress string
	conn         ISockConn

	PackProxy   IPackSendReceiver
	PackHandler PackHandler
}

func (c *SockClientBase) LocalAddress() string {
	return c.conn.LocalAddr().String()
}

func (c *SockClientBase) SetPackHandler(handler PackHandler) error {
	c.PackHandler = handler
	if nil != c.PackProxy {
		c.PackProxy.SetPackHandler(handler)
	}
	return nil
}

func (c *SockClientBase) IsReceiving() bool {
	return c.PackProxy.IsReceiving()
}

func (c *SockClientBase) Opening() bool {
	c.clientMu.RLock()
	defer c.clientMu.RUnlock()
	return c.opening
}

func (c *SockClientBase) SendPackTo(msg []byte, rAddress ...string) error {
	_, err := c.PackProxy.SendPack(msg, rAddress...)
	return err
}

func (c *SockClientBase) SendBytesTo(bytes []byte, rAddress ...string) error {
	_, err := c.PackProxy.SendBytes(bytes, rAddress...)
	return err
}

func (c *SockClientBase) StartReceiving() error {
	logx.Infoln(c.Name + ".StartReceiving()")
	err := c.PackProxy.StartReceiving()
	return err
}

func (c *SockClientBase) StopReceiving() error {
	logx.Infoln(c.Name + ".StopReceiving()")
	err := c.PackProxy.StopReceiving()
	return err
}

func (c *SockClientBase) setMessageProxy(packProxy IPackSendReceiver) {
	c.PackProxy = packProxy
	if nil != c.PackHandler {
		packProxy.SetPackHandler(c.PackHandler)
	}
}