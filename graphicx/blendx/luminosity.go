package blendx

import (
	"image/color"
	"math"

	"github.com/xuzhuoxi/infra-go/graphicx"
)

func init() {
	RegisterBlendFunc(Luminosity, BlendLuminosityColor, BlendLuminosityRGBA)
}

// BlendLuminosityColor uses destination hue/saturation and source luminance.
// BlendLuminosityColor 采用目标色相/饱和度与源亮度（亮度模式）。
// HSV = Sh, Ss, Dv
func BlendLuminosityColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, _ := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLuminosityRGBA(Sr, Sg, Sb, 0, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendLuminosityRGBA applies luminosity blend on RGBA64 channels.
// BlendLuminosityRGBA 对 RGBA64 通道应用亮度混合。
// HSV = Sh, Ss, Dv
func BlendLuminosityRGBA(Sr, Sg, Sb, _ uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R, G, B = luminosity(Sr, Sg, Sb, Dr, Dg, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = Sh, Ss, Dv
func luminosity(Sr, Sg, Sb uint32, Dr, Dg, Db uint32) (R, G, B uint32) {
	Sh, Ss, _ := graphicx.RGB2HSV(Sr, Sg, Sb)
	_, _, Dv := graphicx.RGB2HSV(Dr, Dg, Db)
	R, G, B = graphicx.HSV2RGB(Sh, Ss, Dv)
	return
}
