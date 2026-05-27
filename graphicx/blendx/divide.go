package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Divide, BlendDivideColor, BlendDivideRGBA)
}

// BlendDivideColor divides destination by source.
// BlendDivideColor 用源色划分目标色（划分模式）。
// R = S / D [0, 1]
// R = 255 * S / D [0, 255]
// R = 65535 * S / D [0, 65535]
func BlendDivideColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDivideRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendDivideRGBA divides RGBA64 channels.
// BlendDivideRGBA 对 RGBA64 通道执行划分混合。
// R = S / D [0, 1]
// R = 255 * S / D [0, 255]
// R = 65535 * S / D [0, 65535]
func BlendDivideRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = divide(Sr, Dr)
	G = divide(Sg, Dg)
	B = divide(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = divide(Sa, Da)
	}
	return
}

// R = S / D [0, 1]
// R = 255 * S / D [0, 255]
// R = 65535 * S / D [0, 65535]
func divide(S, D uint32) uint32 {
	return 65535 * S / D
}
