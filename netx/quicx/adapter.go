package quicx

import "io"

// QUICStreamAdapter adapts a QUIC stream to netx IConnReadWriterAdapter.
// QUICStreamAdapter 将 QUIC 流适配为 netx IConnReadWriterAdapter。
type QUICStreamAdapter struct {
	Reader        io.Reader
	Writer        io.Writer
	remoteAddress string
}

// RemoteAddress returns the remote address string.
// RemoteAddress 返回远程地址字符串。
func (rw *QUICStreamAdapter) RemoteAddress() string {
	return rw.remoteAddress
}

// ReadBytes reads from the QUIC stream.
// ReadBytes 从 QUIC 流读取数据。
func (rw *QUICStreamAdapter) ReadBytes(bytes []byte) (n int, remoteAddress string, err error) {
	n, err = rw.Reader.Read(bytes)
	remoteAddress = rw.remoteAddress
	return
}

// WriteBytes writes to the QUIC stream.
// WriteBytes 向 QUIC 流写入数据。
func (rw *QUICStreamAdapter) WriteBytes(bytes []byte, connId ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
