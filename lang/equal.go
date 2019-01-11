package lang

func Equal(a, b interface{}) bool {
	if &a == &b {
		return true
	}
	switch a.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, complex64, complex128, string, uintptr:
		return a == b
	}
	return false
}
