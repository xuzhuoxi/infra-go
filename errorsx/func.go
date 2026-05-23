package errorsx

// FuncUnavailableError indicates that a function or operation is not available at the given site.
// FuncUnavailableError 表示指定位置的函数或操作不可用。
type FuncUnavailableError string

// Error returns the error message including the site identifier.
// Error 返回包含位置标识的错误信息。
func (e FuncUnavailableError) Error() string { return "Unavailable Function At:" + string(e) }

// Timeout reports whether the error is a timeout; always false for FuncUnavailableError.
// Timeout 报告错误是否为超时；对 FuncUnavailableError 恒为 false。
func (e FuncUnavailableError) Timeout() bool { return false }

// Temporary reports whether the error is temporary; always false for FuncUnavailableError.
// Temporary 报告错误是否为临时性；对 FuncUnavailableError 恒为 false。
func (e FuncUnavailableError) Temporary() bool { return false }

// FuncRepeatedCallError indicates that the same function was invoked more than once when not allowed.
// FuncRepeatedCallError 表示在不允许的情况下对同一函数进行了重复调用。
type FuncRepeatedCallError string

// Error returns the error message including the site identifier.
// Error 返回包含位置标识的错误信息。
func (e FuncRepeatedCallError) Error() string { return "Repeated Call Function At:" + string(e) }

// Timeout reports whether the error is a timeout; always false for FuncRepeatedCallError.
// Timeout 报告错误是否为超时；对 FuncRepeatedCallError 恒为 false。
func (e FuncRepeatedCallError) Timeout() bool { return false }

// Temporary reports whether the error is temporary; always false for FuncRepeatedCallError.
// Temporary 报告错误是否为临时性；对 FuncRepeatedCallError 恒为 false。
func (e FuncRepeatedCallError) Temporary() bool { return false }

// FuncNotPreparedError indicates that a function was called before it was prepared or initialized.
// FuncNotPreparedError 表示在函数尚未准备或初始化完成时被调用。
type FuncNotPreparedError string

// Error returns the error message including the site identifier.
// Error 返回包含位置标识的错误信息。
func (e FuncNotPreparedError) Error() string { return "Calling Function Not Prepared At:" + string(e) }

// Timeout reports whether the error is a timeout; always false for FuncNotPreparedError.
// Timeout 报告错误是否为超时；对 FuncNotPreparedError 恒为 false。
func (e FuncNotPreparedError) Timeout() bool { return false }

// Temporary reports whether the error is temporary; always false for FuncNotPreparedError.
// Temporary 报告错误是否为临时性；对 FuncNotPreparedError 恒为 false。
func (e FuncNotPreparedError) Temporary() bool { return false }
