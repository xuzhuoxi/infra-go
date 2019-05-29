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
	RegisterBlendFunc(HardLight, HardLightBlend)
}

// 强光模式
// 根据绘图色来决定是执行“正片叠底”还是“滤色”模式。当绘图色比50%的灰要亮时，则底色变亮，就执行“滤色”模式一样，这对增加图像的高光非常有帮助；
// 当绘图色比50%的灰要暗时，则底色变暗，就执行“正片叠底”模式一样，可增加图像的暗部。当绘图色是纯白色或黑色时得到的是纯白色和黑色。此效果与耀眼的聚光灯照在图像上相似。像亮则更亮，暗则更暗。
// 这种模式实质上同Soft Lishi模式是一样的。它的效果要比Soft Light模式更强烈一些，同Overlay一样，这种模式 也可以在背景对象的表面模拟图案或文本。
// (D<=0.5): R = 2*S*D
// (D>0.5): R = 1 - 2*(1 - S)*(1 - D)
//
// (D<=128): R = S*D/128
// (D>128): R = 255 - (255 - S) * (255 - D) / 128
func HardLightBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = HardLightUnit(source.A, target.A, factor)
	}
	source.R = HardLightUnit(source.R, target.R, factor)
	source.G = HardLightUnit(source.G, target.G, factor)
	source.B = HardLightUnit(source.B, target.B, factor)
	return source
}

// (D<=0.5): R = 2*S*D
// (D>0.5): R = 1 - 2*(1 - S)*(1 - D)
//
// (D<=128): R = S*D/128
// (D>128): R = 255 - (255 - S) * (255 - D) / 128
func HardLightUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	if D <= 128 {
		return uint8((S16 * D16) >> 7)
	} else {
		return uint8(255 - ((255-S16)*(255-D16))>>7)
	}
}
