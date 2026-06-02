package mgo_v2x

// BulkInsert performs bulk insert operations on the given session.
// BulkInsert 在给定会话上执行批量插入操作。
func BulkInsert(ses IMongoSessionV, docs ...[]interface{}) {
}

// BulkUpdate performs bulk update operations on the given session.
// BulkUpdate 在给定会话上执行批量更新操作。
func BulkUpdate(ses IMongoSessionV, all bool, pairs ...[]interface{}) {
}

// BulkUpsert performs bulk upsert operations on the given session.
// BulkUpsert 在给定会话上执行批量 upsert 操作。
func BulkUpsert(ses IMongoSessionV, pairs ...[]interface{}) {
}

// BulkRemove performs bulk remove operations on the given session.
// BulkRemove 在给定会话上执行批量删除操作。
func BulkRemove(ses IMongoSessionV, all bool, selectors ...interface{}) {
}
