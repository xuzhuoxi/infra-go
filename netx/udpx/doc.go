// Package udpx provides UDP socket client and server implementations for netx.
// Package udpx 提供 netx 的 UDP 套接字客户端与服务端实现。
//
// Created by xuzhuoxi on 2023-08-05.
// @author xuzhuoxi
//
// Capabilities:
//   - UDP/UDP4/UDP6 dial and listen clients, plus connectionless server
//   - UDPConnAdapter for address-aware read/write on *net.UDPConn
//   - One-to-many receive mode using remote address as connection id
//   - Registers UDPNetwork, UDP4Network, and UDP6Network with netx on init
//
// 主要职责：
//   - UDP/UDP4/UDP6 拨号与监听客户端，以及无连接服务端
//   - UDPConnAdapter 在 *net.UDPConn 上按地址读写
//   - 一对多接收模式，以远程地址作为连接 id
//   - init 时向 netx 注册 UDPNetwork、UDP4Network、UDP6Network
package udpx
