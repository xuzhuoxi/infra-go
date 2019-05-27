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
	RegisterBlendFunc(Difference, DifferenceBlend)
}

// 差值模式
// R = R = |S - D|
func DifferenceBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = DifferenceUnit(source.A, target.A, factor)
	}
	source.R = DifferenceUnit(source.R, target.R, factor)
	source.G = DifferenceUnit(source.G, target.G, factor)
	source.B = DifferenceUnit(source.B, target.B, factor)
	return source
}

// R = |S - D| 
func DifferenceUnit(S uint8, D uint8, _ float64) uint8 {
	if S > D {
		return S - D
	} else {
		return D - S
	}
}
