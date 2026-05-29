package mathx

// ClampInt clamps value to the inclusive range [min, max].
// ClampInt 将 value 限制在闭区间 [min, max] 内。
func ClampInt(value int, min int, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampUInt clamps value to the inclusive range [min, max].
// ClampUInt 将 value 限制在闭区间 [min, max] 内。
func ClampUInt(value uint, min uint, max uint) uint {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampUInt32 clamps value to the inclusive range [min, max].
// ClampUInt32 将 value 限制在闭区间 [min, max] 内。
func ClampUInt32(value uint32, min uint32, max uint32) uint32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampFloat01 clamps value to the inclusive range [0, 1].
// ClampFloat01 将 value 限制在闭区间 [0, 1] 内。
func ClampFloat01(value float32) float32 {
	return ClampFloat32(value, 0, 1)
}

// ClampFloat32 clamps value to the inclusive range [min, max].
// ClampFloat32 将 value 限制在闭区间 [min, max] 内。
func ClampFloat32(value float32, min float32, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampFloat64 clamps value to the inclusive range [min, max].
// ClampFloat64 将 value 限制在闭区间 [min, max] 内。
func ClampFloat64(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
