package netx

import (
	"io"
	"net"
)

// IRemoteAddress returns the remote address string for the current connection.
// IRemoteAddress 返回当前连接的远程地址字符串。
type IRemoteAddress interface {
	// RemoteAddress returns the remote address, or empty string on error.
	// RemoteAddress 返回当前连接的远程地址；成功时返回远程地址，错误时返回空字符串。
	RemoteAddress() string
}

type iConnReaderAdapter interface {
	// ReadBytes reads one message.
	// ReadBytes 读取一条消息。
	// On success: returns bytes read and remote address.
	// 成功：返回读取到的字节数与远程地址。
	// On error: returns the error.
	// 错误：返回错误。
	ReadBytes(bytes []byte) (n int, remoteAddress string, err error)
}

type iConnWriterAdapter interface {
	// WriteBytes writes one message to the given remote address; rAddress is ignored for one-to-one connections.
	// WriteBytes 向指定远程地址写入一条消息；若当前是一对一连接，则 rAddress 忽略。
	// On success: returns bytes written.
	// 成功：返回写入的字节数。
	// On error: returns the error.
	// 错误：返回错误。
	// For UDP, RemoteAddress is used as connId because connections are not maintained.
	// UDP 协议中，因为实际上不维护连接，所以使用 RemoteAddress 填充 connId。
	WriteBytes(bytes []byte, connId ...string) (n int, err error)
}

// IConnReaderAdapter adapts a connection reader with remote address reporting.
// IConnReaderAdapter 适配连接读取器并提供远程地址。
type IConnReaderAdapter interface {
	iConnReaderAdapter
	IRemoteAddress
}

// IConnWriterAdapter adapts a connection writer with remote address reporting.
// IConnWriterAdapter 适配连接写入器并提供远程地址。
type IConnWriterAdapter interface {
	iConnWriterAdapter
	IRemoteAddress
}

// IConnReadWriterAdapter combines read and write adaptation.
// IConnReadWriterAdapter 组合读写适配能力。
type IConnReadWriterAdapter interface {
	iConnReaderAdapter
	iConnWriterAdapter
	IRemoteAddress
}

//-------------------------------------------------

// ConnReadWriterAdapter wraps io.Reader and io.Writer with a fixed remote address.
// ConnReadWriterAdapter 使用固定远程地址包装 io.Reader 与 io.Writer。
type ConnReadWriterAdapter struct {
	Reader     io.Reader
	Writer     io.Writer
	RemoteAddr net.Addr
}

// RemoteAddress returns the remote address string.
// RemoteAddress 返回远程地址字符串。
func (rw *ConnReadWriterAdapter) RemoteAddress() string {
	return rw.RemoteAddr.String()
}

// ReadBytes reads bytes and returns the fixed remote address.
// ReadBytes 读取数据并返回固定远程地址。
func (rw *ConnReadWriterAdapter) ReadBytes(bytes []byte) (n int, remoteAddress string, err error) {
	n, err = rw.Reader.Read(bytes)
	remoteAddress = rw.RemoteAddress()
	return
}

// WriteBytes writes bytes to the underlying writer.
// WriteBytes 向底层写入器写入数据。
func (rw *ConnReadWriterAdapter) WriteBytes(bytes []byte, connId ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
