// Package netx
// Created by xuzhuoxi
// on 2019-02-14.
// @author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

type ISockServerSetter interface {
	SetSockServer(server ISockServer)
}

type ISockServerGetter interface {
	GetSockServer() ISockServer
}

type IServerRunning interface {
	IsRunning() bool
}

type IServer interface {
	// StartServer
	// 启动服务，会阻塞
	StartServer(params SockParams) error
	// StopServer
	// 停止服务，会阻塞
	StopServer() error
}

type ISockConnection interface {
	// Connections
	// 连接数
	Connections() int
	// CloseConnection
	// 关闭指定连接
	CloseConnection(address string) (err error, ok bool)
}

type ISockServer interface {
	ISockName
	IServer
	IServerRunning
	lang.IChannelLimit
	ISockConnection

	ISockSender
	IPackHandlerContainerSetter
	IPackHandlerContainerGetter

	logx.ILoggerSupport
}

type SockServerBase struct {
	Name     string
	Network  SockNetwork
	ServerMu sync.RWMutex
	Running  bool

	Logger logx.ILogger

	PackHandlerContainer IPackHandlerContainer
}

func (s *SockServerBase) SetName(name string) {
	s.Name = name
}

func (s *SockServerBase) GetName() string {
	return s.Name
}

func (s *SockServerBase) GetLogger() logx.ILogger {
	return s.Logger
}

func (s *SockServerBase) SetLogger(logger logx.ILogger) {
	s.Logger = logger
}

func (s *SockServerBase) IsRunning() bool {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return s.Running
}

func (s *SockServerBase) GetPackHandlerContainer() IPackHandlerContainer {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return s.PackHandlerContainer
}

func (s *SockServerBase) SetPackHandlerContainer(packHandlerContainer IPackHandlerContainer) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	s.PackHandlerContainer = packHandlerContainer
}

func (s *SockServerBase) DispatchServerStartedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStart, dispatcher, nil)
}

func (s *SockServerBase) DispatchServerStoppedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStop, dispatcher, nil)
}

func (s *SockServerBase) DispatchServerConnOpenEvent(dispatcher eventx.IEventDispatcher, address string) {
	dispatcher.DispatchEvent(ServerEventConnOpened, dispatcher, address)
}

func (s *SockServerBase) DispatchServerConnCloseEvent(dispatcher eventx.IEventDispatcher, address string) {
	dispatcher.DispatchEvent(ServerEventConnClosed, dispatcher, address)
}
