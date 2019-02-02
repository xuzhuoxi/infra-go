package cryptox

type BlockNilError string

func (e BlockNilError) Error() string   { return "Block Nil Error At: " + string(e) }
func (e BlockNilError) Timeout() bool   { return false }
func (e BlockNilError) Temporary() bool { return false }
