package netx

import (
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
	"golang.org/x/net/websocket"
	"net/http"
)

func NewWebSocketServer(maxLinkNum int) IWebSocketServer {
	rs := &WebSocketServer{maxLinkNum: maxLinkNum}
	rs.Network = WSNetwork
	rs.splitHandler = DefaultByteSplitHandler
	rs.messageHandler = DefaultMessageHandler
	return rs
}

type WebSocketServer struct {
	SockServerBase
	maxLinkNum int

	httpServer    *http.Server
	mapProxy      map[string]IMessageSendReceiver
	mapConn       map[string]*websocket.Conn
	serverLinkSem chan bool
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
	s.serverLinkSem = make(chan bool, s.maxLinkNum)
	s.mapConn = make(map[string]*websocket.Conn)
	s.mapProxy = make(map[string]IMessageSendReceiver)
	s.running = true
	s.serverMu.Unlock()
	logx.Infoln(funcName + "()")
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
	defer s.serverMu.Unlock()
	if !s.running {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if nil != s.httpServer {
		s.httpServer.Close()
		s.httpServer = nil
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

func (s *WebSocketServer) SendDataTo(data []byte, rAddress ...string) error {
	if 0 == len(rAddress) {
		return NoAddrError("WebSocketServer.SendDataTo")
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

//通常来说:
//LocalAddr=ws://ip:port+pattern
//RemoteAddr=Origin
func (s *WebSocketServer) onWSConn(conn *websocket.Conn) {
	address := conn.Request().RemoteAddr //最根的地址信息
	s.serverLinkSem <- true
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
	connProxy := &WSConnReadWriter{Reader: conn, Writer: conn, RemoteAddrString: conn.Request().RemoteAddr}
	proxy := NewMessageSendReceiver(connProxy, connProxy, false)
	s.mapProxy[address] = proxy
	proxy.SetSplitHandler(s.splitHandler)
	proxy.SetMessageHandler(s.messageHandler)
	s.serverMu.Unlock()
	logx.Traceln("New WebSocket Connection:", address)
	proxy.StartReceiving()
}

func (s *WebSocketServer) closeLinkChannel(c chan bool) {
	close(c)
	//logx.Traceln("closeLinkChannel.finish")
}