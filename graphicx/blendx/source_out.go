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
	RegisterBlendFunc(SourceOut, SourceOutBlend)
}

//
// R = S*(1 - Da)
// R = S*(255 - Da)/255
func SourceOutBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = SourceOutUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = SourceOutUnit(source.R, target.R, Sa, Da, factor)
	source.G = SourceOutUnit(source.G, target.G, Sa, Da, factor)
	source.B = SourceOutUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = S*(1 - Da)
// R = S*(255 - Da)/255
func SourceOutUnit(S uint8, _ uint8, _ uint8, Da uint8, _ float64) uint8 {
	return uint8(uint16(S) * (255 - uint16(Da)) / 255)
}
