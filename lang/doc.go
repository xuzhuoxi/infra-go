// Package lang provides common language-level utilities: value equality, object pools, channel concurrency limits, and deferred callbacks.
// Package lang 提供通用语言级工具：值相等比较、对象池、channel 并发限制与延迟回调。
//
// Created by xuzhuoxi on 2019-03-21.
// @author xuzhuoxi
//
// Capabilities:
//   - Equal and TypeEqual: compare values by reflect kind, including slices and []byte
//   - IObjectPool, ISizeObjectPool, and constructors with optional mutex and max pool size
//   - IChannelLimit, ChannelLimit, and ChannelLimitNone for semaphore-style concurrency control
//   - Callback and FuncCallback: store a function and arguments for Apply or Invoke
//
// 主要职责：
//   - Equal、TypeEqual：按 reflect 类型比较值，支持切片与 []byte
//   - IObjectPool、ISizeObjectPool 及可选同步、容量上限的构造方法
//   - IChannelLimit、ChannelLimit、ChannelLimitNone：基于信号量的并发限制
//   - Callback、FuncCallback：封装函数与参数，支持 Apply 与 Invoke 调用
package lang
