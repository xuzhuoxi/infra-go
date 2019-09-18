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
// C = Max(0,B-F)
func BlendSubtractColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendSubtractRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 减去模式
// 是将原始图像与混合图像相对应的像素提取出来并将它们相减。
// C = Max(0,B-F)
func BlendSubtractRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = subtract(foreR, backR)
	G = subtract(foreG, backG)
	B = subtract(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = subtract(foreA, backA)
	}
	return
}

// C = Max(0,B-F)
func subtract(F, B uint32) uint32 {
	if B > F {
		return B - F
	} else {
		return 0
	}
}
