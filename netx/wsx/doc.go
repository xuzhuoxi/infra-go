// Package wsx provides WebSocket client and server implementations for netx.
// Package wsx 提供 netx 的 WebSocket 客户端与服务端实现。
//
// Created by xuzhuoxi on 2023-08-05.
// @author xuzhuoxi
//
// Capabilities:
//   - WebSocket client and HTTP-upgraded server with configurable pattern and origin
//   - WSConnAdapter bridging websocket.Conn to netx pack I/O
//   - Connection limit and per-client IServerConn management on the server
//   - Registers WSNetwork with netx on init
//
// 主要职责：
//   - WebSocket 客户端与 HTTP 升级服务端，支持可配置路径与 Origin
//   - WSConnAdapter 将 websocket.Conn 桥接到 netx 封包读写
//   - 服务端连接数限制与按客户端管理 IServerConn
//   - init 时向 netx 注册 WSNetwork
package wsx
