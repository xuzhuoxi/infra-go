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
	RegisterBlendFunc(DestinationIn, BlendDestinationInColor, BlendDestinationInRGBA)
}

// BlendDestinationInColor
// R = D*Sa [0,1]
// R = D*Sa/255 [0,255]
// R = D*Sa/65535 [0,65535]
func BlendDestinationInColor(foreColor, backColor color.Color, _ float64, destinationAlpha bool) color.Color {
	_, _, _, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendDestinationInRGBA(0, 0, 0, fA, bR, bG, bB, bA, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendDestinationInRGBA
// R = D*Sa [0,1]
// R = D*Sa/255 [0,255]
// R = D*Sa/65535 [0,65535]
func BlendDestinationInRGBA(_, _, _, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = destinationIn(Dr, Sa)
	G = destinationIn(Dg, Sa)
	B = destinationIn(Db, Sa)
	if destinationAlpha {
		A = Da
	} else {
		A = destinationIn(Da, Sa)
	}
	return
}

// R = D*Sa [0,1]
// R = D*Sa/255 [0,255]
// R = D*Sa/65535 [0,65535]
func destinationIn(D, Sa uint32) uint32 {
	return D * Sa / 65535
}
