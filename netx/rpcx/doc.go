// Package rpcx provides thin wrappers around Go's net/rpc client and server.
// Package rpcx 提供对 Go 标准库 net/rpc 客户端与服务端的轻量封装。
//
// Capabilities:
//   - RPCClient with dial, call, and close over TCP/UDP networks
//   - RPCServer with service registration, blocking listen/accept, and shutdown
//   - Network constants aligned with netx.SockNetwork values
//
// 主要职责：
//   - 支持 TCP/UDP 网络上拨号、调用与关闭的 RPCClient
//   - 支持服务注册、阻塞监听/接受与停止的 RPCServer
//   - 与 netx.SockNetwork 对齐的网络协议常量
package rpcx
