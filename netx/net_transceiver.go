package netx

import (
	"fmt"
	"github.com/xuzhuoxi/go-util/errsx"
	"log"
	"net"
	"sync"
)

func NewTransceiver(conn net.Conn) ITransceiver {
	rs := &Transceiver{conn: conn}
	rs.messageBuff = NewMessageBuff()
	rs.msgHandler = DefaultMessageHandler
	return rs
}

type ITransceiver interface {
	GetConnection() net.Conn
	SendData(data []byte) error

	SetSplitHandler(handler func(buff []byte) ([]byte, []byte))
	SetMessageHandler(handler func(msgData []byte, conn net.Conn, senderAddress string))
	StartReceiving() error
	StopReceiving() error
}

//Transceiver ------------
type Transceiver struct {
	conn        net.Conn
	messageBuff *MessageBuff
	msgHandler  func(data []byte, conn net.Conn, senderAddress string)
	receiving   bool
	lock        sync.RWMutex
}

func (t *Transceiver) GetConnection() net.Conn {
	return t.conn
}

func (t *Transceiver) SendData(data []byte) error {
	if nil == t.conn {
		return ConnNilError("Transceiver.SendData")
	}
	fmt.Println("SendData:", data, t.conn.LocalAddr().String(), t.conn.RemoteAddr().String())
	return sendData(t.conn, data)
}

func (t *Transceiver) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	t.messageBuff.SetCheckMessageHandler(handler)
}

func (t *Transceiver) SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) {
	t.msgHandler = handler
}
func (t *Transceiver) StartReceiving() error {
	funcName := "Transceiver.StartReceiving"
	t.lock.Lock()
	if t.receiving {
		t.lock.Unlock()
		return errsx.FuncRepeatedCallError(funcName)
	}
	t.receiving = true
	t.lock.Unlock()
	defer t.StopReceiving()
	var buff [1024]byte
	for t.receiving {
		n, err := t.conn.Read(buff[:])
		if err != nil {
			break
		}
		t.handleData(buff[:n])
	}
	return nil
}
func (t *Transceiver) StopReceiving() error {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.receiving {
		t.receiving = false
		return nil
	}
	return errsx.FuncRepeatedCallError("Transceiver.StopReceiving")
}

func (t *Transceiver) handleData(newData []byte) {
	t.messageBuff.AppendBytes(newData)
	for t.messageBuff.CheckMessage() {
		t.msgHandler(t.messageBuff.FrontMessage(), t.conn, t.conn.RemoteAddr().String())
	}
}

func sendData(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	return err
}

//------------------------------------

func DefaultMessageHandler(msgData []byte, conn net.Conn, senderAddress string) {
	log.Println("DefaultMessageHandler[Sender:"+senderAddress+",Receiver:"+conn.LocalAddr().String()+"]msgData:", msgData, "dataLen:", len(msgData), "]")
}
