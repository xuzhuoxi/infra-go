package errorsx

type NoCaseCatchError string

func (e NoCaseCatchError) Error() string   { return "No Case Catch Error At: " + string(e) }
func (e NoCaseCatchError) Timeout() bool   { return false }
func (e NoCaseCatchError) Temporary() bool { return false }
