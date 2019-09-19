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
	RegisterBlendFunc(PlusDarker, BlendPlusDarkerColor, BlendPlusDarkerRGBA)
}

// R = MAX(0, (1 - D) + (1 - S)) [0,1]
// R = MAX(0, (255-D) + (255-S)) [0,255]
// R = MAX(0, (65535-D) + (65535-S)) [0,65535]
func BlendPlusDarkerColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendPlusDarkerRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// R = MAX(0, (1 - D) + (1 - S)) [0,1]
// R = MAX(0, (255-D) + (255-S)) [0,255]
// R = MAX(0, (65535-D) + (65535-S)) [0,65535]
func BlendPlusDarkerRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = plusDarker(Sr, Dr)
	G = plusDarker(Sg, Dg)
	B = plusDarker(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = plusDarker(Sa, Da)
	}
	return
}

// R = MAX(0, (1 - D) + (1 - S)) [0,1]
// R = MAX(0, (255-D) + (255-S)) [0,255]
// R = MAX(0, (65535-D) + (65535-S)) [0,65535]
func plusDarker(S, D uint32) uint32 {
	return 65535 + 65535 - S - D
}
