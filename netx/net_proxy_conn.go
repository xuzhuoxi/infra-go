package netx

import (
	"github.com/lucas-clemente/quic-go"
	"io"
	"net"
)

type ReadWriterProxy struct {
	Reader     io.Reader
	Writer     io.Writer
	RemoteAddr net.Addr
}

func (rw *ReadWriterProxy) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, err := rw.Reader.Read(bytes)
	return n, rw.RemoteAddr.String(), err
}

func (rw *ReadWriterProxy) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	return rw.Writer.Write(bytes)
}

type UDPListenReadWriterProxy struct {
	ReadWriter *net.UDPConn
}

func (rw *UDPListenReadWriterProxy) ReadBytes(bytes []byte) (int, interface{}, error) {
	n, addr, err := rw.ReadWriter.ReadFromUDP(bytes)
	return n, addr.String(), err
}

func (rw *UDPListenReadWriterProxy) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	if len(rAddress) == 0 {
		return 0, NoAddrError("UDPListenReadWriterProxy.ReadBytes")
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

type QUICSessionReadWriter struct {
	RemoteAddr net.Addr
	Reader     quic.ReceiveStream
	Writer     quic.SendStream
}

func (rw *QUICSessionReadWriter) ReadBytes(bytes []byte) (int, interface{}, error) {
	rAddr := rw.RemoteAddr.String()
	n, err := rw.Reader.Read(bytes)
	return n, rAddr, err
}

func (rw *QUICSessionReadWriter) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	n, err := rw.Writer.Write(bytes)
	return n, err
}
