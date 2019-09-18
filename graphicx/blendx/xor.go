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
	RegisterBlendFunc(Xor, BlendXorColor, BlendXorRGBA)
}

// 异或模式
// R = S*(1-Da) + D*(1-Sa)
// R = (B*(255 - Fa)+ F*(255 - Ba)) / 255
// R = (B*(65535 - Fa)+ F*(65535 - Ba)) / 65535
// Note that the Porter-Duff "XOR" mode is only titularly related to the classical bitmap XOR operation (which is unsupported by CoreGraphics)
func BlendXorColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendXorRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 异或模式
// R = S*(1-Da) + D*(1-Sa)
// R = (B*(255 - Fa)+ F*(255 - Ba)) / 255
// R = (B*(65535 - Fa)+ F*(65535 - Ba)) / 65535
// Note that the Porter-Duff "XOR" mode is only titularly related to the classical bitmap XOR operation (which is unsupported by CoreGraphics)
func BlendXorRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = xor(foreR, backR, foreA, backA)
	G = xor(foreG, backG, foreA, backA)
	B = xor(foreB, backB, foreA, backA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = xor(foreA, backA, foreA, backA)
	}
	return
}

// R = S*(1-Da) + D*(1-Sa)
// R = (B*(255 - Fa)+ F*(255 - Ba)) / 255
// R = (B*(65535 - Fa)+ F*(65535 - Ba)) / 65535
func xor(F, B, Fa, Ba uint32) uint32 {
	return (B*(65535-Fa) + F*(65535-Ba)) / 65535
}
