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
	AppendProtocolHandlerExtension(protocolId string, handler ProtocolHandler)
	//增加ProtocolId到Handler的表映射
	AppendProtocolHandlerExtensionBatch(protocolId string, handler ProtocolBatchHandler)
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

func (c *ProtocolContainer) AppendProtocolHandlerExtension(protocolId string, handler ProtocolHandler) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, handler, nil))
}

func (c *ProtocolContainer) AppendProtocolHandlerExtensionBatch(protocolId string, handler ProtocolBatchHandler) {
	if c.CheckExtension(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	c.AppendExtension(newProtocolExtension(protocolId, nil, handler))
}
