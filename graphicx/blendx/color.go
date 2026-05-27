package blendx

import (
	"github.com/xuzhuoxi/infra-go/graphicx"
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(Color, BlendColorColor, BlendColorRGBA)
}

// BlendColorColor uses destination luminance and source hue/saturation.
// BlendColorColor 采用目标亮度与源色相/饱和度（颜色模式）。
// HSV = Dh, Ds, Sv
func BlendColorColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, _ := D.RGBA()
	R, G, B, A := BlendColorRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, 0, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendColorRGBA applies color blend on RGBA64 channels.
// BlendColorRGBA 对 RGBA64 通道应用颜色混合。
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
