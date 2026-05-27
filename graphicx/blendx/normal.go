package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Normal, BlendNormalColor, BlendNormalRGBA)
}

// BlendNormalColor applies normal alpha compositing.
// BlendNormalColor 应用正常（Normal）Alpha 混合。
// R = S*(1-Da) + D*Da
func BlendNormalColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	fR, fG, fB, fA := S.RGBA()
	bR, bG, bB, bA := D.RGBA()
	R, G, B, A := BlendNormalRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendNormalRGBA applies normal alpha compositing on RGBA64 channels.
// BlendNormalRGBA 对 RGBA64 通道应用正常（Normal）Alpha 混合。
// R = S*(1-Da) + D*Da
func BlendNormalRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	factor := float64(Da) / 65535
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
