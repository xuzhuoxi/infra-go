package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Xor, BlendXorColor, BlendXorRGBA)
}

// BlendXorColor applies XOR Porter-Duff compositing.
// BlendXorColor 应用 XOR Porter-Duff 合成。
// R = S*(1 - Da) + D*(1 - Sa)
// R = (S*(255 - Da) + D*(255 - Sa))/255
// R = (S*(65535 - Da) + D*(65535 - Sa))/65535
func BlendXorColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendXorRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendXorRGBA applies XOR compositing on RGBA64 channels.
// BlendXorRGBA 对 RGBA64 通道应用 XOR 合成。
// R = S*(1 - Da) + D*(1 - Sa)
// R = (S*(255 - Da) + D*(255 - Sa))/255
// R = (S*(65535 - Da) + D*(65535 - Sa))/65535
func BlendXorRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = xor(Sr, Dr, Sa, Da)
	G = xor(Sg, Dg, Sa, Da)
	B = xor(Sb, Db, Sa, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = xor(Sa, Da, Sa, Da)
	}
	return
}

// R = S*(1 - Da) + D*(1 - Sa)
// R = (S*(255 - Da) + D*(255 - Sa))/255
// R = (S*(65535 - Da) + D*(65535 - Sa))/65535
func xor(S, D, Sa, Da uint32) uint32 {
	return (S*(65535-Da) + D*(65535-Sa)) / 65535
}
