//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"github.com/xuzhuoxi/infra-go/mathx"
	"image/color"
)

func init() {
	RegisterBlendFunc(PinLight, BlendPinLightColor, BlendPinLightRGBA)
}

// 点光模式
// 根据绘图色替换颜色。如果绘图色比50%的灰要亮，绘图色被替换，比绘图色亮的像素不变化。如果绘图色比50%的灰要暗比绘图色亮的像素被替换，比绘图色暗的像素不变化，点光模式对图像增加特殊效果非常有用。
// R = F>0.5 ? Min(B, 2*F-1) : B
// R = F>128 ? Min(B, 2*F-255) : B
// R = F>32768 ? Min(B, 2*F-65535) : B
func BlendPinLightColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendPinLightRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 点光模式
// 根据绘图色替换颜色。如果绘图色比50%的灰要亮，绘图色被替换，比绘图色亮的像素不变化。如果绘图色比50%的灰要暗比绘图色亮的像素被替换，比绘图色暗的像素不变化，点光模式对图像增加特殊效果非常有用。
// R = F>0.5 ? Min(B, 2*F-1) : B
// R = F>128 ? Min(B, 2*F-255) : B
// R = F>32768 ? Min(B, 2*F-65535) : B
func BlendPinLightRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = pinLight(foreR, backR)
	G = pinLight(foreG, backG)
	B = pinLight(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = pinLight(foreA, backA)
	}
	return
}

// R = F>0.5 ? Min(B, 2*F-1) : B
// R = F>128 ? Min(B, 2*F-255) : B
// R = F>32768 ? Min(B, 2*F-65535) : B
func pinLight(F, B uint32) uint32 {
	if F <= 32768 {
		return B
	}
	return uint32(mathx.MinUint(uint(B), 2*uint(F)-65535))
}
