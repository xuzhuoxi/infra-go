package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(PlusLighter, BlendPlusLighterColor, BlendPlusLighterRGBA)
}

// BlendPlusLighterColor applies plus lighter blend.
// BlendPlusLighterColor 应用高级浅色（Plus Lighter）混合。
// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
// R = MIN(65535, S + D) [0,65535]
func BlendPlusLighterColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendPlusLighterRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendPlusLighterRGBA applies plus lighter on RGBA64 channels.
// BlendPlusLighterRGBA 对 RGBA64 通道应用高级浅色混合。
// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
// R = MIN(65535, S + D) [0,65535]
func BlendPlusLighterRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = plusLighter(Sr, Dr)
	G = plusLighter(Sg, Dg)
	B = plusLighter(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = plusLighter(Sa, Da)
	}
	return
}

// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
// R = MIN(65535, S + D) [0,65535]
func plusLighter(S, D uint32) uint32 {
	Add := D + S
	if Add < 65535 {
		return Add
	} else {
		return 65535
	}
}
