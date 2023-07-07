package eventx

type _CallInfo struct {
	Once bool
	Call EventCall
}

func (ci *_CallInfo) Equal(other *_CallInfo) bool {
	return ci == other || (ci.Once == other.Once && (&ci.Call == &other.Call || ci.Call.Equal(other.Call)))
}

type _EventDelegate struct {
	EventType string
	calls     []*_CallInfo
}

func (d *_EventDelegate) Handle(data *EventData) {
	if d.EventType != data.EventType || len(d.calls) == 0 {
		return
	}
	var onceTempAry []*_CallInfo
	for _, call := range d.calls {
		if data.stopped {
			break
		}
		call.Call(data)
		if call.Once {
			onceTempAry = append(onceTempAry, call)
		}
	}
	if len(onceTempAry) > 0 {
		d.removeCalls(onceTempAry...)
	}
}

// AddListener
// 添加监听函数
func (d *_EventDelegate) AddListener(call EventCall) {
	d.calls = append(d.calls, &_CallInfo{Once: false, Call: call})
}

// OnceListener
// 添加单次监听函数
func (d *_EventDelegate) OnceListener(call EventCall) {
	d.calls = append(d.calls, &_CallInfo{Once: true, Call: call})
}

// RemoveListener
// 删除监听函数
func (d *_EventDelegate) RemoveListener(call EventCall, stopAfterMatch bool) {
	for index := len(d.calls) - 1; index >= 0; index-- {
		if d.calls[index].Call.Equal(call) {
			d.calls = append(d.calls[:index], d.calls[index+1:]...)
			if stopAfterMatch {
				break
			}
		}
	}
}

// RemoveListeners
// 删除全部监听函数
func (d *_EventDelegate) RemoveListeners() {
	d.calls = nil
}

func (d *_EventDelegate) containsCall(call *_CallInfo) bool {
	if len(d.calls) == 0 || nil == call {
		return false
	}
	for index := range d.calls {
		if d.calls[index].Equal(call) {
			return true
		}
	}
	return false
}

func (d *_EventDelegate) removeCalls(removes ...*_CallInfo) {
	size := len(d.calls)
	if size == 0 {
		return
	}
	for index := size - 1; index >= 0; index-- {
		for index2 := range removes {
			if d.calls[index].Equal(removes[index2]) {
				d.calls = append(d.calls[:index], d.calls[index+1:]...)
				removes = append(removes[:index2], removes[index2+1:]...)
				break
			}
		}
	}
}
