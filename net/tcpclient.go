package net

import (
	"net"
)

func NewTCPClient() ITCPClient {
	client := &TCPClient{Network: "tcp"}
	return client
}

type ITCPClient interface {
	Dial(address string) error
	Close()

	Send(data []byte) bool

	SetReceivingHandler(handler func(conn net.Conn, data []byte))
	StartReceiving()
	StopReceiving()
}

type TCPClient struct {
	Network     string
	Conn        net.Conn
	transceiver ITransceiver
}

func (c *TCPClient) Dial(address string) error {
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
	}()
	c.transceiver.StopReceiving()
}

func (c *TCPClient) Send(data []byte) bool {
	return c.transceiver.SendData(data)
}

func (c *TCPClient) SetReceivingHandler(handler func(conn net.Conn, data []byte)) {
	c.transceiver.SetReceivingHandler(handler)
}

func (c *TCPClient) StartReceiving() {
	c.transceiver.StartReceiving()
}

func (c *TCPClient) StopReceiving() {
	c.transceiver.StopReceiving()
}
