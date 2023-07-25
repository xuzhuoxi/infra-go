// Package mathx
// Create on 2023/7/23
// @author xuzhuoxi
package mathx

func MinInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Min3Int(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func MinUint(a, b uint) uint {
	if a > b {
		return b
	} else {
		return a
	}
}

func Min3Uint(a, b, c uint) uint {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func MinFloat64(a, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func Min3Float64(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func MinUint32(a, b uint32) uint32 {
	if a > b {
		return b
	} else {
		return a
	}
}
