package netx

import (
	"github.com/xuzhuoxi/infra-go/logx"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	server := NewTCPServer(5)
	var packHandler = func(msgData []byte, sender interface{}) bool {
		senderAddress := sender.(string)
		logx.Traceln("TestTCPServer.msgHandler[Sender:"+senderAddress+"]msgData:", msgData, "dataLen:", len(msgData), "]")
		rs := []byte{byte(len(msgData))}
		rs = append(rs, msgData...)
		server.SendPackTo(rs, senderAddress)
		return true
	}
	server.GetPackHandler().SetPackHandlers([]FuncPackHandler{packHandler})
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
