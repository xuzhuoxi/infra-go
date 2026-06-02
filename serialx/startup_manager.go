package serialx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
)

// NewStartupManager creates a new StartupManager instance.
// NewStartupManager 创建新的 StartupManager 实例。
func NewStartupManager() IStartupManager {
	return &StartupManager{}
}

// IStartupModule defines a startup module with name, lifecycle, and save hooks.
// IStartupModule 定义具备名称、生命周期与保存钩子的启动模块。
type IStartupModule interface {
	eventx.IEventDispatcher
	// Name returns the module name.
	// Name 返回模块名称。
	Name() string
	// StartModule starts the module asynchronously; dispatches EventOnStartupModuleStarted when done.
	// StartModule 异步启动模块，完成后分发 EventOnStartupModuleStarted 事件。
	StartModule()
	// StopModule stops the module asynchronously; dispatches EventOnStartupModuleStopped when done.
	// StopModule 异步停止模块，完成后分发 EventOnStartupModuleStopped 事件。
	StopModule()
	// SaveModule persists module state asynchronously; dispatches EventOnStartupModuleSaved when done.
	// SaveModule 异步保存模块状态，完成后分发 EventOnStartupModuleSaved 事件。
	SaveModule()
}

// IStartupManager defines sequential module registration and ordered start, save, stop, and reboot orchestration.
// IStartupManager 定义模块的串行注册及按序启动、保存、停止与重启编排。
type IStartupManager interface {
	eventx.IEventDispatcher
	// AppendModule appends a module to the manager while idle.
	// AppendModule 在空闲状态下向管理器追加模块。
	AppendModule(module IStartupModule)
	// StartManager starts all modules in registration order.
	// StartManager 按注册顺序启动全部模块。
	StartManager() bool
	// SaveManager saves all modules in registration order.
	// SaveManager 按注册顺序保存全部模块状态。
	SaveManager() bool
	// StopManager saves state then stops all modules in reverse order.
	// StopManager 先保存状态再按逆序停止全部模块。
	StopManager() bool
	// RebootManager stops then restarts all modules.
	// RebootManager 停止后重新启动全部模块。
	RebootManager() bool
}

// StartupManager orchestrates IStartupModule instances with ordered start, save, stop, and reboot.
// StartupManager 按序编排 IStartupModule 的启动、保存、停止与重启。
type StartupManager struct {
	eventx.EventDispatcher
	modules     []IStartupModule
	index       int
	running     bool
	saving      bool
	rebooting   bool
	shutdowning bool
}

// AppendModule appends a module to the manager while idle.
// AppendModule 在空闲状态下向管理器追加模块。
func (o *StartupManager) AppendModule(module IStartupModule) {
	if nil == module || o.running || o.saving || o.rebooting || o.shutdowning {
		return
	}
	o.modules = append(o.modules, module)
}

// StartManager starts all modules in registration order.
// StartManager 按注册顺序启动全部模块。
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

// SaveManager saves all modules in registration order.
// SaveManager 按注册顺序保存全部模块状态。
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

// StopManager saves state then stops all modules in reverse order.
// StopManager 先保存状态再按逆序停止全部模块。
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

// RebootManager stops then restarts all modules.
// RebootManager 停止后重新启动全部模块。
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
