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
	RegisterBlendFunc(LinearLight, BlendLinearLightColor, BlendLinearLightRGBA)
}

// 线性光模式
// 根据绘图色通过增加或降低“亮度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过增加亮度被照亮，如果绘图色比50%的灰暗，图像通过降低亮度变暗。
// R = B + 2*F - 255
// R = B + 2*F - 65535
func BlendLinearLightColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendLinearLightRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 线性光模式
// 根据绘图色通过增加或降低“亮度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过增加亮度被照亮，如果绘图色比50%的灰暗，图像通过降低亮度变暗。
// R = B + 2*F - 255
// R = B + 2*F - 65535
func BlendLinearLightRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = linearLight(foreR, backR)
	G = linearLight(foreG, backG)
	B = linearLight(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = linearLight(foreA, backA)
	}
	return
}

// R = B + 2*F - 255
// R = B + 2*F - 65535
func linearLight(F, B uint32) uint32 {
	return B + 2*F - 65535
}
