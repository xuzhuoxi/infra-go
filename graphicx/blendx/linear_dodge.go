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
	RegisterBlendFunc(LinearDodge, BlendLinearDodgeColor, BlendLinearDodgeRGBA)
}

// 线性减淡模式
// 查看每个通道的颜色信息，通过增加“亮度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// R = B + F
func BlendLinearDodgeColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendLinearDodgeRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 线性减淡模式
// 查看每个通道的颜色信息，通过增加“亮度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// R = B + F
func BlendLinearDodgeRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = linearDodge(foreR, backR)
	G = linearDodge(foreG, backG)
	B = linearDodge(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = linearDodge(foreA, backA)
	}
	return
}

// R = B + F
func linearDodge(F, B uint32) uint32 {
	Add := B + F
	if Add <= 65535 {
		return Add
	} else {
		return 255
	}
}
