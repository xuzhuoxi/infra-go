// Package tcpx provides TCP socket client and server implementations for netx.
// Package tcpx 提供 netx 的 TCP 套接字客户端与服务端实现。
//
// Created by xuzhuoxi on 2023-08-05.
// @author xuzhuoxi
//
// Capabilities:
//   - TCP/TCP4/TCP6 client and server with netx pack framing and handler chains
//   - Connection pool with max connection limit and per-connection IServerConn management
//   - Cached TCP address resolution via GetTCPAddr
//   - Registers TcpNetwork, Tcp4Network, and Tcp6Network with netx on init
//
// 主要职责：
//   - TCP/TCP4/TCP6 客户端与服务端，集成 netx 封包与处理器链
//   - 连接池与最大连接数限制，按连接管理 IServerConn
//   - GetTCPAddr 缓存 TCP 地址解析结果
//   - init 时向 netx 注册 TcpNetwork、Tcp4Network、Tcp6Network
package tcpx
