// Package extendx
// Created by xuzhuoxi
// on 2019-02-17.
// @author xuzhuoxi
//
package extendx

type IExtension interface {
	// ExtensionName 主键标识
	ExtensionName() string
}

type IInitExtension interface {
	// InitExtension 初始化
	InitExtension() error
	// DestroyExtension 反初始化
	DestroyExtension() error
}

type ISaveExtension interface {
	// SaveExtension 保存数据
	SaveExtension() error
}

type IEnableExtension interface {
	// Enable 是否启用
	Enable() bool
	// EnableExtension 启用
	EnableExtension() error
	// DisableExtension 禁用
	DisableExtension() error
}
