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
	RegisterBlendFunc(Exclusion, BlendExclusionColor, BlendExclusionRGBA)
}

// 排除模式
// 可生成和差值模式相似的效果，但比差值模式生成的颜色对比度较小，因而颜色较柔和。与白色混合将使底色反相；与黑色混合则不产生变化。
// R = B + F - 2*B*F [0,1]
// R = B + F - B*F/128 [0,255]
// R = B + F - B*F/32768 [0,65535]
func BlendExclusionColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendExclusionRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 排除模式
// 可生成和差值模式相似的效果，但比差值模式生成的颜色对比度较小，因而颜色较柔和。与白色混合将使底色反相；与黑色混合则不产生变化。
// R = B + F - 2*B*F [0,1]
// R = B + F - B*F/128 [0,255]
// R = B + F - B*F/32768 [0,65535]
func BlendExclusionRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = exclusion(foreR, backR)
	G = exclusion(foreG, backG)
	B = exclusion(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = exclusion(foreA, backA)
	}
	return
}

// R = B + F - 2*B*F [0,1]
// R = B + F - B*F/128 [0,255]
// R = B + F - B*F/32768 [0,65535]
func exclusion(F, B uint32) uint32 {
	return B + F - B*F/32768
}
