package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(LighterColor, BlendLighterColor, BlendLighterColorRGBA)
}

// BlendLighterColor keeps the lighter RGB tuple of source and destination.
// BlendLighterColor 保留源色与目标色中较亮的 RGB（浅色模式）。
func BlendLighterColor(foreColor, backColor color.Color, _ float64, _ bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	if fR+fG+fB+fA >= bR+bG+bB+bA {
		return foreColor
	} else {
		return backColor
	}
}

// BlendLighterColorRGBA keeps the lighter RGB tuple on RGBA64 channels.
// BlendLighterColorRGBA 对 RGBA64 通道保留较亮 RGB（浅色模式）。
func BlendLighterColorRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, _ bool) (R, G, B, A uint32) {
	if Dr+Dg+Db+Da >= Sr+Sg+Sb+Sa {
		return Dr, Dg, Db, Da
	} else {
		return Sr, Sg, Sb, Sa
	}
}
