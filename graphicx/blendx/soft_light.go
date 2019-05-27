//
//Created by xuzhuoxi
//on 2019-05-26.
//@author xuzhuoxi
//
package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(SoftLight, SoftLightBlend)
}

// 柔光模式
// (D<=0.5): R = 2*S*D + S*S*(1 - 2*D)
// (D>0.5): R = 2*S*(1 - D) + (2*D - 1)*√S
//
// (D<=128): R = S*D/128 + (255-2*D)*S*S/65025
// (D>128): R = S*(255-B)/128 + (2*D-255)*√(S/255)
func SoftLightBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = SoftLightUnit(source.A, target.A, factor)
	}
	source.R = SoftLightUnit(source.R, target.R, factor)
	source.G = SoftLightUnit(source.G, target.G, factor)
	source.B = SoftLightUnit(source.B, target.B, factor)
	return source
}

// (D<=0.5): R = 2*S*D + S*S*(1 - 2*D)
// (D>0.5): R = 2*S*(1 - D) + (2*D - 1)*√S
//
// (D<=128): R = S*D/128 + (255-2*D)*S*S/65025
// (D>128): R = S*(255-B)/128 + (2*D-255)*√(S/255)
func SoftLightUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	var temp uint16
	if D <= 128 {
		temp = S16*D16/128 + (255-2*D16)*S16*S16/65025
	} else {
		temp = S16*(255-D16)/128 + (2*D16-255)*uint16(math.Sqrt(float64(S)/255))
	}
	return uint8(temp)
}
