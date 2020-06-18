package quicx

import "github.com/xuzhuoxi/infra-go/netx"

func init() {
	netx.RegisterNetwork(netx.QuicNetwork, newQuicServer, newQUICClient, newQUICClient)
}
