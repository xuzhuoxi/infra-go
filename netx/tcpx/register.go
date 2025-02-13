package tcpx

import "github.com/xuzhuoxi/infra-go/netx"

func init() {
	netx.RegisterNetwork(netx.TcpNetwork, newTCPServer, newTCPClient, newTCPClient)
	netx.RegisterNetwork(netx.Tcp4Network, newTCP4Server, newTCP4Client, newTCP4Client)
	netx.RegisterNetwork(netx.Tcp6Network, newTCP6Server, newTCP6Client, newTCP6Client)
}
