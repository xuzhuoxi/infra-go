package netx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"net"
)

func NewTCPServer(maxLinkNum int) ITCPServer {
	server := &TCPServer{maxLinkNum: maxLinkNum}
	server.Name = "TCPServer"
	server.Network = TcpNetwork
	server.Logger = logx.DefaultLogger()
	server.PackHandler = DefaultPackHandler
	return server
}

type ITCPServer interface {
	ISockServer
}

type TCPServer struct {
	SockServerBase
	maxLinkNum int
	timeout    int

	listener      *net.TCPListener
	serverLinkSem chan bool
	mapProxy      map[string]IPackSendReceiver
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
	s.Logger.Infoln("[TCPServer] listening on:", params.LocalAddress)
	s.listener = listener
	s.serverLinkSem = make(chan bool, s.maxLinkNum)
	s.mapConn = make(map[string]*net.TCPConn)
	s.mapProxy = make(map[string]IPackSendReceiver)
	s.running = true
	s.serverMu.Unlock()
	s.Logger.Infoln(funcName + "()")

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
	s.Logger.Infoln(funcName + "()")
	return nil
}

func (s *TCPServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := TcpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *TCPServer) SendBytesTo(data []byte, rAddress ...string) error {
	if 0 == len(rAddress) {
		return NoAddrError("TCPServer.SendBytesTo")
	}
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	for _, address := range rAddress {
		ts, ok := s.mapProxy[address]
		if ok {
			ts.SendBytes(data)
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
	connProxy := &ReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	proxy := NewPackSendReceiver(connProxy, connProxy, s.PackHandler, TcpDataBlockHandler, s.Logger, false)
	s.mapProxy[address] = proxy
	s.serverMu.Unlock()
	s.Logger.Traceln("[TCPServer] TCP Connection:", address, "Opened!")
	proxy.StartReceiving()
	s.Logger.Traceln("[TCPServer] TCP Connection:", address, "Closed!")
}

func closeLinkChannel(c chan bool) {
	close(c)
	//s.Logger.Traceln("closeLinkChannel.finish")
}

func listenTCP(network string, address string) (*net.TCPListener, error) {
	tcpAddr, _ := GetTCPAddr(network, address)
	return net.ListenTCP(network, tcpAddr)
}
