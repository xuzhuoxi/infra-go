package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Darken, BlendDarkenColor, BlendDarkenRGBA)
}

// BlendDarkenColor takes the darker of source and destination per channel.
// BlendDarkenColor 逐通道取源色与目标色中较暗者（变暗模式）。
// R = min(S, D)
func BlendDarkenColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDarkenRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendDarkenRGBA takes the darker channel values (R = min(S, D)).
// BlendDarkenRGBA 逐通道取较暗值（R = min(S, D)）。
// R = min(S, D)
func BlendDarkenRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = darken(Sr, Dr)
	G = darken(Sg, Dg)
	B = darken(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = darken(Sa, Da)
	}
	return
}

// R = min(S, D)
func darken(S, D uint32) uint32 {
	if S < D {
		return S
	} else {
		return D
	}
}
