package netx

import (
	"sync"

	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/logx"
)

// ISockClientSetter injects an ISockClient implementation.
// ISockClientSetter 注入 ISockClient 实现。
type ISockClientSetter interface {
	SetClient(client ISockClient)
}

// ISockClientGetter returns the bound ISockClient.
// ISockClientGetter 返回已绑定的 ISockClient。
type ISockClientGetter interface {
	GetClient() ISockClient
}

// IClient defines client open/close lifecycle.
// IClient 定义客户端打开与关闭的生命周期。
type IClient interface {
	// IsOpening reports whether the client connection is open.
	// IsOpening 连接是否打开中。
	IsOpening() bool
	// OpenClient opens the client with the given parameters.
	// OpenClient 打开客户端连接。
	OpenClient(params SockParams) error
	// CloseClient closes the client connection.
	// CloseClient 关闭客户端连接。
	CloseClient() error
}

// ISockClient is a socket client with naming, pack receive, and send capabilities.
// ISockClient 是具备命名、收包与发送能力的套接字客户端。
type ISockClient interface {
	ISockName
	IClient
	IPackReceiver
	ISockSender

	LocalAddress() string
}

// ISockEventClient combines event dispatch with ISockClient.
// ISockEventClient 在 ISockClient 基础上支持事件分发。
type ISockEventClient interface {
	eventx.IEventDispatcher
	ISockClient
}

// SockClientBase is a base implementation shared by protocol-specific clients.
// SockClientBase 是各协议客户端共用的基础实现。
type SockClientBase struct {
	Name     string
	Network  SockNetwork
	ClientMu sync.RWMutex
	Opening  bool

	LocalAddr string
	Conn      ISockConn

	PackProxy IPackSendReceiver
	Logger    logx.ILogger

	PackHandler IPackHandlerContainer
}

// SetName sets the client identity name.
// SetName 设置客户端标识名称。
func (o *SockClientBase) SetName(name string) {
	o.Name = name
}

// GetName returns the client identity name.
// GetName 获取客户端标识名称。
func (o *SockClientBase) GetName() string {
	return o.Name
}

// GetPackHandlerContainer returns the pack handler container.
// GetPackHandlerContainer 获取消息包处理器容器。
func (o *SockClientBase) GetPackHandlerContainer() IPackHandlerContainer {
	o.ClientMu.RLock()
	defer o.ClientMu.RUnlock()
	return o.PackHandler
}

// SetPackHandlerContainer sets the pack handler container.
// SetPackHandlerContainer 设置消息包处理器容器。
func (o *SockClientBase) SetPackHandlerContainer(packHandler IPackHandlerContainer) {
	o.ClientMu.Lock()
	defer o.ClientMu.Unlock()
	o.PackHandler = packHandler
	if o.PackProxy != nil {
		o.PackProxy.SetPackHandlerContainer(o.PackHandler)
	}
}

// LocalAddress returns the local address string of the active connection.
// LocalAddress 返回当前连接的本地地址字符串。
func (o *SockClientBase) LocalAddress() string {
	if nil == o.Conn {
		return o.LocalAddr
	}
	return o.Conn.LocalAddr().String()
}

// IsReceiving reports whether pack receiving is active.
// IsReceiving 是否为数据接收中。
func (o *SockClientBase) IsReceiving() bool {
	return o.PackProxy.IsReceiving()
}

// IsOpening reports whether the client is open.
// IsOpening 连接是否打开中。
func (o *SockClientBase) IsOpening() bool {
	o.ClientMu.RLock()
	defer o.ClientMu.RUnlock()
	return o.Opening
}

// GetLogger returns the logger.
// GetLogger 获取日志记录器。
func (o *SockClientBase) GetLogger() logx.ILogger {
	return o.Logger
}

// SetLogger sets the logger.
// SetLogger 设置日志记录器。
func (o *SockClientBase) SetLogger(logger logx.ILogger) {
	o.Logger = logger
}

// SendPackTo sends a framed pack to the remote endpoint.
// SendPackTo 发送二进制消息包。
func (o *SockClientBase) SendPackTo(msg []byte, rAddress ...string) error {
	_, err := o.PackProxy.SendPack(msg, rAddress...)
	return err
}

// SendBytesTo sends raw bytes to the remote endpoint.
// SendBytesTo 发送二进制数据，不会重新打包。
func (o *SockClientBase) SendBytesTo(bytes []byte, rAddress ...string) error {
	_, err := o.PackProxy.SendBytes(bytes, rAddress...)
	return err
}

// StartReceiving starts the pack receive loop.
// StartReceiving 开始接收数据。
func (o *SockClientBase) StartReceiving() error {
	if nil != o.Logger {
		o.Logger.Infoln("[SockClientBase.StartReceiving]", "ServerName="+o.Name)
	}
	err := o.PackProxy.StartReceiving()
	return err
}

// StopReceiving stops the pack receive loop.
// StopReceiving 停止接收数据。
func (o *SockClientBase) StopReceiving() error {
	if nil != o.Logger {
		o.Logger.Infoln("[SockClientBase.StopReceiving]", "ServerName="+o.Name)
	}
	err := o.PackProxy.StopReceiving()
	return err
}
