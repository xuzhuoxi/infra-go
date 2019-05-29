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
	RegisterBlendFunc(ColorDodge, ColorDodgeBlend)
}

// 颜色减淡模式
// 查看每个通道的颜色信息，通过降低“对比度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// 除了指定在这个模式的层上边缘区域更尖锐，以及在这个模式下着色的笔画之外， Color Dodge模式类似于Screen模式创建的效果。另外，不管何时定义color Dodge模式混合 前景与背景像素，背景图像上的暗区域都将会消失。
// R = S + (S*D)/(1-D)
// R = S + S*D/(255-D)
func ColorDodgeBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ColorDodgeUnit(source.A, target.A, factor)
	}
	source.R = ColorDodgeUnit(source.R, target.R, factor)
	source.G = ColorDodgeUnit(source.G, target.G, factor)
	source.B = ColorDodgeUnit(source.B, target.B, factor)
	return source
}

// R = S + (S*D)/(1-D)
// R = S + S*D/(255-D)
func ColorDodgeUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	return uint8(S16 + S16*D16/(255-D16))
}
