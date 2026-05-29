package rpcx

import (
	"errors"
	"net/rpc"
)

// NewRPCClient creates an RPC client using TCP network.
// NewRPCClient 创建使用 TCP 网络的 RPC 客户端。
func NewRPCClient() IRPCClient {
	return NewRPCClientWithNetwork(TCP.String())
}

// NewRPCClientWithNetwork creates an RPC client for the given network.
// NewRPCClientWithNetwork 为指定网络创建 RPC 客户端。
func NewRPCClientWithNetwork(network string) IRPCClient {
	return &RPCClient{Network: network}
}

// IRPCClient defines remote RPC client operations.
// IRPCClient 定义远程 RPC 客户端操作。
type IRPCClient interface {
	// Dial connects to the remote RPC server.
	// Dial 连接远程 RPC 服务。
	Dial(address string) error
	// IsConnected reports whether the client is connected.
	// IsConnected RPC 是否处于连接状态。
	IsConnected() bool
	// Call invokes a remote RPC method.
	// Call 调用 RPC 服务。
	//   args: remote method arguments
	//   args: 远程 RPC 服务方法参数
	//   reply: pointer to receive the return value
	//   reply: 远程 RPC 服务方法返回值，必须为指针引用
	// Returns nil on success, non-nil on failure.
	// 返回值：nil 表示成功，非 nil 表示失败。
	Call(serviceMethod string, args interface{}, reply interface{}) error
	// Close closes the RPC connection.
	// Close 关闭远程 RPC 连接。
	Close()
}

// RPCClient implements IRPCClient using net/rpc.
// RPCClient 基于 net/rpc 实现 IRPCClient。
type RPCClient struct {
	Network   string
	client    *rpc.Client
	connected bool
}

// Dial connects to the remote RPC server at address.
// Dial 连接远程 RPC 服务。
func (c *RPCClient) Dial(address string) error {
	client, err := rpc.Dial(c.Network, address)
	if nil != err {
		return err
	}
	c.client = client
	c.connected = true
	return nil
}

// IsConnected reports whether the client is connected.
// IsConnected RPC 是否处于连接状态。
func (c *RPCClient) IsConnected() bool {
	return nil != c.client && c.connected
}

// Call invokes a remote RPC method.
// Call 调用 RPC 服务。
func (c *RPCClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	if c.IsConnected() {
		return c.client.Call(serviceMethod, args, reply)
	}
	return errors.New("Client does not connect. ")
}

// Close closes the RPC connection.
// Close 关闭远程 RPC 连接。
func (c *RPCClient) Close() {
	c.connected = false
	if nil == c.client {
		return
	}
	c.client.Close()
}
