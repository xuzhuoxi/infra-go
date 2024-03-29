// Package blendx
// Created by xuzhuoxi
// on 2019-05-25.
// @author xuzhuoxi
//
package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(Behind, BlendBehindColor, BlendBehindRGBA)
}

// BlendBehindColor
// 背后模式
// 最终色和绘图色相同。当在有透明区域的图层上操作时背后模式才会出现，可将绘制的线条放在图层中图像的后面。
// 这模式被用来在一个图层内透明的部分进行涂画；但当图层里的“保持透明区域”选中时就不可用了。
// 它只可以在你用涂画工具（画笔，喷枪，图章，历史记录画笔，油漆桶）或是填充命令在图层内的一个对象之后画上阴影或色彩。
// 当在有透明区域的图层上操作时背后模式才会出现
func BlendBehindColor(S, D color.Color, _ float64, _ bool) color.Color {
	_, _, _, Da := D.RGBA()
	if Da < math.MaxUint16 {
		return S
	} else {
		return D
	}
}

// BlendBehindRGBA
// 背后模式
// 最终色和绘图色相同。当在有透明区域的图层上操作时背后模式才会出现，可将绘制的线条放在图层中图像的后面。
// 这模式被用来在一个图层内透明的部分进行涂画；但当图层里的“保持透明区域”选中时就不可用了。
// 它只可以在你用涂画工具（画笔，喷枪，图章，历史记录画笔，油漆桶）或是填充命令在图层内的一个对象之后画上阴影或色彩。
// 当在有透明区域的图层上操作时背后模式才会出现
func BlendBehindRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if Da < math.MaxUint16 {
		return Sr, Sg, Sb, Sa
	} else {
		return Dr, Dg, Db, Da
	}
}
