//
//Created by xuzhuoxi
//on 2019-05-27.
//@author xuzhuoxi
//
package randx

import (
	"math/rand"
	"math"
)

// 随机bool
func RandBool() bool {
	return rand.Int63() > math.MaxInt32
}
