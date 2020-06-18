package udpx

import (
	"github.com/xuzhuoxi/infra-go/netx"
	"net"
)

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
		return 0, netx.NoAddrError("UDPConnAdapter.ReadBytes")
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
