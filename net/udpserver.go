package net

import (
	"bytes"
	"log"
	"net"
)

const (
	UDPBuffLength = 4096

	UDPNetwork  = "udp"
	UDPNetwork4 = "udp4"
	UDPNetwork6 = "udp6"
)

func NewUDPServer() IUDPServer {
	rs := &UDPServer{Network: "udp", handler: defaultUDPHandler}
	return rs
}

//unconnected
type IUDPServer interface {
	//会阻塞
	StartServer(address string)
	StopServer()
	SendData(data []byte, rAddr *net.UDPAddr)
	SetReceivingHandler(handler func(data []byte, rAddr *net.UDPAddr))
}

type UDPServer struct {
	Network string

	conn    *net.UDPConn
	mapBuff map[string]*bytes.Buffer
	handler func(data []byte, rAddr *net.UDPAddr)
	running bool
}

func (s *UDPServer) StartServer(address string) {
	if s.running {
		return
	}
	s.running = true
	defer s.StopServer()
	conn, _ := listenUDP(s.Network, address)
	s.conn = conn
	s.mapBuff = make(map[string]*bytes.Buffer)
	data := make([]byte, 2048)
	for s.running {
		n, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			break
		}
		s.handleData(data[:n], rAddr)
	}
}

func (s *UDPServer) StopServer() {
	defer func() {
		s.running = false
	}()
	if nil != s.conn {
		s.conn.Close()
	}
}

func (s *UDPServer) SendData(data []byte, rAddr *net.UDPAddr) {
	if s.running {
		s.conn.WriteToUDP(data, rAddr)
	}
}

func (s *UDPServer) SetReceivingHandler(handler func(data []byte, rAddr *net.UDPAddr)) {
	s.handler = handler
}

//private ----------
func (s *UDPServer) handleData(data []byte, rAddr *net.UDPAddr) {
	key := rAddr.String()
	buff, ok := s.mapBuff[key]
	if !ok {
		buff = bytes.NewBuffer(make([]byte, UDPBuffLength))
		buff.Reset()
		s.mapBuff[key] = buff
	}
	buff.Write(data)
	//这里分包
	if nil != s.handler {
		s.handler(buff.Bytes(), rAddr)
	}
}

func listenUDP(network string, address string) (*net.UDPConn, string) {
	udpAddr, _ := net.ResolveUDPAddr(network, address)
	listener, err := net.ListenUDP(network, udpAddr)
	if err != nil {
		log.Fatalln("\tnet.ListenUDP:", network, address, ": %v", err)
	}
	return listener, listener.LocalAddr().String()
}

func defaultUDPHandler(data []byte, rAddr *net.UDPAddr) {
	log.Println("defaultUDPHandler[Sender:", rAddr.String(), "data:", data, "]")
}
