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
	RegisterBlendFunc(PlusDarker, PlusDarkerBlend)
}

//
// R = MAX(0, (1 - D) + (1 - S)) [0,1]
// R = MAX(0, (255 - D) + (255 - S)) [0,255]
func PlusDarkerBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = PlusDarkerUnit(source.A, target.A, factor)
	}
	source.R = PlusDarkerUnit(source.R, target.R, factor)
	source.G = PlusDarkerUnit(source.G, target.G, factor)
	source.B = PlusDarkerUnit(source.B, target.B, factor)
	return source
}

// R = MAX(0, (1 - D) + (1 - S)) [0,1]
// R = MAX(0, (255 - D) + (255 - S)) [0,255]
func PlusDarkerUnit(S uint8, D uint8, factor float64) uint8 {
	temp := 510 - uint16(S) - uint16(D)
	if temp > 0 {
		return uint8(temp)
	} else {
		return 0
	}
}
