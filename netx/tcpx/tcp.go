package tcpx

import (
	"log"
	"net"
	"sync"

	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/netx"
)

var (
	mapTCPAddr = make(map[string]*net.TCPAddr)
	mapTCPLock sync.RWMutex
)

// TcpDataBlockHandler is the default data block handler for TCP pack framing.
// TcpDataBlockHandler 是 TCP 封包默认的数据块处理器。
var TcpDataBlockHandler = bytex.NewDefaultDataBlockHandler()

// TCPAddrEqual reports whether two TCP addresses are equal.
// TCPAddrEqual 判断两个 TCP 地址是否相等。
func TCPAddrEqual(addr1 *net.TCPAddr, addr2 *net.TCPAddr) bool {
	if addr1 == addr2 {
		return true
	}
	return addr1.IP.Equal(addr2.IP) && addr1.Port == addr2.Port && addr1.Zone == addr2.Zone
}

// GetTCPAddr resolves and caches a TCP address for the given network and address string.
// GetTCPAddr 解析并缓存指定网络与地址字符串对应的 TCP 地址。
func GetTCPAddr(network string, address string) (*net.TCPAddr, error) {
	mapTCPLock.Lock()
	defer mapTCPLock.Unlock()
	if "" == address {
		return nil, netx.EmptyAddrError("netx.GetTCPAddr")
	}
	addr, ok := mapTCPAddr[address]
	if ok {
		return addr, nil
	}
	newAddr, err := net.ResolveTCPAddr(network, address)
	if nil != err {
		logResolveTCPAddrErr(address, err)
		return nil, err
	}
	mapTCPAddr[address] = newAddr
	return newAddr, nil
}

func logResolveTCPAddrErr(address string, err error) {
	log.Fatalln("ResolveTCPAddr Error:[addirss="+address+"],error=", err)
}
