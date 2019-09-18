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
	RegisterBlendFunc(VividLight, BlendVividLightColor, BlendVividLightRGBA)
}

// 亮光模式
// 根据绘图色通过增加或降低“对比度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过降低对比度被照亮，如果绘图色比50%的灰暗，图像通过增加对比度变暗。
// (F<=128): R = B - (255-B)*(255-2*F) / (2*F)
// (F>128): R = B + B*(2*F-255)/(2*(255-F))
//
// (F<=32768): R = B - (65535-B)*(65535-2*F) / (2*F)
// (F>32768): R = B + B*(2*F-65535)/(2*(65535-F))
func BlendVividLightColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendVividLightRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 亮光模式
// 根据绘图色通过增加或降低“对比度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过降低对比度被照亮，如果绘图色比50%的灰暗，图像通过增加对比度变暗。
// (F<=128): R = B - (255-B)*(255-2*F) / (2*F)
// (F>128): R = B + B*(2*F-255)/(2*(255-F))
//
// (F<=32768): R = B - (65535-B)*(65535-2*F) / (2*F)
// (F>32768): R = B + B*(2*F-65535)/(2*(65535-F))
func BlendVividLightRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = vividLight(foreR, backR)
	G = vividLight(foreG, backG)
	B = vividLight(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = vividLight(foreA, backA)
	}
	return
}

// (F<=128): R = B - (255-B)*(255-2*F) / (2*F)
// (F>128): R = B + B*(2*F-255)/(2*(255-F))
//
// (F<=32768): R = B - (65535-B)*(65535-2*F) / (2*F)
// (F>32768): R = B + B*(2*F-65535)/(2*(65535-F))
func vividLight(F, B uint32) uint32 {
	if F <= 32768 {
		return B - (65535-B)*(65535-2*F)/(2*F)
	} else {
		return B + B*(2*F-65535)/(2*(65535-F))
	}
}
