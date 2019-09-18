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
	RegisterBlendFunc(DestinationAtop, BlendDestinationAtopColor, BlendDestinationAtopRGBA)
}

//
// R = B*(1 - Fa) + F*Ba
// R = (B*(255-Fa) + F*Ba)/255
// R = (B*(65535-Fa) + F*Ba)/65535
func BlendDestinationAtopColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDestinationAtopRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = B*(1 - Fa) + F*Ba
// R = (B*(255-Fa) + F*Ba)/255
// R = (B*(65535-Fa) + F*Ba)/65535
func BlendDestinationAtopRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = destinationAtop(foreR, backR, foreA, backA)
	G = destinationAtop(foreG, backG, foreA, backA)
	B = destinationAtop(foreB, backB, foreA, backA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = destinationAtop(foreA, backA, foreA, backA)
	}
	return
}

// R = B*(1 - Fa) + F*Ba
// R = (B*(255-Fa) + F*Ba)/255
// R = (B*(65535-Fa) + F*Ba)/65535
func destinationAtop(F, B, Fa, Ba uint32) uint32 {
	return (B*(65535-Fa) + F*Ba) / 65535
}
