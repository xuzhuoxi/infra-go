// Package protox
// Created by xuzhuoxi
// on 2019-05-21.
// @author xuzhuoxi
//
package protox

// FuncStartOnPack
// 响应入口
type FuncStartOnPack func(senderAddress string)

// FuncParseMessage
// 解释二进制数据
type FuncParseMessage func(msgBytes []byte) (name string, pid string, uid string, data [][]byte)

// FuncVerify
// 消息处理入口，这里是并发方法
type FuncVerify func(name string, pid string, uid string) (e IProtocolExtension, ok bool)

// FuncStartOnRequest
// 响应开始
type FuncStartOnRequest func(resp IExtensionResponse, req IExtensionRequest)

// FuncFinishOnRequest
// 响应完成
type FuncFinishOnRequest func(resp IExtensionResponse, req IExtensionRequest)

type IExtensionManagerCustomizeSetting interface {
	// SetCustomStartOnPackFunc
	// 设置自定义响应开始行为
	SetCustomStartOnPackFunc(funcStartOnPack FuncStartOnPack)
	// SetCustomParseFunc
	// 设置自定义数据解释行为
	SetCustomParseFunc(funcParse FuncParseMessage)
	// SetCustomVerifyFunc
	// 设置自定义验证
	SetCustomVerifyFunc(funcVerify FuncVerify)
	// SetCustomStartOnRequestFunc
	// 设置自定义响应前置行为
	SetCustomStartOnRequestFunc(funcStart FuncStartOnRequest)
	// SetCustomFinishOnRequestFunc
	// 设置自定义响应完成行为
	SetCustomFinishOnRequestFunc(funcFinish FuncFinishOnRequest)
	// SetCustom
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

func (o *ExtensionManagerCustomizeSupport) SetCustomStartOnPackFunc(funcStartOnPack FuncStartOnPack) {
	o.FuncStartOnPack = funcStartOnPack
}
func (o *ExtensionManagerCustomizeSupport) SetCustomParseFunc(funcParse FuncParseMessage) {
	o.FuncParseMessage = funcParse
}
func (o *ExtensionManagerCustomizeSupport) SetCustomVerifyFunc(funcVerify FuncVerify) {
	o.FuncVerify = funcVerify
}
func (o *ExtensionManagerCustomizeSupport) SetCustomStartOnRequestFunc(funcStart FuncStartOnRequest) {
	o.FuncStartOnRequest = funcStart
}
func (o *ExtensionManagerCustomizeSupport) SetCustomFinishOnRequestFunc(funcFinish FuncFinishOnRequest) {
	o.FuncFinishOnRequest = funcFinish
}
func (o *ExtensionManagerCustomizeSupport) SetCustom(funcStartOnPack FuncStartOnPack, funcParse FuncParseMessage, funcVerify FuncVerify, funcStart FuncStartOnRequest, funcFinish FuncFinishOnRequest) {
	o.FuncStartOnPack, o.FuncParseMessage, o.FuncVerify, o.FuncStartOnRequest, o.FuncFinishOnRequest = funcStartOnPack, funcParse, funcVerify, funcStart, funcFinish
}

func (o *ExtensionManagerCustomizeSupport) CustomStartOnPack(senderAddress string) {
	if nil != o.FuncStartOnPack {
		o.FuncStartOnPack(senderAddress)
	}
}
func (o *ExtensionManagerCustomizeSupport) CustomParseMessage(msgBytes []byte) (name string, pid string, uid string, data [][]byte) {
	if nil != o.FuncParseMessage {
		return o.FuncParseMessage(msgBytes)
	}
	return
}
func (o *ExtensionManagerCustomizeSupport) CustomVerify(name string, pid string, uid string) (e IProtocolExtension, ok bool) {
	if nil != o.FuncVerify {
		return o.FuncVerify(name, pid, uid)
	}
	return
}
func (o *ExtensionManagerCustomizeSupport) CustomStartOnRequest(resp IExtensionResponse, req IExtensionRequest) {
	if nil != o.FuncStartOnRequest {
		o.FuncStartOnRequest(resp, req)
	}
}
func (o *ExtensionManagerCustomizeSupport) CustomFinishOnRequest(resp IExtensionResponse, req IExtensionRequest) {
	if nil != o.FuncFinishOnRequest {
		o.FuncFinishOnRequest(resp, req)
	}
}
