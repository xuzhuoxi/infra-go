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
	RegisterBlendFunc(HardLight, BlendHardLightColor, BlendHardLightRGBA)
}

// 强光模式
// 根据绘图色来决定是执行“正片叠底”还是“滤色”模式。当绘图色比50%的灰要亮时，则底色变亮，就执行“滤色”模式一样，这对增加图像的高光非常有帮助；
// 当绘图色比50%的灰要暗时，则底色变暗，就执行“正片叠底”模式一样，可增加图像的暗部。当绘图色是纯白色或黑色时得到的是纯白色和黑色。此效果与耀眼的聚光灯照在图像上相似。像亮则更亮，暗则更暗。
// 这种模式实质上同Soft Lishi模式是一样的。它的效果要比Soft Light模式更强烈一些，同Overlay一样，这种模式 也可以在背景对象的表面模拟图案或文本。
// (F<=0.5): R = 2*B*F
// (F>0.5): R = 1 - 2*(1-B) * (1-F)
//
// (F<=128): R = B*F/128
// (F>128): R = 255 - (255-B) * (255-F) / 128
func BlendHardLightColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendHardLightRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 强光模式
// 根据绘图色来决定是执行“正片叠底”还是“滤色”模式。当绘图色比50%的灰要亮时，则底色变亮，就执行“滤色”模式一样，这对增加图像的高光非常有帮助；
// 当绘图色比50%的灰要暗时，则底色变暗，就执行“正片叠底”模式一样，可增加图像的暗部。当绘图色是纯白色或黑色时得到的是纯白色和黑色。此效果与耀眼的聚光灯照在图像上相似。像亮则更亮，暗则更暗。
// 这种模式实质上同Soft Lishi模式是一样的。它的效果要比Soft Light模式更强烈一些，同Overlay一样，这种模式 也可以在背景对象的表面模拟图案或文本。
// (F<=0.5): R = 2*B*F
// (F>0.5): R = 1 - 2*(1-B) * (1-F)
//
// (F<=128): R = B*F/128
// (F>128): R = 255 - (255-B) * (255-F) / 128
func BlendHardLightRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = hardLight(foreR, backR)
	G = hardLight(foreG, backG)
	B = hardLight(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = hardLight(foreA, backA)
	}
	return
}

// (F<=0.5): R = 2*B*F
// (F>0.5): R = 1 - 2*(1-B) * (1-F)
//
// (F<=128): R = B*F/128
// (F>128): R = 255 - (255-B) * (255-F) / 128
func hardLight(F, B uint32) uint32 {
	if F <= 128 {
		return B * F / 32768
	} else {
		return 65535 - (65535-B)*(65535-F)/32768
	}
}
