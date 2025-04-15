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
func RandFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandFloat32 任意范围的浮点数
func RandFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
