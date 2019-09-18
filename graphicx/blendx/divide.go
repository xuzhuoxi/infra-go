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
	RegisterBlendFunc(Divide, BlendDivideColor, BlendDivideRGBA)
}

// 划分模式
// 假设上面图层选择划分，那么所看到的图像是，下面的可见图层根据上面这个图层颜色的纯度，相应减去了同等纯度的该颜色，同时上面颜色的明暗度不同，被减去区域图像明度也不同，上面图层颜色的亮，图像亮度变化就会越小，上面图层越暗，被减区域图像就会越亮。
// 也就是说，如果上面图层是白色，那么也不会减去颜色也不会提高明度，如果上面图层是黑色，那么所有不纯的颜色都会被减去，只留着最纯的光的三原色，及其混合色，青品黄与白色。
// R = 255 * B / F
// R = 65535 * B / F
func BlendDivideColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDivideRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 划分模式
// 假设上面图层选择划分，那么所看到的图像是，下面的可见图层根据上面这个图层颜色的纯度，相应减去了同等纯度的该颜色，同时上面颜色的明暗度不同，被减去区域图像明度也不同，上面图层颜色的亮，图像亮度变化就会越小，上面图层越暗，被减区域图像就会越亮。
// 也就是说，如果上面图层是白色，那么也不会减去颜色也不会提高明度，如果上面图层是黑色，那么所有不纯的颜色都会被减去，只留着最纯的光的三原色，及其混合色，青品黄与白色。
// R = 255 * B / F
// R = 65535 * B / F
func BlendDivideRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = divide(foreR, backR)
	G = divide(foreG, backG)
	B = divide(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = divide(foreA, backA)
	}
	return
}

// R = 255 * B / F
// R = 65535 * B / F
func divide(F, B uint32) uint32 {
	return 65535 * B / F
}
