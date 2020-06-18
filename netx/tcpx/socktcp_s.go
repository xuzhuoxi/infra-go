package netx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"net"
)

func NewTCPServer() ITCPServer {
	return newTCPServer().(ITCPServer)
}

func NewTCP4Server() ITCPServer {
	return newTCP4Server().(ITCPServer)
}

func NewTCP6Server() ITCPServer {
	return newTCP6Server().(ITCPServer)
}

func newTCPServer() netx.ISockServer {
	return newTcpS("TCPServer", netx.TcpNetwork)
}

func newTCP4Server() netx.ISockServer {
	return newTcpS("TCP4Server", netx.TcpNetwork4)
}

func newTCP6Server() netx.ISockServer {
	return newTcpS("TCP6Server", netx.TcpNetwork6)
}

func newTcpS(name string, network netx.SockNetwork) netx.ISockServer {
	server := &TCPServer{}
	server.Name = name
	server.Network = network
	server.Logger = logx.DefaultLogger()
	server.PackHandlerContainer = netx.NewIPackHandler(nil)
	return server
}

//----------------------------

type ITCPServer interface {
	netx.ISockServer
	eventx.IEventDispatcher
}

type TCPServer struct {
	eventx.EventDispatcher
	netx.SockServerBase
	lang.ChannelLimit

	timeout  int
	listener *net.TCPListener
	mapProxy map[string]netx.IPackSendReceiver
	mapConn  map[string]*net.TCPConn
}

func (s *TCPServer) StartServer(params netx.SockParams) error {
	funcName := "TCPServer.StartServer"
	s.ServerMu.Lock()
	if s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		s.Network = params.Network
	}
	listener, err := listenTCP(s.Network.String(), params.LocalAddress)
	if nil != err {
		defer s.ServerMu.Unlock()
		return err
	}
	s.Logger.Infoln("[TCPServer] listening on:", params.LocalAddress)
	s.listener = listener
	s.ChannelLimit.StartLimit()
	s.mapConn = make(map[string]*net.TCPConn)
	s.mapProxy = make(map[string]netx.IPackSendReceiver)
	s.Running = true
	s.ServerMu.Unlock()
	s.Logger.Infoln(funcName + "()")
	s.DispatchServerStartedEvent(s)

	defer s.StopServer()
	for s.Running {
		s.ChannelLimit.Add()
		if !s.Running {
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
	s.ServerMu.Lock()
	if !s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.ServerMu.Unlock()
		s.Logger.Infoln(funcName + "()")
		s.DispatchServerStoppedEvent(s)
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
	s.ChannelLimit.StopLimit()
	s.Running = false
	return nil
}

func (s *TCPServer) Connections() int {
	return len(s.mapConn)
}

func (s *TCPServer) CloseConnection(address string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	if conn, ok := s.mapConn[address]; ok {
		delete(s.mapProxy, address)
		delete(s.mapConn, address)
		return conn.Close(), ok
	}
	return errors.New("TCPServer: No Connection At " + address), false
}

func (s *TCPServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := TcpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *TCPServer) SendBytesTo(data []byte, rAddress ...string) error {
	funcName := "TCPServer.SendBytesTo"
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	if !s.Running || nil == s.mapProxy || nil == s.mapConn {
		return netx.ConnNilError(funcName)
	}
	if 0 == len(rAddress) {
		return netx.NoAddrError(funcName)
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
	s.ServerMu.Lock()
	s.mapConn[address] = conn
	connProxy := &netx.ReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	proxy := netx.NewPackSendReceiver(connProxy, connProxy, s.PackHandlerContainer, TcpDataBlockHandler, s.Logger, false)
	s.mapProxy[address] = proxy
	s.ServerMu.Unlock()
	s.DispatchServerConnOpenEvent(s, address)
	s.Logger.Traceln("[TCPServer] TCP Connection:", address, "Opened!")

	defer func() {
		s.DispatchServerConnCloseEvent(s, address)
		s.Logger.Traceln("[TCPServer] TCP Connection:", address, "Closed!")
	}()
	defer func() {
		s.ServerMu.Lock()
		if nil != conn {
			conn.Close()
		}
		delete(s.mapConn, address)
		delete(s.mapProxy, address)
		s.ChannelLimit.Done()
		s.ServerMu.Unlock()
	}()
	proxy.StartReceiving() //这里会阻塞
}

func listenTCP(network string, address string) (*net.TCPListener, error) {
	tcpAddr, _ := GetTCPAddr(network, address)
	return net.ListenTCP(network, tcpAddr)
}
