//
//Created by xuzhuoxi
//on 2019-02-13.
//@author xuzhuoxi
//
package protocolx

type IExtensionTable interface {
	//增加ProtocolId到Handler的表映射
	MapProtocolHandler(protocolId string, handler IProtocolExtension)
	//增加ProtocolId到Handler的表映射
	MapProtocolMultiHandler(protocolId string, maxGoroutine int, handler HandleMulti)
	//增加ProtocolId到Handler的表映射
	MapProtocolOnceHandler(protocolId string, maxGoroutine int, handler HandleOnce)
	//取ProtocolHandler
	GetProtocolHandler(protocolId string) IProtocolExtension
}

func NewExtensionTable() IExtensionTable {
	return &extensionTable{handlerMap: make(map[string]IProtocolExtension)}
}

type extensionTable struct {
	handlerMap map[string]IProtocolExtension
}

func (m *extensionTable) MapProtocolHandler(protocolId string, handler IProtocolExtension) {
	if m.hasMap(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	if protocolId != handler.ProtocolId() {
		panic("Uncertain ProtocolId : " + protocolId + " and " + handler.ProtocolId())
	}
	m.handlerMap[protocolId] = handler
}

func (m *extensionTable) MapProtocolMultiHandler(protocolId string, maxGoroutine int, handler HandleMulti) {
	if m.hasMap(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	m.handlerMap[protocolId] = newProtocolExtension(protocolId, maxGoroutine, nil, handler)
}

func (m *extensionTable) MapProtocolOnceHandler(protocolId string, maxGoroutine int, handler HandleOnce) {
	if m.hasMap(protocolId) {
		panic("Repeat ProtocolId In Map: " + protocolId)
	}
	m.handlerMap[protocolId] = newProtocolExtension(protocolId, maxGoroutine, handler, nil)
}

func (m *extensionTable) GetProtocolHandler(protocolId string) IProtocolExtension {
	if !m.hasMap(protocolId) {
		return nil
	}
	rs, _ := m.handlerMap[protocolId]
	return rs
}

func (m *extensionTable) hasMap(protocolId string) bool {
	_, ok := m.handlerMap[protocolId]
	return ok
}
