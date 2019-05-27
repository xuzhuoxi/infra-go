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
	RegisterBlendFunc(HardLight, HardLightBlend)
}

// 强光模式
// (D<=0.5): R = 2*S*D
// (D>0.5): R = 1 - 2*(1 - S)*(1 - D)
//
// (D<=128): R = S*D/128
// (D>128): R = 255 - (255 - S) * (255 - D) / 128
func HardLightBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = HardLightUnit(source.A, target.A, factor)
	}
	source.R = HardLightUnit(source.R, target.R, factor)
	source.G = HardLightUnit(source.G, target.G, factor)
	source.B = HardLightUnit(source.B, target.B, factor)
	return source
}

// (D<=0.5): R = 2*S*D
// (D>0.5): R = 1 - 2*(1 - S)*(1 - D)
//
// (D<=128): R = S*D/128
// (D>128): R = 255 - (255 - S) * (255 - D) / 128
func HardLightUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	if D <= 128 {
		return uint8((S16 * D16) >> 7)
	} else {
		return uint8(255 - ((255-S16)*(255-D16))>>7)
	}
}
