package net

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	server := NewTCPServer(5)
	var msgHandler = func(msgData []byte, conn net.Conn) {
		log.Println("msgHandler[Sender:"+conn.RemoteAddr().String()+",Receiver:"+conn.LocalAddr().String()+"]msgData:", msgData, "dataLen:", len(msgData), "]")
		rs := []byte{byte(len(msgData))}
		rs = append(rs, msgData...)
		server.GetTransceiver(conn.RemoteAddr().String()).SendData(rs)
	}
	server.SetMessageHandler(msgHandler)
	go server.StartServer(":9999")

	client := NewTCPClient()
	client.Dial("127.0.0.1:9999")
	go client.GetTransceiver().StartReceiving()
	client.Send([]byte{3, 1, 3, 4})
	client.Send([]byte{3, 2, 0, 0})
	client.Send([]byte{3, 3, 2, 1})
	client.Send([]byte{7, 4, 2, 1})
	client.Send([]byte{3, 3, 2, 1})
	client.Send([]byte{3, 5, 2, 1})
	client.Send([]byte{3, 6, 2, 1})
	client.Send([]byte{3, 7, 1, 1})
	client.Close()

	time.Sleep(3 * time.Second)
}
