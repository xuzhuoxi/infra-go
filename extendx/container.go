//
//Created by xuzhuoxi
//on 2019-02-17.
//@author xuzhuoxi
//
package extendx

type IExtensionContainer interface {
	//增加Extension
	AppendExtension(extension IExtension)
	//检查
	CheckExtension(key string) bool
	//取Extension
	GetExtension(key string) IExtension
	//列表
	Extensions() []IExtension
	//按列表处理
	Range(handler func(index int, extension IExtension))
}

func NewExtensionContainer() IExtensionContainer {
	return &ExtensionContainer{extensionMap: make(map[string]IExtension)}
}

type ExtensionContainer struct {
	extensions   []IExtension
	extensionMap map[string]IExtension
}

func (m *ExtensionContainer) AppendExtension(extension IExtension) {
	key := extension.Key()
	if m.hasMap(key) {
		panic("Repeat Key In Map: " + key)
	}
	m.extensionMap[key] = extension
	m.extensions = append(m.extensions, extension)
}

func (m *ExtensionContainer) CheckExtension(key string) bool {
	_, ok := m.extensionMap[key]
	return ok
}

func (m *ExtensionContainer) GetExtension(key string) IExtension {
	if !m.hasMap(key) {
		return nil
	}
	rs, _ := m.extensionMap[key]
	return rs
}

func (m *ExtensionContainer) Extensions() []IExtension {
	cp := make([]IExtension, len(m.extensions))
	copy(cp, m.extensions)
	return cp
}

func (m *ExtensionContainer) Range(handler func(index int, extension IExtension)) {
	for index, extension := range m.extensions {
		handler(index, extension)
	}
}

func (m *ExtensionContainer) hasMap(key string) bool {
	_, ok := m.extensionMap[key]
	return ok
}
