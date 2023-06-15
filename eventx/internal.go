package eventx

type _CallInfo struct {
	Once bool
	Call EventCall
}

func (ci *_CallInfo) Equal(other *_CallInfo) bool {
	return ci == other || (ci.Once == other.Once && &ci.Call == &other.Call)
}

type _EventDelegate struct {
	EventType string

	calls []*_CallInfo
}

func (d *_EventDelegate) Handle(data *EventData) {
	if d.EventType != data.EventType || len(d.calls) == 0 {
		return
	}
	copyArr := d.calls[:]
	var onceTempAry []*_CallInfo
	for _, call := range copyArr {
		if data.stopped {
			break
		}
		if !containsCall(d.calls, call) {
			continue
		}
		call.Call(data)
		if call.Once {
			onceTempAry = append(onceTempAry, call)
		}
	}
	if len(onceTempAry) > 0 {
		removeCall(d.calls, onceTempAry...)
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
func (d *_EventDelegate) RemoveListener(call EventCall) {
	for index := len(d.calls) - 1; index >= 0; index-- {
		if d.calls[index].Call.Equal(call) {
			d.calls = append(d.calls[:index], d.calls[index+1:]...)
		}
	}
}

// RemoveListeners
// 删除全部监听函数
func (d *_EventDelegate) RemoveListeners() {
	d.calls = nil
}

func containsCall(calls []*_CallInfo, call *_CallInfo) bool {
	if len(calls) == 0 {
		return false
	}
	for _, v := range calls {
		if v.Equal(call) {
			return true
		}
	}
	return false
}

func removeCall(calls []*_CallInfo, removeCalls ...*_CallInfo) []*_CallInfo {
	l := len(calls)
	if l == 0 {
		return calls
	}
	rs := calls[:]
	for index := l - 1; index >= 0; index-- {
		for index2, c := range removeCalls {
			if rs[index].Equal(c) {
				rs = append(rs[:index], rs[index+1:]...)
				removeCalls = append(removeCalls[:index2], removeCalls[index2+1:]...)
				break
			}
		}
	}
	return rs
}
