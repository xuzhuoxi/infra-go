package netx

import (
	"fmt"
	"sync"

	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/logx"
)

// ReceiverBuffLen is the default receive buffer size in bytes.
// ReceiverBuffLen 是默认接收缓冲区大小（字节）。
var ReceiverBuffLen = 8192

type iPackReceiver interface {
	IPackHandlerContainerSetter
	IPackHandlerContainerGetter
	// StartReceiving starts receiving data; usually call from a goroutine because it blocks.
	// StartReceiving 开始接收数据；这里一般会阻塞，应该使用协程进行调用。
	StartReceiving() error
	// StopReceiving stops receiving data.
	// StopReceiving 停止接收数据。
	StopReceiving() error
	// IsReceiving reports whether receiving is active.
	// IsReceiving 是否为数据接收中。
	IsReceiving() bool
}

type iPackSender interface {
	// SendBytes sends raw bytes without framing.
	// SendBytes 发送字节数据，不作任何处理。
	SendBytes(bytes []byte, connId ...string) (int, error)
	// SendPack sends a framed data pack.
	// SendPack 发送数据包，把数据进行打包。
	SendPack(pack []byte, connId ...string) (int, error)
}

// IPackReceiver receives and unpacks framed messages.
// IPackReceiver 接收并解包分帧消息。
type IPackReceiver interface {
	iPackReceiver
	logx.ILoggerSupport
}

// IPackReceiverSetter injects an IPackReceiver implementation.
// IPackReceiverSetter 注入 IPackReceiver 实现。
type IPackReceiverSetter interface {
	SetPackReceiver(receiver IPackReceiver)
}

// IPackReceiverGetter returns the bound IPackReceiver.
// IPackReceiverGetter 返回已绑定的 IPackReceiver。
type IPackReceiverGetter interface {
	GetPackReceiver() IPackReceiver
}

// IPackSender sends raw bytes or framed packs.
// IPackSender 发送原始字节或已封包数据。
type IPackSender interface {
	iPackSender
	logx.ILoggerSupport
}

// IPackSenderSetter injects an IPackSender implementation.
// IPackSenderSetter 注入 IPackSender 实现。
type IPackSenderSetter interface {
	SetPackSender(sender IPackSender)
}

// IPackSenderGetter returns the bound IPackSender.
// IPackSenderGetter 返回已绑定的 IPackSender。
type IPackSenderGetter interface {
	GetPackSender() IPackSender
}

// IPackSendReceiver combines pack send and receive.
// IPackSendReceiver 组合消息包发送与接收能力。
type IPackSendReceiver interface {
	iPackSender
	iPackReceiver
	logx.ILoggerSupport
}

//--------------------------------------------------------------------------

// NewPackSender creates a send-only pack handler.
// NewPackSender 创建仅发送的消息包处理器。
func NewPackSender(connInfo IConnInfo, writer IConnWriterAdapter, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackSender {
	return NewPackSendReceiver(connInfo, nil, writer, nil, dataBlockHandler, logger, false)
}

// NewPackReceiver creates a one-to-one receive-only pack handler.
// NewPackReceiver 创建一对一仅接收的消息包处理器。
func NewPackReceiver(connInfo IConnInfo, reader IConnReaderAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackReceiver {
	return NewPackSendReceiver(connInfo, reader, nil, packHandlerContainer, dataBlockHandler, logger, false)
}

// NewPackMultiReceiver creates a one-to-many receive-only pack handler.
// NewPackMultiReceiver 创建一对多仅接收的消息包处理器。
func NewPackMultiReceiver(connInfo IConnInfo, reader IConnReaderAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) IPackReceiver {
	return NewPackSendReceiver(connInfo, reader, nil, packHandlerContainer, dataBlockHandler, logger, true)
}

// NewPackSendReceiver creates a full pack send/receive handler.
// NewPackSendReceiver 创建完整的消息包收发处理器。
func NewPackSendReceiver(connInfo IConnInfo, reader IConnReaderAdapter, writer IConnWriterAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger, multiReceive bool) IPackSendReceiver {
	if multiReceive {
		rs := newPackSendReceiverMulti(connInfo, reader, writer, packHandlerContainer, dataBlockHandler, logger)
		rs.onReceiveBytes = rs.handleDataMulti
		return rs
	} else {
		rs := newPackSendReceiver(connInfo, reader, writer, packHandlerContainer, dataBlockHandler, logger)
		rs.onReceiveBytes = rs.handleData
		return rs
	}
}

//--------------------------------------------------

func newPackSRBase(connInfo IConnInfo, reader IConnReaderAdapter, writer IConnWriterAdapter, packHandlerContainer IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler, logger logx.ILogger) packSRBase {
	return packSRBase{
		connInfo: connInfo, reader: reader, writer: writer,
		PackHandlerContainer: packHandlerContainer,
		dataBlockHandler:     dataBlockHandler, toBlockBuff: bytex.NewBuffDataBlock(dataBlockHandler),
		Logger: logger}
}

type packSRBase struct {
	connInfo IConnInfo
	reader   IConnReaderAdapter
	writer   IConnWriterAdapter
	mu       sync.RWMutex

	//receive
	receiving            bool
	PackHandlerContainer IPackHandlerContainer
	onReceiveBytes       func(newData []byte, connInfo IConnInfo)

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

func (sr *packSRBase) SendBytes(bytes []byte, connId ...string) (int, error) {
	return sr.writer.WriteBytes(bytes, connId...)
}
func (sr *packSRBase) SendPack(msg []byte, connId ...string) (int, error) {
	sr.toBlockBuff.Reset()
	sr.toBlockBuff.WriteData(msg)
	n, err := sr.writer.WriteBytes(sr.toBlockBuff.ReadBytes(), connId...)
	if nil != err {
		sr.Logger.Warnln("[packSRBase.SendPack]", err)
		return n, err
	}
	return n, nil
}

// StartReceiving
// 开始接收数据入口
func (sr *packSRBase) StartReceiving() error {
	funcName := "packSRBase.StartReceiving"
	//sr.Logger.Debugln(funcName)
	sr.mu.Lock()
	if sr.receiving {
		sr.mu.Unlock()
		return errorsx.FuncRepeatedCallError(funcName)
	}
	sr.receiving = true
	sr.mu.Unlock()
	buff := make([]byte, ReceiverBuffLen)
	for sr.receiving {
		n, address, err := sr.reader.ReadBytes(buff[:])
		if err != nil {
			sr.Logger.Warnln("[packSRBase.StartReceiving][for loop]", err)
			break
		}
		//sr.Logger.Debugln("[packSRBase.StartReceiving] Receiving loop:", n, address, sr.connInfo.One2One())
		if sr.connInfo.One2One() { // 一对一连接
			sr.onReceiveBytes(buff[:n], sr.connInfo)
		} else { // 一对多连接
			sr.onReceiveBytes(buff[:n], NewRemoteOnlyConnInfo(sr.connInfo.GetLocalAddress(), address))
		}
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

// handleReceiveBytes
// 处理接收到的数据
func (sr *packSRBase) handleReceiveBytes(buff bytex.IBuffToData, data []byte, connInfo IConnInfo) {
	buff.WriteBytes(data)
	var unPackData []byte
	//sr.Logger.Debugln("handleReceiveBytes:", len(data), sr.PackHandlerContainer.Size(), data)
	for {
		unPackData = buff.ReadDataCopy()
		//sr.Logger.Debugln("handleReceiveBytes2:", len(unPackData))
		if nil == unPackData || len(unPackData) == 0 {
			break
		}
		sr.PackHandlerContainer.ForEachHandler(func(handler FuncPackHandler) bool {
			//sr.Logger.Debugln("handleReceiveBytes2:", len(unPackData), unPackData)
			return handler(unPackData, connInfo, nil)
		})
	}
}

//--------------------------------------------------

func newPackSendReceiver(connInfo IConnInfo, reader IConnReaderAdapter, writer IConnWriterAdapter,
	packHandler IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler,
	logger logx.ILogger) *packSendReceiver {
	return &packSendReceiver{
		packSRBase: newPackSRBase(connInfo, reader, writer, packHandler, dataBlockHandler, logger),
		toDataBuff: bytex.NewBuffToData(dataBlockHandler),
	}
}

type packSendReceiver struct {
	packSRBase
	toDataBuff bytex.IBuffToData
}

// handleData
// 一对一接收数据入口
func (sr *packSendReceiver) handleData(newData []byte, connInfo IConnInfo) {
	//sr.Logger.Debugln("packSendReceiver.handleData:", newData)
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.handleReceiveBytes(sr.toDataBuff, newData, connInfo)
}

//--------------------------------------------------

func newPackSendReceiverMulti(connInfo IConnInfo, reader IConnReaderAdapter, writer IConnWriterAdapter,
	packHandler IPackHandlerContainer, dataBlockHandler bytex.IDataBlockHandler,
	logger logx.ILogger) *packSendReceiverMulti {
	return &packSendReceiverMulti{
		packSRBase:    newPackSRBase(connInfo, reader, writer, packHandler, dataBlockHandler, logger),
		toDataBuffMap: make(map[string]bytex.IBuffToData)}
}

type packSendReceiverMulti struct {
	packSRBase
	toDataBuffMap map[string]bytex.IBuffToData
}

// handleDataMulti
// 多对一接收数据入口
func (sr *packSendReceiverMulti) handleDataMulti(newData []byte, connInfo IConnInfo) {
	fmt.Println("packSendReceiverMulti.handleData:", newData)
	sr.mu.Lock()
	defer sr.mu.Unlock()
	buffToData, ok := sr.toDataBuffMap[connInfo.GetConnId()]
	if !ok {
		buffToData = bytex.NewBuffToData(sr.dataBlockHandler)
		sr.toDataBuffMap[connInfo.GetConnId()] = buffToData
	}
	sr.handleReceiveBytes(buffToData, newData, connInfo)
}
