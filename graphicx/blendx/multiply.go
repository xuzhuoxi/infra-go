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
	RegisterBlendFunc(Multiply, MultiplyBlend)
}

// 正片叠底
// R = S*D 		[0,1]
// R = S*D/255	[0,255]
func MultiplyBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = MultiplyUnit(source.A, target.A, factor)
	}
	source.R = MultiplyUnit(source.R, target.R, factor)
	source.G = MultiplyUnit(source.G, target.G, factor)
	source.B = MultiplyUnit(source.B, target.B, factor)
	return source
}

// R = S*D 		[0,1]
// R = S*D/255	[0,255]
func MultiplyUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(uint16(S) * uint16(D) / 255)
}
