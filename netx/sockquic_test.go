package netx

import (
	"github.com/xuzhuoxi/util-go/logx"
	"testing"
	"time"
)

func TestQUICServer(t *testing.T) {
	server := NewQuicServer()
	var msgHandler = func(msgData []byte, sender interface{}) {
		senderAddress := sender.(string)
		logx.Traceln("TestQUICServer.msgHandler[Sender:"+senderAddress+"]msgData:", msgData, "dataLen:", len(msgData), "]")
		rs := []byte{byte(len(msgData))}
		rs = append(rs, msgData...)
		server.SendDataTo(rs, senderAddress)
	}
	server.SetMessageHandler(msgHandler)
	go server.StartServer(SockParams{LocalAddress: "127.0.0.1:9999"})

	client := NewQUICClient()
	client.OpenClient(SockParams{RemoteAddress: "127.0.0.1:9999"})
	go client.StartReceiving()
	//b := true
	//go func() {
	//	for b {
	client.SendDataTo([]byte{3, 1, 3, 4})
	client.SendDataTo([]byte{3, 2, 0, 0})
	client.SendDataTo([]byte{3, 3, 2, 1})
	client.SendDataTo([]byte{7, 4, 2, 1})
	client.SendDataTo([]byte{3, 3, 2, 1})
	client.SendDataTo([]byte{3, 5, 2, 1})
	client.SendDataTo([]byte{3, 6, 2, 1})
	client.SendDataTo([]byte{3, 7, 1, 1})
	//	}
	//}()
	time.Sleep(100 * time.Second)
	//b = false
	client.CloseClient()
	server.StopServer()
}

//
//const saddr = "localhost:9999"
//
//func startServer() {
//	listener, err := quic.ListenAddr(saddr, generateTLSConfig(), nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//	for {
//		sess, err := listener.Accept()
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			go dealSession(sess)
//		}
//	}
//}
//
//func dealSession(sess quic.Session) {
//	stream, err := sess.AcceptStream()
//	if err != nil {
//		panic(err)
//	} else {
//		for {
//			_, err = io.Copy(loggingWriter{stream}, stream)
//		}
//	}
//}
//
//type loggingWriter struct{ io.Writer }
//
//func (w loggingWriter) Write(b []byte) (int, error) {
//	fmt.Printf("Server: Got '%s'\n", string(b))
//	return w.Writer.Write(b)
//}
