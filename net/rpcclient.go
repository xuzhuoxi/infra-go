package net

import (
	"net/rpc"
)

type IRPCClient interface {
	Dial(network, address string) error
	IsConnected() bool
	Call(serviceMethod string, args interface{}, reply interface{})
	Close()
}

type RPCClient struct {
	client    *rpc.Client
	connected bool
}

func (c *RPCClient) Dial(network, address string) error {
	client, err := rpc.Dial(network, address)
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

func (c *RPCClient) Call(serviceMethod string, args interface{}, reply interface{}) {
	if c.IsConnected() {
		c.client.Call(serviceMethod, args, reply)
	}
}

func (c *RPCClient) Close() {
	c.connected = false
	if nil == c.client {
		return
	}
	c.client.Close()
}

func NewRPCClient() IRPCClient {
	return &RPCClient{}
}
