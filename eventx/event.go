package eventx

import "reflect"

//监听事件的回调
type EventCall func(evd *EventData)

func (c EventCall) Equal(c1 EventCall) bool {
	return &c == &c1 || reflect.ValueOf(c).Pointer() == reflect.ValueOf(c1).Pointer()
}

//事件数据
//@author xuzhuoxi
//Created  on 2019/01/08.
type EventData struct {
	//事件类型
	EventType string
	//事件传递的数据
	Data interface{}

	currentTarget     interface{}
	currentDispatcher IEventDispatcher
	target            interface{}
	stopped           bool
}

/**
 * 事件的发生器
 * @returns {IEventDispatcher}
 */
func (ed *EventData) CurrentDispatcher() IEventDispatcher {
	return ed.currentDispatcher
}

/**
 * 事件当前目标
 * @returns interface{}
 */
func (ed *EventData) CurrentTarget() interface{} {
	return ed.currentTarget
}

/**
 * 是否设置为停止
 * @returns {boolean}
 */
func (ed *EventData) Stopped() bool {
	return ed.stopped
}

/**
 * 防止对事件流中当前节点中和所有后续节点中的事件侦听器进行处理
 */
func (ed *EventData) StopImmediatePropagation() {
	ed.stopped = true
}

// 创建一个EventDispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{}
}

type IEventDispatcher interface {
	/**
	* 添加事件
	* @param eventType 事件类型
	* @param func 监听函数
	 */
	AddEventListener(eventType string, call EventCall)
	/**
	 * 添加单次执行事件
	 * @param eventType
	 * @param func
	 */
	OnceEventListener(eventType string, call EventCall)
	/**
	 * 删除事件
	 * @param eventType 事件类型
	 * @param func 监听函数
	 */
	RemoveEventListener(eventType string, call EventCall)
	/**
	 * 删除一类事件
	 * @param eventType 事件类型
	 */
	RemoveEventListenerByType(eventType string)
	/**
	 * 清除全部事件
	 */
	RemoveEventListeners()
	/**
	 * 触发某一类型的事件  并传递数据
	 * @param eventType 事件类型
	 * @param currentTarget 当前对象
	 * @param data 事件的数据(可为null)
	 */
	DispatchEvent(eventType string, currentTarget interface{}, data interface{})
}

type EventDispatcher struct {
	dMap map[string]*_EventDelegate
}

func (e *EventDispatcher) AddEventListener(eventType string, call EventCall) {
	e.getDelegate(eventType).AddListener(call)
}

func (e *EventDispatcher) OnceEventListener(eventType string, call EventCall) {
	e.getDelegate(eventType).OnceListener(call)
}

func (e *EventDispatcher) RemoveEventListener(eventType string, call EventCall) {
	if !e.hasType(eventType) {
		return
	}
	e.getDelegate(eventType).RemoveListener(call)
}

func (e *EventDispatcher) RemoveEventListenerByType(eventType string) {
	if !e.hasType(eventType) {
		return
	}
	e.getDelegate(eventType).RemoveListeners()
}

func (e *EventDispatcher) RemoveEventListeners() {
	e.dMap = nil
}

func (e *EventDispatcher) DispatchEvent(eventType string, currentTarget interface{}, data interface{}) {
	if !e.hasType(eventType) {
		return
	}
	d := &EventData{EventType: eventType, Data: data, currentTarget: currentTarget, currentDispatcher: e}
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
