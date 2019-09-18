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
	RegisterBlendFunc(NormalThreshold, BlendNormalThresholdColor, BlendNormalThresholdRGBA)
}

// 阈值模式
// 是默认的状态，其最终色和绘图色相同。可通过改变画笔工具选项栏中的“不透明度”来设定不同的透明度。当图像的颜色模式是“位图”或“索引颜色”时，“正常”模式就变成“阈值”模式。
// 在基色存在透明度a%时，混合的运算方式是：最终色=基色*a% + 混合色*（1-a%）。
// R = B * factor + F * (1-factor)
func BlendNormalThresholdColor(foreColor, backColor color.Color, factor float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendNormalThresholdRGBA(fR, fG, fB, fA, bR, bG, bB, bA, factor, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 阈值模式
// 是默认的状态，其最终色和绘图色相同。可通过改变画笔工具选项栏中的“不透明度”来设定不同的透明度。当图像的颜色模式是“位图”或“索引颜色”时，“正常”模式就变成“阈值”模式。
// 在基色存在透明度a%时，混合的运算方式是：最终色=基色*a% + 混合色*（1-a%）。
// R = B * factor + F * (1-factor)
func BlendNormalThresholdRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, factor float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = normalThreshold(foreR, backR, factor)
	G = normalThreshold(foreG, backG, factor)
	B = normalThreshold(foreB, backB, factor)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = normalThreshold(foreA, backA, factor)
	}
	return
}

// R = B * factor + F * (1-factor)
func normalThreshold(F, B uint32, factor float64) uint32 {
	if factor <= 0 {
		return F
	}
	if factor >= 1 {
		return B
	}
	return uint32(float64(B)*factor + float64(F)*(1-factor))
}
