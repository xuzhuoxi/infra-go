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
	ReadBytes(bytes []byte) (int, interface{}, error)
}

type iConnWriterAdapter interface {
	WriteBytes(bytes []byte, rAddress ...string) (int, error)
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

func (rw *ReadWriterAdapter) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, err := rw.Reader.Read(bytes)
	return n, rw.RemoteAddress(), err
}

func (rw *ReadWriterAdapter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	return rw.Writer.Write(bytes)
}

//-------------------------------------------------

type UDPConnAdapter struct {
	ReadWriter *net.UDPConn
}

func (rw *UDPConnAdapter) RemoteAddress() string {
	return ""
}

func (rw *UDPConnAdapter) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, addr, err := rw.ReadWriter.ReadFromUDP(bytes)
	return n, addr.String(), err
}

func (rw *UDPConnAdapter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	if len(rAddress) == 0 {
		return 0, NoAddrError("UDPConnAdapter.ReadBytes")
	}
	n := 0
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

func (rw *QUICStreamAdapter) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, err := rw.Reader.Read(bytes)
	return n, rw.RemoteAddress(), err
}

func (rw *QUICStreamAdapter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	n, err := rw.Writer.Write(bytes)
	return n, err
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

func (rw *WSConnAdapter) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, err := rw.Reader.Read(bytes)
	return n, rw.remoteAddrString, err
}

func (rw *WSConnAdapter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	return rw.Writer.Write(bytes)
}
