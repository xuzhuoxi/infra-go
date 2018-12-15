package net

import (
	"github.com/xuzhuoxi/util"
	"log"
	"net"
)

func NewUDPClient(connected bool) IUDPClient {
	if connected {
		return &UDPDialClient{Network: "udp"}
	} else {
		return &UDPListenClient{Network: "udp", messageHandler: DefaultMessageHandler}
	}
}

func NewUDPClientForMultiRemote() IUDPClient {
	return &UDPMultiRemoteClient{Network: "udp"}
}

type IUDPClient interface {
	SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error
	SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error

	Connected() bool
	Setup(lAddress string, rAddress string) bool
	Close() bool
	StartReceiving() error
	StopReceiving() error
	SendData(data []byte, rAddress string) error
	SendDataToMulti(data []byte, rAddress ...string) error
}

//UDPDialClient
type UDPDialClient struct {
	Network     string
	conn        *net.UDPConn
	transceiver ITransceiver
}

func (c *UDPDialClient) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	c.transceiver.SetSplitHandler(handler)
	return nil
}

func (c *UDPDialClient) SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error {
	c.transceiver.SetMessageHandler(handler)
	return nil
}

func (c *UDPDialClient) Connected() bool { return true }

func (c *UDPDialClient) Setup(lAddress string, rAddress string) bool {
	if nil != c.conn {
		log.Fatalln("\tUDPDialClient:Repeated Setup!")
		return false
	}
	lAddr, _ := getUDPAddr(c.Network, lAddress)
	rAddr, err := getUDPAddr(c.Network, rAddress)
	if nil != err {
		return false
	}
	conn, cErr := net.DialUDP(c.Network, lAddr, rAddr)
	if nil != cErr {
		log.Fatalln("\tUDPDialClient-net.DialUDP:", c.Network, lAddress, rAddress, cErr)
		return false
	}
	c.conn = conn
	c.transceiver = NewTransceiver(conn)
	return true
}

func (c *UDPDialClient) Close() bool {
	rs := closeConn(c.conn)
	c.conn = nil
	log.Println("UDPDialClient:Close()")
	return rs
}

func (c *UDPDialClient) SendData(data []byte, rAddress string) error {
	if nil == c.conn {
		return ConnNilError("c.conn")
	}
	_, err := c.conn.Write(data)
	return err
}

func (c *UDPDialClient) SendDataToMulti(data []byte, rAddress ...string) error {
	return util.FuncUnavailableError("UDPDialClient.SendDataToMulti")
}

func (c *UDPDialClient) StartReceiving() error {
	return c.transceiver.StartReceiving()
}

func (c *UDPDialClient) StopReceiving() error {
	return c.transceiver.StopReceiving()
}

//UDPListenClient
type UDPListenClient struct {
	Network    string
	conn       *net.UDPConn
	remoteAddr *net.UDPAddr

	messageBuff    *MessageBuff
	messageHandler func(data []byte, conn net.Conn, senderAddress string)
	receiving      bool
}

func (c *UDPListenClient) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	c.messageBuff.SetCheckMessageHandler(handler)
	return nil
}

func (c *UDPListenClient) SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error {
	c.messageHandler = handler
	return nil
}

func (c *UDPListenClient) Connected() bool { return false }

func (c *UDPListenClient) Setup(lAddress string, rAddress string) bool {
	if nil != c.conn {
		log.Fatalln("\tUDPListenClient:Repeated Setup!")
		return false
	}
	lAddr, err := getUDPAddr(c.Network, lAddress)
	if nil != err {
		return false
	}
	conn, cErr := net.ListenUDP(c.Network, lAddr)
	if nil != cErr {
		log.Fatalln("\tUDPListenClient:Setup Error:", c.Network, lAddress, cErr)
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
	rs := closeConn(c.conn)
	c.conn = nil
	log.Println("UDPListenClient:Close()")
	return rs
}

func (c *UDPListenClient) SendData(data []byte, rAddress string) error {
	if nil == c.conn {
		return ConnNilError("UDPListenClient.SendData")
	}
	if nil == c.remoteAddr {
		rAddr, err := getUDPAddr(c.Network, rAddress)
		if nil != err {
			return err
		}
		c.remoteAddr = rAddr
	}
	_, e := c.conn.WriteToUDP(data, c.remoteAddr)
	return e
}

func (c *UDPListenClient) SendDataToMulti(data []byte, rAddress ...string) error {
	return util.FuncUnavailableError("UDPListenClient.SendDataToMulti")
}

func (c *UDPListenClient) StartReceiving() error {
	funcName := "UDPListenClient.StartReceiving"
	if nil == c.conn {
		return ConnNilError(funcName)
	}
	if c.receiving {
		return util.FuncRepeatedCallError(funcName)
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
			c.messageHandler(c.messageBuff.FrontMessage(), c.conn, addr.String())
		}
	}
	return nil
}

func (c *UDPListenClient) StopReceiving() error {
	if c.receiving {
		c.receiving = false
		return nil
	}
	return util.FuncRepeatedCallError("UDPListenClient.StopReceiving")
}

//UDPMultiRemoteClient
type UDPMultiRemoteClient struct {
	Network string
	conn    *net.UDPConn
	mapAddr map[string]*net.UDPAddr
	handler func(data []byte, rAddr *net.UDPAddr)
}

func (c *UDPMultiRemoteClient) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	return util.FuncUnavailableError("UDPMultiRemoteClient.SetSplitHandler")
}

func (c *UDPMultiRemoteClient) SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error {
	return util.FuncUnavailableError("UDPMultiRemoteClient.SetMessageHandler")
}

func (c *UDPMultiRemoteClient) Connected() bool { return false }

func (c *UDPMultiRemoteClient) Setup(lAddress string, rAddress string) bool {
	if nil != c.conn {
		log.Fatalln("\tUDPMultiRemoteClient:Repeated Setup!")
		return false
	}
	lAddr, err := getUDPAddr(c.Network, lAddress)
	if nil != err {
		return false
	}
	conn, cErr := net.ListenUDP(c.Network, lAddr)
	if nil != cErr {
		log.Fatalln("\tUDPMultiRemoteClient:Setup Error:", c.Network, lAddress, cErr)
		return false
	}
	c.conn = conn
	c.mapAddr = make(map[string]*net.UDPAddr)
	return true
}

func (c *UDPMultiRemoteClient) Close() bool {
	rs := closeConn(c.conn)
	c.conn = nil
	log.Println("UDPMultiRemoteClient:Close()")
	return rs
}

func (c *UDPMultiRemoteClient) SendData(data []byte, rAddress string) error {
	if nil == c.conn {
		return ConnNilError("UDPMultiRemoteClient.SendData")
	}
	sendDataFromListen(c.conn, data, rAddress)
	return nil
}

func (c *UDPMultiRemoteClient) SendDataToMulti(data []byte, rAddress ...string) error {
	funcName := "UDPMultiRemoteClient.SendDataToMulti"
	if nil == c.conn {
		return ConnNilError(funcName)
	}
	if len(rAddress) == 0 {
		return NoAddrError(funcName)
	}
	sendDataFromListen(c.conn, data, rAddress...)
	return nil
}

func (c *UDPMultiRemoteClient) StartReceiving() error {
	return util.FuncUnavailableError("UDPMultiRemoteClient.StartReceiving")
}

func (c *UDPMultiRemoteClient) StopReceiving() error {
	return util.FuncUnavailableError("UDPMultiRemoteClient.StopReceiving")
}

//private ---------------

func closeConn(conn *net.UDPConn) bool {
	if nil != conn {
		conn.Close()
		return true
	}
	return false
}
