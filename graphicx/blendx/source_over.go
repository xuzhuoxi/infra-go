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
	RegisterBlendFunc(SourceOver, BlendSourceOverColor, BlendSourceOverRGBA)
}

//
// R = D*(1 - Sa) + S
// R = D*(1 - Sa)/255 + S
// R = D*(1 - Sa)/65535 + S
func BlendSourceOverColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendSourceOverRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = D*(1 - Sa) + S
// R = D*(1 - Sa)/255 + S
// R = D*(1 - Sa)/65535 + S
func BlendSourceOverRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = sourceOver(Sr, Dr, Da)
	G = sourceOver(Sg, Dg, Da)
	B = sourceOver(Sb, Db, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = sourceOver(Sa, Da, Da)
	}
	return
}

// R = D*(1 - Sa) + S
// R = D*(1 - Sa)/255 + S
// R = D*(1 - Sa)/65535 + S
func sourceOver(S, D, Sa uint32) uint32 {
	return D*(1-Sa)/65535 + S
}
