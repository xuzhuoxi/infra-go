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

func NewUDPClientForMultiRemote() IUDPClient {
	return &UDPMultiRemoteClient{Network: "udp"}
}

type IUDPClient interface {
	Connected() bool
	Setup(lAddress string, rAddress string) bool
	Close() bool
	SendData(data []byte, rAddress ...string) bool
	SetReceivingHandler(handler func(data []byte, rAddr *net.UDPAddr))
}

//UDPDialClient
type UDPDialClient struct {
	Network string
	conn    *net.UDPConn
	handler func(data []byte, rAddr *net.UDPAddr)
}

func (c *UDPDialClient) Connected() bool {
	return true
}

func (c *UDPDialClient) Setup(lAddress string, rAddress string) bool {
	if nil != c.conn {
		log.Fatalln("UDPDialClient-Repeated Setup!")
		return false
	}
	lAddr, _ := net.ResolveUDPAddr(c.Network, lAddress)
	rAddr, err := net.ResolveUDPAddr(c.Network, rAddress)
	if nil != err {
		logCreateAddrErr(rAddress, err)
		return false
	}
	conn, cErr := net.DialUDP(c.Network, lAddr, rAddr)
	if nil != cErr {
		log.Fatalln("\tUDPDialClient-net.DialUDP:", c.Network, lAddress, rAddress, ": %v", cErr)
		return false
	}
	c.conn = conn
	return true
}

func (c *UDPDialClient) Close() bool {
	return closeConn(c.conn)
}

func (c *UDPDialClient) SendData(data []byte, rAddress ...string) bool {
	if len(rAddress) > 0 {
		log.Fatalln("UDPDialClient can net SendData to rAddress!")
		return false
	}
	_, err := c.conn.Write(data)
	if nil != err {
		log.Fatalln(err)
		return false
	}
	return true
}

func (c *UDPDialClient) SetReceivingHandler(handler func(data []byte, rAddr *net.UDPAddr)) {
	c.handler = handler
}

//UDPListenClient
type UDPListenClient struct {
	Network string
	conn    *net.UDPConn
	handler func(data []byte, rAddr *net.UDPAddr)

	remoteAddress string
	remoteAddr    *net.UDPAddr
}

func (c *UDPListenClient) Connected() bool {
	return false
}

func (c *UDPListenClient) Setup(lAddress string, rAddress string) bool {
	lAddr, err := net.ResolveUDPAddr(c.Network, lAddress)
	if nil != err {
		logCreateAddrErr(lAddress, err)
		return false
	}
	conn, cErr := net.ListenUDP(c.Network, lAddr)
	if nil != cErr {
		log.Fatalln("\tUDPListenClient-Setup:", c.Network, lAddress, ": %v", cErr)
		return false
	}
	c.conn = conn
	if "" != rAddress {
		rAddr, err := net.ResolveUDPAddr(c.Network, rAddress)
		if nil != err {
			logCreateAddrErr(rAddress, err)
			c.remoteAddr = nil
		} else {
			c.remoteAddress = rAddress
			c.remoteAddr = rAddr
		}
	}
	return true
}

func (c *UDPListenClient) Close() bool {
	return closeConn(c.conn)
}

func (c *UDPListenClient) SendData(data []byte, rAddress ...string) bool {
	rLen := len(rAddress)
	if rLen > 1 {
		log.Fatalln("UDPListenClient can not SendData to multi address!")
		return false
	}
	if nil == c.remoteAddr {
		if rLen == 0 {
			log.Fatalln("UDPListenClient: No remote address!")
			return false
		}
		rAddr, err := net.ResolveUDPAddr(c.Network, rAddress[0])
		if nil != err {
			logCreateAddrErr(rAddress[0], err)
			return false
		}
		c.remoteAddr = rAddr
		c.remoteAddress = rAddress[0]
	}
	c.conn.WriteToUDP(data, c.remoteAddr)
	return true
}

func (c *UDPListenClient) SetReceivingHandler(handler func(data []byte, rAddr *net.UDPAddr)) {
	c.handler = handler
}

type UDPMultiRemoteClient struct {
	Network string
	conn    *net.UDPConn
	mapAddr map[string]*net.UDPAddr
	handler func(data []byte, rAddr *net.UDPAddr)
}

func (c *UDPMultiRemoteClient) Connected() bool {
	return false
}

func (c *UDPMultiRemoteClient) Setup(lAddress string, rAddress string) bool {
	lAddr, err := net.ResolveUDPAddr(c.Network, lAddress)
	if nil != err {
		logCreateAddrErr(lAddress, err)
		return false
	}
	conn, cErr := net.ListenUDP(c.Network, lAddr)
	if nil != cErr {
		log.Fatalln("\tUDPMultiRemoteClient-Setup:", c.Network, lAddress, ": %v", cErr)
		return false
	}
	c.conn = conn
	c.mapAddr = make(map[string]*net.UDPAddr)
	return true
}

func (c *UDPMultiRemoteClient) Close() bool {
	return closeConn(c.conn)
}

func (c *UDPMultiRemoteClient) SendData(data []byte, rAddress ...string) bool {
	if len(rAddress) == 0 {
		return false
	}
	var err error
	count := 0
	for _, address := range rAddress {
		addr, ok := c.mapAddr[address]
		if !ok {
			addr, err = net.ResolveUDPAddr(c.Network, address)
			if nil != err {
				logCreateAddrErr(address, err)
				continue
			}
		}
		c.conn.WriteToUDP(data, addr)
		count++
	}
	return count > 0
}

func (c *UDPMultiRemoteClient) SetReceivingHandler(handler func(data []byte, rAddr *net.UDPAddr)) {
	log.Fatalln("UDPMultiRemoteClient does not support Receiving!")
}

//private ---------------

func logCreateAddrErr(address string, err error) {
	log.Fatalln("ResolveUDPAddr Error: ", address, ": %v", err)
}

func closeConn(conn *net.UDPConn) bool {
	if nil != conn {
		conn.Close()
		return true
	}
	return false
}
