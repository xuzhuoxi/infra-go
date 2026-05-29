// Package httpx provides HTTP client helpers and a mux-based HTTP server wrapper.
// Package httpx 提供 HTTP 客户端辅助工具与基于 ServeMux 的 HTTP 服务端封装。
//
// Capabilities:
//   - HttpClient with configurable timeout and GET/POST helpers using callback-based response handling
//   - Package-level DefaultClient and convenience functions for common HTTP calls
//   - HttpServer with route mapping, method-limited handlers, and start/stop lifecycle
//
// 主要职责：
//   - 可配置超时的 HttpClient，以及基于回调响应处理的 GET/POST 辅助方法
//   - 包级 DefaultClient 与常用 HTTP 调用的便捷函数
//   - 支持路由映射、方法限制处理器与启停生命周期的 HttpServer
package httpx
