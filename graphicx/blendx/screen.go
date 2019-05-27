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
	RegisterBlendFunc(Screen, ScreenBlend)
}

// 滤色模式
// R = 1 - (1 - S)*(1 - D)
// R = 255 - (255-S)*(255-D) / 255
func ScreenBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ScreenUnit(source.A, target.A, factor)
	}
	source.R = ScreenUnit(source.R, target.R, factor)
	source.G = ScreenUnit(source.G, target.G, factor)
	source.B = ScreenUnit(source.B, target.B, factor)
	return source
}

// R = 1 - (1 - S)*(1 - D)
// R = 255 - (255-S)*(255-D) / 255
func ScreenUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	return uint8(255 - (255-S16)*(255-D16)/255)
}
