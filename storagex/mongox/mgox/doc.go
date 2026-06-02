// Package mgox wraps the official MongoDB Go driver with multi-client lifecycle management.
// Package mgox 封装官方 MongoDB Go 驱动，提供多客户端生命周期管理。
//
// Created by xuzhuoxi on 2019-06-13.
// @author xuzhuoxi
//
// Capabilities:
//   - IMongoDriver: open/close named clients, master client selection, optional concurrency limit
//   - IMongoClient: connect, reconnect, list databases, locate database with context
//   - IMongoDatabase: list/locate collections on a connected database
//   - Default context variables for client, database, and collection read/write operations
//   - NewOptions: build ClientOptions from a URI
//
// 主要职责：
//   - IMongoDriver：按名称打开/关闭客户端、主客户端选择、可选并发限制
//   - IMongoClient：连接、重连、列出数据库、带 context 定位数据库
//   - IMongoDatabase：在已连接数据库上列出/定位集合
//   - 客户端、数据库、集合读写操作的默认 context 变量
//   - NewOptions：从 URI 构建 ClientOptions
package mgox
