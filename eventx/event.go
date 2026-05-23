package eventx

import "reflect"

// EventCall is the callback invoked when an event is dispatched.
// EventCall 是事件分发时被调用的回调函数。
type EventCall func(evd *EventData)

// Equal reports whether two EventCall values refer to the same function value.
// Equal 判断两个 EventCall 是否指向同一函数值。
func (c EventCall) Equal(c1 EventCall) bool {
	return &c == &c1 || reflect.ValueOf(c).Pointer() == reflect.ValueOf(c1).Pointer()
}

// EventData carries type, payload, target, and dispatcher for a dispatched event.
// EventData 携带已分发事件的类型、载荷、目标与分发器信息。
type EventData struct {
	// EventType is the event type identifier.
	// EventType 是事件类型标识。
	EventType string
	// Data is the payload passed with the event.
	// Data 是随事件传递的载荷数据。
	Data interface{}

	// CurrentTarget is the target passed when the event was dispatched.
	// CurrentTarget 是抛出事件时传入的当前目标对象。
	CurrentTarget interface{}
	// CurrentDispatcher is the object that dispatches this event.
	// CurrentDispatcher 是分发该事件的对象（DispatchEvent 的所属对象，可能是内嵌对象）。
	CurrentDispatcher IEventDispatcher

	stopped bool
}

// Stopped reports whether propagation has been stopped for this event.
// Stopped 返回该事件是否已设置为停止传播。
func (ed *EventData) Stopped() bool {
	return ed.stopped
}

// StopImmediatePropagation prevents processing of this event by listeners on the current and subsequent nodes.
// StopImmediatePropagation 阻止对事件流中当前节点及所有后续节点上的侦听器进行处理。
func (ed *EventData) StopImmediatePropagation() {
	ed.stopped = true
}

// NewEventDispatcher creates a new EventDispatcher.
// NewEventDispatcher 创建一个 EventDispatcher。
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{}
}

// IEventDispatcher defines event listener registration and dispatch.
// IEventDispatcher 定义事件监听器的注册与分发行为。
type IEventDispatcher interface {
	// AddEventListener registers a listener for eventType; duplicate registrations are all invoked.
	// AddEventListener 为 eventType 添加监听器；重复添加会多次执行。
	//
	// eventType is the event type identifier.
	// eventType 为事件类型。
	//
	// call is the listener callback.
	// call 为监听函数。
	AddEventListener(eventType string, call EventCall)
	// OnceEventListener registers a one-shot listener for eventType; duplicate registrations are all invoked.
	// OnceEventListener 为 eventType 添加单次执行监听器；重复添加会多次执行。
	//
	// eventType is the event type identifier.
	// eventType 为事件类型。
	//
	// call is the listener callback.
	// call 为监听函数。
	OnceEventListener(eventType string, call EventCall)
	// RemoveEventListener removes listeners matching eventType and call.
	// RemoveEventListener 删除与 eventType、call 匹配的监听器。
	//
	// eventType is the event type identifier.
	// eventType 为事件类型。
	//
	// call is the listener callback.
	// call 为监听函数。
	//
	// args optional: args[0] bool stopAfterMatch stops after the first match when true.
	// args 可选参数：args[0] bool stopAfterMatch 为 true 时在首次匹配后停止移除。
	RemoveEventListener(eventType string, call EventCall, args ...bool)
	// RemoveEventListenerByType removes all listeners for eventType.
	// RemoveEventListenerByType 删除 eventType 下的全部监听器。
	//
	// eventType is the event type identifier.
	// eventType 为事件类型。
	RemoveEventListenerByType(eventType string)
	// RemoveEventListeners clears all registered listeners.
	// RemoveEventListeners 清除全部已注册监听器。
	RemoveEventListeners()
	// DispatchEvent dispatches eventType with currentTarget and data to registered listeners.
	// DispatchEvent 触发 eventType 事件并向已注册监听器传递 data。
	//
	// eventType is the event type identifier.
	// eventType 为事件类型。
	//
	// currentTarget is the dispatch target object.
	// currentTarget 为当前对象。
	//
	// data is the event payload (maybe nil).
	// data 为事件数据（可为 nil）。
	DispatchEvent(eventType string, currentTarget interface{}, data interface{})
}

// EventDispatcher manages listeners grouped by event type and dispatches EventData to them.
// EventDispatcher 按事件类型管理监听器并向其分发 EventData。
type EventDispatcher struct {
	dMap map[string]*_EventDelegate
}

// AddEventListener registers a listener for eventType; duplicate registrations are all invoked.
// AddEventListener 为 eventType 添加监听器；重复添加会多次执行。
//
// eventType is the event type identifier.
// eventType 为事件类型。
//
// call is the listener callback.
// call 为监听函数。
func (e *EventDispatcher) AddEventListener(eventType string, call EventCall) {
	e.getDelegate(eventType).AddListener(call)
}

// OnceEventListener registers a one-shot listener for eventType; duplicate registrations are all invoked.
// OnceEventListener 为 eventType 添加单次执行监听器；重复添加会多次执行。
//
// eventType is the event type identifier.
// eventType 为事件类型。
//
// call is the listener callback.
// call 为监听函数。
func (e *EventDispatcher) OnceEventListener(eventType string, call EventCall) {
	e.getDelegate(eventType).OnceListener(call)
}

// RemoveEventListener removes listeners matching eventType and call.
// RemoveEventListener 删除与 eventType、call 匹配的监听器。
//
// eventType is the event type identifier.
// eventType 为事件类型。
//
// call is the listener callback.
// call 为监听函数。
//
// args optional: args[0] bool stopAfterMatch stops after the first match when true.
// args 可选参数：args[0] bool stopAfterMatch 为 true 时在首次匹配后停止移除。
func (e *EventDispatcher) RemoveEventListener(eventType string, call EventCall, args ...bool) {
	if !e.hasType(eventType) {
		return
	}
	stopAfterMatch := false
	if len(args) > 0 {
		stopAfterMatch = args[0]
	}
	e.getDelegate(eventType).RemoveListener(call, stopAfterMatch)
}

// RemoveEventListenerByType removes all listeners for eventType.
// RemoveEventListenerByType 删除 eventType 下的全部监听器。
//
// eventType is the event type identifier.
// eventType 为事件类型。
func (e *EventDispatcher) RemoveEventListenerByType(eventType string) {
	if !e.hasType(eventType) {
		return
	}
	e.getDelegate(eventType).RemoveListeners()
}

// RemoveEventListeners clears all registered listeners.
// RemoveEventListeners 清除全部已注册监听器。
func (e *EventDispatcher) RemoveEventListeners() {
	e.dMap = nil
}

// DispatchEvent dispatches eventType with currentTarget and data to registered listeners.
// DispatchEvent 触发 eventType 事件并向已注册监听器传递 data。
//
// eventType is the event type identifier.
// eventType 为事件类型。
//
// currentTarget is the dispatch target object.
// currentTarget 为当前对象。
//
// data is the event payload (maybe nil).
// data 为事件数据（可为 nil）。
func (e *EventDispatcher) DispatchEvent(eventType string, currentTarget interface{}, data interface{}) {
	if !e.hasType(eventType) {
		return
	}
	d := &EventData{EventType: eventType, Data: data, CurrentTarget: currentTarget, CurrentDispatcher: e}
	e.getDelegate(eventType).Handle(d)
}

func (e *EventDispatcher) hasType(eventType string) bool {
	if nil == e.dMap {
		return false
	}
	_, ok := e.dMap[eventType]
	return ok
}

func (e *EventDispatcher) getDelegate(eventType string) *_EventDelegate {
	if nil == e.dMap {
		e.dMap = make(map[string]*_EventDelegate)
	}
	if !e.hasType(eventType) {
		e.dMap[eventType] = &_EventDelegate{EventType: eventType}
	}
	d, _ := e.dMap[eventType]
	return d
}
