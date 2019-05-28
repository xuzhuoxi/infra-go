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
	RegisterBlendFunc(SourceIn, SourceInBlend)
}

//
// R = S*Da
// R = S*Da/255
func SourceInBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = SourceInUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = SourceInUnit(source.R, target.R, Sa, Da, factor)
	source.G = SourceInUnit(source.G, target.G, Sa, Da, factor)
	source.B = SourceInUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = S*Da
// R = S*Da/255
func SourceInUnit(S uint8, _ uint8, _ uint8, Da uint8, _ float64) uint8 {
	return uint8(uint16(S) * uint16(Da) / 255)
}
