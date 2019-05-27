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
	RegisterBlendFunc(Divide, DivideBlend)
}

// 划分模式
// C = 255 * A / B
func DivideBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = DivideUnit(source.A, target.A, factor)
	}
	source.R = DivideUnit(source.R, target.R, factor)
	source.G = DivideUnit(source.G, target.G, factor)
	source.B = DivideUnit(source.B, target.B, factor)
	return source
}

// C = 255 * A / B
func DivideUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(255 * uint16(S) / uint16(D))
}
