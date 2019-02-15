package netx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"log"
	"net"
	"sync"
)

const (
	TcpNetwork  = "tcp"
	TcpNetwork4 = "tcp4"
	TcpNetwork6 = "tcp6"
)

var (
	mapTCPAddr = make(map[string]*net.TCPAddr)
	mapTCPLock sync.RWMutex
)

var TcpDataBlockHandler = bytex.NewDefaultDataBlockHandler()

type ITCPServer interface {
	ISockServer
}

type ITCPClient interface {
	ISockClient
}

func TCPAddrEqual(addr1 *net.TCPAddr, addr2 *net.TCPAddr) bool {
	if addr1 == addr2 {
		return true
	}
	return addr1.IP.Equal(addr2.IP) && addr1.Port == addr2.Port && addr1.Zone == addr2.Zone
}

func GetTCPAddr(network string, address string) (*net.TCPAddr, error) {
	mapTCPLock.Lock()
	defer mapTCPLock.Unlock()
	if "" == address {
		return nil, EmptyAddrError("netx.GetTCPAddr")
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
