package rpcx

import (
	"net"
	"net/rpc"

	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
)

// Network constants for RPC transport aligned with netx.SockNetwork.
// 与 netx.SockNetwork 对齐的 RPC 传输网络常量。
const (
	TCP  = netx.TcpNetwork
	TCP4 = netx.Tcp4Network
	TCP6 = netx.Tcp6Network
	UDP  = netx.UDPNetwork
)

// NewRPCServer creates an RPC server using TCP network.
// NewRPCServer 创建使用 TCP 网络的 RPC 服务端。
func NewRPCServer() IRPCServer {
	return NewRPCServerWithNetwork(TCP.String())
}

// NewRPCServerWithNetwork creates an RPC server for the given network.
// NewRPCServerWithNetwork 为指定网络创建 RPC 服务端。
func NewRPCServerWithNetwork(network string) IRPCServer {
	rs := &RPCServer{Network: network, Server: rpc.NewServer()}
	return rs
}

// IRPCServer defines RPC service registration and server lifecycle.
// IRPCServer 定义 RPC 服务注册与服务端生命周期。
type IRPCServer interface {
	// Register publishes exported methods of rcvr on the server.
	// Register 在服务端发布接收者的方法集。
	//   rcvr: e.g. new(Arith) where Arith is a struct type
	//   rcvr: 例如 new(Arith)，Arith 为结构体
	Register(rcvr interface{}) error
	// RegisterName is like Register but uses the provided service name.
	// RegisterName 与 Register 类似，但使用提供的名称代替接收者的具体类型。
	//   rcvr: e.g. new(Arith) where Arith is a struct type
	//   rcvr: 例如 new(Arith)，Arith 为结构体
	RegisterName(name string, rcvr interface{}) error
	// StartServer starts the RPC server on addr and blocks.
	// StartServer 启动 RPC 服务，会阻塞。
	StartServer(addr string)
	// StopServer stops the RPC server.
	// StopServer 停止 RPC 服务。
	StopServer()
}

// RPCServer implements IRPCServer using net/rpc.
// RPCServer 基于 net/rpc 实现 IRPCServer。
type RPCServer struct {
	logx.LoggerSupport

	Network  string
	Server   *rpc.Server
	Listener net.Listener
}

// Register publishes exported methods of rcvr on the server.
// Register 在服务端发布接收者的方法集。
func (s *RPCServer) Register(rcvr interface{}) error {
	return s.Server.Register(rcvr)
}

// RegisterName publishes rcvr under the given service name.
// RegisterName 以指定服务名发布接收者。
func (s *RPCServer) RegisterName(name string, rcvr interface{}) error {
	return s.Server.RegisterName(name, rcvr)
}

// StartServer listens on addr and accepts RPC connections; blocks until stopped.
// StartServer 在 addr 上监听并接受 RPC 连接，会阻塞。
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

// StopServer closes the listener and stops accepting connections.
// StopServer 关闭监听器并停止接受连接。
func (s *RPCServer) StopServer() {
	s.Listener.Close()
}
