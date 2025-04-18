// Package serialx
// Create on 2023/7/2
// @author xuzhuoxi
package serialx

// ISerialManager Events
const (
	// Module Events ---------- ---------- ---------- ---------- ----------

	// EventOnSerialModuleStarted
	// Serial module start finish event
	EventOnSerialModuleStarted = "Serial-Module:OnObserverStarted"
	// EventOnSerialModuleStopped
	// Serial module stop finish event
	EventOnSerialModuleStopped = "Serial-Module:OnObserverStopped"

	// Manager Events ---------- ---------- ---------- ---------- ----------

	// EventOnSerialManagerStarted
	// Serial manger start finish event
	EventOnSerialManagerStarted = "Serial-Manager:OnManagerStarted"
	// EventOnSerialManagerStopped
	// Serial manger stop finish event
	EventOnSerialManagerStopped = "Serial-Manager:OnManagerStopped"
)

// IStartupManager Events
const (
	// Module Events ---------- ---------- ---------- ---------- ----------

	// EventOnStartupModuleSaved
	// Startup module save finish event
	EventOnStartupModuleSaved = "Startup-Module:OnStartupModuleSaved"
	// EventOnStartupModuleStarted
	// Startup module start finish event
	EventOnStartupModuleStarted = "Startup-Module:OnStartupModuleStarted"
	// EventOnStartupModuleStopped
	// Startup module stop finish event
	EventOnStartupModuleStopped = "Startup-Module:OnStartupModuleStopped"

	// Manager Events ---------- ---------- ---------- ---------- ----------

	// EventOnStartupManagerSaved
	// Startup manger save finish event
	EventOnStartupManagerSaved = "Startup-Manager:OnStartupManagerSaved"
	// EventOnStartupManagerStarted
	// Startup manger start finish event
	EventOnStartupManagerStarted = "Startup-Manager:OnStartupManagerStarted"
	// EventOnStartupManagerStopped
	// Startup manger stop finish event
	EventOnStartupManagerStopped = "Startup-Manager:OnStartupManagerStopped"
	// EventOnStartupManagerRebooted
	// Startup manger reboot finish event
	EventOnStartupManagerRebooted = "Startup-Manager:OnStartupManagerRebooted"
)
