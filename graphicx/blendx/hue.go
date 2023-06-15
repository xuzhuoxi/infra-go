// Package blendx
// Created by xuzhuoxi
// on 2019-05-25.
// @author xuzhuoxi
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

// BlendHueColor
// 色相模式
// 是采用底色的亮度、饱和度以及绘图色的色相来创建最终色。
// HSV = Dh, Ss, Sv
func BlendHueColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendHueRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendHueRGBA
// 色相模式
// 是采用底色的亮度、饱和度以及绘图色的色相来创建最终色。
// HSV = Dh, Ss, Sv
func BlendHueRGBA(Sr, Sg, Sb, _ uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R, G, B = hueUnit(Sr, Sg, Sb, Dr, Dg, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = Dh, Ss, Sv
func hueUnit(Sr, Sg, Sb uint32, Dr, Dg, Db uint32) (R, G, B uint32) {
	_, Ss, Sv := graphicx.RGB2HSV(Sr, Sg, Sb)
	Dh, _, _ := graphicx.RGB2HSV(Dr, Dg, Db)
	R, G, B = graphicx.HSV2RGB(Dh, Ss, Sv)
	return
}
