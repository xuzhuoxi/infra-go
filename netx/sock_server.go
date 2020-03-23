//
//Created by xuzhuoxi
//on 2019-02-14.
//@author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

const (
	ServerEventStart = "netx.ServerEventStart"
	ServerEventStop  = "netx.ServerEventStop"

	ServerEventConnOpened = "netx.ServerEventConnOpened"
	ServerEventConnClosed = "netx.ServerEventConnClosed"
)

type ISockServerSetter interface {
	SetSockServer(server ISockServer)
}

type ISockServerGetter interface {
	GetSockServer() ISockServer
}

type IServerRunning interface {
	Running() bool
}

type IServer interface {
	// 启动服务，会阻塞
	StartServer(params SockParams) error
	// 停止服务，会阻塞
	StopServer() error
}

type ISockConnection interface {
	// 连接数
	Connections() int
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
	serverMu sync.RWMutex
	running  bool

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

func (s *SockServerBase) Running() bool {
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	return s.running
}

func (s *SockServerBase) GetPackHandlerContainer() IPackHandlerContainer {
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	return s.PackHandlerContainer
}

func (s *SockServerBase) SetPackHandlerContainer(packHandlerContainer IPackHandlerContainer) {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	s.PackHandlerContainer = packHandlerContainer
}

func (s *SockServerBase) dispatchServerStartedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStart, dispatcher, nil)
}

func (s *SockServerBase) dispatchServerStoppedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStop, dispatcher, nil)
}

func (s *SockServerBase) dispatchServerConnOpenEvent(dispatcher eventx.IEventDispatcher, address string) {
	dispatcher.DispatchEvent(ServerEventConnOpened, dispatcher, address)
}

func (s *SockServerBase) dispatchServerConnCloseEvent(dispatcher eventx.IEventDispatcher, address string) {
	dispatcher.DispatchEvent(ServerEventConnClosed, dispatcher, address)
}
