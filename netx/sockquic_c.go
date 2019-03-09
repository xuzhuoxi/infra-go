package netx

import (
	"crypto/tls"
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
)

func NewQUICClient() IQuicClient {
	client := &QUICClient{}
	client.Name = "QUICClient"
	client.Network = QuicNetwork
	client.Logger = logx.DefaultLogger()
	client.PackHandler = NewIPackHandler(nil)
	return client
}

type IQuicClient interface {
	ISockClient
}

type QUICClient struct {
	SockClientBase
	stream quic.Stream
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
		c.Logger.Warnln(funcName, err)
		return err
	}
	c.conn = session
	stream, err := session.OpenStreamSync()
	if nil != err {
		c.Logger.Warnln(funcName, err)
		return err
	}
	c.stream = stream
	connProxy := &QUICStreamAdapter{Reader: stream, Writer: stream, RemoteAddr: session.RemoteAddr()}
	c.PackProxy = NewPackSendReceiver(connProxy, connProxy, c.PackHandler, QuicDataBlockHandler, c.Logger, false)
	c.opening = true
	c.Logger.Infoln(funcName + "()")
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
	if nil != c.stream {
		c.stream.Close()
		c.stream = nil
	}
	if nil != c.conn {
		c.conn.Close()
		c.conn = nil
	}
	c.Logger.Infoln(funcName + "()")
	return nil
}
