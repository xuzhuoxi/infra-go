// Package protox
// Created by xuzhuoxi
// on 2019-05-18.
// @author xuzhuoxi
//
package protox

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/bytex"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/netx"
	"sync"
)

// IExtensionManager
// Extension管理接口
type IExtensionManager interface {
	logx.ILoggerSetter
	netx.IAddressProxySetter

	// InitManager
	// 初始化
	// handlerContainer: 解包处理
	// extensionContainer： 服务扩展
	// sockSender: 消息发送器
	InitManager(handlerContainer netx.IPackHandlerContainer, extensionContainer IProtocolExtensionContainer, sockSender netx.ISockSender)

	// StartManager
	// 开始运行
	StartManager()
	// StopManager
	// 停止运行
	StopManager()

	// SaveExtension
	// 保存指定Extension的临时数据
	SaveExtension(name string)
	// SaveExtensions
	// 保存全部Extension的临时数据
	SaveExtensions()

	// EnableExtension
	// 启用指定Extension的临时数据
	EnableExtension(name string)
	// EnableExtensions
	// 启用全部Extension的临时数据
	EnableExtensions()
	// DisableExtension
	// 禁用指定Extension的临时数据
	DisableExtension(name string)
	// DisableExtensions
	// 禁用全部Extension的临时数据
	DisableExtensions()

	// OnMessageUnpack
	// 消息处理入口，这里是并发方法
	OnMessageUnpack(msgData []byte, senderAddress string, other interface{}) bool
	// DoRequest
	// 消息处理入口，这里是并发方法
	DoRequest(extension IProtocolExtension, req IExtensionRequest, resp IExtensionResponse)
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
	m.HandlerContainer.AppendPackHandler(m.OnMessageUnpack)
}

func (m *ExtensionManager) StopManager() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	m.HandlerContainer.ClearHandler(m.OnMessageUnpack)
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

// OnMessageUnpack
// 消息处理入口，这里是并发方法
// 并发注意:本方法是否并发，取决于SockServer的并发性
// 在netx中，TCP,Quic,WebSocket为并发响应，UDP为非并发响应
// msgData非共享的，但在parsePackMessage后这部分数据会发生变化
func (m *ExtensionManager) OnMessageUnpack(msgData []byte, senderAddress string, other interface{}) bool {
	//m.Logger.Infoln("ExtensionManager.onPack", senderAddress, msgData)
	m.CustomStartOnPack(senderAddress)
	name, pid, uid, data := m.ParseMessage(msgData)
	extension, rsCode := m.Verify(name, pid, uid)
	if CodeSuc != rsCode {
		// 这里可以直接响应失败
		return false
	}
	//参数处理
	response, request := m.GetRecycleParams(extension, senderAddress, name, pid, uid, data)
	defer func() {
		DefaultRequestPool.Recycle(request)
		DefaultResponsePool.Recycle(response)
	}()
	//响应处理
	m.DoRequest(extension, request, response)
	return true
}

// ParseMessage
// block0 : eName utf8
// block1 : pid	utf8
// block2 : uid	utf8
// block3 : data [][]byte
// [n]其它信息
func (m *ExtensionManager) ParseMessage(msgBytes []byte) (extName string, pid string, uid string, data [][]byte) {
	if nil != m.FuncParseMessage {
		return m.FuncParseMessage(msgBytes)
	}
	buffToData := bytex.DefaultPoolBuffToData.GetInstance()
	defer bytex.DefaultPoolBuffToData.Recycle(buffToData)

	buffToData.WriteBytes(msgBytes)
	extName = buffToData.ReadString()
	pid = buffToData.ReadString()
	uid = buffToData.ReadString()
	if buffToData.Len() <= 0 {
		return
	}
	index := 0
	for buffToData.Len() > 0 {
		n, d := buffToData.ReadDataTo(msgBytes[index:]) //由于msgBytes前部分数据已经处理完成，可以利用这部分空间
		//h.singleCase.GetLogger().Traceln("parsePackMessage", uid, d)
		if n == 0 { // 没有读到字节，注意 n!=0时, d是有可能是nil的
			//h.singleCase.GetLogger().Warnln("data is nil")
			break
		}
		data = append(data, d)
		index += n
	}
	return extName, pid, uid, data
}

func (m *ExtensionManager) Verify(name string, pid string, uid string) (e IProtocolExtension, rsCode int32) {
	if nil != m.FuncVerify {
		return m.FuncVerify(name, pid, uid)
	}
	extension, ok := m.GetProtocolExtension(name)
	//有效性验证
	if !ok {
		if nil != m.Logger {
			m.Logger.Warnln(fmt.Sprintf("Undefined Extension(%s)! Sender(%s)", name, uid))
		}
		return nil, CodeProtoFail
	}
	if !extension.CheckProtocolId(pid) { //有效性检查
		if nil != m.Logger {
			m.Logger.Warnln(fmt.Sprintf("Undefined ProtoId(%s) Send to Extension(%s)! Sender(%s)", pid, name, uid))
		}
		return nil, CodeProtoFail
	}
	return extension, CodeSuc
}

// GetRecycleParams
// 构造响应参数
func (m *ExtensionManager) GetRecycleParams(extension IProtocolExtension, senderAddress string, name string, pid string, uid string, data [][]byte) (resp IExtensionResponse, req IExtensionRequest) {
	t, h := extension.GetParamInfo(pid)
	response := DefaultResponsePool.GetInstance().(iExtResponse)
	response.SetHeader(name, pid, uid, senderAddress)
	response.SetSockSender(m.SockSender)
	response.SetAddressProxy(m.AddressProxy)
	response.SetResultCode(CodeSuc)
	response.SetParamInfo(t, h)
	request := DefaultRequestPool.GetInstance()
	request.SetHeader(name, pid, uid, senderAddress)
	request.SetRequestData(t, h, data)
	return response, request
}

// GetRecycleResponse
// 构造响应参数
func (m *ExtensionManager) GetRecycleResponse(extension IProtocolExtension, senderAddress string, name string, pid string, uid string, data [][]byte) (resp IExtensionResponse) {
	t, h := extension.GetParamInfo(pid)
	response := DefaultResponsePool.GetInstance().(iExtResponse)
	response.SetHeader(name, pid, uid, senderAddress)
	response.SetSockSender(m.SockSender)
	response.SetAddressProxy(m.AddressProxy)
	response.SetResultCode(CodeSuc)
	response.SetParamInfo(t, h)
	return response
}

// GetRecycleRequest
// 获取可回收的请求结构
func (m *ExtensionManager) GetRecycleRequest(extension IProtocolExtension, senderAddress string, name string, pid string, uid string, data [][]byte) (req IExtensionRequest) {
	t, h := extension.GetParamInfo(pid)
	request := DefaultRequestPool.GetInstance()
	request.SetHeader(name, pid, uid, senderAddress)
	request.SetRequestData(t, h, data)
	return request
}

func (m *ExtensionManager) DoRequest(extension IProtocolExtension, request IExtensionRequest, response IExtensionResponse) {
	// 响应处理
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
}

func (m *ExtensionManager) GetProtocolExtension(eName string) (pe IProtocolExtension, ok bool) {
	if pe, ok := m.ExtensionContainer.GetExtension(eName).(IProtocolExtension); ok {
		return pe, true
	}
	return nil, false
}
