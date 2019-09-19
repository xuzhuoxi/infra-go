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
	RegisterBlendFunc(PinLight, BlendPinLightColor, BlendPinLightRGBA)
}

// 点光模式
// 根据绘图色替换颜色。如果绘图色比50%的灰要亮，绘图色被替换，比绘图色亮的像素不变化。如果绘图色比50%的灰要暗比绘图色亮的像素被替换，比绘图色暗的像素不变化，点光模式对图像增加特殊效果非常有用。
// R = D>0.5 ? Min(S, 2*D-1) : S
// R = D>128 ? Min(S, 2*D-255) : S
// R = D>32768 ? Min(S, 2*D-65535) : S
func BlendPinLightColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendPinLightRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 点光模式
// 根据绘图色替换颜色。如果绘图色比50%的灰要亮，绘图色被替换，比绘图色亮的像素不变化。如果绘图色比50%的灰要暗比绘图色亮的像素被替换，比绘图色暗的像素不变化，点光模式对图像增加特殊效果非常有用。
// R = D>0.5 ? Min(S, 2*D-1) : S
// R = D>128 ? Min(S, 2*D-255) : S
// R = D>32768 ? Min(S, 2*D-65535) : S
func BlendPinLightRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = pinLight(Sr, Dr)
	G = pinLight(Sg, Dg)
	B = pinLight(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = pinLight(Sa, Da)
	}
	return
}

// R = D>0.5 ? Min(S, 2*D-1) : S
// R = D>128 ? Min(S, 2*D-255) : S
// R = D>32768 ? Min(S, 2*D-65535) : S
func pinLight(S, D uint32) uint32 {
	if D > 32768 {
		return S
	} else {
		result := 2*D - 65535
		if S < result {
			return S
		} else {
			return result
		}
	}
}
