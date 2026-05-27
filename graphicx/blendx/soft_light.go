package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(SoftLight, BlendSoftLightColor, BlendSoftLightRGBA)
}

// BlendSoftLightColor applies soft light blend.
// BlendSoftLightColor 应用柔光（Soft Light）混合。
func BlendSoftLightColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendSoftLightRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendSoftLightRGBA applies soft light on RGBA64 channels.
// BlendSoftLightRGBA 对 RGBA64 通道应用柔光混合。
func BlendSoftLightRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = softLight(Sr, Dr)
	G = softLight(Sg, Dg)
	B = softLight(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = softLight(Sa, Da)
	}
	return
}

// (F<=0.5): R = 2*S*D + S*S*(1-2*D)
// (F>0.5) : R = 2*S*(1-D) + (2*D - 1)*√S
//
// (D<=128): R = S*D/128 + (255-2*D)*S*S/65025
// (D>128) : R = S*(255-D)/128 + (2*D-255)*√(S/255)
//
// (D<=128): R = S*D/32768 + (65535-2*D)*S*S/4294836225
// (D>128) : R = S*(65535-D)/32768 + (2*D-65535)*√(S/65535)
func softLight(S, D uint32) uint32 {
	if D <= 32768 {
		return S*D/32768 + (65535-2*D)*S*S/4294836225
	} else {
		return S*(65535-D)/32768 + (2*D-65535)*uint32(math.Sqrt(float64(S)/65535))
	}
}
