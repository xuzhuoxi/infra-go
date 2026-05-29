package mathx

import "math"

// AbsInt returns the absolute value of value.
// AbsInt 返回 value 的绝对值。
func AbsInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// AbsInt8 returns the absolute value of value.
// AbsInt8 返回 value 的绝对值。
func AbsInt8(value int8) int8 {
	if value < 0 {
		return -value
	}
	return value
}

// AbsInt16 returns the absolute value of value.
// AbsInt16 返回 value 的绝对值。
func AbsInt16(value int16) int16 {
	if value < 0 {
		return -value
	}
	return value
}

// AbsInt32 returns the absolute value of value.
// AbsInt32 返回 value 的绝对值。
func AbsInt32(value int32) int32 {
	if value < 0 {
		return -value
	}
	return value
}

// AbsInt64 returns the absolute value of value.
// AbsInt64 返回 value 的绝对值。
func AbsInt64(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

// AbsFloat32 returns the absolute value of value.
// AbsFloat32 返回 value 的绝对值。
func AbsFloat32(value float32) float32 {
	return float32(math.Abs(float64(value)))
}

// AbsFloat64 returns the absolute value of value.
// AbsFloat64 返回 value 的绝对值。
func AbsFloat64(value float64) float64 {
	return math.Abs(value)
}
