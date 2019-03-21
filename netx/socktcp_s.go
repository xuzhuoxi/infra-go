package netx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/logx"
	"net"
)

func NewTCPServer() ITCPServer {
	server := &TCPServer{}
	server.Name = "TCPServer"
	server.Network = TcpNetwork
	server.Logger = logx.DefaultLogger()
	server.PackHandler = NewIPackHandler(nil)
	return server
}

type ITCPServer interface {
	ISockServer
	eventx.IEventDispatcher
}

type TCPServer struct {
	eventx.EventDispatcher
	SockServerBase
	LinkLimit

	timeout  int
	listener *net.TCPListener
	mapProxy map[string]IPackSendReceiver
	mapConn  map[string]*net.TCPConn
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
	s.LinkLimit.StartLimit()
	s.mapConn = make(map[string]*net.TCPConn)
	s.mapProxy = make(map[string]IPackSendReceiver)
	s.running = true
	s.serverMu.Unlock()
	s.dispatchServerStartedEvent(s)
	s.Logger.Infoln(funcName + "()")

	defer s.StopServer()
	for s.running {
		s.LinkLimit.Add()
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
	if !s.running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.serverMu.Unlock()
		s.dispatchServerStoppedEvent(s)
		s.Logger.Infoln(funcName + "()")
	}()
	if nil != s.listener {
		s.listener.Close()
		s.listener = nil
	}
	for _, value := range s.mapProxy {
		value.StopReceiving()
	}
	s.mapProxy = nil
	for _, value := range s.mapConn {
		value.Close()
	}
	s.mapConn = nil
	s.LinkLimit.StopLimit()
	s.running = false
	return nil
}

func (s *TCPServer) CloseConnection(address string) error {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	if conn, ok := s.mapConn[address]; ok {
		delete(s.mapProxy, address)
		delete(s.mapConn, address)
		return conn.Close()
	}
	return errors.New("TCPServer: No Connection At " + address)
}

func (s *TCPServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := TcpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *TCPServer) SendBytesTo(data []byte, rAddress ...string) error {
	funcName := "TCPServer.SendBytesTo"
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	if !s.running || nil == s.mapProxy || nil == s.mapConn {
		return ConnNilError(funcName)
	}
	if 0 == len(rAddress) {
		return NoAddrError(funcName)
	}
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
	s.serverMu.Lock()
	s.mapConn[address] = conn
	connProxy := &ReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	proxy := NewPackSendReceiver(connProxy, connProxy, s.PackHandler, TcpDataBlockHandler, s.Logger, false)
	s.mapProxy[address] = proxy
	s.serverMu.Unlock()
	s.dispatchServerConnOpenEvent(s, address)
	s.Logger.Traceln("[TCPServer] TCP Connection:", address, "Opened!")

	defer func() {
		s.dispatchServerConnCloseEvent(s, address)
		s.Logger.Traceln("[TCPServer] TCP Connection:", address, "Closed!")
	}()
	defer func() {
		s.serverMu.Lock()
		if nil != conn {
			conn.Close()
		}
		delete(s.mapConn, address)
		delete(s.mapProxy, address)
		s.LinkLimit.Done()
		s.serverMu.Unlock()
	}()
	proxy.StartReceiving() //这里会阻塞
}

func closeLinkChannel(c chan struct{}) {
	close(c)
	//s.Logger.Traceln("closeLinkChannel.finish")
}

func listenTCP(network string, address string) (*net.TCPListener, error) {
	tcpAddr, _ := GetTCPAddr(network, address)
	return net.ListenTCP(network, tcpAddr)
}
