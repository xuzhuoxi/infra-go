package netx

import (
	"net"
)

type IReaderProxy interface {
	ReadBytes(bytes []byte) (int, interface{}, error)
}

type IWriterProxy interface {
	WriteBytes(bytes []byte, rAddress ...string) (int, error)
}

type IReadWriterProxy interface {
	IReaderProxy
	IWriterProxy
}

type HandlerForSplit func(buff []byte) ([]byte, []byte)

type ISplitHandler interface {
	SetSplitHandler(handler HandlerForSplit) error
}

type HandlerForMessage func(msgBytes []byte, info interface{})

type IMessageHandler interface {
	SetMessageHandler(handler HandlerForMessage) error
}

type IByteSplitter interface {
	ISplitHandler
	AppendBytes(data []byte)
	CheckSplit() bool
	FrontBytes() []byte
}

type IMessageReceiver interface {
	ISplitHandler
	IMessageHandler
	StartReceiving() error
	StopReceiving() error
	IsReceiving() bool
}

type IMessageSender interface {
	SendMessage(msg []byte, rAddress ...string) (int, error)
}

type IMessageSendReceiver interface {
	IMessageSender
	IMessageReceiver
}

type ISendData interface {
	SendDataTo(data []byte, rAddress ...string) error
}

type ISockConn interface {
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

type IServer interface {
	StartServer(params SockParams) error //会阻塞
	StopServer() error
	Running() bool
}

type ISockServer interface {
	ISplitHandler
	IMessageHandler
	ISendData

	IServer
}

type IClient interface {
	OpenClient(params SockParams) error
	CloseClient() error
	Opening() bool
}

type ISockClient interface {
	IMessageReceiver
	ISendData
	IClient

	LocalAddress() string
}

type IUDPServer interface {
	ISockServer
}

type ITCPServer interface {
	ISockServer
}

type IQUICServer interface {
	ISockServer
}

type IWebSocketServer interface {
	ISockServer
}

type IUDPClient interface {
	ISockClient
}

type ITCPClient interface {
	ISockClient
}

type IQuicClient interface {
	ISockClient
}

type IWebSocketClient interface {
	ISockClient
}
