//
//Created by xuzhuoxi
//on 2019-02-12.
//@author xuzhuoxi
//
package protox

import "github.com/xuzhuoxi/infra-go/extendx"

type RequestDataType uint16

const (
	None RequestDataType = iota
	ByteArray
	StructValue
)

type IProtocolExtension interface {
	extendx.IExtension
	//当前处理器所属ProtocolId
	ProtocolId() string
}

type IGoroutineExtension interface {
	//最大并发处理个数
	MaxGo() int
}

type IBatchExtension interface {
	//是否批量处理
	Batch() bool
}

type IRequestExtension interface {
	//响应结数据类型
	RequestDataType() RequestDataType
	//响应结构体
	RequestData() interface{}
	//请求响应
	OnRequest(pId string, data interface{}, data2 ...interface{})
}

type IBeforeRequestExtension interface {
	//执行响应前的一些处理
	BeforeRequest()
}

type IAfterRequestExtension interface {
	//响应结束前的一些处理
	AfterRequest()
}
