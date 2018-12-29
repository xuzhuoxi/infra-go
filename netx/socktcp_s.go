package netx

import (
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
	"net"
)

func NewTCPServer(maxLinkNum int) ITCPServer {
	rs := &TCPServer{maxLinkNum: maxLinkNum}
	rs.Name = "TCPServer"
	rs.Network = TcpNetwork
	rs.splitHandler = DefaultByteSplitHandler
	rs.messageHandler = DefaultMessageHandler
	return rs
}

type TCPServer struct {
	SockServerBase
	maxLinkNum int
	timeout    int

	listener      *net.TCPListener
	serverLinkSem chan bool
	mapProxy      map[string]IMessageSendReceiver
	mapConn       map[string]*net.TCPConn
}

func (s *TCPServer) StartServer(params SockParams) error {
	funcName := "TCPServer.StartServer"
	s.serverMu.Lock()
	if s.running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		s.Network = params.Network
	}
	listener, err := listenTCP(s.Network, params.LocalAddress)
	if nil != err {
		defer s.serverMu.Unlock()
		return err
	}
	s.listener = listener
	s.serverLinkSem = make(chan bool, s.maxLinkNum)
	s.mapConn = make(map[string]*net.TCPConn)
	s.mapProxy = make(map[string]IMessageSendReceiver)
	s.running = true
	s.serverMu.Unlock()
	logx.Infoln(funcName + "()")

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
		go s.processTCPConn(rAddress, tcpConn)
	}
	return nil
}

func (s *TCPServer) StopServer() error {
	funcName := "TCPServer.StopServer"
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	if !s.running {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if nil != s.listener {
		s.listener.Close()
		s.listener = nil
	}
	for _, value := range s.mapConn {
		value.Close()
	}
	s.mapConn = nil
	closeLinkChannel(s.serverLinkSem)
	s.running = false
	logx.Infoln(funcName + "()")
	return nil
}

func (s *TCPServer) SendDataTo(data []byte, rAddress ...string) error {
	if 0 == len(rAddress) {
		return NoAddrError("TCPServer.SendData")
	}
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	for _, address := range rAddress {
		ts, ok := s.mapProxy[address]
		if ok {
			ts.SendMessage(data)
		}
	}
	return nil
}

//private -----------------

func (s *TCPServer) processTCPConn(address string, conn *net.TCPConn) {
	defer func() {
		s.serverMu.Lock()
		defer s.serverMu.Unlock()
		if nil != conn {
			conn.Close()
		}
		delete(s.mapConn, address)
		delete(s.mapProxy, address)
		<-s.serverLinkSem
	}()
	s.serverMu.Lock()
	s.mapConn[address] = conn
	connProxy := &ReadWriterProxy{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	proxy := NewMessageSendReceiver(connProxy, connProxy, TcpRW, s.Network)
	s.mapProxy[address] = proxy
	proxy.SetSplitHandler(s.splitHandler)
	proxy.SetMessageHandler(s.messageHandler)
	s.serverMu.Unlock()
	logx.Traceln("New TCP Connection:", address)
	proxy.StartReceiving()
}

func closeLinkChannel(c chan bool) {
	close(c)
	//logx.Traceln("closeLinkChannel.finish")
}

func listenTCP(network string, address string) (*net.TCPListener, error) {
	tcpAddr, _ := GetTCPAddr(network, address)
	return net.ListenTCP(network, tcpAddr)
}
