package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(SourceIn, BlendSourceInColor, BlendSourceInRGBA)
}

// BlendSourceInColor applies source in Porter-Duff compositing.
// BlendSourceInColor 应用 SourceIn Porter-Duff 合成。
// R = S*Da
// R = S*Da/255
// R = S*Da/65535
func BlendSourceInColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	fR, fG, fB, fA := S.RGBA()
	bR, bG, bB, bA := D.RGBA()
	R, G, B, A := BlendSourceInRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendSourceInRGBA applies source in on RGBA64 channels.
// BlendSourceInRGBA 对 RGBA64 通道应用 SourceIn 合成。
// R = S*Da
// R = S*Da/255
// R = S*Da/65535
func BlendSourceInRGBA(Sr, Sg, Sb, Sa uint32, _, _, _, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = sourceIn(Sr, Da)
	G = sourceIn(Sg, Da)
	B = sourceIn(Sb, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = sourceIn(Sa, Da)
	}
	return
}

// R = S*Da
// R = S*Da/255
// R = S*Da/65535
func sourceIn(S, Da uint32) uint32 {
	return S * Da / 65535
}
