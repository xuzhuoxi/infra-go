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
	RegisterBlendFunc(Hue, BlendHueColor, BlendHueRGBA)
}

// 色相模式
// 是采用底色的亮度、饱和度以及绘图色的色相来创建最终色。
// HSV = fH, bS, bV
func BlendHueColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendHueRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 色相模式
// 是采用底色的亮度、饱和度以及绘图色的色相来创建最终色。
// HSV = fH, bS, bV
func BlendHueRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, _ uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R, G, B = hueUnit(foreR, foreG, foreB, backR, backG, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = fH, bS, bV
func hueUnit(foreR, foreG, foreB uint32, backR, backG, backB uint32) (R, G, B uint32) {
	_, backS, backV := graphicx.RGB2HSV(backR, backG, backB)
	foreH, _, _ := graphicx.RGB2HSV(foreR, foreG, foreB)
	R, G, B = graphicx.HSV2RGB(foreH, backS, backV)
	return
}
