package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(DarkerColor, BlendDarkerColorColor, BlendDarkerColorRGBA)
}

// BlendDarkerColorColor keeps the darker RGB tuple of source and destination.
// BlendDarkerColorColor 保留源色与目标色中较暗的 RGB（深色模式）。
func BlendDarkerColorColor(S, D color.Color, _ float64, _ bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	if Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa {
		return S
	} else {
		return D
	}
}

// BlendDarkerColorRGBA keeps the darker RGB tuple on RGBA64 channels.
// BlendDarkerColorRGBA 对 RGBA64 通道保留较暗 RGB（深色模式）。
func BlendDarkerColorRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa {
		return Sr, Sg, Sb, Sa
	} else {
		return Dr, Dg, Db, Da
	}
}
