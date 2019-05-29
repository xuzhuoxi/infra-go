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
	RegisterBlendFunc(Darken, DarkenBlend)
}

// 变暗模式
// 与Lighten相反，将两个图像中更暗的那个被选来作为结果。
// 用于查找各颜色通道内的颜色信息，并按照像素对比底色和绘图色，那个更暗，便以这种颜色作为此图像最终的颜色，也就是取两个颜色中的暗色作为最终色。亮于底色的颜色被替换，暗于底色的颜色保持不变。
// R = min(S, D)
func DarkenBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = DarkenUnit(source.A, target.A, factor)
	}
	source.R = DarkenUnit(source.R, target.R, factor)
	source.G = DarkenUnit(source.G, target.G, factor)
	source.B = DarkenUnit(source.B, target.B, factor)
	return source
}

// R = min(S, D)
func DarkenUnit(S uint8, D uint8, _ float64) uint8 {
	if S < D{
		return S
	}else {
		return D
	}
}
