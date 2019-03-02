//
//Created by xuzhuoxi
//on 2019-02-26.
//@author xuzhuoxi
//
package protox

func NewProtocolExtensionSupport(Name string) ProtocolExtensionSupport {
	return ProtocolExtensionSupport{Name: Name, ProtoIdToValue: make(map[string]interface{})}
}

type ProtocolExtensionSupport struct {
	Name           string
	ProtoIdToValue map[string]interface{}
}

func (s *ProtocolExtensionSupport) ExtensionName() string {
	return s.Name
}

func (s *ProtocolExtensionSupport) CheckProtocolId(ProtoId string) bool {
	_, ok := s.ProtoIdToValue[ProtoId]
	return ok
}

//---------------------------------------

func NewGoroutineExtensionSupport(ProtoId string, MaxGo int) GoroutineExtensionSupport {
	return GoroutineExtensionSupport{ProtoId: ProtoId, MaxGoroutine: MaxGo}
}

type GoroutineExtensionSupport struct {
	ProtoId      string
	MaxGoroutine int
}

func (s *GoroutineExtensionSupport) Key() string {
	return s.ProtoId
}

func (s *GoroutineExtensionSupport) ProtocolId() string {
	return s.ProtoId
}

func (s *GoroutineExtensionSupport) MaxGo() int {
	return s.MaxGoroutine
}

//type ProtocolExtensionSupport struct {
//	Channel chan bool
//}
//
//func (e *ProtocolExtensionSupport) MaxGo() int {
//	fmt.Println("ProtocolExtensionSupport.MaxGo")
//	return 1
//}
//
//func (e *ProtocolExtensionSupport) Batch() bool {
//	panic("implement me")
//}
//
//func (e *ProtocolExtensionSupport) OnRequest(pId string, data interface{}, data2 ...interface{}) {
//	panic("implement me")
//}
//
//func (e *ProtocolExtensionSupport) HandleRequest(pId string, data interface{}, data2 ...interface{}) {
//	if e.MaxGo() > 1 {
//		e.addGoroutine()
//		go func() {
//			defer e.doneGoroutine()
//			e.doOnRequest(pId, data, data2...)
//		}()
//	} else {
//		e.doOnRequest(pId, data, data2...)
//	}
//}
//
//func (e *ProtocolExtensionSupport) doOnRequest(pId string, data interface{}, data2 ...interface{}) {
//	if e.Batch() {
//		e.OnRequest(pId, data, data2...)
//	} else {
//		e.OnRequest(pId, data)
//		len2 := len(data2)
//		if len2 > 0 {
//			for index := 0; index < len2; index++ {
//				e.OnRequest(pId, data2[index])
//			}
//		}
//	}
//}
//
//func (e *ProtocolExtensionSupport) addGoroutine() {
//	if nil == e.Channel {
//		e.Channel = make(chan bool, e.MaxGo())
//	}
//	e.Channel <- true
//}
//
//func (e *ProtocolExtensionSupport) doneGoroutine() {
//	if nil != e.Channel {
//		<-e.Channel
//	}
//}
