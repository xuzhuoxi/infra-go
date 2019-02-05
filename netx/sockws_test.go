package netx

import (
	"github.com/xuzhuoxi/util-go/logx"
	"golang.org/x/net/websocket"
	"net/http"
	"testing"
	"time"
)

func TestWSServer(t *testing.T) {
	server := NewWebSocketServer(5)
	var msgHandler = func(msgData []byte, sender interface{}) {
		senderAddress := sender.(string)
		logx.Traceln("TestWSServer.msgHandler[Sender:"+senderAddress+"]msgData:", msgData, "dataLen:", len(msgData), "]")
		rs := []byte{byte(len(msgData))}
		rs = append(rs, msgData...)
		server.SendDataTo(rs, senderAddress)
	}
	server.SetMessageHandler(msgHandler)
	go server.StartServer(SockParams{LocalAddress: "127.0.0.1:9999", WSPattern: "/"})

	client := NewWebSocketClient()
	client.OpenClient(SockParams{RemoteAddress: "ws://127.0.0.1:9999", WSPattern: "/", WSOrigin: "http://127.0.0.1:9999/"})
	go client.StartReceiving()
	client.SendDataTo([]byte{3, 1, 3, 4})
	client.SendDataTo([]byte{3, 2, 0, 0})
	client.SendDataTo([]byte{3, 3, 2, 1})
	client.SendDataTo([]byte{7, 4, 2, 1})
	client.SendDataTo([]byte{3, 3, 2, 1})
	client.SendDataTo([]byte{3, 5, 2, 1})
	client.SendDataTo([]byte{3, 6, 2, 1})
	client.SendDataTo([]byte{3, 7, 1, 1})

	time.Sleep(1 * time.Second)
	client.CloseClient()

	time.Sleep(1 * time.Second)
	server.StopServer()
}

func TestWSServer2(t *testing.T) {
	go server()
	client()
	time.Sleep(100 * time.Second)
}

func server() {
	//http.Handle("/echo", websocket.Handler(svrConnHandler))
	//http.ListenAndServe(":6666", nil)

	httpMux := http.NewServeMux()
	httpMux.Handle("/echo", websocket.Handler(svrConnHandler))
	httpServer := &http.Server{Addr: ":6666", Handler: httpMux}
	httpServer.ListenAndServe()
}

func svrConnHandler(conn *websocket.Conn) {
	logx.Traceln("Server New Conn")
}

var origin333 = "http://127.0.0.1:6666/"
var url333 = "ws://127.0.0.1:6666/echo"

func client() {
	conn, _ := websocket.Dial(url333, "", origin333)
	logx.Traceln(conn)
	go clientConnHandler(conn)
}

func clientConnHandler(conn *websocket.Conn) {
	logx.Traceln("Client New Conn")
}