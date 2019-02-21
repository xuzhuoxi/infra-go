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

type IServer interface {
	StartServer(params SockParams) error //会阻塞
	StopServer() error
	Running() bool
}

type ISockServer interface {
	IServer
	ISockSender
	IPackHandlerSetter
	logx.ILoggerSupport
}

type SockServerBase struct {
	Name     string
	Network  string
	serverMu sync.RWMutex
	running  bool

	PackHandler PackHandler
	Logger      logx.ILogger
}

func (s *SockServerBase) SetPackHandler(handler PackHandler) error {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	s.PackHandler = handler
	return nil
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
