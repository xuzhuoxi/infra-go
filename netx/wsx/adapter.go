package wsx

import "io"

// WSConnAdapter adapts a WebSocket connection to netx IConnReadWriterAdapter.
// WSConnAdapter 将 WebSocket 连接适配为 netx IConnReadWriterAdapter。
type WSConnAdapter struct {
	Reader        io.Reader
	Writer        io.Writer
	remoteAddress string
}

// RemoteAddress returns the remote address string.
// RemoteAddress 返回远程地址字符串。
func (rw *WSConnAdapter) RemoteAddress() string {
	return rw.remoteAddress
}

// ReadBytes reads from the WebSocket connection.
// ReadBytes 从 WebSocket 连接读取数据。
func (rw *WSConnAdapter) ReadBytes(bytes []byte) (n int, remoteAddress string, err error) {
	n, err = rw.Reader.Read(bytes)
	remoteAddress = rw.remoteAddress
	return
}

// WriteBytes writes to the WebSocket connection.
// WriteBytes 向 WebSocket 连接写入数据。
func (rw *WSConnAdapter) WriteBytes(bytes []byte, connId ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
