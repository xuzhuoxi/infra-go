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
	RegisterBlendFunc(HardMix, HardMixBlend)
}

// 实色混合模式
// S+D>=255 : R = 255
// S+D<255 : R = 0
func HardMixBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = HardMixUnit(source.A, target.A, factor)
	}
	source.R = HardMixUnit(source.R, target.R, factor)
	source.G = HardMixUnit(source.G, target.G, factor)
	source.B = HardMixUnit(source.B, target.B, factor)
	return source
}

// S+D>=255 : R = 255
// S+D<255 : R = 0
func HardMixUnit(S uint8, D uint8, _ float64) uint8 {
	if uint16(S)+uint16(D) >= 255 {
		return 255
	} else {
		return 0
	}
}
