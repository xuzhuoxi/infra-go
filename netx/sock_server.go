//
//Created by xuzhuoxi
//on 2019-02-14.
//@author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
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
	SetServer(server ISockServer)
}

type ISockServerGetter interface {
	GetServer() ISockServer
}

type IServerRunning interface {
	Running() bool
}

type IServer interface {
	StartServer(params SockParams) error //会阻塞
	StopServer() error
}

type ISockServer interface {
	IServer
	IServerRunning
	ISockSender
	IPackHandlerSetter
	IPackHandlerGetter
	logx.ILoggerSupport
}

type SockServerBase struct {
	Name     string
	Network  string
	serverMu sync.RWMutex
	running  bool

	Logger logx.ILogger

	PackHandler IPackHandler
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

func (s *SockServerBase) GetPackHandler() IPackHandler {
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	return s.PackHandler
}

func (s *SockServerBase) SetPackHandler(packHandler IPackHandler) {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	s.PackHandler = packHandler
}

func (s *SockServerBase) dispatchServerStartedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStart, nil)
}

func (s *SockServerBase) dispatchServerStoppedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStop, nil)
}

func (s *SockServerBase) dispatchServerConnOpenEvent(dispatcher eventx.IEventDispatcher, address string) {
	dispatcher.DispatchEvent(ServerEventConnOpened, address)
}

func (s *SockServerBase) dispatchServerConnCloseEvent(dispatcher eventx.IEventDispatcher, address string) {
	dispatcher.DispatchEvent(ServerEventConnClosed, address)
}
