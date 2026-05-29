package netx

import "errors"

// EmptyAddrError indicates an empty address at the given location.
// EmptyAddrError 表示指定位置地址为空。
type EmptyAddrError string

// Error returns the error message.
// Error 返回错误信息。
func (e EmptyAddrError) Error() string   { return "EmptyAddrError Address At:" + string(e) }
// Timeout reports whether the error is a timeout.
// Timeout 是否为超时错误。
func (e EmptyAddrError) Timeout() bool   { return false }
// Temporary reports whether the error is temporary.
// Temporary 是否为临时错误。
func (e EmptyAddrError) Temporary() bool { return false }

// NoAddrError indicates a missing target address at the given location.
// NoAddrError 表示指定位置缺少目标地址。
type NoAddrError string

// Error returns the error message.
// Error 返回错误信息。
func (e NoAddrError) Error() string   { return "NoAddrError At:" + string(e) }
// Timeout reports whether the error is a timeout.
// Timeout 是否为超时错误。
func (e NoAddrError) Timeout() bool   { return false }
// Temporary reports whether the error is temporary.
// Temporary 是否为临时错误。
func (e NoAddrError) Temporary() bool { return false }

// ConnNilError indicates a nil connection at the given location.
// ConnNilError 表示指定位置连接为空。
type ConnNilError string

// Error returns the error message.
// Error 返回错误信息。
func (e ConnNilError) Error() string   { return "Conn Is Nil At:" + string(e) }
// Timeout reports whether the error is a timeout.
// Timeout 是否为超时错误。
func (e ConnNilError) Timeout() bool   { return false }
// Temporary reports whether the error is temporary.
// Temporary 是否为临时错误。
func (e ConnNilError) Temporary() bool { return false }

// ReaderUnSupportError indicates an unsupported reader at the given location.
// ReaderUnSupportError 表示指定位置读取器不受支持。
type ReaderUnSupportError string

// Error returns the error message.
// Error 返回错误信息。
func (e ReaderUnSupportError) Error() string   { return "Reader Is Not Support At:" + string(e) }
// Timeout reports whether the error is a timeout.
// Timeout 是否为超时错误。
func (e ReaderUnSupportError) Timeout() bool   { return false }
// Temporary reports whether the error is temporary.
// Temporary 是否为临时错误。
func (e ReaderUnSupportError) Temporary() bool { return false }

// ErrHttpServerStarted is returned when an HTTP server is already running.
// ErrHttpServerStarted 在 HTTP 服务已启动时返回。
var ErrHttpServerStarted = errors.New("http: Server started")

// ErrSockServerStarted is returned when a socket server is already running.
// ErrSockServerStarted 在套接字服务已启动时返回。
var ErrSockServerStarted = errors.New("socket: Server started")
