package netx

import (
	"fmt"
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
	"sync"
)

const ReceiverBuffLen = 2048

func NewMessageSender(writer IWriterProxy, t ReadWriterType, Network string) IMessageSender {
	rs := &msgSendReceiver{writer: writer, messageHandler: DefaultMessageHandler}
	rs.init()
	return rs
}

func NewMessageReceiver(reader IReaderProxy, t ReadWriterType, Network string) IMessageReceiver {
	rs := &msgSendReceiver{reader: reader, messageHandler: DefaultMessageHandler}
	rs.init()
	return rs
}

func NewMessageSendReceiver(reader IReaderProxy, writer IWriterProxy, t ReadWriterType, Network string) IMessageSendReceiver {
	rs := &msgSendReceiver{reader: reader, writer: writer, messageHandler: DefaultMessageHandler}
	rs.init()
	return rs
}

type msgSendReceiver struct {
	rwType    ReadWriterType
	receiving bool
	mu        sync.Mutex

	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(msgBytes []byte, info interface{})

	reader IReaderProxy
	writer IWriterProxy

	mapSplitter map[string]IByteSplitter
	splitter    IByteSplitter
}

func (sr *msgSendReceiver) init() {
	switch {
	case sr.rwType == TcpRW || sr.rwType == UdpDialRW || sr.rwType == QuicRW:
		sr.splitter = NewByteSplitter()
		if nil != sr.splitHandler {
			sr.splitter.SetSplitHandler(sr.splitHandler)
		}
	case sr.rwType == UdpListenRW:
		sr.mapSplitter = make(map[string]IByteSplitter)
	}
}

func (sr *msgSendReceiver) SendMessage(msg []byte, rAddress ...string) (int, error) {
	n, err := sr.writer.WriteBytes(msg, rAddress...)
	if nil != err {
		logx.Warnln("msgSendReceiver.SendMessage", err)
		return n, err
	}
	return n, nil
}

func (sr *msgSendReceiver) SetSplitHandler(handler HandlerForSplit) error {
	funcName := "msgSendReceiver.SetSplitHandler"
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.splitHandler = handler
	switch {
	case sr.rwType == TcpRW || sr.rwType == UdpDialRW || sr.rwType == QuicRW:
		sr.splitter.SetSplitHandler(handler)
	case sr.rwType == UdpListenRW:
		for _, value := range sr.mapSplitter {
			value.SetSplitHandler(handler)
		}
	default:
		return errorsx.NoCaseCatchError(funcName)
	}
	return nil
}

func (sr *msgSendReceiver) SetMessageHandler(handler HandlerForMessage) error {
	sr.messageHandler = handler
	return nil
}

func (sr *msgSendReceiver) StartReceiving() error {
	funcName := "msgSendReceiver.StartReceiving"
	if sr.receiving {
		return errorsx.FuncRepeatedCallError(funcName)
	}

	defer sr.StopReceiving()
	sr.receiving = true
	var buff [ReceiverBuffLen]byte
	for sr.receiving {
		n, address, err := sr.reader.ReadBytes(buff[:])
		if err != nil {
			break
		}
		sr.handleData(buff[:n], address)
	}
	return nil
}

func (sr *msgSendReceiver) StopReceiving() error {
	funcName := "msgSendReceiver.StopReceiving"
	if !sr.receiving {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	sr.receiving = false
	return nil
}

func (sr *msgSendReceiver) IsReceiving() bool {
	return sr.receiving
}

func (sr *msgSendReceiver) handleData(newData []byte, address interface{}) {
	strAddress := address.(string)
	switch {
	case sr.rwType == TcpRW || sr.rwType == UdpDialRW || sr.rwType == QuicRW:
		sr.handleSplitterData(sr.splitter, newData, strAddress)
	case sr.rwType == UdpListenRW:
		sr.mu.Lock()
		defer sr.mu.Unlock()
		splitter, ok := sr.mapSplitter[strAddress]
		if !ok {
			splitter = NewByteSplitter()
			splitter.SetSplitHandler(sr.splitHandler)
			sr.mapSplitter[strAddress] = splitter
		}
		sr.handleSplitterData(splitter, newData, strAddress)
	}
}

func (sr *msgSendReceiver) handleSplitterData(splitter IByteSplitter, data []byte, address string) {
	splitter.AppendBytes(data)
	for splitter.CheckSplit() {
		sr.messageHandler(splitter.FrontBytes(), address)
	}
}

func DefaultMessageHandler(msgData []byte, info interface{}) {
	logx.Traceln("DefaultMessageHandler[Sender:"+fmt.Sprint(info)+"]msgData:", msgData, "dataLen:", len(msgData), "]")
}
