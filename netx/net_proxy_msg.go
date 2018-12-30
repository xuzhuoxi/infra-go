package netx

import (
	"fmt"
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
	"sync"
)

const ReceiverBuffLen = 2048

func NewMessageSender(writer IWriterProxy) IMessageSender {
	return NewMessageSendReceiver(nil, writer, false)
}

func NewMessageReceiver(reader IReaderProxy, multiReceive bool) IMessageReceiver {
	return NewMessageSendReceiver(reader, nil, multiReceive)
}

func NewMessageSendReceiver(reader IReaderProxy, writer IWriterProxy, multiReceive bool) IMessageSendReceiver {
	base := msgsr{reader: reader, writer: writer, splitHandler: DefaultByteSplitHandler, messageHandler: DefaultMessageHandler}
	if multiReceive {
		rs := &msgSendReceiverMulti{msgsr: base, mapSplitter: make(map[string]IByteSplitter)}
		rs.dispatchHandler = rs.handleDataMulti
		return rs
	} else {
		rs := &msgSendReceiver{msgsr: base, splitter: NewByteSplitter()}
		rs.dispatchHandler = rs.handleData
		return rs
	}
}

type msgsr struct {
	receiving bool
	mu        sync.Mutex

	splitHandler   HandlerForSplit
	messageHandler HandlerForMessage

	reader          IReaderProxy
	writer          IWriterProxy
	dispatchHandler func(newData []byte, address interface{})
}

func (sr *msgsr) SetSplitHandler(handler HandlerForSplit) error {
	panic("implement me")
}
func (sr *msgsr) SendMessage(msg []byte, rAddress ...string) (int, error) {
	n, err := sr.writer.WriteBytes(msg, rAddress...)
	if nil != err {
		logx.Warnln("msgsr.SendMessage", err)
		return n, err
	}
	return n, nil
}
func (sr *msgsr) SetMessageHandler(handler HandlerForMessage) error {
	sr.messageHandler = handler
	return nil
}
func (sr *msgsr) StartReceiving() error {
	funcName := "msgsr.StartReceiving"
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
		sr.dispatchHandler(buff[:n], address)
	}
	return nil
}
func (sr *msgsr) StopReceiving() error {
	funcName := "msgsr.StopReceiving"
	if !sr.receiving {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	sr.receiving = false
	return nil
}
func (sr *msgsr) IsReceiving() bool {
	return sr.receiving
}
func (sr *msgsr) handleSplitterData(splitter IByteSplitter, data []byte, address string) {
	splitter.AppendBytes(data)
	for splitter.CheckSplit() {
		sr.messageHandler(splitter.FrontBytes(), address)
	}
}

type msgSendReceiver struct {
	msgsr
	splitter IByteSplitter
}

func (sr *msgSendReceiver) SetSplitHandler(handler HandlerForSplit) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.splitHandler = handler
	if nil == sr.splitter {
		sr.splitter = NewByteSplitter()
	}
	sr.splitter.SetSplitHandler(handler)
	return nil
}
func (sr *msgSendReceiver) handleData(newData []byte, address interface{}) {
	strAddress := address.(string)
	sr.handleSplitterData(sr.splitter, newData, strAddress)
}

type msgSendReceiverMulti struct {
	msgsr
	mapSplitter map[string]IByteSplitter
}

func (sr *msgSendReceiverMulti) SetSplitHandler(handler HandlerForSplit) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.splitHandler = handler
	if nil == sr.mapSplitter {
		sr.mapSplitter = make(map[string]IByteSplitter)
		return nil
	}
	for _, value := range sr.mapSplitter {
		value.SetSplitHandler(handler)
	}
	return nil
}
func (sr *msgSendReceiverMulti) handleDataMulti(newData []byte, address interface{}) {
	strAddress := address.(string)
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

func DefaultMessageHandler(msgData []byte, info interface{}) {
	logx.Traceln("DefaultMessageHandler[Sender:"+fmt.Sprint(info)+"]msgData:", msgData, "dataLen:", len(msgData), "]")
}
