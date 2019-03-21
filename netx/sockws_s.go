package netx

import (
	"github.com/pkg/errors"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/logx"
	"golang.org/x/net/websocket"
	"net/http"
)

func NewWebSocketServer() IWebSocketServer {
	server := &WebSocketServer{}
	server.Name = "WSServer"
	server.Network = WSNetwork
	server.Logger = logx.DefaultLogger()
	server.PackHandler = NewIPackHandler(nil)
	return server
}

type IWebSocketServer interface {
	ISockServer
	eventx.IEventDispatcher
}

type WebSocketServer struct {
	eventx.EventDispatcher
	SockServerBase
	LinkLimit

	httpServer *http.Server
	mapProxy   map[string]IPackSendReceiver
	mapConn    map[string]*websocket.Conn
}

func (s *WebSocketServer) StartServer(params SockParams) error {
	funcName := "WebSocketServer.StartServer"
	s.serverMu.Lock()
	if s.running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer s.StopServer()
	httpMux := http.NewServeMux()
	httpMux.Handle(params.WSPattern, websocket.Handler(s.onWSConn))
	s.httpServer = &http.Server{Addr: params.LocalAddress, Handler: httpMux}
	s.LinkLimit.StartLimit()
	s.mapConn = make(map[string]*websocket.Conn)
	s.mapProxy = make(map[string]IPackSendReceiver)
	s.running = true
	s.serverMu.Unlock()
	s.dispatchServerStartedEvent(s)
	s.Logger.Infoln(funcName + "()")
	err := s.httpServer.ListenAndServe()
	if nil != err {
		s.running = false
		return err
	}
	return nil
}

func (s *WebSocketServer) StopServer() error {
	funcName := "WebSocketServer.StopServer"
	s.serverMu.Lock()
	if !s.running {
		defer s.serverMu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	defer func() { //应该解锁后抛出事件
		s.serverMu.Unlock()
		s.dispatchServerStoppedEvent(s)
		s.Logger.Infoln(funcName + "()")
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
	s.LinkLimit.StopLimit()
	s.running = false
	return nil
}

func (s *WebSocketServer) CloseConnection(address string) error {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	if conn, ok := s.mapConn[address]; ok {
		delete(s.mapProxy, address)
		delete(s.mapConn, address)
		return conn.Close()
	} else {
		return errors.New("WebSocketServer: No Connection At " + address)
	}
}

func (s *WebSocketServer) SendPackTo(pack []byte, rAddress ...string) error {
	bytes := WsDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, rAddress...)
}

func (s *WebSocketServer) SendBytesTo(bytes []byte, rAddress ...string) error {
	funcName := "WebSocketServer.SendPackTo"
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	if !s.running || nil == s.mapProxy {
		return ConnNilError(funcName)
	}
	if 0 == len(rAddress) {
		return NoAddrError(funcName)
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
	s.LinkLimit.Add()
	s.serverMu.Lock()
	s.mapConn[address] = conn
	connProxy := &WSConnAdapter{Reader: conn, Writer: conn, RemoteAddrString: conn.Request().RemoteAddr}
	proxy := NewPackSendReceiver(connProxy, connProxy, s.PackHandler, WsDataBlockHandler, s.Logger, false)
	s.mapProxy[address] = proxy
	s.serverMu.Unlock()
	s.dispatchServerConnOpenEvent(s, address)
	s.Logger.Traceln("[WebSocketServer] WebSocket Connection:", address, "Opened!")

	defer func() {
		s.dispatchServerConnCloseEvent(s, address)
		s.Logger.Traceln("[WebSocketServer] WebSocket Connection:", address, "Closed!")
	}()
	defer func() {
		s.serverMu.Lock()
		if nil != conn {
			conn.Close()
			conn = nil
		}
		delete(s.mapConn, address)
		delete(s.mapProxy, address)
		s.LinkLimit.Done()
		s.serverMu.Unlock()
	}()
	proxy.StartReceiving() //这里会阻塞
}

func (s *WebSocketServer) closeLinkChannel(c chan struct{}) {
	close(c)
	//s.Logger.Traceln("closeLinkChannel.finish")
}
