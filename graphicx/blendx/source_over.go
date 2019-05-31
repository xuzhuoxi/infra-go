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
	RegisterBlendFunc(SourceOver, SourceOverBlend)
}

//
// R = S + D*(1 - Sa)
// R = S + D*(255 - Sa)/255
func SourceOverBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = SourceOverUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = SourceOverUnit(source.R, target.R, Sa, Da, factor)
	source.G = SourceOverUnit(source.G, target.G, Sa, Da, factor)
	source.B = SourceOverUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = S + D*(1 - Sa)
// R = S + D*(255 - Sa)/255
func SourceOverUnit(S uint8, D uint8, Sa uint8, Da uint8, _ float64) uint8 {
	return uint8(uint16(S) + uint16(D)*(255-uint16(Sa))/255)
}
