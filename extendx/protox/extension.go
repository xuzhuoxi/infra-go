// Package protox
// Created by xuzhuoxi
// on 2019-02-12.
// @author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/extendx"
)

// ExtensionParamType
// Extension响应参数类型
type ExtensionParamType int

const (
	None ExtensionParamType = iota
	Binary
	Json
	Object
)

// ExtensionHandlerNoneParam
// Extension响应函数－无参数
type ExtensionHandlerNoneParam func(resp IExtensionResponse, req IExtensionRequest)

// ExtensionHandlerBinaryParam
// Extension响应函数－二进制参数
type ExtensionHandlerBinaryParam func(resp IExtensionBinaryResponse, req IExtensionBinaryRequest)

// ExtensionHandlerJsonParam
// Extension响应函数－Json参数
type ExtensionHandlerJsonParam func(resp IExtensionJsonResponse, req IExtensionJsonRequest)

// ExtensionHandlerObjectParam
// Extension响应函数－具体对象参数
type ExtensionHandlerObjectParam func(resp IExtensionObjectResponse, req IExtensionObjectRequest)

type IProtocolExtension interface {
	extendx.IExtension
	extendx.IInitExtension
	// CheckProtocolId
	// 检查ProtoId是否为被当前扩展支持
	CheckProtocolId(protoId string) bool
	// GetParamInfo
	// 检查ProtoId对应的设置
	GetParamInfo(protoId string) (paramType ExtensionParamType, handler IProtocolParamsHandler)
}

//-------------------------------------------------------------

type IGoroutineExtension interface {
	// MaxGo
	// 最大并发处理个数
	MaxGo() int
}

//-------------------------------------------------------------

type IBeforeRequestExtension interface {
	// BeforeRequest
	// 执行响应前的一些处理
	BeforeRequest(req IExtensionRequest)
}

type IRequestExtension interface {
	// OnRequest
	// 请求响应
	OnRequest(resp IExtensionResponse, req IExtensionRequest)
}

type IRequestExtensionSetter interface {
	// SetRequestHandler
	// 设置请求响应处理(无参数)
	SetRequestHandler(protoId string, handler ExtensionHandlerNoneParam)
	// SetRequestHandlerBinary
	// 设置请求响应处理(字节数组参数)
	SetRequestHandlerBinary(protoId string, handler ExtensionHandlerBinaryParam)
	// SetRequestHandlerJson
	// 设置请求响应处理(Json参数)
	SetRequestHandlerJson(protoId string, handler ExtensionHandlerJsonParam)
	// SetRequestHandlerObject
	//设置请求响应处理(对象参数)
	SetRequestHandlerObject(protoId string, handler ExtensionHandlerObjectParam, paramOrigin interface{}, paramHandler IProtocolParamsHandler)
	// ClearRequestHandler
	// 清除设置
	ClearRequestHandler(protoId string)
}

type IAfterRequestExtension interface {
	// AfterRequest
	// 响应结束前的一些处理
	AfterRequest(resp IExtensionResponse, req IExtensionRequest)
}
