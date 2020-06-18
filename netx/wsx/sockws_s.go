package wsx

import (
	"github.com/pkg/errors"
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
	server := &WebSocketServer{}
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
	lang.ChannelLimit

	httpServer *http.Server
	mapProxy   map[string]netx.IPackSendReceiver
	mapConn    map[string]*websocket.Conn
}

func (s *WebSocketServer) StartServer(params netx.SockParams) error {
	funcName := "WebSocketServer.StartServer"
	s.ServerMu.Lock()
	if s.Running {
		defer s.ServerMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer s.StopServer()
	httpMux := http.NewServeMux()
	httpMux.Handle(params.WSPattern, websocket.Handler(s.onWSConn))
	s.httpServer = &http.Server{Addr: params.LocalAddress, Handler: httpMux}
	s.ChannelLimit.StartLimit()
	s.mapConn = make(map[string]*websocket.Conn)
	s.mapProxy = make(map[string]netx.IPackSendReceiver)
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
	funcName := "WebSocketServer.StopServer"
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

func (s *WebSocketServer) Connections() int {
	return len(s.mapConn)
}

func (s *WebSocketServer) CloseConnection(address string) (err error, ok bool) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	if conn, ok := s.mapConn[address]; ok {
		delete(s.mapProxy, address)
		delete(s.mapConn, address)
		return conn.Close(), true
	} else {
		return errors.New("WebSocketServer: No Connection At " + address), false
	}
}

func (s *WebSocketServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := WsDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *WebSocketServer) SendBytesTo(bytes []byte, rAddress ...string) error {
	funcName := "WebSocketServer.SendPackTo"
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	if !s.Running || nil == s.mapProxy {
		return netx.ConnNilError(funcName)
	}
	if 0 == len(rAddress) {
		return netx.NoAddrError(funcName)
	}
	for _, address := range rAddress {
		ts, ok := s.mapProxy[address]
		if ok {
			ts.SendBytes(bytes)
		}
	}
	return nil
}

//通常来说:
//LocalAddr=ws://ip:port+pattern
//RemoteAddr=Origin
func (s *WebSocketServer) onWSConn(conn *websocket.Conn) {
	address := conn.Request().RemoteAddr //最根的地址信息
	s.ChannelLimit.Add()
	s.ServerMu.Lock()
	s.mapConn[address] = conn
	connProxy := &WSConnAdapter{Reader: conn, Writer: conn, remoteAddrString: conn.Request().RemoteAddr}
	proxy := netx.NewPackSendReceiver(connProxy, connProxy, s.PackHandlerContainer, WsDataBlockHandler, s.Logger, false)
	s.mapProxy[address] = proxy
	s.ServerMu.Unlock()
	s.DispatchServerConnOpenEvent(s, address)
	s.Logger.Traceln("[WebSocketServer] WebSocket Connection:", address, "Opened!")

	defer func() {
		s.DispatchServerConnCloseEvent(s, address)
		s.Logger.Traceln("[WebSocketServer] WebSocket Connection:", address, "Closed!")
	}()
	defer func() {
		s.ServerMu.Lock()
		if nil != conn {
			conn.Close()
			conn = nil
		}
		delete(s.mapConn, address)
		delete(s.mapProxy, address)
		s.ChannelLimit.Done()
		s.ServerMu.Unlock()
	}()
	proxy.StartReceiving() //这里会阻塞
}
