package net

import (
	"github.com/xuzhuoxi/util/errs"
	"log"
	"net"
	"sync"
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
	SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error
	SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error
	StartServer(address string) error //会阻塞
	StopServer() error
	Running() bool
	GetTransceiver(key string) ITransceiver
}

type TCPServer struct {
	Network    string
	maxLinkNum int
	timeout    int

	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(data []byte, conn net.Conn, senderAddress string)

	listener       *net.TCPListener
	serverLinkSem  chan bool
	mapTransceiver map[string]ITransceiver
	mapLock        sync.RWMutex
	running        bool
	runningLock    sync.Mutex
}

func (s *TCPServer) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	s.splitHandler = handler
	return nil
}

func (s *TCPServer) SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error {
	s.messageHandler = handler
	return nil
}

func (s *TCPServer) StartServer(address string) error {
	funcName := "TCPServer.StartServer"
	s.runningLock.Lock()
	if s.running {
		return errs.FuncRepeatedCallError(funcName)
	}
	log.Println(funcName + "()")
	s.running = true
	listener, err := listenTCP(s.Network, address)
	if nil != err {
		log.Fatalln(err)
		s.runningLock.Unlock()
		return err
	}
	s.listener = listener
	s.serverLinkSem = make(chan bool, s.maxLinkNum)
	s.mapTransceiver = make(map[string]ITransceiver)
	s.runningLock.Unlock()
	defer s.StopServer()
	for s.running {
		s.serverLinkSem <- true
		if !s.running {
			break
		}
		tcpConn, err := listener.AcceptTCP()
		if nil != err { //Listener已经关闭
			return err
		}
		rAddress := tcpConn.RemoteAddr().String()
		log.Println("New Connection:", rAddress)
		go s.processTCPConn(rAddress, tcpConn)
	}
	return nil
}

func (s *TCPServer) StopServer() error {
	funcName := "TCPServer.StopServer"
	s.runningLock.Lock()
	defer s.runningLock.Unlock()
	if !s.running {
		return errs.FuncRepeatedCallError(funcName)
	}
	log.Println(funcName + "()")
	defer func() {
		for _, value := range s.mapTransceiver {
			value.GetConnection().Close()
		}
		s.running = false
		close(s.serverLinkSem)
	}()
	if nil != s.listener {
		s.listener.Close()
	}
	return nil
}

func (s *TCPServer) Running() bool {
	return s.running
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
		s.setMapValue(key, nil)
		<-s.serverLinkSem
	}()
	transceiver := NewTransceiver(conn)
	s.setMapValue(key, transceiver)
	transceiver.SetSplitHandler(s.splitHandler)
	transceiver.SetMessageHandler(s.messageHandler)
	transceiver.StartReceiving()
}

func (s *TCPServer) setMapValue(key string, value ITransceiver) {
	s.mapLock.Lock()
	defer s.mapLock.Unlock()
	if nil == value {
		delete(s.mapTransceiver, key)
	} else {
		s.mapTransceiver[key] = value
	}
}

func listenTCP(network string, address string) (*net.TCPListener, error) {
	tcpAddr, _ := getTCPAddr(network, address)
	return net.ListenTCP(network, tcpAddr)
}
