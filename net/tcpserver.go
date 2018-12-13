package net

import (
	"log"
	"net"
)

const (
	TcpNetwork  = "tcp"
	TcpNetwork4 = "tcp4"
	TcpNetwork6 = "tcp6"
)

func NewTCPServer(maxLinkNum int) ITCPServer {
	rs := &TCPServer{Network: TcpNetwork, maxLinkNum: maxLinkNum}
	rs.splitHandler = DefaultSplitHandler
	rs.messageHandler = DefaultMessageHandler
	return rs
}

type ITCPServer interface {
	SetSplitHandler(handler func(buff []byte) ([]byte, []byte))
	SetMessageHandler(handler func(data []byte, sender string, receiver string))
	StartServer(address string) //会阻塞
	StopServer()
	GetTransceiver(key string) ITransceiver
}

type TCPServer struct {
	Network    string
	maxLinkNum int
	timeout    int

	listener       *net.TCPListener
	mapTransceiver map[string]ITransceiver
	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(data []byte, sender string, receiver string)
	running        bool
	serverSem      chan bool
}

func (s *TCPServer) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) {
	s.splitHandler = handler
}

func (s *TCPServer) SetMessageHandler(handler func(data []byte, sender string, receiver string)) {
	s.messageHandler = handler
}

func (s *TCPServer) StartServer(address string) {
	if s.running {
		return
	}
	defer s.StopServer()
	s.running = true
	listener, _ := listenTCP(s.Network, address)
	s.listener = listener
	s.serverSem = make(chan bool, s.maxLinkNum)
	s.mapTransceiver = make(map[string]ITransceiver)
	for s.running {
		s.serverSem <- true
		tcpConn, err := listener.AcceptTCP()
		if nil != err { //Listener已经关闭
			log.Fatalln(err)
			break
		}
		key := tcpConn.RemoteAddr().String()
		log.Println("New Connection:", key)
		go s.processTCPConn(key, tcpConn)
	}
}

func (s *TCPServer) StopServer() {
	defer func() {
		for _, value := range s.mapTransceiver {
			value.GetConnection().Close()
		}
		s.mapTransceiver = nil
		close(s.serverSem)
	}()
	if nil != s.listener {
		s.listener.Close()
	}
}

func (s *TCPServer) GetTransceiver(key string) ITransceiver {
	ts, ok := s.mapTransceiver[key]
	if ok {
		return ts
	}
	return nil
}

//private -----------------

func (s *TCPServer) processTCPConn(key string, conn *net.TCPConn) {
	defer func() {
		conn.Close()
		delete(s.mapTransceiver, key)
		<-s.serverSem
	}()
	transceiver := NewTransceiver(conn)
	s.mapTransceiver[key] = transceiver
	transceiver.SetSplitHandler(s.splitHandler)
	transceiver.SetMessageHandler(s.messageHandler)
	transceiver.StartReceiving()
}

func listenTCP(network string, address string) (*net.TCPListener, string) {
	tcpAddr, _ := getTCPAddr(network, address)
	listener, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		log.Fatalln("\tnet.ListenTCP:", network, address, ": %v", err)
	}
	return listener, listener.Addr().String()
}
