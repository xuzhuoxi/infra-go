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
// HSV = bH, fS, bV
func BlendSaturationColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	foreR, foreG, foreB, foreA := foreColor.RGBA()
	backR, backG, backB, _ := backColor.RGBA()
	R, G, B, A := BlendSaturationRGBA(foreR, foreG, foreB, foreA, backR, backG, backB, 0, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 饱和度模式
// 是采用底色的亮度、色相以及绘图色的饱和度来创建最终色。如果绘图色的饱和度为0，则原图没有变化。
// 在把纯蓝色应用到一个灰暗的背景图像中时，显出了背景中 的原始纯色，但蓝色并未加入到合成图像中。如果选择一种中性颜色（一种并不显示主流色 度的颜色），对背景图像不发生任何变化。
// Saturation模式可用来显出图像中颜色强度已经由 于岁月变得灰暗的底层颜色。
// HSV = bH, fS, bV
func BlendSaturationRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, _ uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R, G, B = saturation(foreR, foreG, foreB, backR, backG, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = bH, fS, bV
func saturation(foreR, foreG, foreB uint32, backR, backG, backB uint32) (R, G, B uint32) {
	backH, _, backV := graphicx.RGB2HSV(backR, backG, backB)
	_, foreS, _ := graphicx.RGB2HSV(foreR, foreG, foreB)
	R, G, B = graphicx.HSV2RGB(backH, foreS, backV)
	return
}
