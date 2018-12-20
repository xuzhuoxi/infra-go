package netx

import (
	"github.com/xuzhuoxi/go-util/errsx"
	"log"
	"net"
	"sync"
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
	SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error
	SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error
	StartServer(address string) error //会阻塞
	StopServer() error
	SendData(data []byte, rAddress ...string) error
}

type UDPServer struct {
	Network string

	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(data []byte, conn net.Conn, senderAddress string)

	conn        *net.UDPConn
	mapBuff     map[string]*MessageBuff
	mapLock     sync.RWMutex
	running     bool
	runningLock sync.Mutex
}

func (s *UDPServer) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	s.splitHandler = handler
	return nil
}

func (s *UDPServer) SetMessageHandler(handler func(data []byte, conn net.Conn, senderAddress string)) error {
	s.messageHandler = handler
	return nil
}

func (s *UDPServer) StartServer(address string) error {
	s.runningLock.Lock()
	if s.running {
		s.runningLock.Unlock()
		return errsx.FuncRepeatedCallError("UDPServer.StartServer")
	}
	s.running = true
	conn, err := listenUDP(s.Network, address)
	if nil != err {
		log.Fatalln(err)
		s.runningLock.Unlock()
		return err
	}
	s.conn = conn
	s.mapBuff = make(map[string]*MessageBuff)
	data := make([]byte, 1024)
	s.runningLock.Unlock()
	defer s.StopServer()
	for s.running {
		n, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			break
		}
		s.handleData(data[:n], rAddr)
	}
	return nil
}

func (s *UDPServer) StopServer() error {
	funcName := "UDPServer.StopServer"
	s.runningLock.Lock()
	defer s.runningLock.Unlock()
	if !s.running {
		return errsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.running = false
	}()
	if nil != s.conn {
		s.conn.Close()
	}
	log.Println(funcName + "()")
	return nil
}

func (s *UDPServer) SendData(data []byte, rAddress ...string) error {
	funcName := "UDPServer.SendData"
	if !s.running {
		return errsx.FuncNotPreparedError(funcName)
	}
	if len(rAddress) == 0 {
		return NoAddrError(funcName)
	}
	sendDataFromListen(s.conn, data, rAddress...)
	return nil
}

//private ----------
func (s *UDPServer) handleData(data []byte, rAddr *net.UDPAddr) {
	//fmt.Println("handleData:", data, rAddr)
	senderAddress := rAddr.String()
	buff, ok := s.mapBuff[senderAddress]
	if !ok {
		buff = NewMessageBuff()
		s.setMapValue(senderAddress, buff)
		buff.SetCheckMessageHandler(s.splitHandler)
	}
	buff.AppendBytes(data)
	for buff.CheckMessage() {
		s.messageHandler(buff.FrontMessage(), s.conn, senderAddress)
	}
}

func (s *UDPServer) setMapValue(key string, value *MessageBuff) {
	s.mapLock.Lock()
	defer s.mapLock.Unlock()
	if nil == value {
		delete(s.mapBuff, key)
	} else {
		s.mapBuff[key] = value
	}
}

func listenUDP(network string, address string) (*net.UDPConn, error) {
	udpAddr, _ := getUDPAddr(network, address)
	return net.ListenUDP(network, udpAddr)
}

func (s *UDPServer) defaultUDPHandler(data []byte, rAddr *net.UDPAddr) {
	log.Println("defaultUDPHandler[S:", rAddr.String(), "R:", s.conn.LocalAddr().String(), "data:", data, "]")
}
