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
	RegisterBlendFunc(SourceOut, BlendSourceOutColor, BlendSourceOutRGBA)
}

//
// R = S*(1 - Da)
// R = S*(255 - Da)/255
// R = S*(65535 - Da)/65535
func BlendSourceOutColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendSourceOutRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = S*(1 - Da)
// R = S*(255 - Da)/255
// R = S*(65535 - Da)/65535
func BlendSourceOutRGBA(Sr, Sg, Sb, Sa uint32, _, _, _, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = sourceOut(Sr, Da)
	G = sourceOut(Sg, Da)
	B = sourceOut(Sb, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = sourceOut(Sa, Da)
	}
	return
}

// R = S*(1 - Da)
// R = S*(255 - Da)/255
// R = S*(65535 - Da)/65535
func sourceOut(S, Da uint32) uint32 {
	return S * (65535 - Da) / 65535
}
