package netx

import (
	"github.com/lucas-clemente/quic-go"
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
	remoteAddr net.Addr
}

func (rw *ReadWriterAdapter) RemoteAddress() string {
	return rw.remoteAddr.String()
}

func (rw *ReadWriterAdapter) ReadBytes(bytes []byte) (n int, address string, err error) {
	n, err = rw.Reader.Read(bytes)
	address = rw.RemoteAddress()
	return
}

func (rw *ReadWriterAdapter) WriteBytes(bytes []byte, rAddress ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}

//-------------------------------------------------

type UDPConnAdapter struct {
	ReadWriter *net.UDPConn
}

func (rw *UDPConnAdapter) RemoteAddress() string {
	return ""
}

func (rw *UDPConnAdapter) ReadBytes(bytes []byte) (n int, address string, err error) {
	n, addr, err := rw.ReadWriter.ReadFromUDP(bytes)
	return n, addr.String(), err
}

func (rw *UDPConnAdapter) WriteBytes(bytes []byte, rAddress ...string) (n int, err error) {
	if len(rAddress) == 0 {
		return 0, NoAddrError("UDPConnAdapter.ReadBytes")
	}
	n = 0
	network := rw.ReadWriter.LocalAddr().Network()
	for _, address := range rAddress {
		uAddr, err := GetUDPAddr(network, address)
		if nil != err {
			continue
		}
		rw.ReadWriter.WriteToUDP(bytes, uAddr)
		n++
	}
	return n, nil
}

//-------------------------------------------------

type QUICStreamAdapter struct {
	remoteAddr net.Addr
	Reader     quic.ReceiveStream
	Writer     quic.SendStream
}

func (rw *QUICStreamAdapter) RemoteAddress() string {
	return rw.remoteAddr.String()
}

func (rw *QUICStreamAdapter) ReadBytes(bytes []byte) (n int, address string, err error) {
	n, err = rw.Reader.Read(bytes)
	address = rw.RemoteAddress()
	return
}

func (rw *QUICStreamAdapter) WriteBytes(bytes []byte, rAddress ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}

//-------------------------------------------------

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
