package netx

import (
	"io"
	"net"
)

type IRemoteAddress interface {
	// RemoteAddress
	// 返回当前连接的远程地址，
	// 成功: 返回远程地址，
	// 错误: 返回空字符串，
	RemoteAddress() string
}

type iConnReaderAdapter interface {
	// ReadBytes
	// 读取一条消息，
	// 成功: 返回读取到的字节数，来远程地址，
	// 错误: 返回错误，
	ReadBytes(bytes []byte) (n int, remoteAddress string, err error)
}

type iConnWriterAdapter interface {
	// WriteBytes
	// 向指定远程地址写入一条消息，如果当前是一对一的连接，则 rAddress 忽略，
	// 成功: 返回写入的字节数，
	// 错误: 返回错误，
	// UDP协议中，因为实际上不维护连接，所以使用RemoteAddress填充connId
	WriteBytes(bytes []byte, connId ...string) (n int, err error)
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

type ConnReadWriterAdapter struct {
	Reader     io.Reader
	Writer     io.Writer
	RemoteAddr net.Addr
}

func (rw *ConnReadWriterAdapter) RemoteAddress() string {
	return rw.RemoteAddr.String()
}

func (rw *ConnReadWriterAdapter) ReadBytes(bytes []byte) (n int, remoteAddress string, err error) {
	n, err = rw.Reader.Read(bytes)
	remoteAddress = rw.RemoteAddress()
	return
}

func (rw *ConnReadWriterAdapter) WriteBytes(bytes []byte, connId ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
