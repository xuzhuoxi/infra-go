package udpx

import "github.com/xuzhuoxi/infra-go/netx"

func init() {
	netx.RegisterNetwork(netx.UDPNetwork, newUDPServer, newUDPDialClient, newUDPListenClient)
	netx.RegisterNetwork(netx.UDP4Network, newUDP4Server, newUDP4DialClient, newUDP4ListenClient)
	netx.RegisterNetwork(netx.UDP6Network, newUDP6Server, newUDP6DialClient, newUDP6ListenClient)
}
