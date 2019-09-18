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
	RegisterBlendFunc(DestinationOver, BlendDestinationOverColor, BlendDestinationOverRGBA)
}

//
// R = F*(1 - Ba) + B
// R = F*(255 - Ba)/255 + B
// R = F*(65535 - Ba)/65535 + B
func BlendDestinationOverColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDestinationOverRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = F*(1 - Ba) + B
// R = F*(255 - Ba)/255 + B
// R = F*(65535 - Ba)/65535 + B
func BlendDestinationOverRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = destinationOver(foreR, backR, backA)
	G = destinationOver(foreG, backG, backA)
	B = destinationOver(foreB, backB, backA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = destinationOver(foreA, backA, backA)
	}
	return
}

// R = F*(1 - Ba) + B
// R = F*(255 - Ba)/255 + B
// R = F*(65535 - Ba)/65535 + B
func destinationOver(F, B, Ba uint32) uint32 {
	return F*(65535-Ba)/65535 + B
}
