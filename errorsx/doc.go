// Package errorsx provides domain-specific error types for infra-go components.
// Package errorsx 为 infra-go 各组件提供领域相关的错误类型。
//
// Created by xuzhuoxi on 2018-12-26.
// @author xuzhuoxi
//
// Capabilities:
//   - String-backed error types that implement error and net.Error (Timeout/Temporary always false)
//   - NoCaseCatchError: a switch or branch reached default without handling the case
//   - FuncUnavailableError, FuncRepeatedCallError, FuncNotPreparedError: function lifecycle/state errors, typically keyed by a function or site name via string conversion
//
// 主要职责：
//   - 基于 string 的错误类型，实现 error 与 net.Error（Timeout/Temporary 恒为 false）
//   - NoCaseCatchError：switch 或分支落入 default 且未处理该情况
//   - FuncUnavailableError、FuncRepeatedCallError、FuncNotPreparedError：函数生命周期/状态类错误，通常通过对函数名或位置名的 string 转换构造
package errorsx
