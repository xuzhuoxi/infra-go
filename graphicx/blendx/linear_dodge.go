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
	RegisterBlendFunc(LinearDodge, BlendLinearDodgeColor, BlendLinearDodgeRGBA)
}

// BlendLinearDodgeColor
// 线性减淡模式
// 查看每个通道的颜色信息，通过增加“亮度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// R = S + D
func BlendLinearDodgeColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLinearDodgeRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendLinearDodgeRGBA
// 线性减淡模式
// 查看每个通道的颜色信息，通过增加“亮度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// R = S + D
func BlendLinearDodgeRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = linearDodge(Sr, Dr)
	G = linearDodge(Sg, Dg)
	B = linearDodge(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = linearDodge(Sa, Da)
	}
	return
}

// R = S + D
func linearDodge(S, D uint32) uint32 {
	Add := S + D
	if Add <= 65535 {
		return Add
	} else {
		return 255
	}
}
