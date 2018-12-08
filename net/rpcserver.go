package net

import (
	"log"
	"net"
	"net/rpc"
)

const (
	NetworkTCP = "tcp"
)

type IRPCServer interface {
	Register(rcvr interface{}) error
	RegisterName(name string, rcvr interface{}) error
	//会阻塞
	StartServer(addr string)
	StopServer()
}

type RPCServer struct {
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
	l, newServerAddr := listenRPC(NetworkTCP, addr)
	log.Println("\tRPC server listening on:", newServerAddr)
	s.Listener = l
	s.Server.Accept(l)
}

func (s *RPCServer) StopServer() {
	s.Listener.Close()
}

func NewRPCServer() IRPCServer {
	rs := &RPCServer{Server: rpc.NewServer()}
	return rs
}

func listenRPC(network string, address string) (net.Listener, string) {
	l, e := net.Listen(network, address) // any available address
	if e != nil {
		log.Fatalln("\tnet.Listen "+network+" "+address+": %v", e)
	}
	return l, l.Addr().String()
}
