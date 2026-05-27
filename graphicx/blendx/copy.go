package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Copy, BlendCopyColor, BlendCopyRGBA)
}

// BlendCopyColor copies source over destination (R = S).
// BlendCopyColor 用源色覆盖目标色（R = S）。
// R = S
func BlendCopyColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	if !destinationAlpha {
		return S
	}
	Sr, Sg, Sb, _ := S.RGBA()
	_, _, _, Da := D.RGBA()
	return &color.RGBA64{R: uint16(Sr), G: uint16(Sg), B: uint16(Sb), A: uint16(Da)}
}

// BlendCopyRGBA copies source on RGBA64 channels (R = S).
// BlendCopyRGBA 对 RGBA64 通道复制源色（R = S）。
// R = S
func BlendCopyRGBA(Sr, Sg, Sb, Sa uint32, _, _, _, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R, G, B = Sr, Sg, Sb
	if destinationAlpha {
		A = Da
	} else {
		A = Sa
	}
	return
}
