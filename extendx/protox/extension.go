//
//Created by xuzhuoxi
//on 2019-02-12.
//@author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/extendx"
)

type IProtocolExtension interface {
	extendx.IExtension
	//初始化支持的ProtocolId
	InitProtocolId()
	//检查ProtoId是否为被当前扩展支持
	CheckProtocolId(protoId string) bool
}

type IGoroutineExtension interface {
	//最大并发处理个数
	MaxGo() int
}

type IBatchExtension interface {
	//是否批量处理
	Batch() bool
}

type IBeforeRequestExtension interface {
	//执行响应前的一些处理
	BeforeRequest(protoId string)
}

type IOnNoneRequestExtension interface {
	//请求响应
	OnRequest(resp extendx.IExtensionResponse, protoId string, uid string)
}

type IOnBinaryRequestExtension interface {
	//请求响应
	OnRequest(resp extendx.IExtensionBinaryResponse, protoId string, uid string, data []byte, data2 ...[]byte)
}

type IOnObjectRequestExtension interface {
	//响应结构体
	GetRequestData(pProtoId string) (dataCopy interface{})
	//请求响应
	OnRequest(resp extendx.IExtensionObjectResponse, protoId string, uid string, data interface{}, data2 ...interface{})
}

type IAfterRequestExtension interface {
	//响应结束前的一些处理
	AfterRequest(protoId string)
}
