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
// 公式：fR+fG+fB>=bR+bG+bB 则 R=F
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
// 公式：fR+fG+fB>=bR+bG+bB 则 R=F
// 比较混合色和基色的所有通道值的总和并显示值较大的颜色。“浅色”不会生成第三种颜色（可以通过“变亮”混合获得），因为它将从基色和混合色中选取最大的通道值来创建结果色。
func BlendLighterColorRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if foreR+foreG+foreB+foreA >= backR+backG+backB+backA {
		return foreR, foreG, foreB, foreA
	} else {
		return backR, backG, backB, backA
	}
}
