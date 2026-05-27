package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Multiply, BlendMultiplyColor, BlendMultiplyRGBA)
}

// BlendMultiplyColor multiplies source and destination (multiply mode).
// BlendMultiplyColor 正片叠底混合源色与目标色。
// R = S*D 			[0,1]
// R = S*D/255		[0,255]
// R = S*D/65535	[0,65535]
func BlendMultiplyColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendMultiplyRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendMultiplyRGBA multiplies RGBA64 channels (R = S*D/65535).
// BlendMultiplyRGBA 对 RGBA64 通道正片叠底（R = S*D/65535）。
// R = S*D 			[0,1]
// R = S*D/255		[0,255]
// R = S*D/65535	[0,65535]
func BlendMultiplyRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = multiply(Sr, Dr)
	G = multiply(Sg, Dg)
	B = multiply(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = multiply(Sa, Da)
	}
	return
}

// R = S*D 			[0,1]
// R = S*D/255		[0,255]
// R = S*D/65535	[0,65535]
func multiply(S, D uint32) uint32 {
	return S * D / 65535
}
