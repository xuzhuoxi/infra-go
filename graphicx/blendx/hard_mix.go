package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(HardMix, BlendHardMixColor, BlendHardMixRGBA)
}

// BlendHardMixColor applies hard mix blend.
// BlendHardMixColor 应用实色混合（Hard Mix）模式。
// R = S+D>=1 ? 1 : 0
// R = S+D>=255 ? 255 : 0
// R = S+D>=65535 ? 65535 : 0
func BlendHardMixColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendHardMixRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendHardMixRGBA applies hard mix on RGBA64 channels.
// BlendHardMixRGBA 对 RGBA64 通道应用实色混合。
// R = S+D>=1 ? 1 : 0
// R = S+D>=255 ? 255 : 0
// R = S+D>=65535 ? 65535 : 0
func BlendHardMixRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = hardMix(Sr, Dr)
	G = hardMix(Sg, Dg)
	B = hardMix(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = hardMix(Sa, Da)
	}
	return
}

// R = S+D>=1 ? 1 : 0
// R = S+D>=255 ? 255 : 0
// R = S+D>=65535 ? 65535 : 0
func hardMix(S, D uint32) uint32 {
	if S+D >= 65535 {
		return 65535
	} else {
		return 0
	}
}
