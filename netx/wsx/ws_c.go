package wsx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"golang.org/x/net/websocket"
)

// NewWebSocketClient creates a WebSocket client.
// NewWebSocketClient 创建 WebSocket 客户端。
func NewWebSocketClient() IWebSocketClient {
	return newWebSocketClient().(IWebSocketClient)
}

func newWebSocketClient() netx.ISockClient {
	client := &WebSocketClient{}
	client.Name = "WSClient"
	client.Network = netx.WSNetwork
	client.Logger = logx.DefaultLogger()
	client.PackHandler = netx.NewIPackHandler(nil)
	return client
}

//---------------------------

// IWebSocketClient is a WebSocket-specific ISockClient.
// IWebSocketClient 是 WebSocket 专用的 ISockClient。
type IWebSocketClient interface {
	netx.ISockClient
}

// WebSocketClient implements IWebSocketClient over WebSocket connections.
// WebSocketClient 是基于 WebSocket 连接的 IWebSocketClient 实现。
type WebSocketClient struct {
	netx.SockClientBase
}

// OpenClient dials the remote WebSocket endpoint.
// OpenClient 拨号连接远程 WebSocket 端点。
func (c *WebSocketClient) OpenClient(params netx.SockParams) error {
	funcName := "[WebSocketClient.OpenClient]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if "" != params.Network {
		c.Network = params.Network
	}
	conn, err := websocket.Dial(params.RemoteAddress+params.WSPattern, params.WSProtocol, params.WSOrigin)
	if nil != err {
		return err
	}
	c.Conn = conn //LocalAddr=Origin
	connProxy := &WSConnAdapter{Reader: conn, Writer: conn, remoteAddress: params.RemoteAddress}
	connInfo := netx.NewRemoteOnlyConnInfo(conn.LocalAddr().String(), conn.RemoteAddr().String())
	c.PackProxy = netx.NewPackSendReceiver(connInfo, connProxy, connProxy, c.PackHandler, WsDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

// CloseClient closes the WebSocket connection.
// CloseClient 关闭 WebSocket 连接。
func (c *WebSocketClient) CloseClient() error {
	funcName := "[WebSocketClient.CloseClient]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if !c.Opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	c.Opening = false
	if nil != c.Conn {
		c.Conn.Close()
		c.Conn = nil
	}
	c.Logger.Infoln(funcName, "()")
	return nil
}
