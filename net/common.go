package net

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
)

const (
	TransceiverBuffLength = 4096
)

func NewTransceiver(conn net.Conn) ITransceiver {
	rs := &Transceiver{conn: conn}
	rs.buff = bytes.NewBuffer(make([]byte, TransceiverBuffLength))
	rs.buff.Reset()
	rs.handler = defaultHandler
	return rs
}

type ITransceiver interface {
	SendData(data []byte) bool
	SetReceivingHandler(handler func(conn net.Conn, data []byte))
	StartReceiving()
	StopReceiving()
}

//Transceiver ------------
type Transceiver struct {
	conn      net.Conn
	buff      *bytes.Buffer
	handler   func(conn net.Conn, data []byte)
	receiving bool
}

func (t *Transceiver) SendData(data []byte) bool {
	return sendMsg(t.conn, data)
}

func (t *Transceiver) SetReceivingHandler(handler func(conn net.Conn, data []byte)) {
	t.handler = handler
}
func (t *Transceiver) StartReceiving() {
	t.receiving = true
	var buff [128]byte
	for t.receiving {
		n, err := t.conn.Read(buff[:])
		if err != nil {
			break
		}
		t.handleData(buff[:n])
	}
}
func (t *Transceiver) StopReceiving() {
	t.receiving = false
}

func (t *Transceiver) handleData(newData []byte) {
	t.buff.Write(newData)
	t.handler(t.conn, t.buff.Bytes())
}

//private ------------------------
func sendMsg(c net.Conn, data []byte) bool {
	n, err := c.Write(data)
	log.Println("Send Data[Sender:", c.LocalAddr(), "Receiver:", c.RemoteAddr(), "data:", data, "len:", n, "]")
	if n != len(data) || err != nil {
		return false
	}
	return true
}

func recvMsg(c net.Conn, buff []byte) bool {
	for read := 0; read != len(buff); {
		n, err := c.Read(buff)
		read += n
		if err != nil {
			return false
		}
	}
	return true
}

func recvMsg2(c net.Conn, buff *[]byte) bool {
	data, err := ioutil.ReadAll(c)
	if nil != err {
		return false
	}
	buff = &data
	return true
}

func defaultHandler(conn net.Conn, buffData []byte) {
	log.Println("Receive Data[Sender:", conn.RemoteAddr(), "Receiver:", conn.LocalAddr(), "buff:", buffData, "buffLen:", len(buffData), "]")
}
