package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(HardLight, BlendHardLightColor, BlendHardLightRGBA)
}

// BlendHardLightColor applies hard light blend.
// BlendHardLightColor 应用强光（Hard Light）混合。
func BlendHardLightColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendHardLightRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendHardLightRGBA applies hard light on RGBA64 channels.
// BlendHardLightRGBA 对 RGBA64 通道应用强光混合。
func BlendHardLightRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = hardLight(Sr, Dr)
	G = hardLight(Sg, Dg)
	B = hardLight(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = hardLight(Sa, Da)
	}
	return
}

// (D<=0.5): R = 2*S*D
// (D>0.5) : R = 1 - 2*(1-S) * (1-D)
//  ------------
// (D<=128): R = S * D / 128
// (D >128): R = 255 - (255-S)*(255-D)/128
//  ------------
// (D<=128): R = S * D / 32768
// (D>128) : R = 65535 - (65535-S)*(65535-D)/32768
func hardLight(S, D uint32) uint32 {
	if D <= 128 {
		return S * D / 32768
	} else {
		return 65535 - (65535-S)*(65535-D)/32768
	}
}
