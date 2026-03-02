package quicx

import "io"

type QUICStreamAdapter struct {
	Reader        io.Reader
	Writer        io.Writer
	remoteAddress string
}

func (rw *QUICStreamAdapter) RemoteAddress() string {
	return rw.remoteAddress
}

func (rw *QUICStreamAdapter) ReadBytes(bytes []byte) (n int, remoteAddress string, err error) {
	n, err = rw.Reader.Read(bytes)
	remoteAddress = rw.remoteAddress
	return
}

func (rw *QUICStreamAdapter) WriteBytes(bytes []byte, connId ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
