//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"fmt"
	"testing"
)

func TestMode(t *testing.T) {
	fmt.Println(255 / 128)
	fmt.Println(511 / 10)
}

func TestSubtract(t *testing.T) {
	var b uint32 = 12
	var f uint32 = 23
	fmt.Println(b - f)
	fmt.Println(f - b)
}
