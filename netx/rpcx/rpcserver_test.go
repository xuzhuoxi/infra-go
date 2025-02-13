package rpcx

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
	fmt.Println("Add:", reply.C, args.A, args.B)
	return nil
}

func TestRPC(t *testing.T) {
	s := NewRPCServer()
	s.Register(new(Arith))
	go s.StartServer(":9999")
	clientRPCClient()
}

func clientRPCClient() {
	client := NewRPCClient()
	err := client.Dial(":9999")
	if nil != err {
		fmt.Println("Dial Fail:", err)
		return
	}

	fmt.Println("——————————————————————————————————————————————")
	args1 := &Args{A: 1, B: 1}
	var reply1 = &Reply{}
	fmt.Println(fmt.Sprintf("Before call-1: C=%d, A=%d, B=%d", reply1.C, args1.A, args1.B))
	err1 := client.Call("Arith.Add", args1, reply1)
	if nil != err1 {
		fmt.Println("Call Arith.Add Fail:", err1)
	} else {
		fmt.Println(fmt.Sprintf("After call-1: C=%d, A=%d, B=%d", reply1.C, args1.A, args1.B))
	}

	fmt.Println("——————————————————————————————————————————————")

	args2 := Args{A: 2, B: 2}
	var reply2 = Reply{}
	fmt.Println(fmt.Sprintf("Before call-2: C=%d, A=%d, B=%d", reply2.C, args2.A, args2.B))
	err2 := client.Call("Arith.Add", args2, &reply2)
	if nil != err2 {
		fmt.Println("Call Arith.Add Fail:", err2)
	} else {
		fmt.Println(fmt.Sprintf("After call-2: C=%d, A=%d, B=%d", reply2.C, args2.A, args2.B))
	}
	fmt.Println("——————————————————————————————————————————————")
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
