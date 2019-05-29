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
	RegisterBlendFunc(PinLight, PinLightBlend)
}

// 点光模式
// 根据绘图色替换颜色。如果绘图色比50%的灰要亮，绘图色被替换，比绘图色亮的像素不变化。如果绘图色比50%的灰要暗比绘图色亮的像素被替换，比绘图色暗的像素不变化，点光模式对图像增加特殊效果非常有用。
// D <=128 : R = D
// D >128 : R = Min(S, 2*D-255)
func PinLightBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = PinLightUnit(source.A, target.A, factor)
	}
	source.R = PinLightUnit(source.R, target.R, factor)
	source.G = PinLightUnit(source.G, target.G, factor)
	source.B = PinLightUnit(source.B, target.B, factor)
	return source
}

// D <=128 : R = D
// D >128 : R = Min(S, 2*D-255)
func PinLightUnit(S uint8, D uint8, _ float64) uint8 {
	if D <= 128 {
		return D
	} else {
		temp := 2*uint16(D) - 255
		if uint16(S) < temp {
			return S
		} else {
			return uint8(temp)
		}
	}
}
