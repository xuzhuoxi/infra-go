//
//Created by xuzhuoxi
//on 2019-02-26.
//@author xuzhuoxi
//
package protox

import (
	"github.com/xuzhuoxi/infra-go/extendx"
)

type IProtocolExtensionContainer interface {
	extendx.IExtensionContainer
	// 初始化全部Extension
	InitExtensions() []error
	// 反初始化全部Extension
	DestroyExtensions() []error

	// 保存
	SaveExtensions() []error
	// 保存指定
	SaveExtension(name string) error

	// 设置启用全部Extension
	EnableExtensions(enable bool) []error
	// 设置启用Extension
	EnableExtension(name string, enable bool) error
}

func NewIProtocolExtensionContainer() IProtocolExtensionContainer {
	return &ProtocolContainer{ExtensionContainer: extendx.NewExtensionContainer()}
}

func NewProtocolExtensionContainer() ProtocolContainer {
	return ProtocolContainer{ExtensionContainer: extendx.NewExtensionContainer()}
}

//-----------------------------------------------

type ProtocolContainer struct {
	extendx.ExtensionContainer
}

func (c *ProtocolContainer) InitExtensions() []error {
	ln := c.Len()
	if ln == 0 {
		return nil
	}
	var rs []error
	c.Range(func(_ int, extension extendx.IExtension) {
		if e, ok := extension.(extendx.IInitExtension); ok {
			err := e.InitExtension()
			rs = appendError(rs, err)
		}
	})
	return rs
}

func (c *ProtocolContainer) DestroyExtensions() []error {
	ln := c.Len()
	if ln == 0 {
		return nil
	}
	var rs []error
	c.RangeReverse(func(_ int, extension extendx.IExtension) {
		if e, ok := extension.(extendx.IInitExtension); ok {
			err := e.DestroyExtension()
			rs = appendError(rs, err)
		}
	})
	return rs
}

func (c *ProtocolContainer) SaveExtensions() []error {
	ln := c.Len()
	if ln == 0 {
		return nil
	}
	var rs []error
	c.Range(func(_ int, extension extendx.IExtension) {
		if e, ok := extension.(extendx.ISaveExtension); ok {
			err := e.SaveExtension()
			if nil != err {
				rs = append(rs, err)
			}
		}
	})
	return rs
}

func (c *ProtocolContainer) SaveExtension(name string) error {
	var err error
	c.HandleAtName(name, func(_ string, extension extendx.IExtension) {
		if e, ok := extension.(extendx.ISaveExtension); ok {
			err = e.SaveExtension()
		}
	})
	return err
}

func (c *ProtocolContainer) EnableExtensions(enable bool) []error {
	ln := c.Len()
	if ln == 0 {
		return nil
	}
	var rs []error
	if enable {
		c.Range(func(_ int, extension extendx.IExtension) {
			if e, ok := extension.(extendx.IEnableExtension); ok && !e.Enable() {
				err := e.EnableExtension()
				rs = appendError(rs, err)
			}
		})
	} else {
		c.RangeReverse(func(_ int, extension extendx.IExtension) {
			if e, ok := extension.(extendx.IEnableExtension); ok && e.Enable() {
				err := e.DisableExtension()
				rs = appendError(rs, err)
			}
		})
	}
	return rs
}

func (c *ProtocolContainer) EnableExtension(name string, enable bool) error {
	var err error
	c.HandleAtName(name, func(_ string, extension extendx.IExtension) {
		if e, ok := extension.(extendx.IEnableExtension); ok {
			if e.Enable() != enable {
				if enable {
					err = e.EnableExtension()
				} else {
					err = e.DisableExtension()
				}
			}
		}
	})
	return err
}

func appendError(errs []error, err error) []error {
	if nil != err {
		return append(errs, err)
	} else {
		return errs
	}
}
