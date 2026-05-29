package netx

import (
	"sync"

	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/logx"
)

// ISockServerSetter injects an ISockServer implementation.
// ISockServerSetter 注入 ISockServer 实现。
type ISockServerSetter interface {
	SetSockServer(server ISockServer)
}

// ISockServerGetter returns the bound ISockServer.
// ISockServerGetter 返回已绑定的 ISockServer。
type ISockServerGetter interface {
	GetSockServer() ISockServer
}

// IServer defines server start/stop lifecycle.
// IServer 定义服务端启动与停止的生命周期。
type IServer interface {
	// IsRunning reports whether the server is running.
	// IsRunning 服务是否启动中。
	IsRunning() bool
	// StartServer starts the server and blocks until it stops.
	// StartServer 启动服务，会阻塞。
	StartServer(params SockParams) error
	// StopServer stops the server and blocks until shutdown completes.
	// StopServer 停止服务，会阻塞。
	StopServer() error
}

// ISockConnSetSetter injects an IServerConnSet implementation.
// ISockConnSetSetter 注入 IServerConnSet 实现。
type ISockConnSetSetter interface {
	SetSockConnSet(set IServerConnSet)
}

// ISockConnSetGetter returns the bound IServerConnSet.
// ISockConnSetGetter 返回已绑定的 IServerConnSet。
type ISockConnSetGetter interface {
	GetSockConnSet() IServerConnSet
}

// IServerConnSet manages server-side connections.
// IServerConnSet 管理服务端连接集合。
type IServerConnSet interface {
	// SetMaxConn sets the maximum number of concurrent connections.
	// SetMaxConn 设置最大连接数量。
	SetMaxConn(max int)
	// Connections returns the current connection count.
	// Connections 返回当前连接数。
	Connections() int
	// CloseConnection closes the connection identified by connId.
	// CloseConnection 关闭指定连接。
	CloseConnection(connId string) (err error, ok bool)
	// FindConnection looks up a connection by connId.
	// FindConnection 查找连接。
	FindConnection(connInd string) (conn IServerConn, ok bool)
}

// IServerConn represents a single accepted server connection.
// IServerConn 表示服务端接受的单个连接。
type IServerConn interface {
	// ClientAddress returns the client address string.
	// ClientAddress 返回客户端连接地址。
	ClientAddress() string
	// SendBytes sends raw bytes without repacking.
	// SendBytes 发送二进制数据，不会重新打包。
	SendBytes(bytes []byte) error
	// SendPack sends a framed message pack.
	// SendPack 发送二进制消息包（把数据打包，补充长度信息）。
	SendPack(data []byte) error
	// CloseConn closes the connection.
	// CloseConn 关闭连接。
	CloseConn() error
}

// ISockServer is a socket server with connection management and pack handling.
// ISockServer 是具备连接管理与消息包处理的套接字服务端。
//
// Typical structure:
//
//	ISockServer
//	 ∟ map[address:IServerConn]
//	    Each new connection creates an IServerConn.
//	    UDP implements the interface but effectively uses a single IServerConn.
//	IServerConn
//	 ∟ ISockConn
//	 ∟ IPackSendReceiver
//	      ∟ IConnReadWriterAdapter
//	           ∟ Writer
//	           ∟ Reader
//	      ∟ IPackHandlerContainer
//	           ∟ FuncPackHandler
//
// 一般结构如下：
//
//	ISockServer
//	 ∟ map[address:IServerConn]
//	    每接收到一个新连接，创建一个 IServerConn。
//	    UDP 兼容了接口，但实际上只有一个 IServerConn。
//	IServerConn
//	 ∟ ISockConn
//	 ∟ IPackSendReceiver
//	      ∟ IConnReadWriterAdapter
//	           ∟ Writer
//	           ∟ Reader
//	      ∟ IPackHandlerContainer
//	           ∟ FuncPackHandler（数据包响应函数）
type ISockServer interface {
	ISockName
	IServer
	IServerConnSet

	ISockSender
	IPackHandlerContainerSetter
	IPackHandlerContainerGetter

	logx.ILoggerSupport
}

// ISockEventServer combines event dispatch with ISockServer.
// ISockEventServer 在 ISockServer 基础上支持事件分发。
type ISockEventServer interface {
	eventx.IEventDispatcher
	ISockServer
}

// SockServerBase is a base implementation shared by protocol-specific servers.
// SockServerBase 是各协议服务端共用的基础实现。
type SockServerBase struct {
	Name     string
	Network  SockNetwork
	ServerMu sync.RWMutex
	Running  bool

	Logger logx.ILogger

	PackHandlerContainer IPackHandlerContainer
}

// SetName sets the server identity name.
// SetName 设置服务端标识名称。
func (s *SockServerBase) SetName(name string) {
	s.Name = name
}

// GetName returns the server identity name.
// GetName 获取服务端标识名称。
func (s *SockServerBase) GetName() string {
	return s.Name
}

// GetLogger returns the logger.
// GetLogger 获取日志记录器。
func (s *SockServerBase) GetLogger() logx.ILogger {
	return s.Logger
}

// SetLogger sets the logger.
// SetLogger 设置日志记录器。
func (s *SockServerBase) SetLogger(logger logx.ILogger) {
	s.Logger = logger
}

// IsRunning reports whether the server is running.
// IsRunning 服务是否启动中。
func (s *SockServerBase) IsRunning() bool {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return s.Running
}

// GetPackHandlerContainer returns the pack handler container.
// GetPackHandlerContainer 获取消息包处理器容器。
func (s *SockServerBase) GetPackHandlerContainer() IPackHandlerContainer {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return s.PackHandlerContainer
}

// SetPackHandlerContainer sets the pack handler container.
// SetPackHandlerContainer 设置消息包处理器容器。
func (s *SockServerBase) SetPackHandlerContainer(packHandlerContainer IPackHandlerContainer) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	s.PackHandlerContainer = packHandlerContainer
}

// DispatchServerStartedEvent dispatches the server started event.
// DispatchServerStartedEvent 分发服务启动事件。
func (s *SockServerBase) DispatchServerStartedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStart, dispatcher, nil)
}

// DispatchServerStoppedEvent dispatches the server stopped event.
// DispatchServerStoppedEvent 分发服务停止事件。
func (s *SockServerBase) DispatchServerStoppedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStop, dispatcher, nil)
}

// DispatchServerConnOpenEvent dispatches the connection opened event.
// DispatchServerConnOpenEvent 分发连接打开事件。
func (s *SockServerBase) DispatchServerConnOpenEvent(dispatcher eventx.IEventDispatcher, connInfo IConnInfo) {
	dispatcher.DispatchEvent(ServerEventConnOpened, dispatcher, connInfo)
}

// DispatchServerConnCloseEvent dispatches the connection closed event.
// DispatchServerConnCloseEvent 分发连接关闭事件。
func (s *SockServerBase) DispatchServerConnCloseEvent(dispatcher eventx.IEventDispatcher, connInfo IConnInfo) {
	dispatcher.DispatchEvent(ServerEventConnClosed, dispatcher, connInfo)
}
