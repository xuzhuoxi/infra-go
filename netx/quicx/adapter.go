package quicx

import (
	"github.com/lucas-clemente/quic-go"
	"net"
)

type QUICStreamAdapter struct {
	RemoteAddr net.Addr
	Reader     quic.ReceiveStream
	Writer     quic.SendStream
}

func (rw *QUICStreamAdapter) RemoteAddress() string {
	return rw.RemoteAddr.String()
}

func (rw *QUICStreamAdapter) ReadBytes(bytes []byte) (n int, address string, err error) {
	n, err = rw.Reader.Read(bytes)
	address = rw.RemoteAddress()
	return
}

func (rw *QUICStreamAdapter) WriteBytes(bytes []byte, rAddress ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
