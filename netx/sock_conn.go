package netx

import (
	"io"
	"net"
)

type IRemoteAddress interface {
	RemoteAddress() string
}

type iConnReaderAdapter interface {
	ReadBytes(bytes []byte) (n int, address string, err error)
}

type iConnWriterAdapter interface {
	WriteBytes(bytes []byte, rAddress ...string) (n int, err error)
}

type IConnReaderAdapter interface {
	iConnReaderAdapter
	IRemoteAddress
}

type IConnWriterAdapter interface {
	iConnWriterAdapter
	IRemoteAddress
}

type IConnReadWriterAdapter interface {
	iConnReaderAdapter
	iConnWriterAdapter
	IRemoteAddress
}

//-------------------------------------------------

type ReadWriterAdapter struct {
	Reader     io.Reader
	Writer     io.Writer
	RemoteAddr net.Addr
}

func (rw *ReadWriterAdapter) RemoteAddress() string {
	return rw.RemoteAddr.String()
}

func (rw *ReadWriterAdapter) ReadBytes(bytes []byte) (n int, address string, err error) {
	n, err = rw.Reader.Read(bytes)
	address = rw.RemoteAddress()
	return
}

func (rw *ReadWriterAdapter) WriteBytes(bytes []byte, rAddress ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
