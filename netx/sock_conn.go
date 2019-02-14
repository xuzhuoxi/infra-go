package netx

import (
	"github.com/lucas-clemente/quic-go"
	"io"
	"net"
)

type IConnReaderAdapter interface {
	ReadBytes(bytes []byte) (int, interface{}, error)
}

type IConnWriterAdapter interface {
	WriteBytes(bytes []byte, rAddress ...string) (int, error)
}

type IConnReadWriterAdapter interface {
	IConnReaderAdapter
	IConnWriterAdapter
}

//-------------------------------------------------

type ReadWriterAdapter struct {
	Reader     io.Reader
	Writer     io.Writer
	RemoteAddr net.Addr
}

func (rw *ReadWriterAdapter) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, err := rw.Reader.Read(bytes)
	return n, rw.RemoteAddr.String(), err
}

func (rw *ReadWriterAdapter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	return rw.Writer.Write(bytes)
}

//-------------------------------------------------

type UDPConnAdapter struct {
	ReadWriter *net.UDPConn
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
	RemoteAddr net.Addr
	Reader     quic.ReceiveStream
	Writer     quic.SendStream
}

func (rw *QUICStreamAdapter) ReadBytes(bytes []byte) (int, interface{}, error) {
	rAddr := rw.RemoteAddr.String()
	n, err := rw.Reader.Read(bytes)
	return n, rAddr, err
}

func (rw *QUICStreamAdapter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	n, err := rw.Writer.Write(bytes)
	return n, err
}

//-------------------------------------------------

type WSConnAdapter struct {
	Reader           io.Reader
	Writer           io.Writer
	RemoteAddrString string
}

func (rw *WSConnAdapter) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, err := rw.Reader.Read(bytes)
	return n, rw.RemoteAddrString, err
}

func (rw *WSConnAdapter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	return rw.Writer.Write(bytes)
}
