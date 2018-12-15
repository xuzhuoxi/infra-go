package net

import (
	"log"
	"net"
)

var (
	mapTCPAddr map[string]*net.TCPAddr
)

func init() {
	mapTCPAddr = make(map[string]*net.TCPAddr)
}

func TCPAddrEqual(addr1 *net.TCPAddr, addr2 *net.TCPAddr) bool {
	if addr1 == addr2 {
		return true
	}
	return addr1.IP.Equal(addr2.IP) && addr1.Port == addr2.Port && addr1.Zone == addr2.Zone
}

func getTCPAddr(network string, address string) (*net.TCPAddr, error) {
	if "" == address {
		return nil, EmptyAddrError("net.getTCPAddr")
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
	log.Fatalln("ResolveTCPAddr Error:[addirss="+address+"],err=", err)
}
