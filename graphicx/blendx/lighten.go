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
// R = Max(S, D)
func BlendLightenColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLightenRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 变亮模式
// 与Darken相反，取两个像素中更亮的作为结果。
// 查看每个通道的颜色信息，并按照像素对比两个颜色，那个更亮，便以这种颜色作为此像素最终的颜色，也就是取两个颜色中的亮色作为最终色。绘图色中亮于底色的颜色被保留，暗于底色的颜色被替换。
// R = Max(S, D)
func BlendLightenRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = lighten(Sr, Dr)
	G = lighten(Sg, Dg)
	B = lighten(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = lighten(Sa, Da)
	}
	return
}

// R = Max(S, D)
func lighten(S, D uint32) uint32 {
	if S > D {
		return S
	} else {
		return D
	}
}
