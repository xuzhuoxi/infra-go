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
	SetSplitHandler(handler func(buff []byte) ([]byte, []byte))
	SetMessageHandler(handler func(data []byte, conn net.Conn))

	Connected() bool
	Setup(lAddress string, rAddress string) bool
	Close() bool
	StartReceiving()
	StopReceiving()
	SendData(data []byte, rAddress string)
	SendDataToMulti(data []byte, rAddress ...string)
}

//UDPDialClient
type UDPDialClient struct {
	Network     string
	conn        *net.UDPConn
	transceiver ITransceiver
}

func (c *UDPDialClient) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	c.transceiver.SetSplitHandler(handler)
}

func (c *UDPDialClient) SetMessageHandler(handler func(data []byte, conn net.Conn)) {
	c.transceiver.SetMessageHandler(handler)
}

func (c *UDPDialClient) Connected() bool {
	return true
}

func (c *UDPDialClient) Setup(lAddress string, rAddress string) bool {
	if nil != c.conn {
		log.Fatalln("UDPDialClient-Repeated Setup!")
		return false
	}
	lAddr, _ := getUDPAddr(c.Network, lAddress)
	rAddr, err := getUDPAddr(c.Network, rAddress)
	if nil != err {
		return false
	}
	conn, cErr := net.DialUDP(c.Network, lAddr, rAddr)
	if nil != cErr {
		log.Fatalln("\tUDPDialClient-net.DialUDP:", c.Network, lAddress, rAddress, ": %v", cErr)
		return false
	}
	c.conn = conn
	c.transceiver = NewTransceiver(conn)
	return true
}

func (c *UDPDialClient) Close() bool {
	return closeConn(c.conn)
}

func (c *UDPDialClient) SendData(data []byte, rAddress string) {
	_, err := c.conn.Write(data)
	if nil != err {
		log.Fatalln(err)
	}
}

func (c *UDPDialClient) SendDataToMulti(data []byte, rAddress ...string) {
	log.Fatalln("UDPDialClient does not support the method!")
}

func (c *UDPDialClient) StartReceiving() {
	c.transceiver.StartReceiving()
}

func (c *UDPDialClient) StopReceiving() {
	c.transceiver.StopReceiving()
}

//UDPListenClient
type UDPListenClient struct {
	Network    string
	conn       *net.UDPConn
	remoteAddr *net.UDPAddr

	messageBuff    *MessageBuff
	messageHandler func(data []byte, conn net.Conn)
	receiving      bool
}

func (c *UDPListenClient) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	c.messageBuff.SetCheckMessageHandler(handler)
}

func (c *UDPListenClient) SetMessageHandler(handler func(data []byte, conn net.Conn)) {
	c.messageHandler = handler
}

func (c *UDPListenClient) Connected() bool {
	return false
}

func (c *UDPListenClient) Setup(lAddress string, rAddress string) bool {
	lAddr, err := getUDPAddr(c.Network, lAddress)
	if nil != err {
		return false
	}
	conn, cErr := net.ListenUDP(c.Network, lAddr)
	if nil != cErr {
		log.Fatalln("\tUDPListenClient-Setup:", c.Network, lAddress, ": %v", cErr)
		return false
	}
	c.conn = conn
	c.messageBuff = NewMessageBuff()
	rAddr, rErr := getUDPAddr(c.Network, rAddress)
	if nil == rErr {
		c.remoteAddr = rAddr
	}
	return true
}

func (c *UDPListenClient) Close() bool {
	return closeConn(c.conn)
}

func (c *UDPListenClient) SendData(data []byte, rAddress string) {
	if nil == c.remoteAddr {
		rAddr, err := getUDPAddr(c.Network, rAddress)
		if nil != err {
			return
		}
		c.remoteAddr = rAddr
	}
	c.conn.WriteToUDP(data, c.remoteAddr)
}

func (c *UDPListenClient) SendDataToMulti(data []byte, rAddress ...string) {
	log.Fatalln("UDPListenClient does not support the method!")
}

func (c *UDPListenClient) StartReceiving() {
	if nil == c.conn || c.receiving {
		return
	}
	c.receiving = true
	defer c.StopReceiving()
	var buffCache [1024]byte
	for {
		n, addr, err := c.conn.ReadFromUDP(buffCache[:])
		if err != nil {
			break
		}
		if !UDPAddrEqual(addr, c.remoteAddr) {
			continue
		}
		c.messageBuff.AppendBytes(buffCache[:n])
		for c.messageBuff.CheckMessage() {
			c.messageHandler(c.messageBuff.FrontMessage(), c.conn)
		}
	}
}

func (c *UDPListenClient) StopReceiving() {
	if c.receiving {
		c.receiving = false
	}
}

//UDPMultiRemoteClient
type UDPMultiRemoteClient struct {
	Network string
	conn    *net.UDPConn
	mapAddr map[string]*net.UDPAddr
	handler func(data []byte, rAddr *net.UDPAddr)
}

func (c *UDPMultiRemoteClient) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	log.Fatalln("UDPMultiRemoteClient does not support the method[SetSplitHandler]!")
}

func (c *UDPMultiRemoteClient) SetMessageHandler(handler func(data []byte, conn net.Conn)) {
	log.Fatalln("UDPMultiRemoteClient does not support the method[SetMessageHandler]!")
}

func (c *UDPMultiRemoteClient) Connected() bool {
	return false
}

func (c *UDPMultiRemoteClient) Setup(lAddress string, rAddress string) bool {
	lAddr, err := getUDPAddr(c.Network, lAddress)
	if nil != err {
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

func (c *UDPMultiRemoteClient) SendData(data []byte, rAddress string) {
	sendDataFromListen(c.conn, data, rAddress)
}

func (c *UDPMultiRemoteClient) SendDataToMulti(data []byte, rAddress ...string) {
	if len(rAddress) == 0 {
		return
	}
	sendDataFromListen(c.conn, data, rAddress...)
}

func (c *UDPMultiRemoteClient) StartReceiving() {
	log.Fatalln("UDPMultiRemoteClient does not support the method[StartReceiving]!")
}

func (c *UDPMultiRemoteClient) StopReceiving() {
	log.Fatalln("UDPMultiRemoteClient does not support the method[StopReceiving]!")
}

//private ---------------

func closeConn(conn *net.UDPConn) bool {
	if nil != conn {
		conn.Close()
		return true
	}
	return false
}
