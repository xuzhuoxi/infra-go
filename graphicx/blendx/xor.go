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
	RegisterBlendFunc(Xor, BlendXorColor, BlendXorRGBA)
}

// BlendXorColor
// 异或模式
// R = S*(1 - Da) + D*(1 - Sa)
// R = (S*(255 - Da) + D*(255 - Sa))/255
// R = (S*(65535 - Da) + D*(65535 - Sa))/65535
// Note that the Porter-Duff "XOR" mode is only titularly related to the classical bitmap XOR operation (which is unsupported by CoreGraphics)
func BlendXorColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendXorRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendXorRGBA
// 异或模式
// R = S*(1 - Da) + D*(1 - Sa)
// R = (S*(255 - Da) + D*(255 - Sa))/255
// R = (S*(65535 - Da) + D*(65535 - Sa))/65535
// Note that the Porter-Duff "XOR" mode is only titularly related to the classical bitmap XOR operation (which is unsupported by CoreGraphics)
func BlendXorRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = xor(Sr, Dr, Sa, Da)
	G = xor(Sg, Dg, Sa, Da)
	B = xor(Sb, Db, Sa, Da)
	if destinationAlpha {
		A = Da
	} else {
		A = xor(Sa, Da, Sa, Da)
	}
	return
}

// R = S*(1 - Da) + D*(1 - Sa)
// R = (S*(255 - Da) + D*(255 - Sa))/255
// R = (S*(65535 - Da) + D*(65535 - Sa))/65535
func xor(S, D, Sa, Da uint32) uint32 {
	return (S*(65535-Da) + D*(65535-Sa)) / 65535
}
