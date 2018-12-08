package net

import (
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	server := NewTCPServer(5)
	go server.StartServer(":9999")

	client := NewTCPClient()
	client.Dial("tcp", "127.0.0.1:9999")
	time.Sleep(500 * time.Millisecond)
	client.Send([]byte{1, 2, 3, 4})
	client.Send([]byte{0, 0, 0, 0})
	client.Send([]byte{4, 3, 2, 1})
	client.Send([]byte{4, 3, 2, 1})
	client.Send([]byte{4, 3, 2, 1})
	client.Send([]byte{4, 3, 2, 1})
	client.Send([]byte{4, 3, 2, 1})
	client.Send([]byte{1, 1, 1, 1})
	client.Close()

	time.Sleep(50 * time.Second)
}
