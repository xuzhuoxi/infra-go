package net

import (
	"log"
	"testing"
	"time"
)

func TestUDP(t *testing.T) {
	server := NewUDPServer()
	go server.StartServer(":9999")
	log.Println("Server Start!")

	client1 := NewUDPClient(false)
	client1.Setup(":9998")
	go func() {
		for {
			client1.SendData([]byte{1, 3}, "127.0.0.1:9999")
		}
	}()

	client2 := NewUDPClient(true)
	client2.Setup(":9999")
	go func() {
		for {
			client2.SendData([]byte{2, 0}, "")
		}
	}()
	time.Sleep(3 * time.Second)
}
