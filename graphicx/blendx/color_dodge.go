// Package blendx
// Created by xuzhuoxi
// on 2019-05-25.
// @author xuzhuoxi
//
package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(ColorDodge, BlendColorDodgeColor, BlendColorDodgeRGBA)
}

// BlendColorDodgeColor
// 颜色减淡模式
// 查看每个通道的颜色信息，通过降低“对比度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// 除了指定在这个模式的层上边缘区域更尖锐，以及在这个模式下着色的笔画之外， Color Dodge模式类似于Screen模式创建的效果。另外，不管何时定义color Dodge模式混合 前景与背景像素，背景图像上的暗区域都将会消失。
// R = S + S*D/(1-D)
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func BlendColorDodgeColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendAddRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.NRGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendColorDodgeRGBA
// 颜色减淡模式
// 查看每个通道的颜色信息，通过降低“对比度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// 除了指定在这个模式的层上边缘区域更尖锐，以及在这个模式下着色的笔画之外， Color Dodge模式类似于Screen模式创建的效果。另外，不管何时定义color Dodge模式混合 前景与背景像素，背景图像上的暗区域都将会消失。
// R = S + S*D/(1-D)
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func BlendColorDodgeRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = colorDodge(Sr, Dr)
	G = colorDodge(Sg, Dg)
	B = colorDodge(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = colorDodge(Sa, Da)
	}
	return
}

// R = S + S*D/(1-D)
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func colorDodge(S, D uint32) uint32 {
	return S + S*D/(65535-D)
}
