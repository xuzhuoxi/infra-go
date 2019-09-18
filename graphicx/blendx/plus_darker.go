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
	RegisterBlendFunc(PlusDarker, BlendPlusDarkerColor, BlendPlusDarkerRGBA)
}

// R = MAX(0, (1-D) + (1-S)) [0,1]
// R = MAX(0, (255-D) + (255-S)) [0,255]
// R = MAX(0, (65535-D) + (65535-S)) [0,65535]
func BlendPlusDarkerColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendPlusDarkerRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// R = MAX(0, (1-D) + (1-S)) [0,1]
// R = MAX(0, (255-D) + (255-S)) [0,255]
// R = MAX(0, (65535-D) + (65535-S)) [0,65535]
func BlendPlusDarkerRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = plusDarker(foreR, backR)
	G = plusDarker(foreG, backG)
	B = plusDarker(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = plusDarker(foreA, backA)
	}
	return
}

// R = MAX(0, (1-D) + (1-S)) [0,1]
// R = MAX(0, (255-D) + (255-S)) [0,255]
// R = MAX(0, (65535-D) + (65535-S)) [0,65535]
func plusDarker(F, B uint32) uint32 {
	return 65535 + 65535 - F - B
}
