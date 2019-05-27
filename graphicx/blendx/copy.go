//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Copy, CopyBlend)
}

// 覆盖模式
// R = 0
func CopyBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = CopyUnit(source.A, target.A, factor)
	}
	source.R = CopyUnit(source.R, target.R, factor)
	source.G = CopyUnit(source.G, target.G, factor)
	source.B = CopyUnit(source.B, target.B, factor)
	return source
}

// R = D
func CopyUnit(S uint8, D uint8, _ float64) uint8 {
	return D
}
