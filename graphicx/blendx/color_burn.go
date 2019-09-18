//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(ColorBurn, BlendColorBurnColor, BlendColorBurnRGBA)
}

// 颜色加深模式
// 查看每个通道的颜色信息，通过增加“对比度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// 除了背景上的较淡区域消失，且图像区域呈现尖锐的边缘特性之外，这种Color Burn模式创建的效果类似于由MuItiply模式创建的效果。
// R = B - ((1 - B) * (1 - F)) / F
// R = B - (255-B)*(255-F) / F
// R = B - (65535-B)*(65535-F) / F (64位图像)
func BlendColorBurnColor(foreColor, backColor color.Color, factor float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendColorBurnRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.NRGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 颜色加深模式
// 查看每个通道的颜色信息，通过增加“对比度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// 除了背景上的较淡区域消失，且图像区域呈现尖锐的边缘特性之外，这种Color Burn模式创建的效果类似于由MuItiply模式创建的效果。
// R = B - ((1 - B) * (1 - F)) / F
// R = B - (255-B)*(255-F) / F
// R = B - (65535-B)*(65535-F) / F (64位图像)
func BlendColorBurnRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = colorBurn(foreR, backR)
	G = colorBurn(foreG, backG)
	B = colorBurn(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = colorBurn(foreA, backA)
	}
	return
}

// R = B - ((1 - B) * (1 - F)) / F
// R = B - (255-B)*(255-F) / F (32位图像)
// R = B - (65535-B)*(65535-F) / F (64位图像)
func colorBurn(F, B uint32) uint32 {
	return F - (math.MaxUint16-F)*(math.MaxUint16-B)/B
}
