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
	RegisterBlendFunc(DestinationAtop, DestinationAtopBlend)
}

//
// R = S*(1 - Da) + D*Sa
// R = S*(255-Da)/255 + D*Sa/255/255
func DestinationAtopBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = DestinationAtopUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = DestinationAtopUnit(source.R, target.R, Sa, Da, factor)
	source.G = DestinationAtopUnit(source.G, target.G, Sa, Da, factor)
	source.B = DestinationAtopUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = S*(1 - Da) + D*Sa
// R = S*(255-Da)/255 + D*Sa/255/255
func DestinationAtopUnit(S uint8, D uint8, Sa uint8, Da uint8, _ float64) uint8 {
	return uint8(uint16(S)*(255-uint16(Da))/255 + uint16(D)*uint16(Sa)/65025)
}
