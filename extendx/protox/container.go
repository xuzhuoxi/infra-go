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
	AppendProtocolHandlerExtension(protocolId string, maxGoroutine int, handler ProtocolHandler, reqData interface{})
	//增加ProtocolId到Handler的表映射
	AppendProtocolHandlerExtensionBatch(protocolId string, maxGoroutine int, handler ProtocolBatchHandler, reqData interface{})
}

func NewIProtocolExtensionContainer() IProtocolContainer {
	return &ProtocolContainer{ExtensionContainer: extendx.NewExtensionContainer()}
}

func NewProtocolExtensionContainer() ProtocolContainer {
	return ProtocolContainer{ExtensionContainer: extendx.NewExtensionContainer()}
}

type ProtocolContainer struct {
	extendx.ExtensionContainer
}

func (c *ProtocolContainer) AppendProtocolHandlerExtension(protocolId string, maxGoroutine int, handler ProtocolHandler, reqData interface{}) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, maxGoroutine, handler, nil, reqData))
}

func (c *ProtocolContainer) AppendProtocolHandlerExtensionBatch(protocolId string, maxGoroutine int, handler ProtocolBatchHandler, reqData interface{}) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, maxGoroutine, nil, handler, reqData))
}
