package net

import (
	"log"
	"net"
)

func NewTCPServer(maxLinkNum int) ITCPServer {
	rs := &TCPServer{maxLinkNum: maxLinkNum}
	rs.mapTransceiver = make(map[string]ITransceiver)
	return rs
}

type ITCPServer interface {
	//会阻塞
	StartServer(address string)
	StopServer()
	GetTransceiver(key string) ITransceiver
}

type TCPServer struct {
	maxLinkNum int
	timeout    int
	serverSem  chan bool

	listener       *net.TCPListener
	mapTransceiver map[string]ITransceiver
}

func (s *TCPServer) StartServer(address string) {
	listener, _ := listenTCP("tcp", address)
	s.listener = listener
	s.serverSem = make(chan bool, s.maxLinkNum)
	for {
		tcpConn, err := listener.AcceptTCP()
		if nil != err { //Listener已经关闭
			log.Fatalln(err)
			break
		}
		s.serverSem <- true
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
		s.mapTransceiver = make(map[string]ITransceiver)
	}()
	s.listener.Close()
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
		delete(s.mapTransceiver, key)
	}()
	transceiver := NewTransceiver(conn)
	s.mapTransceiver[key] = transceiver
	transceiver.StartReceiving()
}

func listenTCP(network string, address string) (*net.TCPListener, string) {
	tcpAddr, _ := net.ResolveTCPAddr(network, address)
	listener, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		log.Fatalln("\tnet.ListenTCP:", network, address, ": %v", err)
	}
	return listener, listener.Addr().String()
}
