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
	RegisterBlendFunc(Xor, XorBlend)
}

// 异或模式
// R = S*(1 - Da) + D*(1 - Sa)
// R = S*(255 - Da)/255 + D*(255 - Sa)/255
func XorBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = XorUnit(source.A, target.A, Sa, Da, 0)
	}
	source.R = XorUnit(source.R, target.R, Sa, Da, 0)
	source.G = XorUnit(source.G, target.G, Sa, Da, 0)
	source.B = XorUnit(source.B, target.B, Sa, Da, 0)
	return source
}

// R = S*(1 - Da) + D*(1 - Sa)
// R = S*(255 - Da)/255 + D*(255 - Sa)/255
func XorUnit(S uint8, D uint8, Sa uint8, Da uint8, factor float64) uint8 {
	return uint8(uint16(S*(255-Da))/255 + uint16(D*(255-Sa)/255))
}
