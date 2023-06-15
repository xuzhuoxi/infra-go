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
	RegisterBlendFunc(HardLight, BlendHardLightColor, BlendHardLightRGBA)
}

// BlendHardLightColor
// 强光模式
// 根据绘图色来决定是执行“正片叠底”还是“滤色”模式。当绘图色比50%的灰要亮时，则底色变亮，就执行“滤色”模式一样，这对增加图像的高光非常有帮助；
// 当绘图色比50%的灰要暗时，则底色变暗，就执行“正片叠底”模式一样，可增加图像的暗部。当绘图色是纯白色或黑色时得到的是纯白色和黑色。此效果与耀眼的聚光灯照在图像上相似。像亮则更亮，暗则更暗。
// 这种模式实质上同Soft Light模式是一样的。它的效果要比Soft Light模式更强烈一些，同Overlay一样，这种模式 也可以在背景对象的表面模拟图案或文本。
// (D<=0.5): R = 2*S*D
// (D>0.5) : R = 1 - 2*(1-S) * (1-D)
//  ------------
// (D<=128): R = S * D / 128
// (D >128): R = 255 - (255-S)*(255-D)/128
//  ------------
// (D<=128): R = S * D / 32768
// (D>128) : R = 65535 - (65535-S)*(65535-D)/32768
func BlendHardLightColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendHardLightRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendHardLightRGBA
// 强光模式
// 根据绘图色来决定是执行“正片叠底”还是“滤色”模式。当绘图色比50%的灰要亮时，则底色变亮，就执行“滤色”模式一样，这对增加图像的高光非常有帮助；
// 当绘图色比50%的灰要暗时，则底色变暗，就执行“正片叠底”模式一样，可增加图像的暗部。当绘图色是纯白色或黑色时得到的是纯白色和黑色。此效果与耀眼的聚光灯照在图像上相似。像亮则更亮，暗则更暗。
// 这种模式实质上同Soft Light模式是一样的。它的效果要比Soft Light模式更强烈一些，同Overlay一样，这种模式 也可以在背景对象的表面模拟图案或文本。
// (D<=0.5): R = 2*S*D
// (D>0.5) : R = 1 - 2*(1-S) * (1-D)
//  ------------
// (D<=128): R = S * D / 128
// (D >128): R = 255 - (255-S)*(255-D)/128
//  ------------
// (D<=128): R = S * D / 32768
// (D>128) : R = 65535 - (65535-S)*(65535-D)/32768
func BlendHardLightRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = hardLight(Sr, Dr)
	G = hardLight(Sg, Dg)
	B = hardLight(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = hardLight(Sa, Da)
	}
	return
}

// (D<=0.5): R = 2*S*D
// (D>0.5) : R = 1 - 2*(1-S) * (1-D)
//  ------------
// (D<=128): R = S * D / 128
// (D >128): R = 255 - (255-S)*(255-D)/128
//  ------------
// (D<=128): R = S * D / 32768
// (D>128) : R = 65535 - (65535-S)*(65535-D)/32768
func hardLight(S, D uint32) uint32 {
	if D <= 128 {
		return S * D / 32768
	} else {
		return 65535 - (65535-S)*(65535-D)/32768
	}
}
