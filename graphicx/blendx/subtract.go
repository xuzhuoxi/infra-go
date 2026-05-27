package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Subtract, BlendSubtractColor, BlendSubtractRGBA)
}

// BlendSubtractColor subtracts source from destination.
// BlendSubtractColor 从目标色减去源色（减去模式）。
func BlendSubtractColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendSubtractRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendSubtractRGBA subtracts RGBA64 channels.
// BlendSubtractRGBA 对 RGBA64 通道执行减去混合。
func BlendSubtractRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = subtract(Sr, Dr)
	G = subtract(Sg, Dg)
	B = subtract(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = subtract(Sa, Da)
	}
	return
}

// C = Max(0,S-D)
func subtract(S, D uint32) uint32 {
	if S > D {
		return S - D
	} else {
		return 0
	}
}
