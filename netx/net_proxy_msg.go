package netx

import (
	"fmt"
	"github.com/xuzhuoxi/go-util/errorsx"
	"log"
	"sync"
)

const ReceiverBuffLen = 2048

func NewMessageSender(writer interface{}, t ReadWriterType, Network string) IMessageSender {
	proxy := NewWriterProxy(writer, t, Network)
	rs := &msgSendReceiver{writer: proxy, messageHandler: DefaultMessageHandler}
	rs.init()
	return rs
}

func NewMessageReceiver(reader interface{}, t ReadWriterType, Network string) IMessageReceiver {
	proxy := NewReaderProxy(reader, t, Network)
	rs := &msgSendReceiver{reader: proxy, messageHandler: DefaultMessageHandler}
	rs.init()
	return rs
}

func NewMessageSendReceiver(reader interface{}, writer interface{}, t ReadWriterType, Network string) IMessageSendReceiver {
	proxy := NewReadWriterProxy(reader, writer, t, Network)
	rs := &msgSendReceiver{reader: proxy, writer: proxy, messageHandler: DefaultMessageHandler}
	rs.init()
	return rs
}

type msgSendReceiver struct {
	rwType    ReadWriterType
	receiving bool
	mu        sync.Mutex

	reader         IReaderProxy
	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(msgBytes []byte, info interface{})
	mapSplitter    map[string]IByteSplitter
	splitter       IByteSplitter

	writer IWriterProxy
}

func (sr *msgSendReceiver) init() {
	switch {
	case sr.rwType == TcpRW || sr.rwType == UdpDialRW:
		sr.splitter = NewByteSplitter()
		if nil != sr.splitHandler {
			sr.splitter.SetSplitHandler(sr.splitHandler)
		}
	case sr.rwType == UdpListenRW:
		sr.mapSplitter = make(map[string]IByteSplitter)
	}
}

func (sr *msgSendReceiver) SendMessage(msg []byte, rAddress ...string) (int, error) {
	return sr.writer.WriteBytes(msg, rAddress...)
}

func (sr *msgSendReceiver) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	funcName := "msgSendReceiver.SetSplitHandler"
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.splitHandler = handler
	switch {
	case sr.rwType == TcpRW || sr.rwType == UdpDialRW:
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

func (sr *msgSendReceiver) SetMessageHandler(handler func(msgBytes []byte, info interface{})) error {
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
	case sr.rwType == TcpRW || sr.rwType == UdpDialRW:
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
	log.Println("DefaultMessageHandler[Sender:"+fmt.Sprint(info)+"]msgData:", msgData, "dataLen:", len(msgData), "]")
}
