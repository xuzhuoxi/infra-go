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
// R = S*(1-factor) + D*factor
func BlendNormalThresholdColor(S, D color.Color, factor float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendNormalThresholdRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, factor, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 阈值模式
// 是默认的状态，其最终色和绘图色相同。可通过改变画笔工具选项栏中的“不透明度”来设定不同的透明度。当图像的颜色模式是“位图”或“索引颜色”时，“正常”模式就变成“阈值”模式。
// 在基色存在透明度a%时，混合的运算方式是：最终色=基色*a% + 混合色*（1-a%）。
// R = S*(1-factor) + D*factor
func BlendNormalThresholdRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, factor float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = normalThreshold(Sr, Dr, factor)
	G = normalThreshold(Sg, Dg, factor)
	B = normalThreshold(Sb, Db, factor)
	if destinationAlpha {
		A = Da
	} else {
		A = normalThreshold(Sa, Da, factor)
	}
	return
}

// R = S*(1-factor) + D*factor
func normalThreshold(S, D uint32, factor float64) uint32 {
	if factor <= 0 {
		return S
	}
	if factor >= 1 {
		return D
	}
	return uint32(float64(S)*(1-factor) + float64(D)*factor)
}
