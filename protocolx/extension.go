//
//Created by xuzhuoxi
//on 2019-02-12.
//@author xuzhuoxi
//
package protocolx

type HandleOnce func(pId string, data interface{})                        //单个独立性处理
type HandleMulti func(pId string, data interface{}, data2 ...interface{}) //批量处理

type IProtocolExtension interface {
	//当前处理器所属ProtocolId
	ProtocolId() string
	//请求响应
	OnRequest(pId string, data interface{}, data2 ...interface{})
}

func NewProtocolExtensionMulti(pId string, maxGo int, multi HandleMulti) IProtocolExtension {
	return newProtocolExtension(pId, maxGo, nil, multi)
}
func NewProtocolExtension(pId string, maxGo int, once HandleOnce) IProtocolExtension {
	return newProtocolExtension(pId, maxGo, once, nil)
}

//-----------------------------------------------------

func newProtocolExtension(pId string, maxGo int, once HandleOnce, multi HandleMulti) *protocolExtension {
	return &protocolExtension{protocolId: pId, channel: make(chan bool, maxGo), onceHandler: once, multiHandler: multi}
}

type protocolExtension struct {
	protocolId string
	channel    chan bool

	onceHandler  HandleOnce
	multiHandler HandleMulti
}

func (e *protocolExtension) ProtocolId() string {
	return e.protocolId
}

func (e *protocolExtension) OnRequest(pId string, data interface{}, data2 ...interface{}) {
	e.addGoroutine()
	go func() {
		defer e.doneGoroutine()
		if nil != e.multiHandler {
			e.multiHandler(pId, data, data2...)
		} else {
			e.onceHandler(pId, data)
			len2 := len(data2)
			if len2 > 0 {
				for index := 0; index < len2; index++ {
					e.onceHandler(pId, data2[index])
				}
			}
		}
	}()
}

func (e *protocolExtension) addGoroutine() {
	e.channel <- true
}

func (e *protocolExtension) doneGoroutine() {
	<-e.channel
}
