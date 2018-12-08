package net

import (
	"log"
	"net"
)

func NewTCPServer(maxLinkNum int) ITCPServer {
	rs := &TCPServer{maxLinkNum: maxLinkNum}
	rs.mapConn = make(map[string]*net.TCPConn)
	return rs
}

type ITCPServer interface {
	//会阻塞
	StartServer(address string)
	StopServer()
	GetConn(key string) net.Conn
}

type TCPServer struct {
	maxLinkNum int
	timeout    int
	serverSem  chan bool

	listener *net.TCPListener
	mapConn  map[string]*net.TCPConn
}

func (s *TCPServer) StartServer(address string) {
	listener, _ := listenTCP("tcp", address)
	s.listener = listener
	s.serverSem = make(chan bool, s.maxLinkNum)
	for {
		tcpConn, err := listener.AcceptTCP()
		if nil != err {
			log.Fatalln(err)
			continue
		}
		s.serverSem <- true
		key := tcpConn.RemoteAddr().String()
		s.mapConn[key] = tcpConn
		log.Println("New Connection:", key)
		go s.processTCPConn(key, tcpConn)
	}
}

func (s *TCPServer) StopServer() {
	defer func() {
		s.listener.Close()
		s.listener = nil
	}()
	for _, value := range s.mapConn {
		value.Close()
	}
	s.mapConn = make(map[string]*net.TCPConn)
}

func (s *TCPServer) GetConn(key string) net.Conn {
	conn, ok := s.mapConn[key]
	if ok {
		return conn
	}
	return nil
}

//private -----------------

func (s *TCPServer) processTCPConn(key string, conn *net.TCPConn) {
	defer func() {
		delete(s.mapConn, key)
	}()
	transceiver := NewTransceiver(conn)
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
