package tcpx

import "github.com/xuzhuoxi/infra-go/netx"

func init() {
	netx.RegisterNetwork(netx.TcpNetwork, newTCPServer, newTCPClient, newTCPClient)
	netx.RegisterNetwork(netx.TcpNetwork4, newTCP4Server, newTCP4Client, newTCP4Client)
	netx.RegisterNetwork(netx.TcpNetwork6, newTCP6Server, newTCP6Client, newTCP6Client)
}
