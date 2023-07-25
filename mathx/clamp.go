// Package mathx
// Create on 2023/7/23
// @author xuzhuoxi
package mathx

func ClampInt(value int, min int, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ClampUInt(value uint, min uint, max uint) uint {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ClampUInt32(value uint32, min uint32, max uint32) uint32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ClampFloat01(value float32) float32 {
	return ClampFloat32(value, 0, 1)
}

func ClampFloat32(value float32, min float32, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ClampFloat64(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
