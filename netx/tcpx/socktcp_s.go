package tcpx

import (
	"errors"
	"fmt"
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

	channelLimit lang.ChannelLimit
	timeout      int
	listener     *net.TCPListener
	mapConn      map[string]*TcpSockConn
}

func (s *TCPServer) StartServer(params netx.SockParams) error {
	funcName := fmt.Sprintf("TCPServer[%s].StartServer", s.Name)
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
	s.channelLimit.StartLimit()
	s.mapConn = make(map[string]*TcpSockConn)
	s.Running = true
	s.ServerMu.Unlock()
	s.Logger.Infoln(funcName + "()")
	s.DispatchServerStartedEvent(s)

	defer s.StopServer()
	for s.Running {
		s.channelLimit.Add()
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
	funcName := fmt.Sprintf("TCPServer[%s].StopServer", s.Name)
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
	for _, value := range s.mapConn {
		value.CloseConn()
	}
	s.mapConn = nil
	s.channelLimit.StopLimit()
	s.Running = false
	return nil
}

func (s *TCPServer) SetMaxConn(max int) {
	s.channelLimit.SetMax(max)
}

func (s *TCPServer) Connections() int {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return len(s.mapConn)
}

func (s *TCPServer) CloseConnection(address string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	if conn, ok := s.mapConn[address]; ok {
		delete(s.mapConn, address)
		return conn.CloseConn(), ok
	}
	return errors.New("TCPServer: No Connection At " + address), false
}

func (s *TCPServer) FindConnection(address string) (conn netx.IServerConn, ok bool) {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	conn, ok = s.mapConn[address]
	return
}

func (s *TCPServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := TcpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *TCPServer) SendBytesTo(data []byte, rAddress ...string) error {
	funcName := fmt.Sprintf("TCPServer[%s].SendBytesTo", s.Name)
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	if !s.Running || nil == s.mapConn {
		return netx.ConnNilError(funcName)
	}
	if 0 == len(rAddress) {
		return netx.NoAddrError(funcName)
	}
	for _, address := range rAddress {
		ts, ok := s.mapConn[address]
		if ok {
			ts.SendBytes(data)
		}
	}
	return nil
}

//private -----------------

func (s *TCPServer) processTCPConn(address string, conn *net.TCPConn) {
	s.ServerMu.Lock()
	connProxy := &netx.ReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	proxy := netx.NewPackSendReceiver(connProxy, connProxy, s.PackHandlerContainer, TcpDataBlockHandler, s.Logger, false)
	s.mapConn[address] = &TcpSockConn{Address: address, Conn: conn, SRProxy: proxy}
	s.ServerMu.Unlock()
	s.DispatchServerConnOpenEvent(s, address)
	s.Logger.Infoln("[TCPServer] TCP Connection:", address, "Opened!")

	defer func() {
		s.DispatchServerConnCloseEvent(s, address)
		s.Logger.Infoln("[TCPServer] TCP Connection:", address, "Closed!")
	}()
	defer func() {
		s.ServerMu.Lock()
		delete(s.mapConn, address)
		if nil != conn {
			conn.Close()
		}
		s.channelLimit.Done()
		s.ServerMu.Unlock()
	}()
	proxy.StartReceiving() //这里会阻塞
}

func listenTCP(network string, address string) (*net.TCPListener, error) {
	tcpAddr, _ := GetTCPAddr(network, address)
	return net.ListenTCP(network, tcpAddr)
}
