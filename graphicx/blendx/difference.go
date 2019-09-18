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
	RegisterBlendFunc(Difference, BlendDifferenceColor, BlendDifferenceRGBA)
}

// 差值模式
// 查看每个通道中的颜色信息，比较底色和绘图色，用较亮的像素点的像素值减去较暗的像素点的像素值。与白色混合将使底色反相；与黑色混合则不产生变化。
// R = |F - B|
func BlendDifferenceColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDifferenceRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 差值模式
// 查看每个通道中的颜色信息，比较底色和绘图色，用较亮的像素点的像素值减去较暗的像素点的像素值。与白色混合将使底色反相；与黑色混合则不产生变化。
// R = |F - B|
func BlendDifferenceRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = difference(foreR, backR)
	G = difference(foreG, backG)
	B = difference(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = difference(foreA, backA)
	}
	return
}

// R = |F - B|
func difference(F, B uint32) uint32 {
	if F > B {
		return F - B
	} else {
		return B - F
	}
}
