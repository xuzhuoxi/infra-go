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
	RegisterBlendFunc(DestinationOver, DestinationOverBlend)
}

//
// R = S*(1 - Da) + D
// R = S*(255 - Da)/255 + D
func DestinationOverBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = DestinationOverUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = DestinationOverUnit(source.R, target.R, Sa, Da, factor)
	source.G = DestinationOverUnit(source.G, target.G, Sa, Da, factor)
	source.B = DestinationOverUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = S*(1 - Da) + D
// R = S*(255 - Da)/255 + D
func DestinationOverUnit(S uint8, D uint8, Sa uint8, Da uint8, _ float64) uint8 {
	return uint8(uint16(S)*(255-uint16(Da))/255 + uint16(D))
}
