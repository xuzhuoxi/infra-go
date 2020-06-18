package wsx

import "io"

type WSConnAdapter struct {
	Reader           io.Reader
	Writer           io.Writer
	remoteAddrString string
}

func (rw *WSConnAdapter) RemoteAddress() string {
	return rw.remoteAddrString
}

func (rw *WSConnAdapter) ReadBytes(bytes []byte) (n int, address string, err error) {
	n, err = rw.Reader.Read(bytes)
	address = rw.remoteAddrString
	return
}

func (rw *WSConnAdapter) WriteBytes(bytes []byte, rAddress ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
