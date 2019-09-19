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
	RegisterBlendFunc(Subtract, BlendSubtractColor, BlendSubtractRGBA)
}

// 减去模式
// 是将原始图像与混合图像相对应的像素提取出来并将它们相减。
// C = Max(0,S-D)
func BlendSubtractColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendSubtractRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 减去模式
// 是将原始图像与混合图像相对应的像素提取出来并将它们相减。
// C = Max(0,S-D)
func BlendSubtractRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = subtract(Sr, Dr)
	G = subtract(Sg, Dg)
	B = subtract(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = subtract(Sa, Da)
	}
	return
}

// C = Max(0,S-D)
func subtract(S, D uint32) uint32 {
	if S > D {
		return S - D
	} else {
		return 0
	}
}
