package net

import (
	"log"
	"net"
)

var (
	mapUDPAddr map[string]*net.UDPAddr
)

func init() {
	mapUDPAddr = make(map[string]*net.UDPAddr)
}

func UDPAddrEqual(addr1 *net.UDPAddr, addr2 *net.UDPAddr) bool {
	if addr1 == addr2 {
		return true
	}
	return addr1.IP.Equal(addr2.IP) && addr1.Port == addr2.Port && addr1.Zone == addr2.Zone
}

func getUDPAddr(network string, address string) (*net.UDPAddr, error) {
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

func sendDataFromListen(listenConn *net.UDPConn, data []byte, rAddress ...string) {
	if len(rAddress) > 0 {
		for _, address := range rAddress {
			addr, err := getUDPAddr(listenConn.LocalAddr().Network(), address)
			if nil != err {
				continue
			}
			listenConn.WriteToUDP(data, addr)
		}
	}
}

func logResolveUDPAddrErr(address string, err error) {
	log.Fatalln("ResolveUDPAddr Error:[addirss="+address+"],err=", err)
}
