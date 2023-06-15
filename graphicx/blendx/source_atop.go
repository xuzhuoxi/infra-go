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
	RegisterBlendFunc(SourceAtop, BlendSourceAtopColor, BlendSourceAtopRGBA)
}

// BlendSourceAtopColor
// R = S*Da + D*(1 - Sa)
// R = (S*Da + D*(255 - Sa))/255
// R = (S*Da + D*(65535 - Sa))/65535
func BlendSourceAtopColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendSourceAtopRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendSourceAtopRGBA
// R = S*Da + D*(1 - Sa)
// R = (S*Da + D*(255 - Sa))/255
// R = (S*Da + D*(65535 - Sa))/65535
func BlendSourceAtopRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = sourceAtop(Sr, Dr, Sa, Da)
	G = sourceAtop(Sg, Dg, Sa, Da)
	B = sourceAtop(Sb, Db, Sa, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = sourceAtop(Sa, Da, Sa, Da)
	}
	return
}

// R = S*Da + D*(1 - Sa)
// R = (S*Da + D*(255 - Sa))/255
// R = (S*Da + D*(65535 - Sa))/65535
func sourceAtop(S, D, Sa, Da uint32) uint32 {
	return (S*Da + D*(65535-Sa)) / 65535
}
