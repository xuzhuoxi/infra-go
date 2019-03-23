//
//Created by xuzhuoxi
//on 2019-02-26.
//@author xuzhuoxi
//
package protox

import "github.com/xuzhuoxi/infra-go/extendx"

type ProtocolHandler func(protoId string, uid string, data []byte)                       //单个独立性处理
type ProtocolBatchHandler func(protoId string, uid string, data []byte, data2 ...[]byte) //批量处理

func NewHandlerExtensionBatch(protoId string, batchHandler ProtocolBatchHandler) IProtocolExtension {
	return newProtocolExtension(protoId, nil, batchHandler)
}

func NewHandlerExtension(protoId string, handler ProtocolHandler) IProtocolExtension {
	return newProtocolExtension(protoId, handler, nil)
}

func newProtocolExtension(protoId string, once ProtocolHandler, multi ProtocolBatchHandler) *ProtocolHandlerExtension {
	return &ProtocolHandlerExtension{protocolId: protoId, protoHandler: once, batchHandler: multi}
}

type ProtocolHandlerExtension struct {
	protocolId string

	protoHandler ProtocolHandler
	batchHandler ProtocolBatchHandler
}

func (e *ProtocolHandlerExtension) ExtensionName() string {
	return e.protocolId
}

func (e *ProtocolHandlerExtension) InitProtocolId() {
	return
}

func (e *ProtocolHandlerExtension) CheckProtocolId(ProtocolId string) bool {
	return e.protocolId == ProtocolId
}

func (e *ProtocolHandlerExtension) Batch() bool {
	return nil != e.batchHandler
}

func (e *ProtocolHandlerExtension) OnRequest(resp extendx.IExtensionResponse, protoId string, uid string, data []byte, data2 ...[]byte) {
	go func() {
		if e.Batch() {
			e.batchHandler(protoId, uid, data, data2...)
		} else {
			e.protoHandler(protoId, uid, data)
			len2 := len(data2)
			if len2 > 0 {
				for index := 0; index < len2; index++ {
					e.protoHandler(protoId, uid, data2[index])
				}
			}
		}
	}()
}
