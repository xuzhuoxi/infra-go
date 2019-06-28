package mongox

type SessionNilError string

func (e SessionNilError) Error() string   { return "Session Nil At: " + string(e) }
func (e SessionNilError) Timeout() bool   { return false }
func (e SessionNilError) Temporary() bool { return false }

type ClientNilError string

func (e ClientNilError) Error() string   { return "Client Nil At: " + string(e) }
func (e ClientNilError) Timeout() bool   { return false }
func (e ClientNilError) Temporary() bool { return false }
