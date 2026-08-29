//go:build !skiptest

// Package tcpx
// Create on 2025/5/14
// @author xuzhuoxi
package tcpx

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/encodingx"
	"github.com/xuzhuoxi/infra-go/netx"
)

func TestTCPResponseServer(t *testing.T) {
	// 初始化默认配置
	bytex.DefaultOrder = binary.LittleEndian
	bytex.DefaultDataBlockHandler = bytex.NewDataBlockHandler(bytex.DefaultOrder, bytex.DefaultDataToBlockHandler, bytex.DefaultBlockToDataHandler)
	encodingx.DefaultOrder, encodingx.DefaultDataBlockHandler = bytex.DefaultOrder, bytex.DefaultDataBlockHandler
	TcpDataBlockHandler = bytex.DefaultDataBlockHandler
	netx.ReceiverBuffLen = 8192

	server := NewTCPServer()
	server.SetMaxConn(200)
	var packHandler = func(data []byte, connInfo netx.IConnInfo, other interface{}) bool {
		//logx.Traceln(fmt.Sprintf("TestTCPServer.packHandler{Conn=%s,Data=%s,Other=%s]}", connInfo, fmt.Sprint(data), fmt.Sprint(other)))
		rs := fmt.Sprintf("TestTCPResponseServer:%s", string(data))
		server.SendPackTo([]byte(rs), connInfo.GetConnId())
		fmt.Println("ReceiveData:", string(data))
		return true
	}
	server.GetPackHandlerContainer().AppendPackHandler(packHandler)
	go server.StartServer(netx.SockParams{LocalAddress: "127.0.0.1:9999"})

	//client := NewTCPClient()
	//client.OpenClient(netx.SockParams{RemoteAddress: "127.0.0.1:9999"})
	//client.SendPackTo([]byte("Hello World!"))
	//time.Sleep(1 * time.Second)
	//client.CloseClient()

	ch := make(chan interface{})
	<-ch
}
