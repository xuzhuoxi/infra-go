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
	return newTcpS("TCP4Server", netx.Tcp4Network)
}

func newTCP6Server() netx.ISockServer {
	return newTcpS("TCP6Server", netx.Tcp6Network)
}

func newTcpS(name string, network netx.SockNetwork) netx.ISockServer {
	server := &TCPServer{
		logFuncNameSend: fmt.Sprintf("TCPServer[%s].SendBytesTo", name),
	}
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
	logFuncNameSend string

	channelLimit lang.ChannelLimit
	timeout      int
	listener     *net.TCPListener
	mapConn      map[string]netx.IServerConn // [connId:netx.IServerConn]
}

func (s *TCPServer) StartServer(params netx.SockParams) error {
	funcName := fmt.Sprintf("[TCPServer(%s).StartServer]", s.Name)
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
	s.Logger.Infoln(funcName, "listening on:", params.LocalAddress)
	s.listener = listener
	s.channelLimit.StartLimit()
	s.mapConn = make(map[string]netx.IServerConn)
	s.Running = true
	s.ServerMu.Unlock()
	s.Logger.Infoln(funcName, "()")
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
		go s.processTCPConn(tcpConn)
	}
	return nil
}

func (s *TCPServer) StopServer() error {
	funcName := fmt.Sprintf("[TCPServer(%s).StopServer]", s.Name)
	s.ServerMu.Lock()
	if !s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() {
		s.ServerMu.Unlock()
		s.Logger.Infoln(funcName, "()")
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

func (s *TCPServer) CloseConnection(connId string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	if conn, ok := s.mapConn[connId]; ok {
		delete(s.mapConn, connId)
		err = conn.CloseConn()
		return err, nil != err
	}
	return errors.New("TCPServer: No Connection At " + connId), false
}

func (s *TCPServer) FindConnection(connId string) (conn netx.IServerConn, ok bool) {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	conn, ok = s.mapConn[connId]
	return
}

func (s *TCPServer) SendPackTo(pack []byte, connId ...string) error {
	bytes := TcpDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, connId...)
}

func (s *TCPServer) SendBytesTo(data []byte, connId ...string) error {
	funcName := s.logFuncNameSend
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	if !s.Running || nil == s.mapConn {
		return netx.ConnNilError(funcName)
	}
	if 0 == len(connId) {
		return netx.NoAddrError(funcName)
	}
	for _, cId := range connId {
		ts, ok := s.mapConn[cId]
		if ok {
			ts.SendBytes(data)
		}
	}
	return nil
}

//private -----------------

func (s *TCPServer) processTCPConn(conn *net.TCPConn) {
	connInfo := netx.NewConnInfo(conn.LocalAddr().String(), conn.RemoteAddr().String())
	proxy := s.startConn(connInfo, conn)
	proxy.StartReceiving() // 这里会阻塞
	s.endConn(connInfo, conn)
}

func (s *TCPServer) startConn(connInfo netx.IConnInfo, conn *net.TCPConn) netx.IPackSendReceiver {
	s.ServerMu.Lock()
	rwProxy := &netx.ConnReadWriterAdapter{Reader: conn, Writer: conn, RemoteAddr: conn.RemoteAddr()}
	proxy := netx.NewPackSendReceiver(connInfo, rwProxy, rwProxy, s.PackHandlerContainer, TcpDataBlockHandler, s.Logger, false)
	s.mapConn[connInfo.GetConnId()] = &TcpSockConn{TcpConnInfo: connInfo, Conn: conn, SRProxy: proxy}
	s.ServerMu.Unlock()

	s.DispatchServerConnOpenEvent(s, connInfo)
	s.Logger.Infoln("[TCPServer.startConn]", "TCP Connection:", connInfo, "Opened!")
	return proxy
}

func (s *TCPServer) endConn(connInfo netx.IConnInfo, conn *net.TCPConn) {
	s.ServerMu.Lock()
	// 删除连接
	delete(s.mapConn, connInfo.GetConnId())
	s.ServerMu.Unlock()
	if nil != conn {
		conn.Close()
	}
	s.channelLimit.Done()
	// 抛出事件
	s.DispatchServerConnCloseEvent(s, connInfo)
	s.Logger.Infoln("[TCPServer.endConn]", "TCP Connection:", connInfo, "Closed!")
}

func listenTCP(network string, localAddress string) (*net.TCPListener, error) {
	tcpAddr, _ := GetTCPAddr(network, localAddress)
	return net.ListenTCP(network, tcpAddr)
}
