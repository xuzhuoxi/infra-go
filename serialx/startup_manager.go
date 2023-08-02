// Package serialx
// Create on 2023/7/2
// @author xuzhuoxi
package serialx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
)

func NewStartupManager() IStartupManager {
	return &StartupManager{}
}

type IStartupModule interface {
	eventx.IEventDispatcher
	Name() string
	StartModule()
	StopModule()
	SaveModule()
}

type IStartupManager interface {
	eventx.IEventDispatcher
	// AppendModule 添加模块
	AppendModule(module IStartupModule)
	// StartManager 启动
	StartManager() bool
	// SaveManager 保存状态
	SaveManager() bool
	// StopManager 停止
	StopManager() bool
	// RebootManager 重启
	RebootManager() bool
}

type StartupManager struct {
	eventx.EventDispatcher
	modules     []IStartupModule
	index       int
	running     bool
	saving      bool
	rebooting   bool
	shutdowning bool
}

func (o *StartupManager) AppendModule(module IStartupModule) {
	if nil == module || o.running || o.saving || o.rebooting || o.shutdowning {
		return
	}
	o.modules = append(o.modules, module)
}

func (o *StartupManager) StartManager() bool {
	if o.running || o.rebooting {
		return false
	}
	o.running, o.index = true, 0
	o.startModule()
	return true
}

func (o *StartupManager) startModule() {
	if o.index >= len(o.modules) {
		o.DispatchEvent(EventOnStartupManagerStarted, o, nil)
		return
	}
	m := o.modules[o.index]
	m.OnceEventListener(EventOnStartupModuleStarted, o.onModuleStarted)
	m.StartModule()
}

func (o *StartupManager) onModuleStarted(evd *eventx.EventData) {
	o.index += 1
	o.startModule()
}

func (o *StartupManager) SaveManager() bool {
	if !o.running || o.saving {
		return false
	}
	o.saving, o.index = true, 0
	o.saveModule()
	return true
}

func (o *StartupManager) saveModule() {
	if o.index == len(o.modules) {
		o.saving = false
		o.DispatchEvent(EventOnStartupManagerSaved, o, nil)
		return
	}
	m := o.modules[o.index]
	m.OnceEventListener(EventOnStartupModuleSaved, o.onModuleSaved)
	m.SaveModule()
}

func (o *StartupManager) onModuleSaved(evd *eventx.EventData) {
	o.index += 1
	o.saveModule()
}

func (o *StartupManager) StopManager() bool {
	if !o.running {
		return false
	}
	o.shutdowning = true
	o.prepareShutdown()
	return true
}

func (o *StartupManager) prepareShutdown() {
	o.OnceEventListener(EventOnStartupManagerSaved, o.onManagerShutdownSaved)
	if !o.saving {
		o.saving, o.index = true, 0
		o.saveModule()
	}
}

func (o *StartupManager) onManagerShutdownSaved(evd *eventx.EventData) {
	o.index = len(o.modules) - 1
	o.shutdownModule()
}

func (o *StartupManager) shutdownModule() {
	if o.index < 0 {
		o.running, o.shutdowning = false, false
		o.DispatchEvent(EventOnStartupManagerStopped, o, nil)
		return
	}
	m := o.modules[o.index]
	m.OnceEventListener(EventOnStartupModuleStopped, o.onModuleStopped)
	m.StopModule()
}

func (o *StartupManager) onModuleStopped(evd *eventx.EventData) {
	o.index -= 1
	o.shutdownModule()
}

func (o *StartupManager) RebootManager() bool {
	if !o.running || o.rebooting || o.shutdowning {
		return false
	}
	o.rebooting = true
	o.reboot()
	return true
}

func (o *StartupManager) reboot() {
	o.OnceEventListener(EventOnStartupManagerStopped, o.onManagerShutdownFinish)
	if !o.shutdowning {
		o.shutdowning = true
		o.prepareShutdown()
	}
}

func (o *StartupManager) onManagerShutdownFinish(evd *eventx.EventData) {
	o.OnceEventListener(EventOnStartupManagerStarted, o.onManagerStartupFinish)
	o.running, o.index = true, 0
	o.startModule()
}

func (o *StartupManager) onManagerStartupFinish(evd *eventx.EventData) {
	o.rebooting = false
	o.DispatchEvent(EventOnStartupManagerRebooted, o, nil)
}
