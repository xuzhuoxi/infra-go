//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"github.com/xuzhuoxi/infra-go/graphicx"
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(Saturation, BlendSaturationColor, BlendSaturationRGBA)
}

// 饱和度模式
// 是采用底色的亮度、色相以及绘图色的饱和度来创建最终色。如果绘图色的饱和度为0，则原图没有变化。
// 在把纯蓝色应用到一个灰暗的背景图像中时，显出了背景中 的原始纯色，但蓝色并未加入到合成图像中。如果选择一种中性颜色（一种并不显示主流色 度的颜色），对背景图像不发生任何变化。
// Saturation模式可用来显出图像中颜色强度已经由 于岁月变得灰暗的底层颜色。
// HSV = Sh, Ds, Sv
func BlendSaturationColor(foreColor, backColor color.Color, _ float64, destinationAlpha bool) color.Color {
	foreR, foreG, foreB, foreA := foreColor.RGBA()
	backR, backG, backB, _ := backColor.RGBA()
	R, G, B, A := BlendSaturationRGBA(foreR, foreG, foreB, foreA, backR, backG, backB, 0, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 饱和度模式
// 是采用底色的亮度、色相以及绘图色的饱和度来创建最终色。如果绘图色的饱和度为0，则原图没有变化。
// 在把纯蓝色应用到一个灰暗的背景图像中时，显出了背景中 的原始纯色，但蓝色并未加入到合成图像中。如果选择一种中性颜色（一种并不显示主流色 度的颜色），对背景图像不发生任何变化。
// Saturation模式可用来显出图像中颜色强度已经由 于岁月变得灰暗的底层颜色。
// HSV = Sh, Ds, Sv
func BlendSaturationRGBA(Sr, Sg, Sb, _ uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R, G, B = saturation(Sr, Sg, Sb, Dr, Dg, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = Sh, Ds, Sv
func saturation(Sr, Sg, Sb uint32, Dr, Dg, Db uint32) (R, G, B uint32) {
	Sh, _, Sv := graphicx.RGB2HSV(Sr, Sg, Sb)
	_, Ds, _ := graphicx.RGB2HSV(Dr, Dg, Db)
	R, G, B = graphicx.HSV2RGB(Sh, Ds, Sv)
	return
}
