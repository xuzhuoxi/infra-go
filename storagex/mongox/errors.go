package mongox

// SessionNilError indicates a missing or invalid MongoDB session.
// SessionNilError 表示 MongoDB 会话缺失或无效。
type SessionNilError string

// Error returns the error message with the location context.
// Error 返回包含位置上下文的错误信息。
func (e SessionNilError) Error() string { return "Session Nil At: " + string(e) }

// Timeout reports whether the error is a timeout.
// Timeout 报告该错误是否为超时错误。
func (e SessionNilError) Timeout() bool { return false }

// Temporary reports whether the error is temporary.
// Temporary 报告该错误是否为临时错误。
func (e SessionNilError) Temporary() bool { return false }

// ClientNilError indicates a missing or invalid MongoDB client.
// ClientNilError 表示 MongoDB 客户端缺失或无效。
type ClientNilError string

// Error returns the error message with the location context.
// Error 返回包含位置上下文的错误信息。
func (e ClientNilError) Error() string { return "Client Nil At: " + string(e) }

// Timeout reports whether the error is a timeout.
// Timeout 报告该错误是否为超时错误。
func (e ClientNilError) Timeout() bool { return false }

// Temporary reports whether the error is temporary.
// Temporary 报告该错误是否为临时错误。
func (e ClientNilError) Temporary() bool { return false }
