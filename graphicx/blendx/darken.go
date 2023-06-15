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
	RegisterBlendFunc(Darken, BlendDarkenColor, BlendDarkenRGBA)
}

// BlendDarkenColor
// 变暗模式
// 与Lighten相反，将两个图像中更暗的那个被选来作为结果。
// 用于查找各颜色通道内的颜色信息，并按照像素对比底色和绘图色，那个更暗，便以这种颜色作为此图像最终的颜色，也就是取两个颜色中的暗色作为最终色。亮于底色的颜色被替换，暗于底色的颜色保持不变。
// R = min(S, D)
func BlendDarkenColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDarkenRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendDarkenRGBA
// 变暗模式
// 与Lighten相反，将两个图像中更暗的那个被选来作为结果。
// 用于查找各颜色通道内的颜色信息，并按照像素对比底色和绘图色，那个更暗，便以这种颜色作为此图像最终的颜色，也就是取两个颜色中的暗色作为最终色。亮于底色的颜色被替换，暗于底色的颜色保持不变。
// R = min(S, D)
func BlendDarkenRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = darken(Sr, Dr)
	G = darken(Sg, Dg)
	B = darken(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = darken(Sa, Da)
	}
	return
}

// R = min(S, D)
func darken(S, D uint32) uint32 {
	if S < D {
		return S
	} else {
		return D
	}
}
