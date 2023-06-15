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
	RegisterBlendFunc(DarkerColor, BlendDarkerColorColor, BlendDarkerColorRGBA)
}

// BlendDarkerColorColor
// 深色模式
// 公式：Dr+Dg+Db>=Sr+Sg+Sb 则 R=S
// 比较混合色和基色的所有通道值的总和并显示值较小的颜色。“深色”不会生成第三种颜色（可以通过“变暗”混合获得），因为它将从基色和混合色中选取最小的通道值来创建结果色。
func BlendDarkerColorColor(S, D color.Color, _ float64, _ bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	if Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa {
		return S
	} else {
		return D
	}
}

// BlendDarkerColorRGBA
// 深色模式
// 公式：Dr+Dg+Db>=Sr+Sg+Sb 则 R=S
// 比较混合色和基色的所有通道值的总和并显示值较小的颜色。“深色”不会生成第三种颜色（可以通过“变暗”混合获得），因为它将从基色和混合色中选取最小的通道值来创建结果色。
func BlendDarkerColorRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa {
		return Sr, Sg, Sb, Sa
	} else {
		return Dr, Dg, Db, Da
	}
}
