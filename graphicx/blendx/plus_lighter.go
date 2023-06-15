// Package blendx
// Created by xuzhuoxi
// on 2019-05-25.
// @author xuzhuoxi
//
package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(PlusLighter, BlendPlusLighterColor, BlendPlusLighterRGBA)
}

// BlendPlusLighterColor
// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
// R = MIN(65535, S + D) [0,65535]
func BlendPlusLighterColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendPlusLighterRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendPlusLighterRGBA
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
