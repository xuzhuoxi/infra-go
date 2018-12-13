package net

import (
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
	rs := &UDPServer{Network: "udp"}
	rs.splitHandler = DefaultSplitHandler
	rs.messageHandler = DefaultMessageHandler
	return rs
}

//unconnected
type IUDPServer interface {
	SetSplitHandler(handler func(buff []byte) ([]byte, []byte))
	SetMessageHandler(handler func(data []byte, sender string, receiver string))
	StartServer(address string) //会阻塞
	StopServer()
	SendData(data []byte, rAddress ...string)
}

type UDPServer struct {
	Network        string
	conn           *net.UDPConn
	mapBuff        map[string]*MessageBuff
	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(data []byte, sender string, receiver string)
	running        bool
}

func (s *UDPServer) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	s.splitHandler = handler
}

func (s *UDPServer) SetMessageHandler(handler func(data []byte, sender string, receiver string)) {
	s.messageHandler = handler
}

func (s *UDPServer) StartServer(address string) {
	if s.running {
		return
	}
	s.running = true
	defer s.StopServer()
	conn, _ := listenUDP(s.Network, address)
	s.conn = conn
	s.mapBuff = make(map[string]*MessageBuff)
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

func (s *UDPServer) SendData(data []byte, rAddress ...string) {
	if s.running && len(rAddress) > 0 {
		sendDataFromListen(s.conn, data, rAddress...)
	}
}

//private ----------
func (s *UDPServer) handleData(data []byte, rAddr *net.UDPAddr) {
	key := rAddr.String()
	buff, ok := s.mapBuff[key]
	if !ok {
		buff = NewMessageBuff()
		s.mapBuff[key] = buff
		buff.SetCheckMessageHandler(s.splitHandler)
	}
	buff.AppendBytes(data)
	for buff.CheckMessage() {
		s.messageHandler(buff.FrontMessage(), rAddr.String(), s.conn.LocalAddr().String())
	}
}

func listenUDP(network string, address string) (*net.UDPConn, string) {
	udpAddr, _ := getUDPAddr(network, address)
	listener, err := net.ListenUDP(network, udpAddr)
	if err != nil {
		log.Fatalln("\tnet.ListenUDP:", network, address, ": %v", err)
	}
	return listener, listener.LocalAddr().String()
}

func (s *UDPServer) defaultUDPHandler(data []byte, rAddr *net.UDPAddr) {
	log.Println("defaultUDPHandler[S:", rAddr.String(), "R:", s.conn.LocalAddr().String(), "data:", data, "]")
}
