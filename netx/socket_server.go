// Package netx
// Created by xuzhuoxi
// on 2019-02-14.
// @author xuzhuoxi
//
package netx

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

type ISockServerSetter interface {
	SetSockServer(server ISockServer)
}

type ISockServerGetter interface {
	GetSockServer() ISockServer
}

type IServer interface {
	// IsRunning
	// 服务是否启动中
	IsRunning() bool
	// StartServer
	// 启动服务，会阻塞
	StartServer(params SockParams) error
	// StopServer
	// 停止服务，会阻塞
	StopServer() error
}

type ISockConnSetSetter interface {
	SetSockConnSet(set IServerConnSet)
}

type ISockConnSetGetter interface {
	GetSockConnSet() IServerConnSet
}

type IServerConnSet interface {
	// SetMaxConn
	// 设置最大连接数量
	SetMaxConn(max int)
	// Connections
	// 连接数
	Connections() int
	// CloseConnection
	// 关闭指定连接
	CloseConnection(connId string) (err error, ok bool)
	// FindConnection
	// 查找连接
	FindConnection(connInd string) (conn IServerConn, ok bool)
}

type IServerConn interface {
	// ClientAddress
	// 客户端连接地址
	ClientAddress() string
	// SendBytes
	// 发送二进制数据, 不会重新打包
	SendBytes(bytes []byte) error
	// SendPack
	// 发送二进制消息包(把数据打包，补充长度信息)
	SendPack(data []byte) error
	// CloseConn
	// 关闭连接
	CloseConn() error
}

// ISockServer Socket连接服务器
//
// 一般结构如下：/n
//
//  ISocketServer
//   ∟  map[address:IServerConn]
//  	每接收到一个新连接，创建一个IServerConn
// 		udp兼容了接口，但实际上只有一个IServerConn
//  IServerConn
//   ∟  ISockConn
//   ∟  IPackSendReceiver
//        ∟  IConnReadWriterAdapter
//             ∟  Writer
//             ∟  Reader
//        ∟  IPackHandlerContainer
//             ∟ - FuncPackHandler (数据包响应函数)
type ISockServer interface {
	ISockName
	IServer
	IServerConnSet

	ISockSender
	IPackHandlerContainerSetter
	IPackHandlerContainerGetter

	logx.ILoggerSupport
}

type ISockEventServer interface {
	eventx.IEventDispatcher
	ISockServer
}

type SockServerBase struct {
	Name     string
	Network  SockNetwork
	ServerMu sync.RWMutex
	Running  bool

	Logger logx.ILogger

	PackHandlerContainer IPackHandlerContainer
}

func (s *SockServerBase) SetName(name string) {
	s.Name = name
}

func (s *SockServerBase) GetName() string {
	return s.Name
}

func (s *SockServerBase) GetLogger() logx.ILogger {
	return s.Logger
}

func (s *SockServerBase) SetLogger(logger logx.ILogger) {
	s.Logger = logger
}

func (s *SockServerBase) IsRunning() bool {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return s.Running
}

func (s *SockServerBase) GetPackHandlerContainer() IPackHandlerContainer {
	s.ServerMu.RLock()
	defer s.ServerMu.RUnlock()
	return s.PackHandlerContainer
}

func (s *SockServerBase) SetPackHandlerContainer(packHandlerContainer IPackHandlerContainer) {
	s.ServerMu.Lock()
	defer s.ServerMu.Unlock()
	s.PackHandlerContainer = packHandlerContainer
}

func (s *SockServerBase) DispatchServerStartedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStart, dispatcher, nil)
}

func (s *SockServerBase) DispatchServerStoppedEvent(dispatcher eventx.IEventDispatcher) {
	dispatcher.DispatchEvent(ServerEventStop, dispatcher, nil)
}

func (s *SockServerBase) DispatchServerConnOpenEvent(dispatcher eventx.IEventDispatcher, connInfo IConnInfo) {
	dispatcher.DispatchEvent(ServerEventConnOpened, dispatcher, connInfo)
}

func (s *SockServerBase) DispatchServerConnCloseEvent(dispatcher eventx.IEventDispatcher, connInfo IConnInfo) {
	dispatcher.DispatchEvent(ServerEventConnClosed, dispatcher, connInfo)
}
