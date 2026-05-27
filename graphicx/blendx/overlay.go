package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Overlay, BlendOverlayColor, BlendOverlayRGBA)
}

// BlendOverlayColor applies overlay blend.
// BlendOverlayColor 应用叠加（Overlay）混合。
// R = S<=0.5 ? 2*S*D : 1-2*(1-S)*(1-D)
// R = S<=128 ? S*D/128 : 255-(255-S)*(255-d)/128
// R = S<=32768 ? S*D/32768 : 65535-(65535-S)*(65534-D)/32768
func BlendOverlayColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendOverlayRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendOverlayRGBA applies overlay on RGBA64 channels.
// BlendOverlayRGBA 对 RGBA64 通道应用叠加混合。
// R = S<=0.5 ? 2*S*D : 1-2*(1-S)*(1-D)
// R = S<=128 ? S*D/128 : 255-(255-S)*(255-d)/128
// R = S<=32768 ? S*D/32768 : 65535-(65535-S)*(65534-D)/32768
func BlendOverlayRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = overlay(Sr, Dr)
	G = overlay(Sg, Dg)
	B = overlay(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = overlay(Sa, Da)
	}
	return
}

// R = S<=0.5 ? 2*S*D : 1-2*(1-S)*(1-D)
// R = S<=128 ? S*D/128 : 255-(255-S)*(255-d)/128
// R = S<=32768 ? S*D/32768 : 65535-(65535-S)*(65534-D)/32768
func overlay(S, D uint32) uint32 {
	if S <= 32768 {
		return S * D / 32768
	} else {
		return 65535 - (65535-S)*(65534-D)/32768
	}
}
