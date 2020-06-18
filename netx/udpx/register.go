package udpx

import "github.com/xuzhuoxi/infra-go/netx"

func init() {
	netx.RegisterNetwork(netx.UDPNetwork, newUDPServer, newUDPDialClient, newUDPListenClient)
	netx.RegisterNetwork(netx.UDPNetwork4, newUDP4Server, newUDP4DialClient, newUDP4ListenClient)
	netx.RegisterNetwork(netx.UDPNetwork6, newUDP6Server, newUDP6DialClient, newUDP6ListenClient)
}
