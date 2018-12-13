package net

import (
	"log"
	"net"
)

func NewTransceiver(conn net.Conn) ITransceiver {
	rs := &Transceiver{conn: conn}
	rs.messageBuff = NewMessageBuff()
	rs.msgHandler = DefaultMessageHandler
	return rs
}

type ITransceiver interface {
	GetConnection() net.Conn
	SendData(data []byte) bool

	SetSplitHandler(handler func(buff []byte) ([]byte, []byte))
	SetMessageHandler(handler func(msgData []byte, sender string, receiver string))
	StartReceiving()
	StopReceiving()
}

//Transceiver ------------
type Transceiver struct {
	conn        net.Conn
	messageBuff *MessageBuff
	msgHandler  func(data []byte, sender string, receiver string)
	receiving   bool
}

func (t *Transceiver) GetConnection() net.Conn {
	return t.conn
}

func (t *Transceiver) SendData(data []byte) bool {
	return sendData(t.conn, data)
}

func (t *Transceiver) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	t.messageBuff.SetCheckMessageHandler(handler)
}

func (t *Transceiver) SetMessageHandler(handler func(data []byte, sender string, receiver string)) {
	t.msgHandler = handler
}
func (t *Transceiver) StartReceiving() {
	if t.receiving {
		return
	}
	t.receiving = true
	defer t.StopReceiving()
	var buff [1024]byte
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
	t.messageBuff.AppendBytes(newData)
	for t.messageBuff.CheckMessage() {
		t.msgHandler(t.messageBuff.FrontMessage(), t.conn.RemoteAddr().String(), t.conn.LocalAddr().String())
	}
}

func sendData(c net.Conn, data []byte) bool {
	n, err := c.Write(data)
	log.Println("Send Data[Sender:", c.LocalAddr(), "Receiver:", c.RemoteAddr(), "data:", data, "len:", n, "]")
	if n != len(data) || err != nil {
		return false
	}
	return true
}

//------------------------------------

func DefaultMessageHandler(msgData []byte, sender string, receiver string) {
	log.Println("DefaultMessageHandler[Sender:"+sender+",Receiver:"+receiver+"]msgData:", msgData, "dataLen:", len(msgData), "]")
}
