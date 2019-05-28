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
	RegisterBlendFunc(DestinationOut, DestinationOutBlend)
}

//
// R = D*(1 - Sa)
// R = D*(255-Sa)/255
func DestinationOutBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	Sa := source.A
	Da := target.A
	if !keepAlpha {
		source.A = DestinationOutUnit(source.A, target.A, Sa, Da, factor)
	}
	source.R = DestinationOutUnit(source.R, target.R, Sa, Da, factor)
	source.G = DestinationOutUnit(source.G, target.G, Sa, Da, factor)
	source.B = DestinationOutUnit(source.B, target.B, Sa, Da, factor)
	return source
}

// R = D*(1 - Sa) [0,1]
// R = D*(255-Sa)/255 [0,255]
func DestinationOutUnit(_ uint8, D uint8, Sa uint8, _ uint8, _ float64) uint8 {
	return uint8(uint16(D) * (255 - uint16(Sa)) / 255)
}
