package quicx

import (
	"crypto/tls"
	"github.com/lucas-clemente/quic-go"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
)

func NewQUICClient() IQuicClient {
	return newQUICClient().(IQuicClient)
}

func newQUICClient() netx.ISockClient {
	client := &QUICClient{}
	client.Name = "QUICClient"
	client.Network = netx.QuicNetwork
	client.Logger = logx.DefaultLogger()
	client.PackHandler = netx.NewIPackHandler(nil)
	return client
}

//----------------------------

type IQuicClient interface {
	netx.ISockClient
}

type QUICClient struct {
	netx.SockClientBase
	stream quic.Stream
}

func (c *QUICClient) OpenClient(params netx.SockParams) error {
	funcName := "[QUICClient.OpenClient]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if c.Opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	if "" != params.Network {
		c.Network = params.Network
	}
	session, err1 := quic.DialAddr(params.RemoteAddress, &tls.Config{InsecureSkipVerify: true}, nil)
	if nil != err1 {
		c.Logger.Warnln(funcName, err1)
		return err1
	}
	c.Conn = session
	stream, err2 := session.OpenStreamSync()
	if nil != err2 {
		c.Logger.Warnln(funcName, err2)
		return err2
	}
	c.stream = stream
	connProxy := &QUICStreamAdapter{Reader: stream, Writer: stream, RemoteAddr: session.RemoteAddr()}
	c.PackProxy = netx.NewPackSendReceiver(connProxy, connProxy, c.PackHandler, QuicDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

func (c *QUICClient) CloseClient() error {
	funcName := "[QUICClient.CloseClient]"
	c.ClientMu.Lock()
	defer c.ClientMu.Unlock()
	if !c.Opening {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	c.Opening = false
	if nil != c.stream {
		c.stream.Close()
		c.stream = nil
	}
	if nil != c.Conn {
		c.Conn.Close()
		c.Conn = nil
	}
	c.Logger.Infoln(funcName, "()")
	return nil
}
