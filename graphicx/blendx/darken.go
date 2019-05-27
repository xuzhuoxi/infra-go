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
	RegisterBlendFunc(Darken, DarkenBlend)
}

// 变暗模式
// R = min(S, D)
func DarkenBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = DarkenUnit(source.A, target.A, factor)
	}
	source.R = DarkenUnit(source.R, target.R, factor)
	source.G = DarkenUnit(source.G, target.G, factor)
	source.B = DarkenUnit(source.B, target.B, factor)
	return source
}

// R = min(S, D)
func DarkenUnit(S uint8, D uint8, _ float64) uint8 {
	if S < D{
		return S
	}else {
		return D
	}
}
