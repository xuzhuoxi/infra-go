package udpx

import (
	"net"

	"github.com/xuzhuoxi/infra-go/netx"
)

// UDPConnAdapter adapts *net.UDPConn to netx IConnReadWriterAdapter with per-datagram addresses.
// UDPConnAdapter 将 *net.UDPConn 适配为按数据报地址读写的 netx IConnReadWriterAdapter。
type UDPConnAdapter struct {
	ReadWriter *net.UDPConn
}

// RemoteAddress always returns empty for UDP adapter.
// RemoteAddress 对 UDP 适配器始终返回空字符串。
func (rw *UDPConnAdapter) RemoteAddress() string {
	return ""
}

// ReadBytes reads a UDP datagram and returns the sender address.
// ReadBytes 读取 UDP 数据报并返回发送方地址。
func (rw *UDPConnAdapter) ReadBytes(bytes []byte) (n int, address string, err error) {
	n1, addr, err1 := rw.ReadWriter.ReadFromUDP(bytes)
	return n1, addr.String(), err1
}

// WriteBytes sends a datagram to each address in rAddress.
// WriteBytes 向 rAddress 中的每个地址发送数据报。
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
