package wsx

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"golang.org/x/net/websocket"
)

// NewWebSocketServer creates a WebSocket server.
// NewWebSocketServer 创建 WebSocket 服务端。
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

// IWebSocketServer is a WebSocket-specific ISockServer with event dispatch.
// IWebSocketServer 是支持事件分发的 WebSocket 专用 ISockServer。
type IWebSocketServer interface {
	netx.ISockServer
	eventx.IEventDispatcher
}

// WebSocketServer implements IWebSocketServer over HTTP WebSocket upgrade.
// WebSocketServer 基于 HTTP WebSocket 升级实现 IWebSocketServer。
type WebSocketServer struct {
	eventx.EventDispatcher
	netx.SockServerBase
	logFuncNameSend string

	channelLimit lang.ChannelLimit
	httpServer   *http.Server
	mapConn      map[string]netx.IServerConn // [connId:netx.IServerConn]
}

// StartServer starts the HTTP server and accepts WebSocket connections.
// StartServer 启动 HTTP 服务并接受 WebSocket 连接。
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

// StopServer stops the HTTP server and closes all WebSocket connections.
// StopServer 停止 HTTP 服务并关闭所有 WebSocket 连接。
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

// SetMaxConn sets the maximum number of concurrent WebSocket connections.
// SetMaxConn 设置最大 WebSocket 连接数量。
func (s *WebSocketServer) SetMaxConn(max int) {
	s.channelLimit.SetMax(max)
}

// Connections returns the current connection count.
// Connections 返回当前连接数。
func (s *WebSocketServer) Connections() int {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return len(s.mapConn)
}

// CloseConnection closes the connection identified by connId.
// CloseConnection 关闭指定连接。
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

// FindConnection looks up a connection by connId.
// FindConnection 查找连接。
func (s *WebSocketServer) FindConnection(connId string) (conn netx.IServerConn, ok bool) {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	conn, ok = s.mapConn[connId]
	return
}

// SendPackTo sends a framed pack to the given connection ids.
// SendPackTo 向指定连接发送二进制消息包。
func (s *WebSocketServer) SendPackTo(pack []byte, connId ...string) error {
	bytes := WsDataBlockHandler.DataToBlock(pack)
	return s.SendBytesTo(bytes, connId...)
}

// SendBytesTo sends raw bytes to the given connection ids.
// SendBytesTo 向指定连接发送二进制数据。
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
