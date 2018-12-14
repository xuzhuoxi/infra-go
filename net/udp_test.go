package net

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestUDP(t *testing.T) {
	server := NewUDPServer()
	go server.StartServer(":9999")
	log.Println("Server Start!")

	client1 := NewUDPClient(false)
	client1.Setup(":9998", "")
	go func() {
		for {
			client1.SendData([]byte{1, 3, 3, 21, 5, 6, 7}, "127.0.0.1:9999")
		}
	}()

	client2 := NewUDPClient(true)
	client2.Setup("", "127.0.0.1:9999")
	go func() {
		for {
			client2.SendData([]byte{2, 0}, "")
		}
	}()
	time.Sleep(3 * time.Second)
}

func TestUDP2(t *testing.T) {
	//ports := []int{9990, 9991, 9992, 9993, 9994, 9995, 9996, 9997, 9998, 9999}
	ports := []int{9990}
	addrs := []string{}
	for _, port := range ports {
		server := NewUDPServer()
		address := "127.0.0.1:" + strconv.Itoa(port)
		addrs = append(addrs, address)
		go server.StartServer(address)
		log.Println("Server Start!")
	}
	fmt.Println(addrs)

	client1 := NewUDPClientForMultiRemote()
	client1.Setup(":9900", "")
	go func() {
		for {
			client1.SendDataToMulti([]byte{2, 4}, addrs...)
		}
	}()
	time.Sleep(10 * time.Millisecond)
}

func TestUDP3(t *testing.T) {
	address := ":9999"
	addt, _ := getTCPAddr("tcp", address)
	addu, _ := getUDPAddr("udp", address)
	fmt.Println(addt, addt.IP, addt.Port, addt.String())
	fmt.Println(addu, addu.IP, addu.Port, addu.String())
}
