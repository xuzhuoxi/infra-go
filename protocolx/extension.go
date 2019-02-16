//
//Created by xuzhuoxi
//on 2019-02-12.
//@author xuzhuoxi
//
package protocolx

import "github.com/xuzhuoxi/infra-go/extendx"

type HandlerProtocol func(pId string, data interface{})                            //单个独立性处理
type HandlerProtocolBatch func(pId string, data interface{}, data2 ...interface{}) //批量处理

type IProtocolExtension interface {
	extendx.IExtension
	//当前处理器所属ProtocolId
	ProtocolId() string
	//是否批量处理
	Batch() bool
	//请求响应
	HandleRequest(pId string, data interface{}, data2 ...interface{})
}

type IProtocolContainer interface {
	extendx.IExtensionContainer
	//增加ProtocolId到Handler的表映射
	AppendProtocolExtension(protocolId string, maxGoroutine int, handler HandlerProtocol)
	//增加ProtocolId到Handler的表映射
	AppendProtocolExtensionBatch(protocolId string, maxGoroutine int, handler HandlerProtocolBatch)
}

func NewProtocolExtensionMulti(pId string, maxGo int, multi HandlerProtocolBatch) IProtocolExtension {
	return newProtocolExtension(pId, maxGo, nil, multi)
}

func NewProtocolExtension(pId string, maxGo int, once HandlerProtocol) IProtocolExtension {
	return newProtocolExtension(pId, maxGo, once, nil)
}

func NewProtocolExtensionContainer() IProtocolContainer {
	return &ProtocolContainer{IExtensionContainer: extendx.NewExtensionContainer()}
}

//-----------------------------------------------------

type ProtocolContainer struct {
	extendx.IExtensionContainer
}

func (c *ProtocolContainer) AppendProtocolExtension(protocolId string, maxGoroutine int, handler HandlerProtocol) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, maxGoroutine, handler, nil))
}

func (c *ProtocolContainer) AppendProtocolExtensionBatch(protocolId string, maxGoroutine int, handler HandlerProtocolBatch) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, maxGoroutine, nil, handler))
}

//-----------------------------------------------------

func newProtocolExtension(pId string, maxGo int, once HandlerProtocol, multi HandlerProtocolBatch) *ProtocolExtension {
	return &ProtocolExtension{protocolId: pId, channel: make(chan bool, maxGo), protoHandler: once, batchHandler: multi}
}

type ProtocolExtension struct {
	protocolId string
	channel    chan bool

	protoHandler HandlerProtocol
	batchHandler HandlerProtocolBatch
}

func (e *ProtocolExtension) Key() string {
	return e.protocolId
}

func (e *ProtocolExtension) ProtocolId() string {
	return e.protocolId
}

func (e *ProtocolExtension) Batch() bool {
	return nil != e.batchHandler
}

func (e *ProtocolExtension) HandleRequest(pId string, data interface{}, data2 ...interface{}) {
	e.addGoroutine()
	go func() {
		defer e.doneGoroutine()
		if e.Batch() {
			e.batchHandler(pId, data, data2...)
		} else {
			e.protoHandler(pId, data)
			len2 := len(data2)
			if len2 > 0 {
				for index := 0; index < len2; index++ {
					e.protoHandler(pId, data2[index])
				}
			}
		}
	}()
}

func (e *ProtocolExtension) addGoroutine() {
	e.channel <- true
}

func (e *ProtocolExtension) doneGoroutine() {
	<-e.channel
}
