package wsx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"golang.org/x/net/websocket"
	"net/http"
)

func NewWebSocketServer() IWebSocketServer {
	return newWebSocketServer().(IWebSocketServer)
}

func newWebSocketServer() netx.ISockServer {
	server := &WebSocketServer{
		logFuncNameSend: "WebSocketServer[WSServer].SendBytesTo",
	}
	server.Name = "WSServer"
	server.Network = netx.WSNetwork
	server.Logger = logx.DefaultLogger()
	server.PackHandlerContainer = netx.NewIPackHandler(nil)
	return server
}

//-------------------------

type IWebSocketServer interface {
	netx.ISockServer
	eventx.IEventDispatcher
}

type WebSocketServer struct {
	eventx.EventDispatcher
	netx.SockServerBase
	logFuncNameSend string

	channelLimit lang.ChannelLimit
	httpServer   *http.Server
	mapConn      map[string]netx.IServerConn
}

func (s *WebSocketServer) StartServer(params netx.SockParams) error {
	funcName := fmt.Sprintf("WebSocketServer[%s].StartServer", s.Name)
	s.ServerMu.Lock()
	if s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer s.StopServer()
	httpMux := http.NewServeMux()
	httpMux.Handle(params.WSPattern, websocket.Handler(s.onWSConn))
	s.httpServer = &http.Server{Addr: params.LocalAddress, Handler: httpMux}
	s.channelLimit.StartLimit()
	s.mapConn = make(map[string]netx.IServerConn)
	s.Running = true
	s.ServerMu.Unlock()
	s.Logger.Infoln(funcName + "()")
	s.DispatchServerStartedEvent(s)
	err := s.httpServer.ListenAndServe()
	if nil != err {
		s.Running = false
		return err
	}
	return nil
}

func (s *WebSocketServer) StopServer() error {
	funcName := fmt.Sprintf("WebSocketServer[%s].StopServer", s.Name)
	s.ServerMu.Lock()
	if !s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() { //应该解锁后抛出事件
		s.ServerMu.Unlock()
		s.Logger.Infoln(funcName + "()")
		s.DispatchServerStoppedEvent(s)
	}()
	if nil != s.httpServer {
		s.httpServer.Close()
		s.httpServer = nil
	}
	for _, value := range s.mapConn {
		value.CloseConn()
	}
	s.mapConn = nil
	s.channelLimit.StopLimit()
	s.Running = false
	return nil
}

func (s *WebSocketServer) SetMaxConn(max int) {
	s.channelLimit.SetMax(max)
}

func (s *WebSocketServer) Connections() int {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return len(s.mapConn)
}

func (s *WebSocketServer) CloseConnection(address string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	if conn, ok := s.mapConn[address]; ok {
		delete(s.mapConn, address)
		err = conn.CloseConn()
		return err, nil != err
	}
	return errors.New("WebSocketServer: No Connection At " + address), false
}

func (s *WebSocketServer) FindConnection(address string) (conn netx.IServerConn, ok bool) {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	conn, ok = s.mapConn[address]
	return
}

func (s *WebSocketServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := WsDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *WebSocketServer) SendBytesTo(bytes []byte, rAddress ...string) error {
	funcName := s.logFuncNameSend
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
			ts.SendBytes(bytes)
		}
	}
	return nil
}

// 通常来说:
// LocalAddr=ws://ip:port+pattern
// RemoteAddr=Origin
func (s *WebSocketServer) onWSConn(conn *websocket.Conn) {
	address := conn.Request().RemoteAddr //客户端地址信息
	proxy := s.startConn(address, conn)
	proxy.StartReceiving() // 这里会阻塞
	s.endConn(address, conn)
}

func (s *WebSocketServer) startConn(address string, conn *websocket.Conn) netx.IPackSendReceiver {
	s.channelLimit.Add()
	s.ServerMu.Lock()
	rwProxy := &WSConnAdapter{Reader: conn, Writer: conn, remoteAddrString: conn.Request().RemoteAddr}
	proxy := netx.NewPackSendReceiver(rwProxy, rwProxy, s.PackHandlerContainer, WsDataBlockHandler, s.Logger, false)
	s.mapConn[address] = &WsSockConn{Address: address, Conn: conn, SRProxy: proxy}
	s.ServerMu.Unlock()

	s.DispatchServerConnOpenEvent(s, address)
	s.Logger.Infoln("[WebSocketServer] WebSocket Connection:", address, "Opened!")
	return proxy
}

func (s *WebSocketServer) endConn(address string, conn *websocket.Conn) {
	// 删除连接
	s.ServerMu.Lock()
	delete(s.mapConn, address)
	s.ServerMu.Unlock()

	if nil != conn {
		conn.Close()
	}
	s.channelLimit.Done()
	// 抛出事件
	s.DispatchServerConnCloseEvent(s, address)
	s.Logger.Infoln("[WebSocketServer] WebSocket Connection:", address, "Closed!")
}
