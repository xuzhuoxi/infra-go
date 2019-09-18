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
	RegisterBlendFunc(DestinationOut, BlendDestinationOutColor, BlendDestinationOutRGBA)
}

//
// R = B*(1 - Fa) [0,1]
// R = B*(255-Fa)/255 [0,255]
// R = B*(65535-Fa)/65535 [0,65535]
func BlendDestinationOutColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	_, _, _, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDestinationOutRGBA(0, 0, 0, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = B*(1 - Fa) [0,1]
// R = B*(255-Fa)/255 [0,255]
// R = B*(65535-Fa)/65535 [0,65535]
func BlendDestinationOutRGBA(_, _, _, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = destinationOut(backR, foreA)
	G = destinationOut(backG, foreA)
	B = destinationOut(backB, foreA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = destinationOut(backA, foreA)
	}
	return
}

// R = B*(1 - Fa) [0,1]
// R = B*(255-Fa)/255 [0,255]
// R = B*(65535-Fa)/65535 [0,65535]
func destinationOut(B, Fa uint32) uint32 {
	return B * (65535 - Fa) / 65535
}
