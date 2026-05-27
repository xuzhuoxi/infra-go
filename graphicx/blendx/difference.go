package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Difference, BlendDifferenceColor, BlendDifferenceRGBA)
}

// BlendDifferenceColor applies difference blend.
// BlendDifferenceColor 应用差值（Difference）混合。
// R = |S - D|
func BlendDifferenceColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDifferenceRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendDifferenceRGBA applies difference on RGBA64 channels.
// BlendDifferenceRGBA 对 RGBA64 通道应用差值混合。
// R = |S - D|
func BlendDifferenceRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = difference(Sr, Dr)
	G = difference(Sg, Dg)
	B = difference(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = difference(Sa, Da)
	}
	return
}

// R = |S - D|
func difference(S, D uint32) uint32 {
	if S > D {
		return S - D
	} else {
		return D - S
	}
}
