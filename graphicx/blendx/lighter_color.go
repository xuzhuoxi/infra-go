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
	RegisterBlendFunc(LighterColor, LighterColorBlend)
}

// 线性变浅模式
// 公式：rB+gB+bBrB+gB+bB>=rA+gA+bA 则 C=B
// 比较混合色和基色的所有通道值的总和并显示值较大的颜色。“浅色”不会生成第三种颜色（可以通过“变亮”混合获得），因为它将从基色和混合色中选取最大的通道值来创建结果色。
func LighterColorBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = LighterColorUnit(source.A, target.A, factor)
	}
	source.R = LighterColorUnit(source.R, target.R, factor)
	source.G = LighterColorUnit(source.G, target.G, factor)
	source.B = LighterColorUnit(source.B, target.B, factor)
	return source
}

// 公式：rB+gB+bBrB+gB+bB>=rA+gA+bA 则 C=B
// 比较混合色和基色的所有通道值的总和并显示值较大的颜色。“浅色”不会生成第三种颜色（可以通过“变亮”混合获得），因为它将从基色和混合色中选取最大的通道值来创建结果色。
func LighterColorUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(uint16(S) + uint16(D) - 255)
}
