// Package mgo_v2x wraps gopkg.in/mgo.v2 with multi-session lifecycle management.
// Package mgo_v2x 封装 gopkg.in/mgo.v2，提供多会话生命周期管理。
//
// Capabilities:
//   - IMongoDriverV: open/close/clone named sessions, master session selection
//   - IMongoSessionV: locate database and collection, expose underlying mgo.Session
//   - BulkInsert, BulkUpdate, BulkUpsert, BulkRemove: bulk operation helpers (stubs)
//
// 主要职责：
//   - IMongoDriverV：按名称打开/关闭/克隆会话、主会话选择
//   - IMongoSessionV：定位数据库与集合，暴露底层 mgo.Session
//   - BulkInsert、BulkUpdate、BulkUpsert、BulkRemove：批量操作辅助函数（当前为占位实现）
package mgo_v2x
