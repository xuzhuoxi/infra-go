// Package serialx provides sequential module lifecycle managers with event-driven orchestration.
// Package serialx 提供基于事件驱动的顺序模块生命周期管理器。
//
// Created by xuzhuoxi on 2023-07-02.
// @author xuzhuoxi
//
// Capabilities:
//   - ISerialManager / SerialManager: append modules and start or stop them one by one in order
//   - IStartupManager / StartupManager: ordered start, save, stop, and reboot with persistence hooks
//   - ISerialModule and IStartupModule: module contracts integrated with eventx.IEventDispatcher
//   - Event name constants for manager and module lifecycle notifications
//
// 主要职责：
//   - ISerialManager / SerialManager：按序追加模块并逐个启动或停止
//   - IStartupManager / StartupManager：按序启动、保存、停止与重启，并支持持久化钩子
//   - ISerialModule 与 IStartupModule：与 eventx.IEventDispatcher 集成的模块契约
//   - 管理器与模块生命周期事件名称常量
package serialx
