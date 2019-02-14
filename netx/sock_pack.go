package netx

import (
	"fmt"
	"github.com/xuzhuoxi/util-go/bytex"
	"github.com/xuzhuoxi/util-go/errorsx"
	"github.com/xuzhuoxi/util-go/logx"
	"sync"
)

const ReceiverBuffLen = 2048

type PackHandler func(msgBytes []byte, info interface{})

func DefaultPackHandler(msgData []byte, info interface{}) {
	logx.Traceln("DefaultMessageHandler[Sender:"+fmt.Sprint(info)+"]msgData:", msgData, "dataLen:", len(msgData), "]")
}

type IPackHandlerSetter interface {
	SetPackHandler(handler PackHandler) error
}

type IPackReceiver interface {
	IPackHandlerSetter
	StartReceiving() error
	StopReceiving() error
	IsReceiving() bool
}

type IPackSender interface {
	//发送字节数据，不作任何处理
	SendBytes(bytes []byte, rAddress ...string) (int, error)
	//发送数据包，把数据进行打包
	SendPack(pack []byte, rAddress ...string) (int, error)
}

type IPackSendReceiver interface {
	IPackSender
	IPackReceiver
}

func NewPackSender(writer IConnWriterAdapter, dataBlockHandler bytex.IDataBlockHandler) IPackSender {
	return NewPackSendReceiver(nil, writer, nil, dataBlockHandler, false)
}

func NewPackReceiver(reader IConnReaderAdapter, packHandler PackHandler, dataBlockHandler bytex.IDataBlockHandler) IPackReceiver {
	return NewPackSendReceiver(reader, nil, packHandler, dataBlockHandler, false)
}

func NewPackMultiReceiver(reader IConnReaderAdapter, packHandler PackHandler, dataBlockHandler bytex.IDataBlockHandler) IPackReceiver {
	return NewPackSendReceiver(reader, nil, packHandler, dataBlockHandler, true)
}

func NewPackSendReceiver(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler PackHandler, dataBlockHandler bytex.IDataBlockHandler, multiReceive bool) IPackSendReceiver {
	if multiReceive {
		rs := newPackSendReceiverMulti(reader, writer, packHandler, dataBlockHandler)
		rs.onReceiveBytes = rs.handleDataMulti
		return rs
	} else {
		rs := newPackSendReceiver(reader, writer, packHandler, dataBlockHandler)
		rs.onReceiveBytes = rs.handleData
		return rs
	}
}

//--------------------------------------------------

func newPackSRBase(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler PackHandler, dataBlockHandler bytex.IDataBlockHandler) packSRBase {
	return packSRBase{reader: reader, writer: writer, packHandler: packHandler, dataBlockHandler: dataBlockHandler, toBlockBuff: bytex.NewBuffDataBlock(dataBlockHandler)}
}

type packSRBase struct {
	reader IConnReaderAdapter
	writer IConnWriterAdapter
	mu     sync.Mutex

	//receive
	receiving      bool
	packHandler    PackHandler
	onReceiveBytes func(newData []byte, address interface{})
	//send
	dataBlockHandler bytex.IDataBlockHandler
	toBlockBuff      bytex.IBuffToBlock
}

func (sr *packSRBase) SetPackHandler(handler PackHandler) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.packHandler = handler
	return nil
}
func (sr *packSRBase) SendBytes(bytes []byte, rAddress ...string) (int, error) {
	return sr.writer.WriteBytes(bytes, rAddress...)
}
func (sr *packSRBase) SendPack(msg []byte, rAddress ...string) (int, error) {
	sr.toBlockBuff.Reset()
	sr.toBlockBuff.WriteData(msg)
	n, err := sr.writer.WriteBytes(sr.toBlockBuff.ReadBytes(), rAddress...)
	if nil != err {
		logx.Warnln("packSRBase.SendMessage", err)
		return n, err
	}
	return n, nil
}
func (sr *packSRBase) StartReceiving() error {
	funcName := "packSRBase.StartReceiving"
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
		sr.onReceiveBytes(buff[:n], address)
	}
	return nil
}
func (sr *packSRBase) StopReceiving() error {
	funcName := "packSRBase.StopReceiving"
	if !sr.receiving {
		return errorsx.FuncRepeatedCallError(funcName)
	}
	sr.receiving = false
	return nil
}
func (sr *packSRBase) IsReceiving() bool {
	return sr.receiving
}
func (sr *packSRBase) handleReceiveBytes(buff bytex.IBuffToData, data []byte, address string) {
	buff.WriteBytes(data)
	var unPackData []byte
	for {
		unPackData = buff.ReadData()
		if nil == unPackData {
			break
		}
		sr.packHandler(unPackData, address)
	}
}

//--------------------------------------------------

func newPackSendReceiver(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler PackHandler, dataBlockHandler bytex.IDataBlockHandler) *packSendReceiver {
	return &packSendReceiver{packSRBase: newPackSRBase(reader, writer, packHandler, dataBlockHandler), toDataBuff: bytex.NewBuffToData(dataBlockHandler)}
}

type packSendReceiver struct {
	packSRBase
	toDataBuff bytex.IBuffToData
}

func (sr *packSendReceiver) handleData(newData []byte, address interface{}) {
	//fmt.Println("packSendReceiver.handleData:", newData)
	strAddress := address.(string)
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.handleReceiveBytes(sr.toDataBuff, newData, strAddress)
}

//--------------------------------------------------

func newPackSendReceiverMulti(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler PackHandler, dataBlockHandler bytex.IDataBlockHandler) *packSendReceiverMulti {
	return &packSendReceiverMulti{packSRBase: newPackSRBase(reader, writer, packHandler, dataBlockHandler), toDataBuffMap: make(map[string]bytex.IBuffToData)}
}

type packSendReceiverMulti struct {
	packSRBase
	toDataBuffMap map[string]bytex.IBuffToData
}

func (sr *packSendReceiverMulti) handleDataMulti(newData []byte, address interface{}) {
	//fmt.Println("packSendReceiverMulti.handleData:", newData)
	strAddress := address.(string)
	sr.mu.Lock()
	defer sr.mu.Unlock()
	buffToData, ok := sr.toDataBuffMap[strAddress]
	if !ok {
		buffToData = bytex.NewBuffToData(sr.dataBlockHandler)
		sr.toDataBuffMap[strAddress] = buffToData
	}
	sr.handleReceiveBytes(buffToData, newData, strAddress)
}
