package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(Clear, BlendClearColor, BlendClearRGBA)
}

// BlendClearColor clears to transparent black (R = 0).
// BlendClearColor 清除为透明黑色（R = 0）。
// R = 0
func BlendClearColor(S, D color.Color, _ float64, _ bool) color.Color {
	return &color.RGBA64{}
}

// BlendClearRGBA clears to transparent black on RGBA64 channels (R = 0).
// BlendClearRGBA 对 RGBA64 通道清除为透明黑色（R = 0）。
// R = 0
func BlendClearRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, _ bool) (R, G, B, A uint32) {
	return 0, 0, 0, 0
}
