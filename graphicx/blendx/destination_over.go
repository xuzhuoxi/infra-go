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
// R = S*(1 - Da) + D
// R = S*(255 - Da)/255 + D
// R = S*(65535 - Da)/65535 + D
func BlendDestinationOverColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendDestinationOverRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = S*(1 - Da) + D
// R = S*(255 - Da)/255 + D
// R = S*(65535 - Da)/65535 + D
func BlendDestinationOverRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = destinationOver(Sr, Dr, Da)
	G = destinationOver(Sg, Dg, Da)
	B = destinationOver(Sb, Db, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = destinationOver(Sa, Da, Da)
	}
	return
}

// R = S*(1 - Da) + D
// R = S*(255 - Da)/255 + D
// R = S*(65535 - Da)/65535 + D
func destinationOver(S, D, Da uint32) uint32 {
	return S*(65535-Da)/65535 + D
}
