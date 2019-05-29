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
	RegisterBlendFunc(VividLight, VividLightBlend)
}

// 亮光模式
// 根据绘图色通过增加或降低“对比度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过降低对比度被照亮，如果绘图色比50%的灰暗，图像通过增加对比度变暗。
// (D<=128): R = S - (255-S)*(255-2*D) / (2*D)
// (D>128): R = S + S * (2*D - 255) / (2*(255-D))
func VividLightBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = VividLightUnit(source.A, target.A, factor)
	}
	source.R = VividLightUnit(source.R, target.R, factor)
	source.G = VividLightUnit(source.G, target.G, factor)
	source.B = VividLightUnit(source.B, target.B, factor)
	return source
}

// (D<=128): R = S - (255-S)*(255-2*D) / (2*D)
// (D>128): R = S + S * (2*D - 255) / (2*(255-D))
func VividLightUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	if D <= 128 {
		return uint8(S16 - (255-S16)*(255-2*D16)/(2*D16))
	} else {
		return uint8(S16 + (S16*(2*D16-255)/(255-D16))>>1)
	}
}
