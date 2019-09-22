package mathx

func MinUint32(a, b uint32) uint32 {
	if a > b {
		return b
	} else {
		return a
	}
}
func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	} else {
		return b
	}
}
