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
	RegisterBlendFunc(Darken, BlendDarkenColor, BlendDarkenRGBA)
}

// 变暗模式
// 与Lighten相反，将两个图像中更暗的那个被选来作为结果。
// 用于查找各颜色通道内的颜色信息，并按照像素对比底色和绘图色，那个更暗，便以这种颜色作为此图像最终的颜色，也就是取两个颜色中的暗色作为最终色。亮于底色的颜色被替换，暗于底色的颜色保持不变。
// R = min(S, D)
func BlendDarkenColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDarkenRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 变暗模式
// 与Lighten相反，将两个图像中更暗的那个被选来作为结果。
// 用于查找各颜色通道内的颜色信息，并按照像素对比底色和绘图色，那个更暗，便以这种颜色作为此图像最终的颜色，也就是取两个颜色中的暗色作为最终色。亮于底色的颜色被替换，暗于底色的颜色保持不变。
// R = min(S, D)
func BlendDarkenRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = darken(foreR, backR)
	G = darken(foreG, backG)
	B = darken(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = darken(foreA, backA)
	}
	return
}

// R = min(S, D)
func darken(F, B uint32) uint32 {
	if F < B {
		return F
	} else {
		return B
	}
}
