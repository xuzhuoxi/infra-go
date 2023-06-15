// Package blendx
// Created by xuzhuoxi
// on 2019-05-25.
// @author xuzhuoxi
//
package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(ColorBurn, BlendColorBurnColor, BlendColorBurnRGBA)
}

// BlendColorBurnColor
// 颜色加深模式
// 查看每个通道的颜色信息，通过增加“对比度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// 除了背景上的较淡区域消失，且图像区域呈现尖锐的边缘特性之外，这种Color Burn模式创建的效果类似于由MuItiply模式创建的效果。
// R = S - (1-S) * (1-D) / D
// R = S - (255-S) * (255-D) / D
// R = S - (65535-S) * (65535-D) / D (64位图像)
func BlendColorBurnColor(S, D color.Color, factor float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendColorBurnRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.NRGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendColorBurnRGBA
// 颜色加深模式
// 查看每个通道的颜色信息，通过增加“对比度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// 除了背景上的较淡区域消失，且图像区域呈现尖锐的边缘特性之外，这种Color Burn模式创建的效果类似于由MuItiply模式创建的效果。
// R = S - (1-S) * (1-D) / D
// R = S - (255-S) * (255-D) / D
// R = S - (65535-S) * (65535-D) / D (64位图像)
func BlendColorBurnRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = colorBurn(Sr, Dr)
	G = colorBurn(Sg, Dg)
	B = colorBurn(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = colorBurn(Sa, Da)
	}
	return
}

// R = S - (1-S) * (1-D) / D
// R = S - (255-S) * (255-D) / D
// R = S - (65535-S) * (65535-D) / D (64位图像)
func colorBurn(S, D uint32) uint32 {
	return S - (math.MaxUint16-S)*(math.MaxUint16-D)/D
}
