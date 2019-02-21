package netx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"golang.org/x/net/websocket"
)

func NewWebSocketClient() IWebSocketClient {
	client := &WebSocketClient{}
	client.Name = "WSClient"
	client.Network = WSNetwork
	client.Logger = logx.DefaultLogger()
	client.PackHandler = DefaultPackHandler
	return client
}

type IWebSocketClient interface {
	ISockClient
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
	connProxy := &WSConnAdapter{Reader: conn, Writer: conn, RemoteAddrString: params.RemoteAddress}
	c.PackProxy = NewPackSendReceiver(connProxy, connProxy, c.PackHandler, WsDataBlockHandler, c.Logger, false)
	c.opening = true
	c.Logger.Infoln(funcName + "()")
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
	c.Logger.Infoln(funcName + "()")
	return nil
}
