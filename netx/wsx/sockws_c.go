package wsx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"golang.org/x/net/websocket"
)

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

type IWebSocketClient interface {
	netx.ISockClient
}

type WebSocketClient struct {
	netx.SockClientBase
}

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
	connProxy := &WSConnAdapter{Reader: conn, Writer: conn, remoteAddrString: params.RemoteAddress}
	c.PackProxy = netx.NewPackSendReceiver(connProxy, connProxy, c.PackHandler, WsDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

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
