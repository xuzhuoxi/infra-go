package quicx

import (
	"github.com/lucas-clemente/quic-go"
)

type QUICStreamAdapter struct {
	Reader        quic.ReceiveStream
	Writer        quic.SendStream
	remoteAddress string
}

func (rw *QUICStreamAdapter) RemoteAddress() string {
	return rw.remoteAddress
}

func (rw *QUICStreamAdapter) ReadBytes(bytes []byte) (n int, remoteAddress string, err error) {
	n, err = rw.Reader.Read(bytes)
	remoteAddress = rw.remoteAddress
	return
}

func (rw *QUICStreamAdapter) WriteBytes(bytes []byte, connId ...string) (n int, err error) {
	return rw.Writer.Write(bytes)
}
