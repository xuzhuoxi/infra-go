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
	InitManager(sock netx.ISockServer, container IProtocolContainer)

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

func NewExtensionManager() IExtensionManager {
	return &ExtensionManager{}
}

type ExtensionManager struct {
	sock         netx.ISockServer
	container    IProtocolContainer
	logger       logx.ILogger
	addressProxy netx.IAddressProxy
	mutex        sync.RWMutex
}

func (m *ExtensionManager) SetAddressProxy(proxy netx.IAddressProxy) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.addressProxy = proxy
}

func (m *ExtensionManager) InitManager(sock netx.ISockServer, container IProtocolContainer) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.sock, m.container = sock, container
}

func (m *ExtensionManager) SetLogger(logger logx.ILogger) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.logger = logger
}

func (m *ExtensionManager) StartManager() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container.InitExtensions()
	m.sock.GetPackHandler().AppendPackHandler(m.onPack)
}

func (m *ExtensionManager) StopManager() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.sock.GetPackHandler().ClearHandler(m.onPack)
	m.container.DestroyExtensions()
}

func (m *ExtensionManager) SaveExtensions() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container.SaveExtensions()
}

func (m *ExtensionManager) SaveExtension(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container.SaveExtension(name)
}

func (m *ExtensionManager) EnableExtension(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container.EnableExtension(name, true)
}

func (m *ExtensionManager) DisableExtension(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container.EnableExtension(name, false)
}

func (m *ExtensionManager) EnableExtensions() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container.EnableExtensions(true)
}

func (m *ExtensionManager) DisableExtensions() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.container.EnableExtensions(false)
}

//---------------------------------

//消息处理入口，这里是并发方法
//msgData非共享的，但在parsePackMessage后这部分数据会发生变化
func (m *ExtensionManager) onPack(msgData []byte, senderAddress string, other interface{}) bool {
	//m.logger.Infoln("ExtensionManager.onPack", senderAddress, msgData)
	name, pid, uid, data := m.parsePackMessage(msgData)
	extension := m.getProtocolExtension(name)

	//有效性验证
	if nil == extension {
		if nil != m.logger {
			m.logger.Warnln(fmt.Sprintf("Undefined Extension(%s)! Sender(%s)", name, uid))
		}
		return false
	}
	if !extension.CheckProtocolId(pid) { //有效性检查
		if nil != m.logger {
			m.logger.Warnln(fmt.Sprintf("Undefined ProtoId(%s) Send to Extension(%s)! Sender(%s)", pid, name, uid))
		}
		return false
	}
	//m.logger.Infoln("ExtensionManager.onPack:Params:", name, pid, uid, data)

	//构造响应参数
	t, h := extension.GetParamInfo(pid)
	response := DefaultResponsePool.GetInstance()
	response.SetHeader(name, pid, uid, senderAddress)
	response.SetSockServer(m.sock)
	response.SetAddressProxy(m.addressProxy)
	response.SetParamInfo(t, h)
	defer DefaultResponsePool.Recycle(response)
	request := DefaultRequestPool.GetInstance()
	request.SetHeader(name, pid, uid, senderAddress)
	request.SetRequestData(t, h, data)
	defer DefaultRequestPool.Recycle(request)

	//响应处理
	if be, ok := extension.(IBeforeRequestExtension); ok { //前置处理
		be.BeforeRequest(request)
	}
	if re, ok := extension.(IRequestExtension); ok {
		re.OnRequest(response, request)
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
func (m *ExtensionManager) parsePackMessage(msgBytes []byte) (name string, pid string, uid string, data [][]byte) {
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

func (m *ExtensionManager) getProtocolExtension(pid string) IProtocolExtension {
	e := m.container.GetExtension(pid)
	if pe, ok := e.(IProtocolExtension); ok {
		return pe
	}
	return nil
}
