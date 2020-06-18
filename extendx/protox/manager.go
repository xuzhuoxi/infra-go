//
//Created by xuzhuoxi
//on 2019-05-18.
//@author xuzhuoxi
//
package protox

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"sync"
)

// Extension管理接口
type IExtensionManager interface {
	logx.ILoggerSetter
	netx.IAddressProxySetter

	// 初始化
	InitManager(handlerContainer netx.IPackHandlerContainer, extensionContainer IProtocolExtensionContainer, sockSender netx.ISockSender)

	// 开始运行
	StartManager()
	// 停止运行
	StopManager()

	// 保存指定Extension的临时数据
	SaveExtension(name string)
	// 保存全部Extension的临时数据
	SaveExtensions()

	// 启用指定Extension的临时数据
	EnableExtension(name string)
	// 启用全部Extension的临时数据
	EnableExtensions()
	// 禁用指定Extension的临时数据
	DisableExtension(name string)
	// 禁用全部Extension的临时数据
	DisableExtensions()

	// 消息处理入口，这里是并发方法
	onPack(msgData []byte, senderAddress string, other interface{}) bool
}

//---------------------------------------------

func NewIExtensionManager() IExtensionManager {
	return NewExtensionManager()
}

func NewExtensionManager() *ExtensionManager {
	return &ExtensionManager{}
}

type ExtensionManager struct {
	HandlerContainer   netx.IPackHandlerContainer
	ExtensionContainer IProtocolExtensionContainer
	SockSender         netx.ISockSender

	Logger       logx.ILogger
	AddressProxy netx.IAddressProxy
	Mutex        sync.RWMutex

	ExtensionManagerCustomizeSupport
}

func (m *ExtensionManager) InitManager(handlerContainer netx.IPackHandlerContainer, extensionContainer IProtocolExtensionContainer,
	sockSender netx.ISockSender) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.HandlerContainer, m.ExtensionContainer, m.SockSender = handlerContainer, extensionContainer, sockSender
}

func (m *ExtensionManager) SetAddressProxy(proxy netx.IAddressProxy) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.AddressProxy = proxy
}

func (m *ExtensionManager) SetLogger(logger logx.ILogger) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.Logger = logger
}

func (m *ExtensionManager) StartManager() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.ExtensionContainer.InitExtensions()
	m.HandlerContainer.AppendPackHandler(m.onPack)
}

func (m *ExtensionManager) StopManager() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.HandlerContainer.ClearHandler(m.onPack)
	m.ExtensionContainer.DestroyExtensions()
}

func (m *ExtensionManager) SaveExtensions() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.ExtensionContainer.SaveExtensions()
}

func (m *ExtensionManager) SaveExtension(name string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.ExtensionContainer.SaveExtension(name)
}

func (m *ExtensionManager) EnableExtension(name string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.ExtensionContainer.EnableExtension(name, true)
}

func (m *ExtensionManager) DisableExtension(name string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.ExtensionContainer.EnableExtension(name, false)
}

func (m *ExtensionManager) EnableExtensions() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.ExtensionContainer.EnableExtensions(true)
}

func (m *ExtensionManager) DisableExtensions() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.ExtensionContainer.EnableExtensions(false)
}

//---------------------------------

// 消息处理入口，这里是并发方法
// 并发注意:本方法是否并发，取决于SockServer的并发性
// 在netx中，TCP,Quic,WebSocket为并发响应，UDP为非并发响应
// msgData非共享的，但在parsePackMessage后这部分数据会发生变化
func (m *ExtensionManager) onPack(msgData []byte, senderAddress string, other interface{}) bool {
	//m.logger.Infoln("ExtensionManager.onPack", senderAddress, msgData)
	m.CustomStartOnPack(senderAddress)
	name, pid, uid, data := m.ParseMessage(msgData)
	extension, ok := m.Verify(name, pid, uid)
	if !ok {
		return false
	}
	//参数处理
	response, request := m.GenParams(extension, senderAddress, name, pid, uid, data)
	defer func() {
		DefaultRequestPool.Recycle(request)
		DefaultResponsePool.Recycle(response)
	}()
	//响应处理
	if be, ok := extension.(IBeforeRequestExtension); ok { //前置处理
		be.BeforeRequest(request)
	}
	if re, ok := extension.(IRequestExtension); ok {
		m.CustomStartOnRequest(response, request)
		re.OnRequest(response, request)
		m.CustomFinishOnRequest(response, request)
	}
	if ae, ok := extension.(IAfterRequestExtension); ok { //后置处理
		ae.AfterRequest(response, request)
	}
	return true
}

//block0 : eName utf8
//block1 : pid	utf8
//block2 : uid	utf8
//[n]其它信息
func (m *ExtensionManager) ParseMessage(msgBytes []byte) (name string, pid string, uid string, data [][]byte) {
	if nil != m.FuncParseMessage {
		return m.FuncParseMessage(msgBytes)
	}
	index := 0
	buffToData := bytex.DefaultPoolBuffToData.GetInstance()
	defer bytex.DefaultPoolBuffToData.Recycle(buffToData)

	buffToData.WriteBytes(msgBytes)
	name = string(buffToData.ReadData())
	pid = string(buffToData.ReadData())
	uid = string(buffToData.ReadData())
	if buffToData.Len() > 0 {
		for buffToData.Len() > 0 {
			n, d := buffToData.ReadDataTo(msgBytes[index:]) //由于msgBytes前部分数据已经处理完成，可以利用这部分空间
			//h.singleCase.GetLogger().Traceln("parsePackMessage", uid, d)
			if nil == d {
				//h.singleCase.GetLogger().Warnln("data is nil")
				break
			}
			data = append(data, d)
			index += n
		}
	}
	return name, pid, uid, data
}

func (m *ExtensionManager) Verify(name string, pid string, uid string) (e IProtocolExtension, ok bool) {
	if nil != m.FuncVerify {
		return m.FuncVerify(name, pid, uid)
	}
	extension, ok := m.GetProtocolExtension(name)
	//有效性验证
	if !ok {
		if nil != m.Logger {
			m.Logger.Warnln(fmt.Sprintf("Undefined Extension(%s)! Sender(%s)", name, uid))
		}
		return nil, false
	}
	if !extension.CheckProtocolId(pid) { //有效性检查
		if nil != m.Logger {
			m.Logger.Warnln(fmt.Sprintf("Undefined ProtoId(%s) Send to Extension(%s)! Sender(%s)", pid, name, uid))
		}
		return nil, false
	}
	return extension, true
}

// 构造响应参数
func (m *ExtensionManager) GenParams(extension IProtocolExtension, senderAddress string, name string, pid string, uid string, data [][]byte) (resp IExtensionResponse, req IExtensionRequest) {
	t, h := extension.GetParamInfo(pid)
	response := DefaultResponsePool.GetInstance()
	response.SetHeader(name, pid, uid, senderAddress)
	response.SetSockSender(m.SockSender)
	response.SetAddressProxy(m.AddressProxy)
	response.SetParamInfo(t, h)
	request := DefaultRequestPool.GetInstance()
	request.SetHeader(name, pid, uid, senderAddress)
	request.SetRequestData(t, h, data)
	return response, request
}

func (m *ExtensionManager) GetProtocolExtension(pid string) (pe IProtocolExtension, ok bool) {
	if pe, ok := m.ExtensionContainer.GetExtension(pid).(IProtocolExtension); ok {
		return pe, true
	}
	return nil, false
}
