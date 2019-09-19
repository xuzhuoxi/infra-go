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
// R = S*(1 - Da) + D*Sa
// R = (S*(255 - Da) + D*Sa) / 255
// R = (S*(65535 - Da) + D*Sa) / 65535
func BlendDestinationAtopColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDestinationAtopRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = S*(1 - Da) + D*Sa
// R = (S*(255 - Da) + D*Sa) / 255
// R = (S*(65535 - Da) + D*Sa) / 65535
func BlendDestinationAtopRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = destinationAtop(Sr, Dr, Sa, Da)
	G = destinationAtop(Sg, Dg, Sa, Da)
	B = destinationAtop(Sb, Db, Sa, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = destinationAtop(Sa, Da, Sa, Da)
	}
	return
}

// R = S*(1 - Da) + D*Sa
// R = (S*(255 - Da) + D*Sa) / 255
// R = (S*(65535 - Da) + D*Sa) / 65535
func destinationAtop(S, D, Sa, Da uint32) uint32 {
	return (S*(65535-Da) + D*Sa) / 65535
}
