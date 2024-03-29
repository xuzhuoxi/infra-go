package netx

import (
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

var ReceiverBuffLen = 2048

type iPackReceiver interface {
	IPackHandlerContainerSetter
	IPackHandlerContainerGetter
	// StartReceiving
	// 开始接收数据
	// 这里一般会阻塞，应该使用协程进行调用
	StartReceiving() error
	// StopReceiving
	// 停止接收数据
	StopReceiving() error
	// IsReceiving
	// 是否为数据接收中
	IsReceiving() bool
}

type iPackSender interface {
	// SendBytes
	// 发送字节数据，不作任何处理
	SendBytes(bytes []byte, rAddress ...string) (int, error)
	// SendPack
	// 发送数据包，把数据进行打包
	SendPack(pack []byte, rAddress ...string) (int, error)
}

type IPackReceiver interface {
	iPackReceiver
	logx.ILoggerSupport
}

type IPackReceiverSetter interface {
	SetPackReceiver(receiver IPackReceiver)
}

type IPackReceiverGetter interface {
	GetPackReceiver() IPackReceiver
}

type IPackSender interface {
	iPackSender
	logx.ILoggerSupport
}

type IPackSenderSetter interface {
	SetPackSender(sender IPackSender)
}

type IPackSenderGetter interface {
	GetPackSender() IPackSender
}

type IPackSendReceiver interface {
	iPackSender
	iPackReceiver
	logx.ILoggerSupport
}

//--------------------------------------------------------------------------

func NewPackSender(writer IConnWriterAdapter, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackSender {
	return NewPackSendReceiver(nil, writer, nil, dataBlockHandler, logger, false)
}

func NewPackReceiver(reader IConnReaderAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackReceiver {
	return NewPackSendReceiver(reader, nil, packHandlerContainer, dataBlockHandler, logger, false)
}

func NewPackMultiReceiver(reader IConnReaderAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackReceiver {
	return NewPackSendReceiver(reader, nil, packHandlerContainer, dataBlockHandler, logger, true)
}

func NewPackSendReceiver(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger, multiReceive bool) IPackSendReceiver {
	if multiReceive {
		rs := newPackSendReceiverMulti(reader, writer, packHandlerContainer, dataBlockHandler, logger)
		rs.onReceiveBytes = rs.handleDataMulti
		return rs
	} else {
		rs := newPackSendReceiver(reader, writer, packHandlerContainer, dataBlockHandler, logger)
		rs.onReceiveBytes = rs.handleData
		return rs
	}
}

//--------------------------------------------------

func newPackSRBase(reader IConnReaderAdapter, writer IConnWriterAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) packSRBase {
	return packSRBase{reader: reader, writer: writer, PackHandlerContainer: packHandlerContainer, dataBlockHandler: dataBlockHandler, toBlockBuff: bytex.NewBuffDataBlock(dataBlockHandler), Logger: logger}
}

type packSRBase struct {
	reader IConnReaderAdapter
	writer IConnWriterAdapter
	mu     sync.RWMutex

	//receive
	receiving            bool
	PackHandlerContainer IPackHandlerContainer
	onReceiveBytes       func(newData []byte, address string)

	//send
	dataBlockHandler bytex.IDataBlockHandler
	toBlockBuff      bytex.IBuffToBlock
	Logger           logx.ILogger
}

func (sr *packSRBase) GetPackHandlerContainer() IPackHandlerContainer {
	return sr.PackHandlerContainer
}

func (sr *packSRBase) SetPackHandlerContainer(packHandlerContainer IPackHandlerContainer) {
	sr.PackHandlerContainer = packHandlerContainer
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
		sr.Logger.Warnln("[packSRBase.SendPack]", "packSRBase.SendPack", err)
		return n, err
	}
	return n, nil
}
func (sr *packSRBase) StartReceiving() error {
	funcName := "packSRBase.StartReceiving"
	sr.mu.Lock()
	if sr.receiving {
		sr.mu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	sr.receiving = true
	sr.mu.Unlock()
	buff := make([]byte, ReceiverBuffLen, ReceiverBuffLen)
	for sr.receiving {
		n, address, err := sr.reader.ReadBytes(buff[:])
		if err != nil {
			sr.Logger.Warnln("[packSRBase.StartReceiving]", err)
			break
		}
		sr.onReceiveBytes(buff[:n], address)
	}
	_ = sr.StopReceiving()
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
		unPackData = buff.ReadDataCopy()
		if nil == unPackData || len(unPackData) == 0 {
			break
		}
		sr.PackHandlerContainer.ForEachHandler(func(handler FuncPackHandler) bool {
			return handler(unPackData, address, nil)
		})
	}
}

//--------------------------------------------------

func newPackSendReceiver(reader IConnReaderAdapter, writer IConnWriterAdapter,
	packHandler IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler,
	logger logx.ILogger) *packSendReceiver {
	return &packSendReceiver{
		packSRBase: newPackSRBase(reader, writer, packHandler, dataBlockHandler, logger),
		toDataBuff: bytex.NewBuffToData(dataBlockHandler),
	}
}

type packSendReceiver struct {
	packSRBase
	toDataBuff bytex.IBuffToData
}

func (sr *packSendReceiver) handleData(newData []byte, address string) {
	//fmt.Println("packSendReceiver.handleData:", newData)
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.handleReceiveBytes(sr.toDataBuff, newData, address)
}

//--------------------------------------------------

func newPackSendReceiverMulti(reader IConnReaderAdapter, writer IConnWriterAdapter,
	packHandler IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler,
	logger logx.ILogger) *packSendReceiverMulti {
	return &packSendReceiverMulti{
		packSRBase:    newPackSRBase(reader, writer, packHandler, dataBlockHandler, logger),
		toDataBuffMap: make(map[string]bytex.IBuffToData)}
}

type packSendReceiverMulti struct {
	packSRBase
	toDataBuffMap map[string]bytex.IBuffToData
}

func (sr *packSendReceiverMulti) handleDataMulti(newData []byte, address string) {
	//fmt.Println("packSendReceiverMulti.handleData:", newData)
	sr.mu.Lock()
	defer sr.mu.Unlock()
	buffToData, ok := sr.toDataBuffMap[address]
	if !ok {
		buffToData = bytex.NewBuffToData(sr.dataBlockHandler)
		sr.toDataBuffMap[address] = buffToData
	}
	sr.handleReceiveBytes(buffToData, newData, address)
}
