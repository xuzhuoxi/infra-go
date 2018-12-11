package net

import (
	"bytes"
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
	rs.handler = defaultTransceiverHandler
	return rs
}

type ITransceiver interface {
	GetConnection() net.Conn
	SendData(data []byte) bool
	SetReceivingHandler(handler func(data []byte, conn net.Conn))
	StartReceiving()
	StopReceiving()
}

//Transceiver ------------
type Transceiver struct {
	conn      net.Conn
	buff      *bytes.Buffer
	handler   func(data []byte, conn net.Conn)
	receiving bool
}

func (t *Transceiver) GetConnection() net.Conn {
	return t.conn
}

func (t *Transceiver) SendData(data []byte) bool {
	return sendMsg(t.conn, data)
}

func (t *Transceiver) SetReceivingHandler(handler func(data []byte, conn net.Conn)) {
	t.handler = handler
}
func (t *Transceiver) StartReceiving() {
	if t.receiving {
		return
	}
	t.receiving = true
	defer t.StopReceiving()
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
	if t.receiving {
		t.receiving = false
	}
}

func (t *Transceiver) handleData(newData []byte) {
	t.buff.Write(newData)
	//这里分包
	if nil != t.handler {
		t.handler(t.buff.Bytes(), t.conn)
	}
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

//
//func recvMsg(c net.Conn, buff []byte) bool {
//	for read := 0; read != len(buff); {
//		n, err := c.Read(buff)
//		read += n
//		if err != nil {
//			return false
//		}
//	}
//	return true
//}
//
//func recvMsg2(c net.Conn, buff *[]byte) bool {
//	data, err := ioutil.ReadAll(c)
//	if nil != err {
//		return false
//	}
//	buff = &data
//	return true
//}

func defaultTransceiverHandler(buffData []byte, conn net.Conn) {
	log.Println("Receive Data[Sender:", conn.RemoteAddr(), "Receiver:", conn.LocalAddr(), "buff:", buffData, "buffLen:", len(buffData), "]")
}
