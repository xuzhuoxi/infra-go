package graphicx

import (
	"math"

	"github.com/xuzhuoxi/infra-go/mathx"
)

// HSV2RGB converts HSV to 16-bit RGB channel values.
// HSV2RGB 将 HSV 转换为 16 位 RGB 通道值。
// H: hue in [0, 360].
// H：色调（色相），范围 [0, 360]。
// S: saturation in [0, 1].
// S：饱和度，范围 [0, 1]。
// V: value (brightness) in [0, 1].
// V：明度（亮度），范围 [0, 1]。
func HSV2RGB(H, S, V float64) (R, G, B uint32) {
	if 0 == S {
		return 0, 0, 0
	}
	H /= 60
	i := int(H)
	f := H - float64(i)

	a := uint32(V * (1 - S))
	b := uint32(V * (1 - S*f))
	c := uint32(V * (1 - S*(1-f)))

	switch i {
	case 0:
		R, G, B = uint32(V), c, a
	case 1:
		R, G, B = b, uint32(V), a
	case 2:
		R, G, B = a, uint32(V), c
	case 3:
		R, G, B = a, b, uint32(V)
	case 4:
		R, G, B = c, a, uint32(V)
	case 5:
		R, G, B = uint32(V), a, b
	}
	return
}

// RGB2HSV converts 16-bit RGB channel values to HSV.
// RGB2HSV 将 16 位 RGB 通道值转换为 HSV。
// H: hue in [0, 360].
// H：色调（色相），范围 [0, 360]。
// S: saturation in [0, 1].
// S：饱和度，范围 [0, 1]。
// V: value (brightness) in [0, 1].
// V：明度（亮度），范围 [0, 1]。
func RGB2HSV(R, G, B uint32) (H, S, V float64) {
	fr, fg, fb := float64(R), float64(G), float64(B)
	max := mathx.Max3Float64(fr, fg, fb)
	min := mathx.Min3Float64(fr, fg, fb)
	S = 1 - min/max
	V = max / math.MaxUint16

	diff := max - min
	if 0 == diff {
		H = 0
		return
	}
	switch max {
	case fb:
		H = 240 + 60*(fr-fg)/diff
	case fg:
		H = 120 + 60*(fb-fr)/diff
	case fr:
		H = 0 + 60*(fg-fb)/diff
	}
	if H < 0 {
		H = H + 360
	}
	return
}

// Hue returns the hue component H in [0, 360] from 16-bit RGB channels.
// Hue 从 16 位 RGB 通道计算色相 H，范围 [0, 360]。
func Hue(R, G, B uint32) (H float64) {
	fr, fg, fb := float64(R), float64(G), float64(B)
	max := mathx.Max3Float64(fr, fg, fb)
	min := mathx.Min3Float64(fr, fg, fb)
	diff := max - min
	if 0 == diff {
		return 0
	}
	switch max {
	case fb:
		H = 240 + 60*(fr-fg)/diff
	case fg:
		H = 120 + 60*(fb-fr)/diff
	case fr:
		H = 0 + 60*(fg-fb)/diff
	}
	if H < 0 {
		H = H + 360
	}
	return
}

// Saturation returns the saturation S in [0, 1] from 16-bit RGB channels.
// Saturation 从 16 位 RGB 通道计算饱和度 S，范围 [0, 1]。
func Saturation(R, G, B uint32) (S float64) {
	fr, fg, fb := float64(R), float64(G), float64(B)
	max := mathx.Max3Float64(fr, fg, fb)
	min := mathx.Min3Float64(fr, fg, fb)
	return 1 - min/max
}

// Value returns the value (brightness) V in [0, 1] from 16-bit RGB channels.
// Value 从 16 位 RGB 通道计算明度 V，范围 [0, 1]。
func Value(R, G, B uint32) (V float64) {
	fr, fg, fb := float64(R), float64(G), float64(B)
	max := mathx.Max3Float64(fr, fg, fb)
	return max / math.MaxUint16
}
