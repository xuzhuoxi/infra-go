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
	RegisterBlendFunc(Exclusion, BlendExclusionColor, BlendExclusionRGBA)
}

// BlendExclusionColor
// 排除模式
// 可生成和差值模式相似的效果，但比差值模式生成的颜色对比度较小，因而颜色较柔和。与白色混合将使底色反相；与黑色混合则不产生变化。
// R = S + D - 2*S*D [0,1]
// R = S + D - S*D/128 [0,255]
// R = S + D - S*D/32768 [0,65535]
func BlendExclusionColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendExclusionRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendExclusionRGBA
// 排除模式
// 可生成和差值模式相似的效果，但比差值模式生成的颜色对比度较小，因而颜色较柔和。与白色混合将使底色反相；与黑色混合则不产生变化。
// R = S + D - 2*S*D [0,1]
// R = S + D - S*D/128 [0,255]
// R = S + D - S*D/32768 [0,65535]
func BlendExclusionRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = exclusion(Sr, Dr)
	G = exclusion(Sg, Dg)
	B = exclusion(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = exclusion(Sa, Da)
	}
	return
}

// R = S + D - 2*S*D [0,1]
// R = S + D - S*D/128 [0,255]
// R = S + D - S*D/32768 [0,65535]
func exclusion(S, D uint32) uint32 {
	return S + D - S*D/32768
}
