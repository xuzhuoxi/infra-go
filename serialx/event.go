package serialx

// ISerialManager Events
const (
	// Module Events ---------- ---------- ---------- ---------- ----------

	// EventOnSerialModuleStarted is dispatched when a serial module finishes starting.
	// EventOnSerialModuleStarted 在串行模块完成启动时分发。
	EventOnSerialModuleStarted = "Serial-Module:OnObserverStarted"
	// EventOnSerialModuleStopped is dispatched when a serial module finishes stopping.
	// EventOnSerialModuleStopped 在串行模块完成停止时分发。
	EventOnSerialModuleStopped = "Serial-Module:OnObserverStopped"

	// Manager Events ---------- ---------- ---------- ---------- ----------

	// EventOnSerialManagerStarted is dispatched when all serial modules have started.
	// EventOnSerialManagerStarted 在全部串行模块启动完成时分发。
	EventOnSerialManagerStarted = "Serial-Manager:OnManagerStarted"
	// EventOnSerialManagerStopped is dispatched when all serial modules have stopped.
	// EventOnSerialManagerStopped 在全部串行模块停止完成时分发。
	EventOnSerialManagerStopped = "Serial-Manager:OnManagerStopped"
)

// IStartupManager Events
const (
	// Module Events ---------- ---------- ---------- ---------- ----------

	// EventOnStartupModuleSaved is dispatched when a startup module finishes saving.
	// EventOnStartupModuleSaved 在启动模块完成保存时分发。
	EventOnStartupModuleSaved = "Startup-Module:OnStartupModuleSaved"
	// EventOnStartupModuleStarted is dispatched when a startup module finishes starting.
	// EventOnStartupModuleStarted 在启动模块完成启动时分发。
	EventOnStartupModuleStarted = "Startup-Module:OnStartupModuleStarted"
	// EventOnStartupModuleStopped is dispatched when a startup module finishes stopping.
	// EventOnStartupModuleStopped 在启动模块完成停止时分发。
	EventOnStartupModuleStopped = "Startup-Module:OnStartupModuleStopped"

	// Manager Events ---------- ---------- ---------- ---------- ----------

	// EventOnStartupManagerSaved is dispatched when all startup modules have been saved.
	// EventOnStartupManagerSaved 在全部启动模块保存完成时分发。
	EventOnStartupManagerSaved = "Startup-Manager:OnStartupManagerSaved"
	// EventOnStartupManagerStarted is dispatched when all startup modules have started.
	// EventOnStartupManagerStarted 在全部启动模块启动完成时分发。
	EventOnStartupManagerStarted = "Startup-Manager:OnStartupManagerStarted"
	// EventOnStartupManagerStopped is dispatched when all startup modules have stopped.
	// EventOnStartupManagerStopped 在全部启动模块停止完成时分发。
	EventOnStartupManagerStopped = "Startup-Manager:OnStartupManagerStopped"
	// EventOnStartupManagerRebooted is dispatched when a manager reboot completes.
	// EventOnStartupManagerRebooted 在管理器重启完成时分发。
	EventOnStartupManagerRebooted = "Startup-Manager:OnStartupManagerRebooted"
)
