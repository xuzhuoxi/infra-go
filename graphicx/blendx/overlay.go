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
	RegisterBlendFunc(Overlay, OverlayBlend)
}

// 叠加模式
// S <= 128 : R = S * D / 128
// S > 128 : R = 255 - (255 - S) * (255 - D) / 128
func OverlayBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = OverlayUnit(source.A, target.A, factor)
	}
	source.R = OverlayUnit(source.R, target.R, factor)
	source.G = OverlayUnit(source.G, target.G, factor)
	source.B = OverlayUnit(source.B, target.B, factor)
	return source
}

// S <= 128 : R = S * D / 128
// S > 128 : R = 255 - (255 - S) * (255 - D) / 128
func OverlayUnit(S uint8, D uint8, _ float64) uint8 {
	if S <= 128 {
		return uint8((uint16(S) * uint16(D)) >> 7)
	} else {
		return uint8(255 - (uint16(255-S)*uint16(255-D))>>7)
	}
}
