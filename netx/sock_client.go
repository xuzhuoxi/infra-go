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

type ISockClientSetter interface {
	SetClient(client ISockClient)
}

type ISockClientGetter interface {
	GetClient() ISockClient
}

type IClientOpening interface {
	Opening() bool
}

type IClient interface {
	// 打开
	OpenClient(params SockParams) error
	// 关闭
	CloseClient() error
}

type ISockClient interface {
	ISockName
	IClient
	IClientOpening
	IPackReceiver
	ISockSender

	LocalAddress() string
}

type SockClientBase struct {
	Name     string
	Network  SockNetwork
	clientMu sync.RWMutex
	opening  bool

	localAddress string
	conn         ISockConn

	PackProxy IPackSendReceiver
	Logger    logx.ILogger

	PackHandler IPackHandlerContainer
}

func (c *SockClientBase) SetName(name string) {
	c.Name = name
}

func (c *SockClientBase) GetName() string {
	return c.Name
}

func (c *SockClientBase) GetPackHandlerContainer() IPackHandlerContainer {
	c.clientMu.RLock()
	defer c.clientMu.RUnlock()
	return c.PackHandler
}

func (c *SockClientBase) SetPackHandlerContainer(packHandler IPackHandlerContainer) {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	c.PackHandler = packHandler
	if c.PackProxy != nil {
		c.PackProxy.SetPackHandlerContainer(c.PackHandler)
	}
}

func (c *SockClientBase) LocalAddress() string {
	return c.conn.LocalAddr().String()
}

func (c *SockClientBase) IsReceiving() bool {
	return c.PackProxy.IsReceiving()
}

func (c *SockClientBase) Opening() bool {
	c.clientMu.RLock()
	defer c.clientMu.RUnlock()
	return c.opening
}

func (c *SockClientBase) GetLogger() logx.ILogger {
	return c.Logger
}

func (s *SockClientBase) SetLogger(logger logx.ILogger) {
	s.Logger = logger
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
	c.Logger.Infoln(c.Name + ".StartReceiving()")
	err := c.PackProxy.StartReceiving()
	return err
}

func (c *SockClientBase) StopReceiving() error {
	c.Logger.Infoln(c.Name + ".StopReceiving()")
	err := c.PackProxy.StopReceiving()
	return err
}
