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
	RegisterBlendFunc(Clear, ClearBlend)
}

// 清除模式
// R = 0
func ClearBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ClearUnit(source.A, target.A, factor)
	}
	source.R = ClearUnit(source.R, target.R, factor)
	source.G = ClearUnit(source.G, target.G, factor)
	source.B = ClearUnit(source.B, target.B, factor)
	return source
}

// R = 0
func ClearUnit(S uint8, D uint8, _ float64) uint8 {
	return 0
}
