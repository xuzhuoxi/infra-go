package netx

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/util-go/errorsx"
	"log"
	"net"
)

type ReadWriterType int

const (
	TcpRW ReadWriterType = iota
	UdpDialRW
	UdpListenRW
	QuicRW
)

func NewReaderProxy(reader interface{}, readerType ReadWriterType, Network string) IReaderProxy {
	return &readWriterProxy{RWType: readerType, Network: Network, Reader: reader}
}

func NewWriterProxy(writer interface{}, writerType ReadWriterType, Network string) IWriterProxy {
	return &readWriterProxy{RWType: writerType, Network: Network, Writer: writer}
}

func NewReadWriterProxy(reader interface{}, writer interface{}, rwType ReadWriterType, Network string) IReadWriterProxy {
	return &readWriterProxy{RWType: rwType, Network: Network, Reader: reader, Writer: writer}
}

type readWriterProxy struct {
	RWType  ReadWriterType
	Network string
	Reader  interface{}
	Writer  interface{}

	quicStream quic.Stream
}

func (p *readWriterProxy) ReadBytes(bytes []byte) (int, interface{}, error) {
	funcName := "readWriterProxy.ReadBytes"
	if nil == p.Reader {
		return 0, nil, ConnNilError(funcName)
	}
	switch r := p.Reader.(type) {
	case *net.TCPConn:
		n, err := r.Read(bytes)
		return n, r.RemoteAddr().String(), err
	case *net.UDPConn:
		switch p.RWType {
		case UdpDialRW:
			n, err := r.Read(bytes)
			return n, r.RemoteAddr().String(), err
		case UdpListenRW:
			n, addr, err := r.ReadFromUDP(bytes)
			return n, addr.String(), err
		}
	case quic.Session:
		rAddr := r.RemoteAddr().String()
		stream, err := r.AcceptStream()
		if nil != err {
			return 0, rAddr, err
		}
		n, err := stream.Read(bytes)
		return n, rAddr, err
		//err := p.initQuicStream(r)
		//if nil != err {
		//	return 0, rAddr, err
		//}
		//n, err2 := p.quicStream.Read(bytes)
		//return n, rAddr, err2
	}
	return 0, nil, errorsx.NoCaseCatchError(funcName)
}

func (p *readWriterProxy) WriteBytes(bytes []byte, rAddress ...string) (int, error) {
	funcName := "readWriterProxy.WriteBytes"
	if nil == p.Writer {
		return 0, ConnNilError(funcName)
	}
	switch w := p.Writer.(type) {
	case *net.TCPConn:
		return w.Write(bytes)
	case *net.UDPConn:
		switch p.RWType {
		case UdpDialRW:
			return w.Write(bytes)
		case UdpListenRW:
			if len(rAddress) == 0 {
				return 0, NoAddrError(funcName)
			}
			n := 0
			for _, address := range rAddress {
				uAddr, err := GetUDPAddr(p.Network, address)
				if nil != err {
					continue
				}
				w.WriteToUDP(bytes, uAddr)
				n++
			}
			return n, nil
		}
	case quic.Session:
		stream, err := w.OpenStreamSync()
		if nil != err {
			return 0, err
		}
		n, err := stream.Write(bytes)
		return n, err
		//err := p.initQuicStream(w)
		//if nil != err {
		//	return 0, err
		//}
		//n, err2 := p.quicStream.Write(bytes)
		//return n, err2
	}

	return 0, errorsx.NoCaseCatchError(funcName)
}

func (p *readWriterProxy) initQuicStream(session quic.Session) error {
	if nil == p.quicStream {
		stream, err := session.OpenStreamSync()
		if nil != err {
			log.Println("readWriterProxy.initQuicStream:", err)
			return err
		}
		p.quicStream = stream
	}
	return nil
}
