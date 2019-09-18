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
	RegisterBlendFunc(SourceIn, BlendSourceInColor, BlendSourceInRGBA)
}

//
// R = B*Fa
// R = B*Fa/255
// R = B*Fa/65535
func BlendSourceInColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendSourceInRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = B*Fa
// R = B*Fa/255
// R = B*Fa/65535
func BlendSourceInRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = sourceIn(foreR, backR, foreA)
	G = sourceIn(foreG, backG, foreA)
	B = sourceIn(foreB, backB, foreA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = sourceIn(foreA, backA, foreA)
	}
	return
}

// R = B*Fa
// R = B*Fa/255
// R = B*Fa/65535
func sourceIn(_, B, Fa uint32) uint32 {
	return B * Fa / 65535
}
