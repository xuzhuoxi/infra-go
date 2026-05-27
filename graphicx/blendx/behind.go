package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(Behind, BlendBehindColor, BlendBehindRGBA)
}

// BlendBehindColor composites source behind destination where destination is transparent.
// BlendBehindColor 在目标透明区域将源色绘制于目标背后（背后模式）。
func BlendBehindColor(S, D color.Color, _ float64, _ bool) color.Color {
	_, _, _, Da := D.RGBA()
	if Da < math.MaxUint16 {
		return S
	} else {
		return D
	}
}

// BlendBehindRGBA composites source behind destination on RGBA64 channels.
// BlendBehindRGBA 对 RGBA64 通道应用背后（Behind）混合。
func BlendBehindRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if Da < math.MaxUint16 {
		return Sr, Sg, Sb, Sa
	} else {
		return Dr, Dg, Db, Da
	}
}
