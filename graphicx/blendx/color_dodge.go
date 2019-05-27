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
	RegisterBlendFunc(ColorDodge, ColorDodgeBlend)
}

// 颜色减淡模式
// R = S + (S*D)/(1-D)
// R = S + S*D/(255-D)
func ColorDodgeBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ColorDodgeUnit(source.A, target.A, factor)
	}
	source.R = ColorDodgeUnit(source.R, target.R, factor)
	source.G = ColorDodgeUnit(source.G, target.G, factor)
	source.B = ColorDodgeUnit(source.B, target.B, factor)
	return source
}

// R = S + (S*D)/(1-D)
// R = S + S*D/(255-D)
func ColorDodgeUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	return uint8(S16 + S16*D16/(255-D16))
}
