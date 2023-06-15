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
	RegisterBlendFunc(Add, BlendAddColor, BlendAddRGBA)
}

// BlendAddColor
// 增加模式
// 是将原始图像及混合图像的对应像素取出来并加在一起；
// R = Min(1, S+D))
// R = Min(255, S+D)
// R = Min(65535, S+D)
func BlendAddColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendAddRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendAddRGBA
// 增加模式
// 是将原始图像及混合图像的对应像素取出来并加在一起；
// R = Min(1, S+D))
// R = Min(255, S+D)
// R = Min(65535, S+D)
func BlendAddRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = blendAdd(Sr, Dr)
	G = blendAdd(Sg, Dg)
	B = blendAdd(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = blendAdd(Sa, Da)
	}
	return
}

// R = Min(1, S+D))
// R = Min(255, S+D)
// R = Min(65535, S+D)
func blendAdd(S, D uint32) uint32 {
	add := S + D
	if add < 65535 {
		return add
	} else {
		return 65535
	}
}
