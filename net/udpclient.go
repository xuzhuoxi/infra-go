package net

import (
	"log"
	"net"
)

func NewUDPClient(connected bool) IUDPClient {
	if connected {
		return &UDPDialClient{Network: "udp"}
	} else {
		return &UDPListenClient{Network: "udp"}
	}
}

type IUDPClient interface {
	Connected() bool
	Setup(address string) bool
	Close() bool
	SendData(data []byte, rAddress string) bool
}

//UDPDialClient
type UDPDialClient struct {
	Network string
	conn    *net.UDPConn
}

func (c *UDPDialClient) Connected() bool {
	return true
}

func (c *UDPDialClient) Setup(address string) bool {
	if nil != c.conn {
		log.Fatalln("Repeated Setup!")
		return false
	}
	uAddr, err := net.ResolveUDPAddr(c.Network, address)
	if nil != err {
		log.Fatalln("\tnet.ResolveUDPAddr:", c.Network, address, ": %v", err)
		return false
	}
	conn, cErr := net.DialUDP(c.Network, nil, uAddr)
	if nil != cErr {
		log.Fatalln("\net.DialUDP:", c.Network, address, ": %v", cErr)
		return false
	}
	c.conn = conn
	return true
}

func (c *UDPDialClient) Close() bool {
	return closeConn(c.conn)
}

func (c *UDPDialClient) SendData(data []byte, rAddress string) bool {
	_, err := c.conn.Write(data)
	if nil != err {
		log.Fatalln(err)
		return false
	}
	return true
}

//UDPListenClient
type UDPListenClient struct {
	Network       string
	conn          *net.UDPConn
	mapRemoteAddr map[string]*net.UDPAddr
}

func (c *UDPListenClient) Connected() bool {
	return false
}

func (c *UDPListenClient) Setup(address string) bool {
	uAddr, err := net.ResolveUDPAddr(c.Network, address)
	if nil != err {
		log.Fatalln("\tnet.ResolveUDPAddr:", c.Network, address, ": %v", err)
		return false
	}
	conn, cErr := net.ListenUDP(c.Network, uAddr)
	if nil != cErr {
		log.Fatalln("\tSetup:", c.Network, address, ": %v", cErr)
		return false
	}
	c.conn = conn
	c.mapRemoteAddr = make(map[string]*net.UDPAddr)
	return true
}

func (c *UDPListenClient) Close() bool {
	return closeConn(c.conn)
}

func (c *UDPListenClient) SendData(data []byte, rAddress string) bool {
	rAddr, ok := c.mapRemoteAddr[rAddress]
	if !ok {
		var err error
		rAddr, err = net.ResolveUDPAddr(c.Network, rAddress)
		if nil != err {
			log.Fatalln(err)
			return false
		}
		c.mapRemoteAddr[rAddress] = rAddr
	}
	c.conn.WriteToUDP(data, rAddr)
	return true
}

//private ---------------

func closeConn(conn *net.UDPConn) bool {
	if nil != conn {
		conn.Close()
		return true
	}
	return false
}
