package blendx

import (
	"image/color"
	"math"

	"github.com/xuzhuoxi/infra-go/graphicx"
)

func init() {
	RegisterBlendFunc(Saturation, BlendSaturationColor, BlendSaturationRGBA)
}

// BlendSaturationColor uses destination hue/value and source saturation.
// BlendSaturationColor 采用目标色相/明度与源饱和度（饱和度模式）。
// HSV = Sh, Ds, Sv
func BlendSaturationColor(foreColor, backColor color.Color, _ float64, destinationAlpha bool) color.Color {
	foreR, foreG, foreB, foreA := foreColor.RGBA()
	backR, backG, backB, _ := backColor.RGBA()
	R, G, B, A := BlendSaturationRGBA(foreR, foreG, foreB, foreA, backR, backG, backB, 0, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendSaturationRGBA applies saturation blend on RGBA64 channels.
// BlendSaturationRGBA 对 RGBA64 通道应用饱和度混合。
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
