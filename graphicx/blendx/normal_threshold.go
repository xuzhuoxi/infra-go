package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(NormalThreshold, BlendNormalThresholdColor, BlendNormalThresholdRGBA)
}

// BlendNormalThresholdColor applies normal blend with a threshold factor.
// BlendNormalThresholdColor 按阈值因子应用正常混合。
// R = S*(1-factor) + D*factor
func BlendNormalThresholdColor(S, D color.Color, factor float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendNormalThresholdRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, factor, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendNormalThresholdRGBA applies normal blend with a threshold factor on RGBA64 channels.
// BlendNormalThresholdRGBA 对 RGBA64 通道按阈值因子应用正常混合。
// R = S*(1-factor) + D*factor
func BlendNormalThresholdRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, factor float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = normalThreshold(Sr, Dr, factor)
	G = normalThreshold(Sg, Dg, factor)
	B = normalThreshold(Sb, Db, factor)
	if destinationAlpha {
		A = Da
	} else {
		A = normalThreshold(Sa, Da, factor)
	}
	return
}

// R = S*(1-factor) + D*factor
func normalThreshold(S, D uint32, factor float64) uint32 {
	if factor <= 0 {
		return S
	}
	if factor >= 1 {
		return D
	}
	return uint32(float64(S)*(1-factor) + float64(D)*factor)
}
