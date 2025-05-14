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
	mapConn      map[string]netx.IServerConn // [connId:netx.IServerConn]
}

func (s *WebSocketServer) StartServer(params netx.SockParams) error {
	funcName := fmt.Sprintf("[WebSocketServer(%s).StartServer]", s.Name)
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
	s.Logger.Infoln(funcName, "()")
	s.DispatchServerStartedEvent(s)
	err := s.httpServer.ListenAndServe()
	if nil != err {
		s.Running = false
		return err
	}
	return nil
}

func (s *WebSocketServer) StopServer() error {
	funcName := fmt.Sprintf("[WebSocketServer(%s).StopServer]", s.Name)
	s.ServerMu.Lock()
	if !s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() { //应该解锁后抛出事件
		s.ServerMu.Unlock()
		s.Logger.Infoln(funcName, "()")
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

func (s *WebSocketServer) CloseConnection(connId string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	if conn, ok := s.mapConn[connId]; ok {
		delete(s.mapConn, connId)
		err = conn.CloseConn()
		return err, nil != err
	}
	return errors.New("WebSocketServer: No Connection At " + connId), false
}

func (s *WebSocketServer) FindConnection(connId string) (conn netx.IServerConn, ok bool) {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	conn, ok = s.mapConn[connId]
	return
}

func (s *WebSocketServer) SendPackTo(pack []byte, connId ...string) error {
	bytes := WsDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, connId...)
}

func (s *WebSocketServer) SendBytesTo(bytes []byte, connId ...string) error {
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
			ts.SendBytes(bytes)
		}
	}
	return nil
}

// 通常来说:
// LocalAddr=ws://ip:port+pattern
// RemoteAddr=Origin
func (s *WebSocketServer) onWSConn(conn *websocket.Conn) {
	remoteAddress := conn.Request().RemoteAddr //客户端地址信息
	localAddress := conn.LocalAddr().String()
	connInfo := netx.NewConnInfo(localAddress, remoteAddress)
	proxy := s.startConn(connInfo, conn)
	proxy.StartReceiving() // 这里会阻塞
	s.endConn(connInfo, conn)
}

func (s *WebSocketServer) startConn(connInfo netx.IConnInfo, conn *websocket.Conn) netx.IPackSendReceiver {
	s.channelLimit.Add()
	s.ServerMu.Lock()
	rwProxy := &WSConnAdapter{Reader: conn, Writer: conn, remoteAddress: connInfo.GetRemoteAddress()}
	proxy := netx.NewPackSendReceiver(connInfo, rwProxy, rwProxy, s.PackHandlerContainer, WsDataBlockHandler, s.Logger, false)
	s.mapConn[connInfo.GetConnId()] = &WsSockConn{WsConnInfo: connInfo, Conn: conn, SRProxy: proxy}
	s.ServerMu.Unlock()

	s.DispatchServerConnOpenEvent(s, connInfo)
	s.Logger.Infoln("[WebSocketServer.startConn] WebSocket Connection:", connInfo, "Opened!")
	return proxy
}

func (s *WebSocketServer) endConn(connInfo netx.IConnInfo, conn *websocket.Conn) {
	// 删除连接
	s.ServerMu.Lock()
	delete(s.mapConn, connInfo.GetConnId())
	s.ServerMu.Unlock()

	if nil != conn {
		conn.Close()
	}
	s.channelLimit.Done()
	// 抛出事件
	s.DispatchServerConnCloseEvent(s, connInfo)
	s.Logger.Infoln("[WebSocketServer.endConn] WebSocket Connection:", connInfo, "Closed!")
}
