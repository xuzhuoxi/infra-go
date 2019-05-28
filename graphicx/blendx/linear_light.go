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
	RegisterBlendFunc(LinearLight, LinearLightBlend)
}

// 线性光模式
// R = S + 2 * D - 255
func LinearLightBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = LinearLightUnit(source.A, target.A, factor)
	}
	source.R = LinearLightUnit(source.R, target.R, factor)
	source.G = LinearLightUnit(source.G, target.G, factor)
	source.B = LinearLightUnit(source.B, target.B, factor)
	return source
}

// R = S + 2 * D - 255
func LinearLightUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(uint16(S) + 2*uint16(D) - 255)
}
