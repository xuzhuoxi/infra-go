package errs

type FuncUnavailableError string

func (e FuncUnavailableError) Error() string   { return "Unavailable Function At:" + string(e) }
func (e FuncUnavailableError) Timeout() bool   { return false }
func (e FuncUnavailableError) Temporary() bool { return false }

type FuncRepeatedCallError string

func (e FuncRepeatedCallError) Error() string   { return "Repeated Call Function At:" + string(e) }
func (e FuncRepeatedCallError) Timeout() bool   { return false }
func (e FuncRepeatedCallError) Temporary() bool { return false }

type FuncNotPreparedError string

func (e FuncNotPreparedError) Error() string   { return "Calling Function Not Prepared At:" + string(e) }
func (e FuncNotPreparedError) Timeout() bool   { return false }
func (e FuncNotPreparedError) Temporary() bool { return false }
