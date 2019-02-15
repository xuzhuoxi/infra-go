package netx

import (
	"github.com/xuzhuoxi/infra-go/logx"
	"testing"
	"time"
)

func TestQUICServer(t *testing.T) {
	server := NewQuicServer()
	c := 0
	var packHandler = func(msgData []byte, sender interface{}) {
		senderAddress := sender.(string)
		logx.Traceln("TestQUICServer.msgHandler[Sender:"+senderAddress+"]msgData:", msgData, "dataLen:", len(msgData), "]", c)
		c++
		rs := []byte{byte(len(msgData))}
		rs = append(rs, msgData...)
		server.SendPackTo(rs, senderAddress)
	}
	server.SetPackHandler(packHandler)
	go server.StartServer(SockParams{LocalAddress: "127.0.0.1:9999"})
	time.Sleep(1 * time.Second)
	client := NewQUICClient()
	client.OpenClient(SockParams{RemoteAddress: "127.0.0.1:9999"})
	go client.StartReceiving()
	b := true
	go func() {
		for b {
			client.SendPackTo([]byte{3, 1, 3, 4})
			client.SendPackTo([]byte{3, 2, 0, 0})
			client.SendPackTo([]byte{3, 3, 2, 1})
			client.SendPackTo([]byte{7, 4, 2, 1})
			client.SendPackTo([]byte{3, 3, 2, 1})
			client.SendPackTo([]byte{3, 5, 2, 1})
			client.SendPackTo([]byte{3, 6, 2, 1})
			client.SendPackTo([]byte{3, 7, 1, 1})
			//time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(10 * time.Second)
	b = false
	client.CloseClient()
	server.StopServer()
}

//Server
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

//Client
//const addr = "localhost:9999"
//const message = "ccc"
//
//func main() {
//	session, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	stream, err := session.OpenStreamSync()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for {
//		fmt.Printf("Client: Sending '%s'\n", message)
//		_, err = stream.Write([]byte(message))
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		buf := make([]byte, len(message))
//		_, err = io.ReadFull(stream, buf)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Printf("Client: Got '%s'\n", buf)
//		time.Sleep(2 * time.Second)
//	}
//}
