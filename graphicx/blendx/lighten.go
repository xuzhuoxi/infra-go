package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Lighten, BlendLightenColor, BlendLightenRGBA)
}

// BlendLightenColor takes the lighter of source and destination per channel.
// BlendLightenColor 逐通道取较亮者（变亮模式）。
// R = Max(S, D)
func BlendLightenColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLightenRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendLightenRGBA takes the lighter channel values (R = max(S, D)).
// BlendLightenRGBA 逐通道取较亮值（R = max(S, D)）。
// R = Max(S, D)
func BlendLightenRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = lighten(Sr, Dr)
	G = lighten(Sg, Dg)
	B = lighten(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = lighten(Sa, Da)
	}
	return
}

// R = Max(S, D)
func lighten(S, D uint32) uint32 {
	if S > D {
		return S
	} else {
		return D
	}
}
