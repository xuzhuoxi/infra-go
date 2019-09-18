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
	RegisterBlendFunc(Copy, BlendCopyColor, BlendCopyRGBA)
}

// 覆盖模式
// R = D
func BlendCopyColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	if !keepForegroundAlpha {
		return backColor
	}
	_, _, _, fA := foreColor.RGBA()
	bR, bG, bB, _ := backColor.RGBA()
	return &color.RGBA64{R: uint16(bR), G: uint16(bG), B: uint16(bB), A: uint16(fA)}
}

// 覆盖模式
// R = D
func BlendCopyRGBA(_, _, _, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R, G, B = backR, backG, backB
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = backA
	}
	return
}
