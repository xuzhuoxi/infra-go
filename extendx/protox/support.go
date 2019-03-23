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
