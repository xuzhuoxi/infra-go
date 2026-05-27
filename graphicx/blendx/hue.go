package blendx

import (
	"image/color"
	"math"

	"github.com/xuzhuoxi/infra-go/graphicx"
)

func init() {
	RegisterBlendFunc(Hue, BlendHueColor, BlendHueRGBA)
}

// BlendHueColor uses destination luminance/saturation and source hue.
// BlendHueColor 采用目标亮度/饱和度与源色相（色相模式）。
// HSV = Dh, Ss, Sv
func BlendHueColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendHueRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendHueRGBA applies hue blend on RGBA64 channels (HSV = Dh, Ss, Sv).
// BlendHueRGBA 对 RGBA64 通道应用色相混合（HSV = Dh, Ss, Sv）。
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
