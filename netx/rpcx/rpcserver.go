package rpcx

import (
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"net"
	"net/rpc"
)

const (
	TCP  = netx.TcpNetwork
	TCP4 = netx.Tcp4Network
	TCP6 = netx.Tcp6Network
	UDP  = netx.UDPNetwork
)

func NewRPCServer() IRPCServer {
	return NewRPCServerWithNetwork(TCP.String())
}
func NewRPCServerWithNetwork(network string) IRPCServer {
	rs := &RPCServer{Network: network, Server: rpc.NewServer()}
	return rs
}

type IRPCServer interface {
	// Register
	// publishes in the server the set of methods of the receiver value that satisfy the following conditions:
	// 参数格式：
	//   rcvr: new(Arith), Arith为结构体
	Register(rcvr interface{}) error
	// RegisterName
	// is like Register but uses the provided name for the type instead of the receiver's concrete type.
	// 参数格式：
	//   rcvr: new(Arith), Arith为结构体
	RegisterName(name string, rcvr interface{}) error
	// StartServer
	// 启动RPC服务,会阻塞
	StartServer(addr string)
	// StopServer
	// 停止RPC服务
	StopServer()
}

type RPCServer struct {
	logx.LoggerSupport

	Network  string
	Server   *rpc.Server
	Listener net.Listener
}

func (s *RPCServer) Register(rcvr interface{}) error {
	return s.Server.Register(rcvr)
}

func (s *RPCServer) RegisterName(name string, rcvr interface{}) error {
	return s.Server.RegisterName(name, rcvr)
}

func (s *RPCServer) StartServer(addr string) {
	if nil == s.Server {
		return
	}
	l, e := net.Listen(s.Network, addr) // any available address
	if e != nil {
		s.GetLogger().Fatalln("\tnetx.Listen "+s.Network+" "+addr+": %v", e)
		return
	}
	s.GetLogger().Info("[RPCServer] listening on:", l.Addr().String())
	s.Listener = l
	s.Server.Accept(l)
}

func (s *RPCServer) StopServer() {
	s.Listener.Close()
}
