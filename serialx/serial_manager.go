package serialx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
)

// NewSerialManager creates a new SerialManager instance.
// NewSerialManager 创建新的 SerialManager 实例。
func NewSerialManager() ISerialManager {
	return &SerialManager{}
}

// ISerialModule defines a serially managed module with start and stop lifecycle.
// ISerialModule 定义可按序管理的模块，具备启动与停止生命周期。
type ISerialModule interface {
	eventx.IEventDispatcher
	// StartModule starts the module asynchronously; dispatches EventOnSerialModuleStarted when done.
	// StartModule 异步启动模块，完成后分发 EventOnSerialModuleStarted 事件。
	StartModule()
	// StopModule stops the module asynchronously; dispatches EventOnSerialModuleStopped when done.
	// StopModule 异步停止模块，完成后分发 EventOnSerialModuleStopped 事件。
	StopModule()
}

// Status represents the lifecycle state of a SerialManager.
// Status 表示 SerialManager 的生命周期状态。
type Status uint8

const (
	// StatusStopped indicates the manager is idle and accepts module changes.
	// StatusStopped 表示管理器处于空闲状态，可修改模块列表。
	StatusStopped Status = iota
	// StatusStarting indicates the manager is starting modules in order.
	// StatusStarting 表示管理器正在按序启动模块。
	StatusStarting
	// StatusStarted indicates all modules have started and the manager is running.
	// StatusStarted 表示全部模块已启动，管理器处于运行状态。
	StatusStarted
	// StatusStopping indicates the manager is stopping modules in reverse order.
	// StatusStopping 表示管理器正在按逆序停止模块。
	StatusStopping
)

// ISerialManager defines sequential module registration and ordered start/stop orchestration.
// ISerialManager 定义模块的串行注册及按序启动/停止编排。
type ISerialManager interface {
	eventx.IEventDispatcher
	// AppendModule appends a module to the manager while stopped.
	// AppendModule 在停止状态下向管理器追加模块。
	AppendModule(module ISerialModule)
	// StartManager starts all modules in registration order.
	// StartManager 按注册顺序启动全部模块。
	StartManager() bool
	// StopManager stops all modules in reverse order.
	// StopManager 按逆序停止全部模块。
	StopManager() bool
}

// SerialManager orchestrates ISerialModule instances in strict start/stop order.
// SerialManager 按严格顺序编排 ISerialModule 的启动与停止。
type SerialManager struct {
	eventx.EventDispatcher
	modules []ISerialModule
	status  Status
	index   int
}

// AppendModule appends a module to the manager while stopped.
// AppendModule 在停止状态下向管理器追加模块。
func (o *SerialManager) AppendModule(module ISerialModule) {
	if o.status != StatusStopped {
		return
	}
	o.modules = append(o.modules, module)
}

// StartManager starts all modules in registration order.
// StartManager 按注册顺序启动全部模块。
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

// StopManager stops all modules in reverse order.
// StopManager 按逆序停止全部模块。
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
