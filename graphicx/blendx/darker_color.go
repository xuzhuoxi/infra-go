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
	RegisterBlendFunc(DarkerColor, BlendDarkerColorColor, BlendDarkerColorRGBA)
}

// 深色模式
// 公式：fR+fG+fB>=bR+bG+bB 则 R=B
// 比较混合色和基色的所有通道值的总和并显示值较小的颜色。“深色”不会生成第三种颜色（可以通过“变暗”混合获得），因为它将从基色和混合色中选取最小的通道值来创建结果色。
func BlendDarkerColorColor(foreColor, backColor color.Color, _ float64, _ bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	if fR+fG+fB+fA >= bR+bG+bB+bA {
		return backColor
	} else {
		return foreColor
	}
}

// 深色模式
// 公式：fR+fG+fB>=bR+bG+bB 则 R=B
// 比较混合色和基色的所有通道值的总和并显示值较小的颜色。“深色”不会生成第三种颜色（可以通过“变暗”混合获得），因为它将从基色和混合色中选取最小的通道值来创建结果色。
func BlendDarkerColorRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if foreR+foreG+foreB+foreA >= backR+backG+backB+backA {
		return backR, backG, backB, backA
	} else {
		return foreR, foreG, foreB, foreA
	}
}
