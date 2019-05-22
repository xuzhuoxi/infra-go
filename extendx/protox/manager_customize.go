//
//Created by xuzhuoxi
//on 2019-05-21.
//@author xuzhuoxi
//
package protox

// 响应入口
type FuncStartOnPack func(senderAddress string)

// 解释二进制数据
type FuncParseMessage func(msgBytes []byte) (name string, pid string, uid string, data [][]byte)

// 消息处理入口，这里是并发方法
type FuncVerify func(name string, pid string, uid string) (e IProtocolExtension, ok bool)

// 响应开始
type FuncStartOnRequest func(resp IExtensionResponse, req IExtensionRequest)

// 响应完成
type FuncFinishOnRequest func(resp IExtensionResponse, req IExtensionRequest)

type IExtensionManagerCustomizeSetting interface {
	// 设置自定义响应开始行为
	SetCustomStartOnPackFunc(funcStartOnPack FuncStartOnPack)
	// 设置自定义数据解释行为
	SetCustomParseFunc(funcParse FuncParseMessage)
	// 设置自定义验证
	SetCustomVerifyFunc(funcVerify FuncVerify)
	// 设置自定义响应前置行为
	SetCustomStartOnRequestFunc(funcStart FuncStartOnRequest)
	// 设置自定义响应完成行为
	SetCustomFinishOnRequestFunc(funcFinish FuncFinishOnRequest)
	// 设置自定义行为
	SetCustom(funcStartOnPack FuncStartOnPack, funcParse FuncParseMessage, funcVerify FuncVerify, funcStart FuncStartOnRequest, funcFinish FuncFinishOnRequest)
}

type IExtensionManagerCustomizeSupport interface {
	CustomStartOnPack(senderAddress string)
	CustomParseMessage(msgBytes []byte) (name string, pid string, uid string, data [][]byte)
	CustomVerify(name string, pid string, uid string) (e IProtocolExtension, ok bool)
	CustomStartOnRequest(resp IExtensionResponse, req IExtensionRequest)
	CustomFinishOnRequest(resp IExtensionResponse, req IExtensionRequest)
}

type ExtensionManagerCustomizeSupport struct {
	FuncStartOnPack     FuncStartOnPack
	FuncParseMessage    FuncParseMessage
	FuncVerify          FuncVerify
	FuncStartOnRequest  FuncStartOnRequest
	FuncFinishOnRequest FuncFinishOnRequest
}

func (m *ExtensionManagerCustomizeSupport) SetCustomStartOnPackFunc(funcStartOnPack FuncStartOnPack) {
	m.FuncStartOnPack = funcStartOnPack
}
func (m *ExtensionManagerCustomizeSupport) SetCustomParseFunc(funcParse FuncParseMessage) {
	m.FuncParseMessage = funcParse
}
func (m *ExtensionManagerCustomizeSupport) SetCustomVerifyFunc(funcVerify FuncVerify) {
	m.FuncVerify = funcVerify
}
func (m *ExtensionManagerCustomizeSupport) SetCustomStartOnRequestFunc(funcStart FuncStartOnRequest) {
	m.FuncStartOnRequest = funcStart
}
func (m *ExtensionManagerCustomizeSupport) SetCustomFinishOnRequestFunc(funcFinish FuncFinishOnRequest) {
	m.FuncFinishOnRequest = funcFinish
}
func (m *ExtensionManagerCustomizeSupport) SetCustom(funcStartOnPack FuncStartOnPack, funcParse FuncParseMessage, funcVerify FuncVerify, funcStart FuncStartOnRequest, funcFinish FuncFinishOnRequest) {
	m.FuncStartOnPack, m.FuncParseMessage, m.FuncVerify, m.FuncStartOnRequest, m.FuncFinishOnRequest = funcStartOnPack, funcParse, funcVerify, funcStart, funcFinish
}

func (s *ExtensionManagerCustomizeSupport) CustomStartOnPack(senderAddress string) {
	if nil != s.FuncStartOnPack {
		s.FuncStartOnPack(senderAddress)
	}
}
func (s *ExtensionManagerCustomizeSupport) CustomParseMessage(msgBytes []byte) (name string, pid string, uid string, data [][]byte) {
	if nil != s.FuncParseMessage {
		return s.FuncParseMessage(msgBytes)
	}
	return
}
func (s *ExtensionManagerCustomizeSupport) CustomVerify(name string, pid string, uid string) (e IProtocolExtension, ok bool) {
	if nil != s.FuncVerify {
		return s.FuncVerify(name, pid, uid)
	}
	return
}
func (s *ExtensionManagerCustomizeSupport) CustomStartOnRequest(resp IExtensionResponse, req IExtensionRequest) {
	if nil != s.FuncStartOnRequest {
		s.FuncStartOnRequest(resp, req)
	}
}
func (s *ExtensionManagerCustomizeSupport) CustomFinishOnRequest(resp IExtensionResponse, req IExtensionRequest) {
	if nil != s.FuncFinishOnRequest {
		s.FuncFinishOnRequest(resp, req)
	}
}
