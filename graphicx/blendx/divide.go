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
	RegisterBlendFunc(Divide, BlendDivideColor, BlendDivideRGBA)
}

// BlendDivideColor
// 划分模式
// 假设上面图层选择划分，那么所看到的图像是，下面的可见图层根据上面这个图层颜色的纯度，相应减去了同等纯度的该颜色，同时上面颜色的明暗度不同，被减去区域图像明度也不同，上面图层颜色的亮，图像亮度变化就会越小，上面图层越暗，被减区域图像就会越亮。
// 也就是说，如果上面图层是白色，那么也不会减去颜色也不会提高明度，如果上面图层是黑色，那么所有不纯的颜色都会被减去，只留着最纯的光的三原色，及其混合色，青品黄与白色。
// R = S / D [0, 1]
// R = 255 * S / D [0, 255]
// R = 65535 * S / D [0, 65535]
func BlendDivideColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDivideRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendDivideRGBA
// 划分模式
// 假设上面图层选择划分，那么所看到的图像是，下面的可见图层根据上面这个图层颜色的纯度，相应减去了同等纯度的该颜色，同时上面颜色的明暗度不同，被减去区域图像明度也不同，上面图层颜色的亮，图像亮度变化就会越小，上面图层越暗，被减区域图像就会越亮。
// 也就是说，如果上面图层是白色，那么也不会减去颜色也不会提高明度，如果上面图层是黑色，那么所有不纯的颜色都会被减去，只留着最纯的光的三原色，及其混合色，青品黄与白色。
// R = S / D [0, 1]
// R = 255 * S / D [0, 255]
// R = 65535 * S / D [0, 65535]
func BlendDivideRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = divide(Sr, Dr)
	G = divide(Sg, Dg)
	B = divide(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = divide(Sa, Da)
	}
	return
}

// R = S / D [0, 1]
// R = 255 * S / D [0, 255]
// R = 65535 * S / D [0, 65535]
func divide(S, D uint32) uint32 {
	return 65535 * S / D
}
