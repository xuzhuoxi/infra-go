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
	RegisterBlendFunc(DarkerColor, DarkerColorBlend)
}

// 线性加深模式
// 公式：rB+gB+bBrB+gB+bB>=rA+gA+bA 则 C=A
// 比较混合色和基色的所有通道值的总和并显示值较小的颜色。“深色”不会生成第三种颜色（可以通过“变暗”混合获得），因为它将从基色和混合色中选取最小的通道值来创建结果色。
func DarkerColorBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = DarkerColorUnit(source.A, target.A, factor)
	}
	source.R = DarkerColorUnit(source.R, target.R, factor)
	source.G = DarkerColorUnit(source.G, target.G, factor)
	source.B = DarkerColorUnit(source.B, target.B, factor)
	return source
}

// 公式：rB+gB+bBrB+gB+bB>=rA+gA+bA 则 C=A
// 比较混合色和基色的所有通道值的总和并显示值较小的颜色。“深色”不会生成第三种颜色（可以通过“变暗”混合获得），因为它将从基色和混合色中选取最小的通道值来创建结果色。
func DarkerColorUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(uint16(S) + uint16(D) - 255)
}
