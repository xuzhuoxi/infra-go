package netx

import (
	"github.com/xuzhuoxi/go-util/logx"
	"log"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	server := NewTCPServer(5)
	var msgHandler = func(msgData []byte, sender interface{}) {
		senderAddress := sender.(string)
		log.Println("TestTCPServer.msgHandler[Sender:"+senderAddress+"]msgData:", msgData, "dataLen:", len(msgData), "]")
		rs := []byte{byte(len(msgData))}
		rs = append(rs, msgData...)
		server.SendDataTo(rs, senderAddress)
	}
	server.SetMessageHandler(msgHandler)
	go server.StartServer(SockParams{LocalAddress: "127.0.0.1:9999"})

	client := NewTCPClient()
	client.OpenClient(SockParams{RemoteAddress: "127.0.0.1:9999"})
	go client.StartReceiving()
	client.SendDataTo([]byte{3, 1, 3, 4})
	client.SendDataTo([]byte{3, 2, 0, 0})
	client.SendDataTo([]byte{3, 3, 2, 1})
	client.SendDataTo([]byte{7, 4, 2, 1})
	client.SendDataTo([]byte{3, 3, 2, 1})
	client.SendDataTo([]byte{3, 5, 2, 1})
	client.SendDataTo([]byte{3, 6, 2, 1})
	client.SendDataTo([]byte{3, 7, 1, 1})
	time.Sleep(50 * time.Millisecond)
	//server.SendDataTo([]byte{3, 1, 2, 3}, client.LocalAddress())
	//server.SendDataTo([]byte{3, 1, 2, 3}, client.LocalAddress())
	//server.SendDataTo([]byte{3, 1, 2, 3}, client.LocalAddress())
	//server.SendDataTo([]byte{3, 1, 2, 3}, client.LocalAddress())
	//server.SendDataTo([]byte{3, 1, 2, 3}, client.LocalAddress())
	time.Sleep(1 * time.Second)
	client.CloseClient()
	server.StopServer()
	logx.Traceln("不是吧")
}
