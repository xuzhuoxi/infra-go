// Package serialx
// Create on 2023/7/2
// @author xuzhuoxi
package serialx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
)

func NewSerialManager() ISerialManager {
	return &SerialManager{}
}

type ISerialModule interface {
	eventx.IEventDispatcher
	StartModule()
	StopModule()
}

type Status uint8

const (
	StatusStopped Status = iota
	StatusStarting
	StatusStarted
	StatusStopping
)

type ISerialManager interface {
	eventx.IEventDispatcher
	// AppendModule 添加模块
	AppendModule(module ISerialModule)
	// StartManager 启动
	StartManager() bool
	// StopManager 停止
	StopManager() bool
}

type SerialManager struct {
	eventx.EventDispatcher
	modules []ISerialModule
	status  Status
	index   int
}

func (o *SerialManager) AppendModule(module ISerialModule) {
	if o.status != StatusStopped {
		return
	}
	o.modules = append(o.modules, module)
}

func (o *SerialManager) StartManager() bool {
	if o.status != StatusStopped {
		return false
	}
	o.status, o.index = StatusStarting, 0
	o.startModule()
	return true
}

func (o *SerialManager) startModule() {
	if o.index >= len(o.modules) {
		o.status = StatusStarted
		o.DispatchEvent(EventOnSerialManagerStarted, o, nil)
		return
	}
	m := o.modules[o.index]
	m.OnceEventListener(EventOnSerialModuleStarted, o.onModuleStarted)
	m.StartModule()
}

func (o *SerialManager) onModuleStarted(evd *eventx.EventData) {
	o.index += 1
	o.startModule()
}

func (o *SerialManager) StopManager() bool {
	if o.status != StatusStarted {
		return false
	}
	o.status, o.index = StatusStopping, len(o.modules)-1
	o.shutdownModule()
	return true
}

func (o *SerialManager) shutdownModule() {
	if o.index < 0 {
		o.status = StatusStopped
		o.DispatchEvent(EventOnSerialManagerStopped, o, nil)
		return
	}
	m := o.modules[o.index]
	m.OnceEventListener(EventOnSerialModuleStopped, o.onModuleStopped)
	m.StopModule()
}

func (o *SerialManager) onModuleStopped(evd *eventx.EventData) {
	o.index -= 1
	o.shutdownModule()
}
