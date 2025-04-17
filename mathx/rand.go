// Package mathx
// Created by xuzhuoxi
// on 2019-05-27.
// @author xuzhuoxi
package mathx

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandBool 随机bool
func RandBool() bool {
	return rand.Int63() > math.MaxInt32
}

// RandFloat64 任意范围的浮点数
func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandFloat32 任意范围的浮点数
func RandFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// RandInt64 任意范围的整 数
func RandInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// RandInt32 任意范围的整数
func RandInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min)
}

// RandUint64 任意范围的整数
func RandUint64(min, max uint64) uint64 {
	return min + uint64(rand.Int63n(int64(max-min)))
}

// RandUint32 任意范围的整数
func RandUint32(min, max uint32) uint32 {
	return min + uint32(rand.Int31n(int32(max-min)))
}
