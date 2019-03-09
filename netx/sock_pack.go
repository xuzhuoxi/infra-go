package netx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

const ReceiverBuffLen = 2048

type iPackReceiver interface {
	IPackHandlerSetter
	StartReceiving() error
	StopReceiving() error
	IsReceiving() bool
}

type iPackSender interface {
	//发送字节数据，不作任何处理
	SendBytes(bytes []byte, rAddress ...string) (int, error)
	//发送数据包，把数据进行打包
	SendPack(pack []byte, rAddress ...string) (int, error)
}

type IPackReceiver interface {
	iPackReceiver
	logx.ILoggerGetter
}

type IPackSender interface {
	iPackSender
	logx.ILoggerGetter
}

type IPackSendReceiver interface {
	iPackSender
	iPackReceiver
	logx.ILoggerGetter
}

//--------------------------------------------------------------------------

func NewPackSender(writer IConnWriterAdapter, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackSender {
	return NewPackSendReceiver(nil, writer, nil, dataBlockHandler, logger, false)
}

func NewPackReceiver(reader IConnReaderAdapter, packHandler IPackHandler, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackReceiver {
	return NewPackSendReceiver(reader, nil, packHandler, dataBlockHandler, logger, false)
}

func NewPackMultiReceiver(reader IConnReaderAdapter, packHandler IPackHandler, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackReceiver {
	return NewPackSendReceiver(reader, nil, packHandler, dataBlockHandler, logger, true)
}

func NewPackSendReceiver(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler IPackHandler, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger, multiReceive bool) IPackSendReceiver {
	if multiReceive {
		rs := newPackSendReceiverMulti(reader, writer, packHandler, dataBlockHandler, logger)
		rs.onReceiveBytes = rs.handleDataMulti
		return rs
	} else {
		rs := newPackSendReceiver(reader, writer, packHandler, dataBlockHandler, logger)
		rs.onReceiveBytes = rs.handleData
		return rs
	}
}

//--------------------------------------------------

func newPackSRBase(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler IPackHandler, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) packSRBase {
	return packSRBase{reader: reader, writer: writer, PackHandler: packHandler, dataBlockHandler: dataBlockHandler, toBlockBuff: bytex.NewBuffDataBlock(dataBlockHandler), Logger: logger}
}

type packSRBase struct {
	reader IConnReaderAdapter
	writer IConnWriterAdapter
	mu     sync.RWMutex

	//receive
	receiving      bool
	PackHandler    IPackHandler
	onReceiveBytes func(newData []byte, address interface{})

	//send
	dataBlockHandler bytex.IDataBlockHandler
	toBlockBuff      bytex.IBuffToBlock
	Logger           logx.ILogger
}

func (sr *packSRBase) SetPackHandler(packHandler IPackHandler) {
	sr.PackHandler = packHandler
}

func (sr *packSRBase) SetLogger(logger logx.ILogger) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.Logger = logger
}

func (sr *packSRBase) GetLogger() logx.ILogger {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	return sr.Logger
}

func (sr *packSRBase) SendBytes(bytes []byte, rAddress ...string) (int, error) {
	return sr.writer.WriteBytes(bytes, rAddress...)
}
func (sr *packSRBase) SendPack(msg []byte, rAddress ...string) (int, error) {
	sr.toBlockBuff.Reset()
	sr.toBlockBuff.WriteData(msg)
	n, err := sr.writer.WriteBytes(sr.toBlockBuff.ReadBytes(), rAddress...)
	if nil != err {
		sr.Logger.Warnln("packSRBase.SendMessage", err)
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
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	return sr.receiving
}
func (sr *packSRBase) handleReceiveBytes(buff bytex.IBuffToData, data []byte, address string) {
	buff.WriteBytes(data)
	var unPackData []byte
	for {
		unPackData = buff.ReadCopyData()
		if nil == unPackData {
			break
		}
		sr.PackHandler.ForEachHandler(func(handler FuncPackHandler) bool {
			return handler(unPackData, address)
		})
	}
}

//--------------------------------------------------

func newPackSendReceiver(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler IPackHandler, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) *packSendReceiver {
	return &packSendReceiver{packSRBase: newPackSRBase(reader, writer, packHandler, dataBlockHandler, logger), toDataBuff: bytex.NewBuffToData(dataBlockHandler)}
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

func newPackSendReceiverMulti(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandler IPackHandler, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) *packSendReceiverMulti {
	return &packSendReceiverMulti{packSRBase: newPackSRBase(reader, writer, packHandler, dataBlockHandler, logger), toDataBuffMap: make(map[string]bytex.IBuffToData)}
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
