//
//Created by xuzhuoxi
//on 2019-02-17.
//@author xuzhuoxi
//
package extendx

import "github.com/pkg/errors"

type IExtensionContainer interface {
	//增加Extension
	AppendExtension(extension IExtension)
	//检查
	CheckExtension(key string) bool
	//取Extension
	GetExtension(key string) IExtension
	//Extension数量
	Len() int
	//列表
	Extensions() []IExtension
	//反向列表
	ExtensionsReversed() []IExtension
	//按列表处理
	Range(handler func(index int, extension IExtension))
	//按反向列表处理
	RangeReverse(handler func(index int, extension IExtension))
	//对指定Extension执行处理
	HandleAt(index int, handler func(index int, extension IExtension)) error
	//对指定Extension执行处理
	HandleAtKey(key string, handler func(key string, extension IExtension)) error
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

func (m *ExtensionContainer) Len() int {
	return len(m.extensions)
}

func (m *ExtensionContainer) Extensions() []IExtension {
	ln := len(m.extensions)
	if 0 == ln {
		return nil
	}
	cp := make([]IExtension, ln)
	copy(cp, m.extensions)
	return cp
}

func (m *ExtensionContainer) ExtensionsReversed() []IExtension {
	ln := len(m.extensions)
	if 0 == ln {
		return nil
	}
	cp := make([]IExtension, ln)
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		cp[i], cp[j] = m.extensions[j], m.extensions[i]
	}
	return cp
}

func (m *ExtensionContainer) Range(handler func(index int, extension IExtension)) {
	for index, extension := range m.extensions {
		handler(index, extension)
	}
}

func (m *ExtensionContainer) RangeReverse(handler func(index int, extension IExtension)) {
	ln := len(m.extensions)
	for index := ln - 1; index >= 0; index-- {
		handler(index, m.extensions[index])
	}
}

func (m *ExtensionContainer) HandleAt(index int, handler func(index int, extension IExtension)) error {
	if index < 0 || index >= len(m.extensions) {
		return errors.New("HandleAt Error : Out of index!")
	}
	handler(index, m.extensions[index])
	return nil
}

func (m *ExtensionContainer) HandleAtKey(key string, handler func(key string, extension IExtension)) error {
	if !m.CheckExtension(key) {
		return errors.New("HandleAtKey Error : No such key [" + key + "]")
	}
	handler(key, m.extensionMap[key])
	return nil
}

func (m *ExtensionContainer) hasMap(key string) bool {
	_, ok := m.extensionMap[key]
	return ok
}
