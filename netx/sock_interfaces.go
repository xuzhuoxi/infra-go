package netx

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

type ISplitHandler interface {
	SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error
}

type IMessageHandler interface {
	SetMessageHandler(handler func(msgBytes []byte, info interface{})) error
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
	SendDataTo(msg []byte, rAddress ...string) error
}

type ISockServer interface {
	ISplitHandler
	IMessageHandler
	ISendData

	StartServer(params SockParams) error //会阻塞
	StopServer() error
	Running() bool
}

type ISockClient interface {
	IMessageReceiver
	ISendData

	LocalAddress() string
	OpenClient(params SockParams) error
	CloseClient() error
	Opening() bool
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

type IUDPClient interface {
	ISockClient
}

type ITCPClient interface {
	ISockClient
}
