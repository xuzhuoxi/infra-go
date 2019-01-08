package mongox

type SessionNilError string

func (e SessionNilError) Error() string   { return "Session Nil At: " + string(e) }
func (e SessionNilError) Timeout() bool   { return false }
func (e SessionNilError) Temporary() bool { return false }

type SessionLimitError string

func (e SessionLimitError) Error() string   { return "Session Is Reaching Limit At: " + string(e) }
func (e SessionLimitError) Timeout() bool   { return false }
func (e SessionLimitError) Temporary() bool { return false }
