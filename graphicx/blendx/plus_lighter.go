//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(PlusLighter, BlendPlusLighterColor, BlendPlusLighterRGBA)
}

// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
// R = MIN(65535, S + D) [0,65535]
func BlendPlusLighterColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendPlusLighterRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
// R = MIN(65535, S + D) [0,65535]
func BlendPlusLighterRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = plusLighter(foreR, backR)
	G = plusLighter(foreG, backG)
	B = plusLighter(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = plusLighter(foreA, backA)
	}
	return
}

// R = MIN(1, S + D) [0,1]
// R = MIN(255, S + D) [0,255]
// R = MIN(65535, S + D) [0,65535]
func plusLighter(F, B uint32) uint32 {
	Add := B + F
	if Add > 65535 {
		return 65535
	} else {
		return Add
	}
}
