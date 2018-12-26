package netx

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestUDPServer(t *testing.T) {
	server := NewUDPServer()
	var msgHandler = func(msgData []byte, sender interface{}) {
		senderAddress := sender.(string)
		log.Println("TestUDPServer.msgHandler[Sender:", senderAddress, "]msgData:", msgData, "dataLen:", len(msgData), "]")
		rs := []byte{byte(len(msgData))}
		rs = append(rs, msgData...)
		server.SendDataTo(rs, senderAddress)
	}
	server.SetMessageHandler(msgHandler)
	go server.StartServer(SockParams{LocalAddress: "127.0.0.1:9999"})
	defer server.StopServer()
	time.Sleep(10 * time.Millisecond)

	client1 := NewUDPDialClient()
	client1.OpenClient(SockParams{RemoteAddress: "127.0.0.1:9999"})
	defer client1.CloseClient()
	go client1.StartReceiving()
	go func() {
		for {
			client1.SendDataTo([]byte{1, 3, 3, 21, 5, 6, 7})
		}
	}()

	client2 := NewUDPListenClient()
	client2.OpenClient(SockParams{LocalAddress: "127.0.0.1:9998"})
	defer client2.CloseClient()
	go client2.StartReceiving()
	go func() {
		for {
			client2.SendDataTo([]byte{2, 0}, "127.0.0.1:9999")
		}
	}()
	time.Sleep(1 * time.Millisecond)
}

func TestUDP2(t *testing.T) {
	//ports := []int{9990, 9991, 9992, 9993, 9994, 9995, 9996, 9997, 9998, 9999}
	ports := []int{9990}
	addrs := []string{}
	for _, port := range ports {
		server := NewUDPServer()
		address := "127.0.0.1:" + strconv.Itoa(port)
		addrs = append(addrs, address)
		go server.StartServer(SockParams{LocalAddress: address})
		log.Println("Server Start!")
	}
	fmt.Println(addrs)

	client1 := NewUDPListenClient()
	client1.OpenClient(SockParams{LocalAddress: ":9900"})
	go func() {
		for {
			client1.SendDataTo([]byte{2, 4}, addrs...)
		}
	}()
	time.Sleep(10 * time.Millisecond)
}

func TestUDP3(t *testing.T) {
	address := ":9999"
	addt, _ := GetTCPAddr("tcp", address)
	addu, _ := GetUDPAddr("udp", address)
	fmt.Println(addt, addt.IP, addt.Port, addt.String())
	fmt.Println(addu, addu.IP, addu.Port, addu.String())
}
