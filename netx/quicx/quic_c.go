package quicx

import (
	"context"
	"crypto/tls"

	"github.com/quic-go/quic-go"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
)

// NewQUICClient creates a QUIC client.
// NewQUICClient 创建 QUIC 客户端。
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

// IQuicClient is a QUIC-specific ISockClient.
// IQuicClient 是 QUIC 专用的 ISockClient。
type IQuicClient interface {
	netx.ISockClient
}

// QUICClient implements IQuicClient over QUIC connections and streams.
// QUICClient 是基于 QUIC 连接与流的 IQuicClient 实现。
type QUICClient struct {
	netx.SockClientBase
	stream *quic.Stream
}

type quicConnAdapter struct {
	*quic.Conn
}

func (a *quicConnAdapter) Close() error {
	return a.Conn.CloseWithError(0, "client closed")
}

// OpenClient dials the remote QUIC endpoint and opens a stream.
// OpenClient 拨号连接远程 QUIC 端点并打开流。
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
	session, err1 := quic.DialAddr(context.Background(), params.RemoteAddress, &tls.Config{InsecureSkipVerify: true}, nil)
	if nil != err1 {
		c.Logger.Warnln(funcName, err1)
		return err1
	}
	c.Conn = &quicConnAdapter{Conn: session}
	stream, err2 := session.OpenStreamSync(context.Background())
	if nil != err2 {
		c.Logger.Warnln(funcName, err2)
		return err2
	}
	c.stream = stream
	connInfo := netx.NewRemoteOnlyConnInfo(session.LocalAddr().String(), session.RemoteAddr().String())
	connProxy := &QUICStreamAdapter{Reader: stream, Writer: stream, remoteAddress: session.RemoteAddr().String()}
	c.PackProxy = netx.NewPackSendReceiver(connInfo, connProxy, connProxy, c.PackHandler, QuicDataBlockHandler, c.Logger, false)
	c.Opening = true
	c.Logger.Infoln(funcName, "()")
	return nil
}

// CloseClient closes the QUIC stream and connection.
// CloseClient 关闭 QUIC 流与连接。
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
