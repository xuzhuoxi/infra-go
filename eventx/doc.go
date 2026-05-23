// Package eventx provides a lightweight publish-subscribe style event dispatcher.
// Package eventx 提供轻量级的发布-订阅式事件分发器。
//
// Created by xuzhuoxi on 2019-01-08.
// @author xuzhuoxi
//
// Capabilities:
//   - EventCall callbacks and EventData payload with propagation control (StopImmediatePropagation)
//   - IEventDispatcher / EventDispatcher: add, once, remove listeners by callback or type, and dispatch events
//   - Duplicate listeners for the same type are allowed and invoked in registration order
//   - Embed EventDispatcher in custom types to reuse listener management
//
// 主要职责：
//   - EventCall 回调与 EventData 载荷，支持传播控制（StopImmediatePropagation）
//   - IEventDispatcher / EventDispatcher：按回调或类型增删监听器、一次性监听与事件分发
//   - 同一事件类型允许重复注册监听器，按注册顺序依次调用
//   - 可在自定义类型中内嵌 EventDispatcher 复用监听管理
package eventx
