// Package quicx provides QUIC socket client and server implementations for netx.
// Package quicx 提供 netx 的 QUIC 套接字客户端与服务端实现。
//
// Created by xuzhuoxi on 2023-08-05.
// @author xuzhuoxi
//
// Capabilities:
//   - QUIC client and server built on quic-go with stream-level pack I/O
//   - Auto-generated TLS config for local development and testing
//   - QUICStreamAdapter bridging quic streams to netx IConnReadWriterAdapter
//   - Registers QuicNetwork with netx on init
//
// 主要职责：
//   - 基于 quic-go 的 QUIC 客户端与服务端，在流级别进行封包读写
//   - 自动生成 TLS 配置，便于本地开发与测试
//   - QUICStreamAdapter 将 quic 流桥接到 netx IConnReadWriterAdapter
//   - init 时向 netx 注册 QuicNetwork
package quicx
