package errorsx

// NoCaseCatchError indicates that a switch or control-flow branch fell through without a matching case handler.
// NoCaseCatchError 表示 switch 或控制流分支未命中任何 case 处理逻辑。
type NoCaseCatchError string

// Error returns the error message including the site identifier.
// Error 返回包含位置标识的错误信息。
func (e NoCaseCatchError) Error() string { return "No Case Catch Error At: " + string(e) }

// Timeout reports whether the error is a timeout; always false for NoCaseCatchError.
// Timeout 报告错误是否为超时；对 NoCaseCatchError 恒为 false。
func (e NoCaseCatchError) Timeout() bool { return false }

// Temporary reports whether the error is temporary; always false for NoCaseCatchError.
// Temporary 报告错误是否为临时性；对 NoCaseCatchError 恒为 false。
func (e NoCaseCatchError) Temporary() bool { return false }
