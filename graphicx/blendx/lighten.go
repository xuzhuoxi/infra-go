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
	RegisterBlendFunc(Lighten, LightenBlend)
}

// 变亮模式
// R = max(S, D)
func LightenBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = LightenUnit(source.A, target.A, factor)
	}
	source.R = LightenUnit(source.R, target.R, factor)
	source.G = LightenUnit(source.G, target.G, factor)
	source.B = LightenUnit(source.B, target.B, factor)
	return source
}

// R = Max(S, D)
func LightenUnit(S uint8, D uint8, _ float64) uint8 {
	if S > D {
		return S
	} else {
		return D
	}
}
