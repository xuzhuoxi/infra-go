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
	RegisterBlendFunc(PlusLighter, PlusLighterBlend)
}

//
// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
func PlusLighterBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = PlusLighterUnit(source.A, target.A, factor)
	}
	source.R = PlusLighterUnit(source.R, target.R, factor)
	source.G = PlusLighterUnit(source.G, target.G, factor)
	source.B = PlusLighterUnit(source.B, target.B, factor)
	return source
}

// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
func PlusLighterUnit(S uint8, D uint8, _ float64) uint8 {
	Add := uint16(S) + uint16(D)
	if Add > 255 {
		return 255
	} else {
		return uint8(Add)
	}
}
