package mathx

// MinInt returns the smaller of a and b.
// MinInt 返回 a 与 b 中的较小值。
func MinInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// Min3Int returns the smallest of a, b, and c.
// Min3Int 返回 a、b、c 中的最小值。
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

// MinUint returns the smaller of a and b.
// MinUint 返回 a 与 b 中的较小值。
func MinUint(a, b uint) uint {
	if a > b {
		return b
	} else {
		return a
	}
}

// Min3Uint returns the smallest of a, b, and c.
// Min3Uint 返回 a、b、c 中的最小值。
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

// MinFloat64 returns the smaller of a and b.
// MinFloat64 返回 a 与 b 中的较小值。
func MinFloat64(a, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

// Min3Float64 returns the smallest of a, b, and c.
// Min3Float64 返回 a、b、c 中的最小值。
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

// MinUint32 returns the smaller of a and b.
// MinUint32 返回 a 与 b 中的较小值。
func MinUint32(a, b uint32) uint32 {
	if a > b {
		return b
	} else {
		return a
	}
}
