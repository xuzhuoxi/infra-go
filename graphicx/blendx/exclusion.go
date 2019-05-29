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
	RegisterBlendFunc(Exclusion, ExclusionBlend)
}

// 排除模式
// 可生成和差值模式相似的效果，但比差值模式生成的颜色对比度较小，因而颜色较柔和。与白色混合将使底色反相；与黑色混合则不产生变化。
// R = S + D - 2*S*D [0,1]
// R = S + D - S*D/128 [0,255]
func ExclusionBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ExclusionUnit(source.A, target.A, factor)
	}
	source.R = ExclusionUnit(source.R, target.R, factor)
	source.G = ExclusionUnit(source.G, target.G, factor)
	source.B = ExclusionUnit(source.B, target.B, factor)
	return source
}

// R = S + D - 2*S*D [0,1]
// R = S + D - S*D/128 [0,255]
func ExclusionUnit(S uint8, D uint8, _ float64) uint8 {
	temp := uint16(S) + uint16(D) - (uint16(S)*uint16(D))>>7
	return uint8(temp)
}
