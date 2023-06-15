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
	RegisterBlendFunc(DestinationOut, BlendDestinationOutColor, BlendDestinationOutRGBA)
}

// BlendDestinationOutColor
// R = D*(1 - Sa) [0,1]
// R = D*(255 - Sa)/255 [0,255]
// R = D*(65535 - Sa)/65535 [0,65535]
func BlendDestinationOutColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	_, _, _, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDestinationOutRGBA(0, 0, 0, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendDestinationOutRGBA
// R = D*(1 - Sa) [0,1]
// R = D*(255 - Sa)/255 [0,255]
// R = D*(65535 - Sa)/65535 [0,65535]
func BlendDestinationOutRGBA(_, _, _, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = destinationOut(Dr, Sa)
	G = destinationOut(Dg, Sa)
	B = destinationOut(Db, Sa)
	if destinationAlpha {
		A = Da
	} else {
		A = destinationOut(Da, Sa)
	}
	return
}

// R = D*(1 - Sa) [0,1]
// R = D*(255 - Sa)/255 [0,255]
// R = D*(65535 - Sa)/65535 [0,65535]
func destinationOut(D, Sa uint32) uint32 {
	return D * (65535 - Sa) / 65535
}
