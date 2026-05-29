package mathx

// MaxInt returns the larger of a and b.
// MaxInt 返回 a 与 b 中的较大值。
func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Max3Int returns the largest of a, b, and c.
// Max3Int 返回 a、b、c 中的最大值。
func Max3Int(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

// MaxUInt returns the larger of a and b.
// MaxUInt 返回 a 与 b 中的较大值。
func MaxUInt(a, b uint) uint {
	if a > b {
		return a
	} else {
		return b
	}
}

// Max3UInt returns the largest of a, b, and c.
// Max3UInt 返回 a、b、c 中的最大值。
func Max3UInt(a, b, c uint) uint {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

// MaxFloat64 returns the larger of a and b.
// MaxFloat64 返回 a 与 b 中的较大值。
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

// Max3Float64 returns the largest of a, b, and c.
// Max3Float64 返回 a、b、c 中的最大值。
func Max3Float64(a, b, c float64) float64 {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

// MaxUint32 returns the larger of a and b.
// MaxUint32 返回 a 与 b 中的较大值。
func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	} else {
		return b
	}
}
