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
// R = B*(1 - Fa)
// R = B*(255 - Fa)/255
// R = B*(65535 - Fa)/65535
func BlendSourceOutColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendSourceOutRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

//
// R = B*(1 - Fa)
// R = B*(255 - Fa)/255
// R = B*(65535 - Fa)/65535
func BlendSourceOutRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = sourceOut(foreR, backR, foreA)
	G = sourceOut(foreG, backG, foreA)
	B = sourceOut(foreB, backB, foreA)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = sourceOut(foreA, backA, foreA)
	}
	return
}

// R = B*(1 - Fa)
// R = B*(255 - Fa)/255
// R = B*(65535 - Fa)/65535
func sourceOut(F_, B, Fa uint32) uint32 {
	return B * (65535 - Fa) / 65535
}
