package netx

import (
	"github.com/xuzhuoxi/infra-go/logx"
	"net"
	"net/rpc"
)

const (
	RpcNetworkTCP  = "tcp"
	RpcNetworkHttp = "http"
)

func NewRPCServer() IRPCServer {
	rs := &RPCServer{Network: "tcp", Server: rpc.NewServer()}
	return rs
}

type IRPCServer interface {
	// Register publishes in the server the set of methods of the receiver value that satisfy the following conditions:
	Register(rcvr interface{}) error
	// RegisterName is like Register but uses the provided name for the type instead of the receiver's concrete type.
	RegisterName(name string, rcvr interface{}) error
	// 启动RPC服务,会阻塞
	StartServer(addr string)
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
	l, newServerAddr := s.listenRPC(addr)
	s.GetLogger().Info("[RPCServer] listening on:", newServerAddr)
	s.Listener = l
	s.Server.Accept(l)
}

func (s *RPCServer) StopServer() {
	s.Listener.Close()
}

func (s *RPCServer) listenRPC(address string) (net.Listener, string) {
	l, e := net.Listen(s.Network, address) // any available address
	if e != nil {
		s.GetLogger().Fatalln("\tnetxu.Listen "+s.Network+" "+address+": %v", e)
	}
	return l, l.Addr().String()
}
