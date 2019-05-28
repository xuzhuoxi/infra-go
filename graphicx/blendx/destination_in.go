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
	RegisterBlendFunc(DestinationIn, DestinationInBlend)
}

//
// R = D*Sa [0,1]
// R = D*Sa/255 [0,255]
func DestinationInBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = DestinationInUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = DestinationInUnit(source.R, target.R, Sa, Da, factor)
	source.G = DestinationInUnit(source.G, target.G, Sa, Da, factor)
	source.B = DestinationInUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = D*Sa [0,1]
// R = D*Sa/255 [0,255]
func DestinationInUnit(_ uint8, D uint8, Sa uint8, _ uint8, _ float64) uint8 {
	return uint8(uint16(D) * uint16(Sa) / 255)
}
