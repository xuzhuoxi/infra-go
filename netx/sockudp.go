package netx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"log"
	"net"
	"sync"
)

const (
	UDPNetwork  = "udp"
	UDPNetwork4 = "udp4"
	UDPNetwork6 = "udp6"
)

var (
	mapUDPAddr = make(map[string]*net.UDPAddr)
	mapUDPLock sync.RWMutex
)

var UdpDataBlockHandler = bytex.NewDefaultDataBlockHandler()

type IUDPServer interface {
	ISockServer
}

type IUDPClient interface {
	ISockClient
}

func UDPAddrEqual(addr1 *net.UDPAddr, addr2 *net.UDPAddr) bool {
	if addr1 == addr2 {
		return true
	}
	return addr1.IP.Equal(addr2.IP) && addr1.Port == addr2.Port && addr1.Zone == addr2.Zone
}

func GetUDPAddr(network string, address string) (*net.UDPAddr, error) {
	mapUDPLock.Lock()
	defer mapUDPLock.Unlock()
	if "" == address {
		return nil, EmptyAddrError("netx.GetUDPAddr")
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
