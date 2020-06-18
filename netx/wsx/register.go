package wsx

import "github.com/xuzhuoxi/infra-go/netx"

func init() {
	netx.RegisterNetwork(netx.WSNetwork, newWebSocketServer, newWebSocketClient, newWebSocketClient)
}
