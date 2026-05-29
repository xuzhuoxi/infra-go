// Package netx provides unified socket abstractions for TCP, UDP, QUIC, and WebSocket with pack framing and connection management.
// Package netx 提供 TCP、UDP、QUIC、WebSocket 的统一套接字抽象，包含消息封包与连接管理。
//
// Created by xuzhuoxi on 2019-02-14.
// @author xuzhuoxi
//
// Capabilities:
//   - SockNetwork registry and factory for protocol-specific client/server implementations (tcpx, udpx, quicx, wsx)
//   - ISockClient / ISockServer with pack send/receive, handler chains, and lifecycle management
//   - IConnInfo for connection identity; IUserConnMapper / IUserAddressMapper for bidirectional user mappings
//   - Event constants for server lifecycle and mapper changes
//
// 主要职责：
//   - SockNetwork 注册表与工厂，按协议创建客户端/服务端实现（tcpx、udpx、quicx、wsx）
//   - ISockClient / ISockServer：封包收发、处理器链与生命周期管理
//   - IConnInfo 连接标识；IUserConnMapper / IUserAddressMapper 用户双向映射
//   - 服务端生命周期与映射变更相关事件常量
package netx
