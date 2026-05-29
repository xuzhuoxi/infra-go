package netx

import (
	"errors"
	"fmt"
	"sync"

	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
)

// FuncPackHandler handles socket pack messages.
// FuncPackHandler 处理 Sock 消息包。
//   data: message payload
//   data: 消息数据
//   connInfo: sender connection info
//   connInfo: 发送者连接信息
//   other: additional context
//   other: 其它信息
//
// Returns catch indicating whether the message was handled.
// 返回 catch 表示当前是否已经把消息处理。
type FuncPackHandler func(data []byte, connInfo IConnInfo, other interface{}) (catch bool)

//--------------------------

// DefaultPackHandler is the default pack handler that logs and consumes the message.
// DefaultPackHandler 是默认消息处理器，记录日志并消费消息。
func DefaultPackHandler(data []byte, connInfo IConnInfo, other interface{}) bool {
	logx.Traceln(fmt.Sprintf("DefaultMessageHandler{Conn=%s,Data=%s,Other=%s]}", connInfo, data, other))
	return true
}

// NewDefaultIPackHandler creates a default IPackHandlerContainer.
// NewDefaultIPackHandler 创建默认的 IPackHandlerContainer。
func NewDefaultIPackHandler() IPackHandlerContainer {
	return NewDefaultPackHandler()
}

// NewDefaultPackHandler creates a PackHandlerContainer with DefaultPackHandler.
// NewDefaultPackHandler 创建包含 DefaultPackHandler 的 PackHandlerContainer。
func NewDefaultPackHandler() *PackHandlerContainer {
	return &PackHandlerContainer{handlers: []FuncPackHandler{DefaultPackHandler}}
}

// NewIPackHandler creates an IPackHandlerContainer from handlers.
// NewIPackHandler 根据处理器列表创建 IPackHandlerContainer。
func NewIPackHandler(handlers []FuncPackHandler) IPackHandlerContainer {
	return NewPackHandler(handlers)
}

// NewPackHandler creates a PackHandlerContainer from handlers.
// NewPackHandler 根据处理器列表创建 PackHandlerContainer。
func NewPackHandler(handlers []FuncPackHandler) *PackHandlerContainer {
	return &PackHandlerContainer{handlers: handlers}
}

// IPackHandlerContainerSetter injects an IPackHandlerContainer.
// IPackHandlerContainerSetter 注入 IPackHandlerContainer。
type IPackHandlerContainerSetter interface {
	SetPackHandlerContainer(packHandlerContainer IPackHandlerContainer)
}

// IPackHandlerContainerGetter returns the bound IPackHandlerContainer.
// IPackHandlerContainerGetter 返回已绑定的 IPackHandlerContainer。
type IPackHandlerContainerGetter interface {
	GetPackHandlerContainer() IPackHandlerContainer
}

// IPackHandlerContainer manages an ordered chain of pack handlers.
// IPackHandlerContainer 管理有序的消息包处理器链。
type IPackHandlerContainer interface {
	// Size returns the number of handlers.
	// Size 返回处理函数数量。
	Size() int
	// FirstHandler invokes the callback with the first handler.
	// FirstHandler 由第一个处理函数处理。
	FirstHandler(first func(handler FuncPackHandler) bool)
	// LastHandler invokes the callback with the last handler.
	// LastHandler 由最后一个处理函数处理。
	LastHandler(last func(handler FuncPackHandler) bool)
	// ForEachHandler iterates handlers in order; stops when each returns true.
	// ForEachHandler 依顺序遍历处理；each 返回 true 则中断遍历。
	ForEachHandler(each func(handler FuncPackHandler) bool)

	// AppendPackHandler appends a pack handler.
	// AppendPackHandler 追加消息处理函数。
	AppendPackHandler(handler FuncPackHandler) error
	// ClearHandler removes a specific handler.
	// ClearHandler 清除指定消息处理函数。
	ClearHandler(handler FuncPackHandler) error
	// ClearHandlers removes all handlers.
	// ClearHandlers 清除全部消息处理函数。
	ClearHandlers() error
	// SetPackHandlers replaces the handler list.
	// SetPackHandlers 设置消息处理函数列表。
	SetPackHandlers(handlers []FuncPackHandler) error
}

// PackHandlerContainer is the default IPackHandlerContainer implementation.
// PackHandlerContainer 是 IPackHandlerContainer 的默认实现。
type PackHandlerContainer struct {
	handlers []FuncPackHandler
	RWMutex  sync.RWMutex
}

// Size returns the number of handlers.
// Size 返回处理函数数量。
func (ph *PackHandlerContainer) Size() int {
	ph.RWMutex.RLock()
	defer ph.RWMutex.RUnlock()
	return len(ph.handlers)
}

// FirstHandler invokes the callback with the first handler.
// FirstHandler 由第一个处理函数处理。
func (ph *PackHandlerContainer) FirstHandler(first func(handler FuncPackHandler) bool) {
	ph.RWMutex.RLock()
	defer ph.RWMutex.RUnlock()
	if len(ph.handlers) == 0 {
		return
	}
	first(ph.handlers[0])
}

// LastHandler invokes the callback with the last handler.
// LastHandler 由最后一个处理函数处理。
func (ph *PackHandlerContainer) LastHandler(first func(handler FuncPackHandler) bool) {
	ph.RWMutex.RLock()
	defer ph.RWMutex.RUnlock()
	if len(ph.handlers) == 0 {
		return
	}
	first(ph.handlers[len(ph.handlers)-1])
}

// ForEachHandler iterates handlers in order.
// ForEachHandler 依顺序遍历处理。
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

// AppendPackHandler appends a pack handler.
// AppendPackHandler 追加消息处理函数。
func (ph *PackHandlerContainer) AppendPackHandler(handler FuncPackHandler) error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	if nil == handler {
		return errors.New("PackHandlerContainer.AppendPackHandler:handler is nil")
	}
	ph.handlers = append(ph.handlers, handler)
	return nil
}

// ClearHandler removes a specific handler.
// ClearHandler 清除指定消息处理函数。
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

// ClearHandlers removes all handlers.
// ClearHandlers 清除全部消息处理函数。
func (ph *PackHandlerContainer) ClearHandlers() error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	ph.handlers = nil
	return nil
}

// SetPackHandlers replaces the handler list.
// SetPackHandlers 设置消息处理函数列表。
func (ph *PackHandlerContainer) SetPackHandlers(handlers []FuncPackHandler) error {
	ph.RWMutex.Lock()
	defer ph.RWMutex.Unlock()
	ph.handlers = handlers
	return nil
}
