//
//Created by xuzhuoxi
//on 2019-02-17.
//@author xuzhuoxi
//
package extendx

import "errors"

type IExtensionContainer interface {
	//增加Extension
	AppendExtension(extension IExtension)
	//检查
	CheckExtension(name string) bool
	//取Extension
	GetExtension(name string) IExtension
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
	HandleAtName(name string, handler func(name string, extension IExtension)) error
}

func NewIExtensionContainer() IExtensionContainer {
	return &ExtensionContainer{extensionMap: make(map[string]IExtension)}
}

func NewExtensionContainer() ExtensionContainer {
	return ExtensionContainer{extensionMap: make(map[string]IExtension)}
}

type ExtensionContainer struct {
	extensions   []IExtension
	extensionMap map[string]IExtension
}

func (m *ExtensionContainer) AppendExtension(extension IExtension) {
	name := extension.ExtensionName()
	if m.checkMap(name) {
		panic("Repeat Name In Map: " + name)
	}
	m.extensionMap[name] = extension
	m.extensions = append(m.extensions, extension)
}

func (m *ExtensionContainer) CheckExtension(name string) bool {
	_, ok := m.extensionMap[name]
	return ok
}

func (m *ExtensionContainer) GetExtension(name string) IExtension {
	if !m.checkMap(name) {
		return nil
	}
	rs, _ := m.extensionMap[name]
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

func (m *ExtensionContainer) HandleAtName(name string, handler func(name string, extension IExtension)) error {
	if !m.CheckExtension(name) {
		return errors.New("HandleAtName Error : No such name [" + name + "]")
	}
	handler(name, m.extensionMap[name])
	return nil
}

func (m *ExtensionContainer) checkMap(name string) bool {
	_, ok := m.extensionMap[name]
	return ok
}
