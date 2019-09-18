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
	RegisterBlendFunc(LinearBurn, BlendLinearBurnColor, BlendLinearBurnRGBA)
}

// 线性加深模式
// 查看每个通道的颜色信息，通过降低“亮度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// R = B + F - 255
// R = B + F - 65535
func BlendLinearBurnColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendLinearBurnRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 线性加深模式
// 查看每个通道的颜色信息，通过降低“亮度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// R = B + F - 255
// R = B + F - 65535
func BlendLinearBurnRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = linearBurn(foreR, backR)
	G = linearBurn(foreG, backG)
	B = linearBurn(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = linearBurn(foreA, backA)
	}
	return
}

// R = B + F - 255
// R = B + F - 65535
func linearBurn(F, B uint32) uint32 {
	return B + F - 65535
}
