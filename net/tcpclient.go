package net

import (
	"net"
	"sync"
)

func NewTCPClient() ITCPClient {
	client := &TCPClient{Network: "tcp"}
	return client
}

type ITCPClient interface {
	Dial(address string) error
	Close()
	Send(data []byte) error
	GetTransceiver() ITransceiver
}

type TCPClient struct {
	Network     string
	Conn        net.Conn
	transceiver ITransceiver
	clientLock  sync.Mutex
}

func (c *TCPClient) Dial(address string) error {
	c.clientLock.Lock()
	defer c.clientLock.Unlock()
	conn, err := net.Dial(c.Network, address)
	if nil != err {
		return err
	}
	c.Conn = conn
	c.transceiver = NewTransceiver(conn)
	return nil
}

func (c *TCPClient) Close() {
	defer func() {
		c.Conn.Close()
		c.Conn = nil
		c.transceiver = nil
		c.clientLock.Unlock()
	}()
	c.clientLock.Lock()
	c.transceiver.StopReceiving()
}

func (c *TCPClient) Send(data []byte) error {
	return c.transceiver.SendData(data)
}

func (c *TCPClient) GetTransceiver() ITransceiver {
	return c.transceiver
}
