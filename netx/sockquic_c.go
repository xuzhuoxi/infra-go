package netx

import (
	"crypto/tls"
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
)

func NewQUICClient() IQuicClient {
	client := &QUICClient{SockClientBase: SockClientBase{Name: "QUICClient", Network: QuicNetwork}}
	return client
}

type QUICClient struct {
	SockClientBase
}

func (c *QUICClient) OpenClient(params SockParams) error {
	funcName := "QUICClient.OpenClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if c.opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		c.Network = params.Network
	}
	session, err := quic.DialAddr(params.RemoteAddress, &tls.Config{InsecureSkipVerify: true}, nil)
	if nil != err {
		return err
	}
	//stream, err := session.OpenStreamSync()
	c.conn = session
	connProxy := &QUICSessionReadWriter{Session: session}
	c.setMessageProxy(NewMessageSendReceiver(connProxy, connProxy, UdpDialRW, c.Network))
	c.opening = true
	logx.Infoln(funcName + "()")
	return nil
}

func (c *QUICClient) CloseClient() error {
	funcName := "QUICClient.CloseClient"
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if !c.opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	c.opening = false
	if nil != c.conn {
		c.conn.Close()
		c.conn = nil
	}
	logx.Infoln(funcName + "()")
	return nil
}

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
