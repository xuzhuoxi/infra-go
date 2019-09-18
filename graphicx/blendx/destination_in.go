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
	RegisterBlendFunc(DestinationIn, BlendDestinationInColor, BlendDestinationInRGBA)
}

// R = B*Fa [0,1]
// R = B*Fa/255 [0,255]
// R = B*Fa/65535 [0,65535]
func BlendDestinationInColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	_, _, _, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDestinationInRGBA(0, 0, 0, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// R = B*Fa [0,1]
// R = B*Fa/255 [0,255]
// R = B*Fa/65535 [0,65535]
func BlendDestinationInRGBA(_, _, _, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = destinationIn(backR, foreA)
	G = destinationIn(backG, foreA)
	B = destinationIn(backB, foreA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = destinationIn(backA, foreA)
	}
	return
}

// R = B*Fa [0,1]
// R = B*Fa/255 [0,255]
// R = B*Fa/65535 [0,65535]
func destinationIn(B, Fa uint32) uint32 {
	return B * Fa / 65535
}
