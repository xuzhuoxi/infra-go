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
	RegisterBlendFunc(Subtract, SubtractBlend)
}

// 减去模式
// 是将原始图像与混合图像相对应的像素提取出来并将它们相减。
// C = Max(0,A-B)
func SubtractBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = SubtractUnit(source.A, target.A, factor)
	}
	source.R = SubtractUnit(source.R, target.R, factor)
	source.G = SubtractUnit(source.G, target.G, factor)
	source.B = SubtractUnit(source.B, target.B, factor)
	return source
}

// C = Max(0,A-B)
func SubtractUnit(S uint8, D uint8, _ float64) uint8 {
	if S > D {
		return S - D
	} else {
		return 0
	}
}
