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
	RegisterBlendFunc(LighterColor, BlendLighterColor, BlendLighterColorRGBA)
}

// 浅色模式
// 公式：Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa 则 R=D
// 比较混合色和基色的所有通道值的总和并显示值较大的颜色。“浅色”不会生成第三种颜色（可以通过“变亮”混合获得），因为它将从基色和混合色中选取最大的通道值来创建结果色。
func BlendLighterColor(foreColor, backColor color.Color, _ float64, _ bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	if fR+fG+fB+fA >= bR+bG+bB+bA {
		return foreColor
	} else {
		return backColor
	}
}

// 浅色模式
// 公式：Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa 则 R=D
// 比较混合色和基色的所有通道值的总和并显示值较大的颜色。“浅色”不会生成第三种颜色（可以通过“变亮”混合获得），因为它将从基色和混合色中选取最大的通道值来创建结果色。
func BlendLighterColorRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa {
		return Dr, Dg, Db, Da
	} else {
		return Sr, Sg, Sb, Sa
	}
}
