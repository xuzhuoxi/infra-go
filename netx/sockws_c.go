package netx

import (
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
	"golang.org/x/net/websocket"
)

func NewWebSocketClient() IWebSocketClient {
	client := &WebSocketClient{SockClientBase: SockClientBase{Name: "WebSocketClient", Network: WSNetwork}}
	return client
}

type WebSocketClient struct {
	SockClientBase
}

func (c *WebSocketClient) OpenClient(params SockParams) error {
	funcName := "WebSocketClient.OpenClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if "" != params.Network {
		c.Network = params.Network
	}
	conn, err := websocket.Dial(params.RemoteAddress+params.WSPattern, params.WSProtocol, params.WSOrigin)
	if nil != err {
		return err
	}
	c.conn = conn //LocalAddr=Origin
	connProxy := &WSConnReadWriter{Reader: conn, Writer: conn, RemoteAddrString: params.RemoteAddress}
	c.messageProxy = NewMessageSendReceiver(connProxy, connProxy, false)
	c.opening = true
	logx.Infoln(funcName + "()")
	return nil
}

func (c *WebSocketClient) CloseClient() error {
	funcName := "WebSocketClient.CloseClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if !c.opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	c.opening = false
	if nil != c.conn {
		c.conn.Close()
		c.conn = nil
	}
	logx.Infoln(funcName + "()")
	return nil
}
