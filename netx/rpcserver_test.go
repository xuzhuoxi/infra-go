package netx

import (
	"fmt"
	"testing"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Add(args Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func TestRPC(t *testing.T) {
	s := NewRPCServer()
	s.Register(new(Arith))
	go s.StartServer("127.0.0.1:9999")
	clientRPCClient()
}

func clientRPCClient() {
	client := NewRPCClient()
	client.Dial("127.0.0.1:9999")

	args := &Args{A: 1, B: 1}
	//var reply = &Reply{}
	var reply = new(Reply)
	fmt.Println(reply.C)
	client.Call("Arith.Add", args, reply)
	fmt.Println(reply.C)

	args2 := &Args{A: 2, B: 2}
	//var reply = &Reply{}
	var reply2 = new(Reply)
	fmt.Println(reply2.C)
	client.Call("Arith.Add", args2, reply2)
	fmt.Println(reply2.C)
}

//func clientRPC() {
//	client, _ := rpc.Dial(NetworkTCP, "127.0.0.1:9999")
//	args := &Args{A: 1, B: 1}
//	//var reply = &Reply{}
//	var reply = new(Reply)
//	fmt.Println(reply.C)
//	client.Call("Arith.Add", args, reply)
//	fmt.Println(reply.C)
//}
