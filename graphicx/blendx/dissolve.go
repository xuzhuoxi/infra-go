package blendx

import (
	"image/color"
	"math/rand"
)

func init() {
	RegisterBlendFunc(Dissolve, BlendDissolveColor, BlendDissolveRGBA)
}

// BlendDissolveColor randomly picks source or destination by factor.
// BlendDissolveColor 按 factor 随机选择源色或目标色（溶解模式）。
func BlendDissolveColor(S, D color.Color, factor float64, destinationAlpha bool) color.Color {
	if rand.Float64() <= factor {
		return S
	} else {
		return D
	}
}

// BlendDissolveRGBA randomly picks source or destination by factor on RGBA64 channels.
// BlendDissolveRGBA 对 RGBA64 通道按 factor 随机选择源色或目标色（溶解模式）。
func BlendDissolveRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, factor float64, _ bool) (R, G, B, A uint32) {
	if rand.Float64() <= factor {
		return Sr, Sg, Sb, Sa
	} else {
		return Dr, Dg, Db, Da
	}
}
