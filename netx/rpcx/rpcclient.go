package rpcx

import (
	"errors"
	"net/rpc"
)

func NewRPCClient() IRPCClient {
	return NewRPCClientWithNetwork(TCP.String())
}

func NewRPCClientWithNetwork(network string) IRPCClient {
	return &RPCClient{Network: network}
}

type IRPCClient interface {
	// Dial
	// 连接远程RPC服务
	Dial(address string) error
	// IsConnected
	// RPC是否处理连接状态
	IsConnected() bool
	// Call
	// 调用RPC服务
	// args: 远程RPC服务方法参数
	// reply: 远程RPC服务方法返回值, 必须为指针引用
	// 返回值: 错误信息
	// 错误信息: nil: 成功; 非nil: 失败
	Call(serviceMethod string, args interface{}, reply interface{}) error
	// Close
	// 关闭远程RPC服务
	Close()
}

type RPCClient struct {
	Network   string
	client    *rpc.Client
	connected bool
}

func (c *RPCClient) Dial(address string) error {
	client, err := rpc.Dial(c.Network, address)
	if nil != err {
		return err
	}
	c.client = client
	c.connected = true
	return nil
}

func (c *RPCClient) IsConnected() bool {
	return nil != c.client && c.connected
}

func (c *RPCClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	if c.IsConnected() {
		return c.client.Call(serviceMethod, args, reply)
	}
	return errors.New("Client does not connect. ")
}

func (c *RPCClient) Close() {
	c.connected = false
	if nil == c.client {
		return
	}
	c.client.Close()
}
