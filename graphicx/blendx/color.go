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
// HSV = Dh, Ds, Sv
func BlendColorColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, _ := D.RGBA()
	R, G, B, A := BlendColorRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, 0, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 颜色模式
// 是采用底色的亮度以及绘图色的色相、饱和度来创建最终色。它可保护原图的灰阶层次，对于图像的色彩微调、给单色和彩色图像着色都非常有用。
// HSV = Dh, Ds, Sv
func BlendColorRGBA(Sr, Sg, Sb, _ uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R, G, B = colorUnit(Sr, Sg, Sb, Dr, Dg, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = Dh, Ds, Sv
func colorUnit(Sr, Sg, Sb uint32, Dr, Dg, Db uint32) (R, G, B uint32) {
	_, _, Sv := graphicx.RGB2HSV(Sr, Sg, Sb)
	Dh, Ds, _ := graphicx.RGB2HSV(Dr, Dg, Db)
	R, G, B = graphicx.HSV2RGB(Dh, Ds, Sv)
	return
}
