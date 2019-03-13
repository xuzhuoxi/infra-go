//
//Created by xuzhuoxi
//on 2019-03-10.
//@author xuzhuoxi
//
package netx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/logx"
	"sync"
)

type FuncPackHandler func(msgBytes []byte, info interface{}) bool

func DefaultPackHandler(msgData []byte, info interface{}) bool {
	logx.Traceln("DefaultMessageHandler[Sender:"+fmt.Sprint(info)+"]msgData:", msgData, "dataLen:", len(msgData), "]")
	return true
}

func NewDefaultIPackHandler() IPackHandler {
	return NewDefaultPackHandler()
}

func NewDefaultPackHandler() *PackHandler {
	return &PackHandler{handlers: []FuncPackHandler{DefaultPackHandler}}
}

func NewIPackHandler(handlers []FuncPackHandler) IPackHandler {
	return NewPackHandler(handlers)
}

func NewPackHandler(handlers []FuncPackHandler) *PackHandler {
	return &PackHandler{handlers: handlers}
}

type IPackHandlerSetter interface {
	SetPackHandler(packHandler IPackHandler)
}

type IPackHandlerGetter interface {
	GetPackHandler() IPackHandler
}

type IPackHandler interface {
	FirstHandler(first func(handler FuncPackHandler) bool)
	ForEachHandler(each func(handler FuncPackHandler) bool)
	AppendPackHandler(handler FuncPackHandler) error
	SetPackHandlers(handlers []FuncPackHandler) error
	ClearHandlers() error
}

type PackHandler struct {
	handlers []FuncPackHandler
	RWMutex  sync.RWMutex
}

func (ph *PackHandler) FirstHandler(first func(handler FuncPackHandler) bool) {
	ph.RWMutex.RLock()
	defer ph.RWMutex.RUnlock()
	if len(ph.handlers) == 0 {
		return
	}
	first(ph.handlers[0])
}

func (ph *PackHandler) ForEachHandler(each func(handler FuncPackHandler) bool) {
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

func (ph *PackHandler) AppendPackHandler(handler FuncPackHandler) error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	if nil == handler {
		return errors.New("PackHandler:handler is nil")
	}
	ph.handlers = append(ph.handlers, handler)
	return nil
}

func (ph *PackHandler) SetPackHandlers(handlers []FuncPackHandler) error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	ph.handlers = handlers
	return nil
}

func (ph *PackHandler) ClearHandlers() error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	ph.handlers = nil
	return nil
}
