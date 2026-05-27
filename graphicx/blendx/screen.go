package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Screen, BlendScreenColor, BlendScreenRGBA)
}

// BlendScreenColor applies screen blend (inverse of multiply).
// BlendScreenColor 应用滤色（Screen）混合。
// R = 1 - (1 - S)*(1 - D)
// R = 255 - (255-S)*(255-D) / 255
// R = 65535 - (65535-S)*(65535-D) / 65535
func BlendScreenColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendScreenRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendScreenRGBA applies screen blend on RGBA64 channels.
// BlendScreenRGBA 对 RGBA64 通道应用滤色混合。
// R = 1 - (1 - S)*(1 - D)
// R = 255 - (255-S)*(255-D) / 255
// R = 65535 - (65535-S)*(65535-D) / 65535
func BlendScreenRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = screen(Sr, Dr)
	G = screen(Sg, Dg)
	B = screen(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = screen(Sa, Da)
	}
	return
}

// R = 1 - (1 - S)*(1 - D)
// R = 255 - (255-S)*(255-D) / 255
// R = 65535 - (65535-S)*(65535-D) / 65535
func screen(S, D uint32) uint32 {
	return 65535 - (65535-S)*(65535-D)/65535
}
