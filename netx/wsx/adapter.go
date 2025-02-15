package wsx

import "io"

type WSConnAdapter struct {
	Reader        io.Reader
	Writer        io.Writer
	remoteAddress string
}

func (rw *WSConnAdapter) RemoteAddress() string {
	return rw.remoteAddress
}

func (rw *WSConnAdapter) ReadBytes(bytes []byte) (n int, remoteAddress string, err error) {
	n, err = rw.Reader.Read(bytes)
	remoteAddress = rw.remoteAddress
	return
}

func (rw *WSConnAdapter) WriteBytes(bytes []byte, connId ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
