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
	RegisterBlendFunc(Lighten, LightenBlend)
}

// 变亮模式
// 与Darken相反，取两个像素中更亮的作为结果。
// 查看每个通道的颜色信息，并按照像素对比两个颜色，那个更亮，便以这种颜色作为此像素最终的颜色，也就是取两个颜色中的亮色作为最终色。绘图色中亮于底色的颜色被保留，暗于底色的颜色被替换。
// R = max(S, D)
func LightenBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = LightenUnit(source.A, target.A, factor)
	}
	source.R = LightenUnit(source.R, target.R, factor)
	source.G = LightenUnit(source.G, target.G, factor)
	source.B = LightenUnit(source.B, target.B, factor)
	return source
}

// R = Max(S, D)
func LightenUnit(S uint8, D uint8, _ float64) uint8 {
	if S > D {
		return S
	} else {
		return D
	}
}
