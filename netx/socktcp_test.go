package netx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/logx"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	server := NewTCPServer()
	server.SetMax(200)
	var packHandler = func(data []byte, senderAddress string, other interface{}) bool {
		logx.Traceln(fmt.Sprintf("TestTCPServer.packHandler{Sender=%s,Data=%s,Other=%s]}", senderAddress, fmt.Sprint(data), fmt.Sprint(other)))
		rs := []byte{byte(len(data))}
		rs = append(rs, data...)
		server.SendPackTo(rs, senderAddress)
		return true
	}
	server.GetPackHandler().SetPackHandlers([]FuncPackHandler{packHandler})
	server.OnceEventListener(ServerEventStart, func(evd *eventx.EventData) {
		fmt.Println(1111111111111111)
	})
	server.OnceEventListener(ServerEventStart, func(evd *eventx.EventData) {
		fmt.Println(3333333333333333)
	})
	server.OnceEventListener(ServerEventStop, func(evd *eventx.EventData) {
		fmt.Println(2222222222222)
	})
	go server.StartServer(SockParams{LocalAddress: "127.0.0.1:9999"})

	client := NewTCPClient()
	client.OpenClient(SockParams{RemoteAddress: "127.0.0.1:9999"})
	go client.StartReceiving()
	client.SendPackTo([]byte{3, 1, 3, 4})
	client.SendPackTo([]byte{3, 2, 0, 0})
	client.SendPackTo([]byte{3, 3, 2, 1})
	client.SendPackTo([]byte{7, 4, 2, 1})
	client.SendPackTo([]byte{3, 3, 2, 1})
	client.SendPackTo([]byte{3, 5, 2, 1})
	client.SendPackTo([]byte{3, 6, 2, 1})
	client.SendPackTo([]byte{3, 7, 1, 1})
	time.Sleep(1 * time.Second)
	client.CloseClient()
	server.StopServer()
}
