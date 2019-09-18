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
	RegisterBlendFunc(Lighten, BlendLightenColor, BlendLightenRGBA)
}

// 变亮模式
// 与Darken相反，取两个像素中更亮的作为结果。
// 查看每个通道的颜色信息，并按照像素对比两个颜色，那个更亮，便以这种颜色作为此像素最终的颜色，也就是取两个颜色中的亮色作为最终色。绘图色中亮于底色的颜色被保留，暗于底色的颜色被替换。
// R = Max(F, B)
func BlendLightenColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendLightenRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 变亮模式
// 与Darken相反，取两个像素中更亮的作为结果。
// 查看每个通道的颜色信息，并按照像素对比两个颜色，那个更亮，便以这种颜色作为此像素最终的颜色，也就是取两个颜色中的亮色作为最终色。绘图色中亮于底色的颜色被保留，暗于底色的颜色被替换。
// R = Max(F, B)
func BlendLightenRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = lighten(foreR, backR)
	G = lighten(foreG, backG)
	B = lighten(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = lighten(foreA, backA)
	}
	return
}

// R = Max(F, B)
func lighten(F, B uint32) uint32 {
	if F > B {
		return F
	} else {
		return B
	}
}
