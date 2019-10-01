//
//Created by xuzhuoxi
//on 2019-03-10.
//@author xuzhuoxi
//
package netx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

// 消息处理函数，用于处理Sock消息
// @param data: 消息数据
// @param senderAddress: 发送者连接地址
// @param other: 其它信息
// @return catch; 当前是否已经把消息处理
type FuncPackHandler func(data []byte, senderAddress string, other interface{}) (catch bool)

//--------------------------

func DefaultPackHandler(data []byte, senderAddress string, other interface{}) bool {
	logx.Traceln(fmt.Sprintf("DefaultMessageHandler{Sender=%s,Data=%s,Other=%s]}", senderAddress, data, other))
	return true
}

func NewDefaultIPackHandler() IPackHandlerContainer {
	return NewDefaultPackHandler()
}

func NewDefaultPackHandler() *PackHandlerContainer {
	return &PackHandlerContainer{handlers: []FuncPackHandler{DefaultPackHandler}}
}

func NewIPackHandler(handlers []FuncPackHandler) IPackHandlerContainer {
	return NewPackHandler(handlers)
}

func NewPackHandler(handlers []FuncPackHandler) *PackHandlerContainer {
	return &PackHandlerContainer{handlers: handlers}
}

type IPackHandlerContainerSetter interface {
	SetPackHandlerContainer(packHandlerContainer IPackHandlerContainer)
}

type IPackHandlerContainerGetter interface {
	GetPackHandlerContainer() IPackHandlerContainer
}

type IPackHandlerContainer interface {
	// 由第一个处理
	FirstHandler(first func(handler FuncPackHandler) bool)
	// 由最后一个处理
	LastHandler(last func(handler FuncPackHandler) bool)
	// 依顺序遍历处理
	// each返回true，则中断遍历
	ForEachHandler(each func(handler FuncPackHandler) bool)

	// 追加消息处理函数
	AppendPackHandler(handler FuncPackHandler) error
	// 清除消息处理函数
	ClearHandler(handler FuncPackHandler) error
	// 清除全部消息处理函数
	ClearHandlers() error
	// 设置消息处理函数列表
	SetPackHandlers(handlers []FuncPackHandler) error
}

type PackHandlerContainer struct {
	handlers []FuncPackHandler
	RWMutex  sync.RWMutex
}

func (ph *PackHandlerContainer) FirstHandler(first func(handler FuncPackHandler) bool) {
	ph.RWMutex.RLock()
	defer ph.RWMutex.RUnlock()
	if len(ph.handlers) == 0 {
		return
	}
	first(ph.handlers[0])
}

func (ph *PackHandlerContainer) LastHandler(first func(handler FuncPackHandler) bool) {
	ph.RWMutex.RLock()
	defer ph.RWMutex.RUnlock()
	if len(ph.handlers) == 0 {
		return
	}
	first(ph.handlers[len(ph.handlers)-1])
}

func (ph *PackHandlerContainer) ForEachHandler(each func(handler FuncPackHandler) bool) {
	ph.RWMutex.RLock()
	defer ph.RWMutex.RUnlock()
	l := len(ph.handlers)
	switch {
	case 0 == l:
		return
	case 1 == l:
		each(ph.handlers[0])
	default:
		for _, handler := range ph.handlers {
			if each(handler) {
				break
			}
		}
	}
}

func (ph *PackHandlerContainer) AppendPackHandler(handler FuncPackHandler) error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	if nil == handler {
		return errors.New("PackHandlerContainer.AppendPackHandler:handler is nil")
	}
	ph.handlers = append(ph.handlers, handler)
	return nil
}

func (ph *PackHandlerContainer) ClearHandler(handler FuncPackHandler) error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	if nil == handler {
		return errors.New("PackHandlerContainer.ClearHandler:handler is nil")
	}
	for index := range ph.handlers {
		if lang.Equal(ph.handlers[index], handler) {
			ph.handlers = append(ph.handlers[:index], ph.handlers[index+1:]...)
			return nil
		}
	}
	return nil
}

func (ph *PackHandlerContainer) ClearHandlers() error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	ph.handlers = nil
	return nil
}

func (ph *PackHandlerContainer) SetPackHandlers(handlers []FuncPackHandler) error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	ph.handlers = handlers
	return nil
}
