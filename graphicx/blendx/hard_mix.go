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
	RegisterBlendFunc(HardMix, BlendHardMixColor, BlendHardMixRGBA)
}

// 实色混合模式
// 根据绘图颜色与底图颜色的颜色数值相加，当相加的颜色数值大于该颜色模式颜色数值的最大值，混合颜色为最大值；
// 当相加的颜色数值小于该颜色模式颜色数值的最大值，混合颜色为0；
// 当相加的颜色数值等于该颜色模式颜色数值的最大值，混合颜色由底图颜色决定，底图颜色值比绘图颜色的颜色值大，则混合颜色为最大值，相反则为0.实色混合能产生颜色较少、边缘较硬的图像效果。
// F+B>=255 : R = 255
// F+B<255 : R = 0
func BlendHardMixColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendHardMixRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 实色混合模式
// 根据绘图颜色与底图颜色的颜色数值相加，当相加的颜色数值大于该颜色模式颜色数值的最大值，混合颜色为最大值；
// 当相加的颜色数值小于该颜色模式颜色数值的最大值，混合颜色为0；
// 当相加的颜色数值等于该颜色模式颜色数值的最大值，混合颜色由底图颜色决定，底图颜色值比绘图颜色的颜色值大，则混合颜色为最大值，相反则为0.实色混合能产生颜色较少、边缘较硬的图像效果。
// R = F+B>=255 ? 255 : 0
// R = F+B>=65535 ? 65535 : 0
func BlendHardMixRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = hardMix(foreR, backR)
	G = hardMix(foreG, backG)
	B = hardMix(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = hardMix(foreA, backA)
	}
	return
}

// R = F+B>=255 ? 255 : 0
// R = F+B>=65535 ? 65535 : 0
func hardMix(F, B uint32) uint32 {
	if F+B >= 65535 {
		return 65535
	} else {
		return 0
	}
}
