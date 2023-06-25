// Package lang
// Create on 2023/6/25
// @author xuzhuoxi
package lang

type FuncCallback = func(...interface{})

func NewCallback(call FuncCallback, args ...interface{}) *Callback {
	return &Callback{call: call, args: args}
}

type Callback struct {
	call func(...interface{})
	args []interface{}
}

func (c *Callback) SetCall(call FuncCallback) {
	c.call = call
}

func (c *Callback) SetArgs(args ...interface{}) {
	c.args = args
}

func (c *Callback) Apply(args ...interface{}) {
	if nil == c.call {
		return
	}
	c.call(args...)
}

func (c *Callback) Invoke() {
	if nil == c.call {
		return
	}
	c.call(c.args...)
}

func (c *Callback) Clear() {
	c.call, c.args = nil, nil
}
