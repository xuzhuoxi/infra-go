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
	RegisterBlendFunc(SourceAtop, SourceAtopBlend)
}

//
// R = S*Da + D*(1 - Sa)
// R = S*Da/255 + D*(255 - Sa)/255
func SourceAtopBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = SourceAtopUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = SourceAtopUnit(source.R, target.R, Sa, Da, factor)
	source.G = SourceAtopUnit(source.G, target.G, Sa, Da, factor)
	source.B = SourceAtopUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = S*Da + D*(1 - Sa)
// R = S*Da/255 + D*(255 - Sa)/255
func SourceAtopUnit(S uint8, D uint8, Sa uint8, Da uint8, _ float64) uint8 {
	return uint8(uint16(S)*uint16(Da)/255 + uint16(D)*(255-uint16(Sa))/255)
}
