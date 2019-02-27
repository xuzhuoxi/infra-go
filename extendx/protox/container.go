//
//Created by xuzhuoxi
//on 2019-02-26.
//@author xuzhuoxi
//
package protox

import "github.com/xuzhuoxi/infra-go/extendx"

type IProtocolContainer interface {
	extendx.IExtensionContainer
	//增加ProtocolId到Handler的表映射
	AppendProtocolExtension(protocolId string, maxGoroutine int, handler ProtocolHandler, reqData interface{})
	//增加ProtocolId到Handler的表映射
	AppendProtocolExtensionBatch(protocolId string, maxGoroutine int, handler ProtocolBatchHandler, reqData interface{})
}

func NewProtocolExtensionContainer() IProtocolContainer {
	return &ProtoContainer{IExtensionContainer: extendx.NewExtensionContainer()}
}

type ProtoContainer struct {
	extendx.IExtensionContainer
}

func (c *ProtoContainer) AppendProtocolExtension(protocolId string, maxGoroutine int, handler ProtocolHandler, reqData interface{}) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, maxGoroutine, handler, nil, reqData))
}

func (c *ProtoContainer) AppendProtocolExtensionBatch(protocolId string, maxGoroutine int, handler ProtocolBatchHandler, reqData interface{}) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, maxGoroutine, nil, handler, reqData))
}
