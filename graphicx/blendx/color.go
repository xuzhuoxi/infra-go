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
	RegisterBlendFunc(Color, BlendColorColor, BlendColorRGBA)
}

// 颜色模式
// 是采用底色的亮度以及绘图色的色相、饱和度来创建最终色。它可保护原图的灰阶层次，对于图像的色彩微调、给单色和彩色图像着色都非常有用。
// HSV = fH, fS, bV
func BlendColorColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	foreR, foreG, foreB, foreA := foreColor.RGBA()
	backR, backG, backB, _ := backColor.RGBA()
	R, G, B, A := BlendColorRGBA(foreR, foreG, foreB, foreA, backR, backG, backB, 0, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 颜色模式
// 是采用底色的亮度以及绘图色的色相、饱和度来创建最终色。它可保护原图的灰阶层次，对于图像的色彩微调、给单色和彩色图像着色都非常有用。
// HSV = fH, fS, bV
func BlendColorRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, _ uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R, G, B = colorUnit(foreR, foreG, foreB, backR, backG, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = fH, fS, bV
func colorUnit(foreR, foreG, foreB uint32, backR, backG, backB uint32) (R, G, B uint32) {
	_, _, backV := graphicx.RGB2HSV(backR, backG, backB)
	foreH, foreS, _ := graphicx.RGB2HSV(foreR, foreG, foreB)
	R, G, B = graphicx.HSV2RGB(foreH, foreS, backV)
	return
}
