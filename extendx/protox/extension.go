//
//Created by xuzhuoxi
//on 2019-02-12.
//@author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/extendx"
)

// Extension响应参数类型
type ExtensionParamType int

const (
	None ExtensionParamType = iota
	Binary
	Json
	Object
)

// Extension响应函数－无参数
type ExtensionHandlerNoneParam func(resp IExtensionResponse, req IExtensionRequest)

// Extension响应函数－二进制参数
type ExtensionHandlerBinaryParam func(resp IExtensionBinaryResponse, req IExtensionBinaryRequest)

// Extension响应函数－Json参数
type ExtensionHandlerJsonParam func(resp IExtensionJsonResponse, req IExtensionJsonRequest)

// Extension响应函数－具体对象参数
type ExtensionHandlerObjectParam func(resp IExtensionObjectResponse, req IExtensionObjectRequest)

type IProtocolExtension interface {
	extendx.IExtension
	extendx.IInitExtension
	// 检查ProtoId是否为被当前扩展支持
	CheckProtocolId(protoId string) bool
	// 检查ProtoId对应的设置
	GetParamInfo(protoId string) (paramType ExtensionParamType, handler IProtocolParamsHandler)
}

//-------------------------------------------------------------

type IGoroutineExtension interface {
	// 最大并发处理个数
	MaxGo() int
}

//-------------------------------------------------------------

type IBeforeRequestExtension interface {
	// 执行响应前的一些处理
	BeforeRequest(req IExtensionRequest)
}

type IRequestExtension interface {
	// 请求响应
	OnRequest(resp IExtensionResponse, req IExtensionRequest)
}

type IRequestExtensionSetter interface {
	// 设置请求响应处理(无参数)
	SetRequestHandler(protoId string, handler ExtensionHandlerNoneParam)
	// 设置请求响应处理(字节数组参数)
	SetRequestHandlerBinary(protoId string, handler ExtensionHandlerBinaryParam)
	// 设置请求响应处理(Json参数)
	SetRequestHandlerJson(protoId string, handler ExtensionHandlerJsonParam)
	// 设置请求响应处理(对象参数)
	SetRequestHandlerObject(protoId string, handler ExtensionHandlerObjectParam, paramOrigin interface{}, paramHandler IProtocolParamsHandler)
	// 清除设置
	ClearRequestHandler(protoId string)
}

type IAfterRequestExtension interface {
	// 响应结束前的一些处理
	AfterRequest(resp IExtensionResponse, req IExtensionRequest)
}
