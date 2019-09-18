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
	RegisterBlendFunc(SourceAtop, BlendSourceAtopColor, BlendSourceAtopRGBA)
}

//
// R = B*Fa + F*(1-Ba)
// R = B*Fa/255 + F*(255 - Ba)/255
// R = B*Fa/65535 + F*(65535 - Ba)/65535
func BlendSourceAtopColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendSourceAtopRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = B*Fa + F*(1-Ba)
// R = B*Fa/255 + F*(255 - Ba)/255
// R = B*Fa/65535 + F*(65535 - Ba)/65535
func BlendSourceAtopRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = sourceAtop(foreR, backR, foreA, backA)
	G = sourceAtop(foreG, backG, foreA, backA)
	B = sourceAtop(foreB, backB, foreA, backA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = sourceAtop(foreA, backA, foreA, backA)
	}
	return
}

// R = B*Fa + F*(1-Ba)
// R = B*Fa/255 + F*(255 - Ba)/255
// R = B*Fa/65535 + F*(65535 - Ba)/65535
func sourceAtop(F, B, Fa, Ba uint32) uint32 {
	return B*Fa/65535 + F*(65535-Ba)/65535
}
