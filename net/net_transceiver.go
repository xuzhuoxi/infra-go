package net

import (
	"fmt"
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
	SetMessageHandler(handler func(msgData []byte, conn net.Conn))
	StartReceiving()
	StopReceiving()
}

//Transceiver ------------
type Transceiver struct {
	conn        net.Conn
	messageBuff *MessageBuff
	msgHandler  func(data []byte, conn net.Conn)
	receiving   bool
}

func (t *Transceiver) GetConnection() net.Conn {
	return t.conn
}

func (t *Transceiver) SendData(data []byte) bool {
	fmt.Println("SendData:", data, t.conn.LocalAddr().String(), t.conn.RemoteAddr().String())
	return sendData(t.conn, data)
}

func (t *Transceiver) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	t.messageBuff.SetCheckMessageHandler(handler)
}

func (t *Transceiver) SetMessageHandler(handler func(data []byte, conn net.Conn)) {
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
		t.msgHandler(t.messageBuff.FrontMessage(), t.conn)
	}
}

func sendData(conn net.Conn, data []byte) bool {
	n, err := conn.Write(data)
	if n != len(data) || err != nil {
		return false
	}
	return true
}

//------------------------------------

func DefaultMessageHandler(msgData []byte, conn net.Conn) {
	log.Println("DefaultMessageHandler[Sender:"+conn.RemoteAddr().String()+",Receiver:"+conn.LocalAddr().String()+"]msgData:", msgData, "dataLen:", len(msgData), "]")
}
