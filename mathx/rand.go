// Package mathx
// Created by xuzhuoxi
// on 2019-05-27.
// @author xuzhuoxi
package mathx

import (
	"math"
	"math/rand"
)

// RandBool 随机bool
func RandBool() bool {
	return rand.Int63() > math.MaxInt32
}
