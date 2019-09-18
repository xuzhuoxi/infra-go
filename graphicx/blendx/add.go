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
	RegisterBlendFunc(Add, BlendAddColor, BlendAddRGBA)
}

// 增加模式
// 是将原始图像及混合图像的对应像素取出来并加在一起；
// R = Min(1, B+F))
// R = Min(255, B+F)
// R = Min(65535, B+F)
func BlendAddColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendAddRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 增加模式
// 是将原始图像及混合图像的对应像素取出来并加在一起；
// R = Min(1, B+F))
// R = Min(255, B+F)
// R = Min(65535, B+F)
func BlendAddRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = blendAdd(foreR, backR)
	G = blendAdd(foreG, backG)
	B = blendAdd(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = blendAdd(foreA, backA)
	}
	return
}

// R = Min(1, B+F))
// R = Min(255, B+F)
// R = Min(65535, B+F)
func blendAdd(F, B uint32) uint32 {
	add := B + F
	if add < 65535 {
		return add
	} else {
		return 65535
	}
}
