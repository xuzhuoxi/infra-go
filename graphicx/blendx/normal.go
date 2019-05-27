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
	RegisterBlendFunc(Normal, NormalBlend)
}

// 正常模式
// R = S * factor + D * (1-factor)
func NormalBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = NormalUnit(source.A, target.A, factor)
	}
	source.R = NormalUnit(source.R, target.R, factor)
	source.G = NormalUnit(source.G, target.G, factor)
	source.B = NormalUnit(source.B, target.B, factor)
	return source
}

// R = S * factor + D * (1-factor)
func NormalUnit(S uint8, D uint8, factor float64) uint8 {
	if factor >= 0 && factor <= 1 {
		return uint8(float64(S)*factor + float64(D)*(1-factor))
	} else {
		return S
	}
}
