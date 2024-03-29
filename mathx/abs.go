// Package mathx
// Created by xuzhuoxi
// on 2019-04-03.
// @author xuzhuoxi
//
package mathx

import "math"

func AbsInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func AbsInt8(value int8) int8 {
	if value < 0 {
		return -value
	}
	return value
}

func AbsInt16(value int16) int16 {
	if value < 0 {
		return -value
	}
	return value
}

func AbsInt32(value int32) int32 {
	if value < 0 {
		return -value
	}
	return value
}

func AbsInt64(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

func AbsFloat32(value float32) float32 {
	return float32(math.Abs(float64(value)))
}

func AbsFloat64(value float64) float64 {
	return math.Abs(value)
}
