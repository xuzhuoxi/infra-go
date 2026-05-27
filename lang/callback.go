package lang

// FuncCallback is a variadic callback function type.
// FuncCallback 是可变参数回调函数类型。
type FuncCallback = func(...interface{})

// NewCallback wraps call and optional default args into a Callback.
// NewCallback 将 call 与可选默认参数封装为 Callback。
func NewCallback(call FuncCallback, args ...interface{}) *Callback {
	return &Callback{call: call, args: args}
}

// Callback stores a function and arguments for deferred invocation.
// Callback 保存函数与参数，供延迟调用。
type Callback struct {
	call func(...interface{})
	args []interface{}
}

// SetCall replaces the stored callback function.
// SetCall 替换保存的回调函数。
func (c *Callback) SetCall(call FuncCallback) {
	c.call = call
}

// SetArgs replaces the stored default arguments.
// SetArgs 替换保存的默认参数。
func (c *Callback) SetArgs(args ...interface{}) {
	c.args = args
}

// Apply invokes the callback with the given arguments.
// Apply 使用给定参数调用回调。
func (c *Callback) Apply(args ...interface{}) {
	if nil == c.call {
		return
	}
	c.call(args...)
}

// Invoke invokes the callback with the stored default arguments.
// Invoke 使用保存的默认参数调用回调。
func (c *Callback) Invoke() {
	if nil == c.call {
		return
	}
	c.call(c.args...)
}

// Clear resets the callback and arguments to nil.
// Clear 将回调与参数清空为 nil。
func (c *Callback) Clear() {
	c.call, c.args = nil, nil
}
