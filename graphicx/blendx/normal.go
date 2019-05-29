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
	RegisterBlendFunc(Normal, NormalBlend)
}

// 正常模式
// 是默认的状态，其最终色和绘图色相同。可通过改变画笔工具选项栏中的“不透明度”来设定不同的透明度。当图像的颜色模式是“位图”或“索引颜色”时，“正常”模式就变成“阈值”模式。
// 在基色存在透明度a%时，混合的运算方式是：最终色=基色*a% + 混合色*（1-a%）。
// R = S * factor + D * (1-factor)
func NormalBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = NormalUnit(source.A, target.A, factor)
	}
	source.R = NormalUnit(source.R, target.R, factor)
	source.G = NormalUnit(source.G, target.G, factor)
	source.B = NormalUnit(source.B, target.B, factor)
	return source
}

// R = S * factor + D * (1-factor)
func NormalUnit(S uint8, D uint8, factor float64) uint8 {
	if factor >= 0 && factor <= 1 {
		return uint8(float64(S)*factor + float64(D)*(1-factor))
	} else {
		return S
	}
}
