//
//Created by xuzhuoxi
//on 2019-02-26.
//@author xuzhuoxi
//
package protox

type ProtocolHandler func(protoId string, data interface{})                            //单个独立性处理
type ProtocolBatchHandler func(protoId string, data interface{}, data2 ...interface{}) //批量处理

func NewHandlerExtensionBatch(protoId string, maxGo int, batchHandler ProtocolBatchHandler, reqData interface{}) IProtocolExtension {
	return newProtocolExtension(protoId, maxGo, nil, batchHandler, reqData)
}

func NewHandlerExtension(protoId string, maxGo int, handler ProtocolHandler, reqData interface{}) IProtocolExtension {
	return newProtocolExtension(protoId, maxGo, handler, nil, reqData)
}

func newProtocolExtension(protoId string, maxGo int, once ProtocolHandler, multi ProtocolBatchHandler, reqData interface{}) *ProtocolHandlerExtension {
	return &ProtocolHandlerExtension{protocolId: protoId, maxGo: maxGo, channel: make(chan struct{}, maxGo), protoHandler: once, batchHandler: multi, reqData: reqData}
}

type ProtocolHandlerExtension struct {
	protocolId string
	maxGo      int
	channel    chan struct{}

	protoHandler ProtocolHandler
	batchHandler ProtocolBatchHandler
	reqData      interface{}
}

func (e *ProtocolHandlerExtension) ExtensionName() string {
	return e.protocolId
}

func (e *ProtocolHandlerExtension) InitProtocolId() {
}

func (e *ProtocolHandlerExtension) CheckProtocolId(ProtocolId string) bool {
	return e.protocolId == ProtocolId
}

func (e *ProtocolHandlerExtension) Batch() bool {
	return nil != e.batchHandler
}

func (e *ProtocolHandlerExtension) MaxGo() int {
	return e.maxGo
}

func (e *ProtocolHandlerExtension) RequestDataType() RequestDataType {
	if nil == e.reqData {
		return None
	}
	if _, ok := e.reqData.([]byte); ok {
		return ByteArray
	}
	return StructValue
}

func (e *ProtocolHandlerExtension) RequestData() interface{} {
	rs := e.reqData
	return rs
}

func (e *ProtocolHandlerExtension) OnRequest(pId string, data interface{}, data2 ...interface{}) {
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

func (e *ProtocolHandlerExtension) addGoroutine() {
	e.channel <- struct{}{}
}

func (e *ProtocolHandlerExtension) doneGoroutine() {
	<-e.channel
}
