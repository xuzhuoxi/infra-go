package udpx

import (
	"log"
	"net"
	"sync"

	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/netx"
)

var (
	mapUDPAddr = make(map[string]*net.UDPAddr)
	mapUDPLock sync.RWMutex
)

// UdpDataBlockHandler is the default data block handler for UDP pack framing.
// UdpDataBlockHandler 是 UDP 封包默认的数据块处理器。
var UdpDataBlockHandler = bytex.NewDefaultDataBlockHandler()

// UDPAddrEqual reports whether two UDP addresses are equal.
// UDPAddrEqual 判断两个 UDP 地址是否相等。
func UDPAddrEqual(addr1 *net.UDPAddr, addr2 *net.UDPAddr) bool {
	if addr1 == addr2 {
		return true
	}
	return addr1.IP.Equal(addr2.IP) && addr1.Port == addr2.Port && addr1.Zone == addr2.Zone
}

// GetUDPAddr resolves and caches a UDP address for the given network and address string.
// GetUDPAddr 解析并缓存指定网络与地址字符串对应的 UDP 地址。
func GetUDPAddr(network string, address string) (*net.UDPAddr, error) {
	mapUDPLock.Lock()
	defer mapUDPLock.Unlock()
	if "" == address {
		return nil, netx.EmptyAddrError("netx.GetUDPAddr")
	}
	addr, ok := mapUDPAddr[address]
	if ok {
		return addr, nil
	}
	newAddr, err := net.ResolveUDPAddr(network, address)
	if nil != err {
		logResolveUDPAddrErr(address, err)
		return nil, err
	}
	mapUDPAddr[address] = newAddr
	return newAddr, nil
}

func logResolveUDPAddrErr(address string, err error) {
	log.Fatalln("\tResolveUDPAddr Error:[addirss="+address+"],error=", err)
}
