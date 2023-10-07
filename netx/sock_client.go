// Package netx
// Created by xuzhuoxi
// on 2019-02-14.
// @author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

type ISockClientSetter interface {
	SetClient(client ISockClient)
}

type ISockClientGetter interface {
	GetClient() ISockClient
}

type IClient interface {
	// IsOpening
	// 连接是否打开中
	IsOpening() bool
	// OpenClient
	// 打开
	OpenClient(params SockParams) error
	// CloseClient
	// 关闭
	CloseClient() error
}

type ISockClient interface {
	ISockName
	IClient
	IPackReceiver
	ISockSender

	LocalAddress() string
}

type ISockEventClient interface {
	eventx.IEventDispatcher
	ISockClient
}

type SockClientBase struct {
	Name     string
	Network  SockNetwork
	ClientMu sync.RWMutex
	Opening  bool

	LocalAddr string
	Conn      ISockConn

	PackProxy IPackSendReceiver
	Logger    logx.ILogger

	PackHandler IPackHandlerContainer
}

func (o *SockClientBase) SetName(name string) {
	o.Name = name
}

func (o *SockClientBase) GetName() string {
	return o.Name
}

func (o *SockClientBase) GetPackHandlerContainer() IPackHandlerContainer {
	o.ClientMu.RLock()
	defer o.ClientMu.RUnlock()
	return o.PackHandler
}

func (o *SockClientBase) SetPackHandlerContainer(packHandler IPackHandlerContainer) {
	o.ClientMu.Lock()
	defer o.ClientMu.Unlock()
	o.PackHandler = packHandler
	if o.PackProxy != nil {
		o.PackProxy.SetPackHandlerContainer(o.PackHandler)
	}
}

func (o *SockClientBase) LocalAddress() string {
	return o.Conn.LocalAddr().String()
}

func (o *SockClientBase) IsReceiving() bool {
	return o.PackProxy.IsReceiving()
}

func (o *SockClientBase) IsOpening() bool {
	o.ClientMu.RLock()
	defer o.ClientMu.RUnlock()
	return o.Opening
}

func (o *SockClientBase) GetLogger() logx.ILogger {
	return o.Logger
}

func (o *SockClientBase) SetLogger(logger logx.ILogger) {
	o.Logger = logger
}

func (o *SockClientBase) SendPackTo(msg []byte, rAddress ...string) error {
	_, err := o.PackProxy.SendPack(msg, rAddress...)
	return err
}

func (o *SockClientBase) SendBytesTo(bytes []byte, rAddress ...string) error {
	_, err := o.PackProxy.SendBytes(bytes, rAddress...)
	return err
}

func (o *SockClientBase) StartReceiving() error {
	if nil != o.Logger {
		o.Logger.Infoln("[SockClientBase.StartReceiving]", "ServerName="+o.Name)
	}
	err := o.PackProxy.StartReceiving()
	return err
}

func (o *SockClientBase) StopReceiving() error {
	if nil != o.Logger {
		o.Logger.Infoln("[SockClientBase.StopReceiving]", "ServerName="+o.Name)
	}
	err := o.PackProxy.StopReceiving()
	return err
}
