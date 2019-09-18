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
	RegisterBlendFunc(ColorDodge, BlendColorDodgeColor, BlendColorDodgeRGBA)
}

// 颜色减淡模式
// 查看每个通道的颜色信息，通过降低“对比度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// 除了指定在这个模式的层上边缘区域更尖锐，以及在这个模式下着色的笔画之外， Color Dodge模式类似于Screen模式创建的效果。另外，不管何时定义color Dodge模式混合 前景与背景像素，背景图像上的暗区域都将会消失。
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func BlendColorDodgeColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendAddRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.NRGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 颜色减淡模式
// 查看每个通道的颜色信息，通过降低“对比度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// 除了指定在这个模式的层上边缘区域更尖锐，以及在这个模式下着色的笔画之外， Color Dodge模式类似于Screen模式创建的效果。另外，不管何时定义color Dodge模式混合 前景与背景像素，背景图像上的暗区域都将会消失。
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func BlendColorDodgeRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, factor float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = colorDodge(foreR, backR)
	G = colorDodge(foreG, backG)
	B = colorDodge(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = colorDodge(foreA, backA)
	}
	return
}

// R = B + B*F/(1-F)
// R = B + B*F/(255-F) (32位)
// R = B + B*F/(65535-F) (64位)
func colorDodge(F, B uint32) uint32 {
	return B + B*F/(65535-F)
}
