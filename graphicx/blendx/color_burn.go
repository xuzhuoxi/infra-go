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
	RegisterBlendFunc(ColorBurn, ColorBurnBlend)
}

// 颜色加深模式
// R = S - ((1 - S) * (1 - D)) / D
// R = S - (255-S)*(255-D) / D
func ColorBurnBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ColorBurnUnit(source.A, target.A, factor)
	}
	source.R = ColorBurnUnit(source.R, target.R, factor)
	source.G = ColorBurnUnit(source.G, target.G, factor)
	source.B = ColorBurnUnit(source.B, target.B, factor)
	return source
}

// R = S - ((1 - S) * (1 - D)) / D
// R = S - (255-S)*(255-D) / D
func ColorBurnUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	return uint8(S16 - (255-S16)*(255-D16)/D16)
}
