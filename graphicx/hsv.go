package graphicx

import (
	"github.com/xuzhuoxi/infra-go/mathx"
	"math"
)

// 求RGB 64位
// H: 色调(色相) [0,360]
// S: 饱和度 [0,1]
// V: 明度(亮度) [0,1]
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

// 取HSV
// H: 色调(色相) [0,360]
// S: 饱和度 [0,1]
// V: 明度(亮度) [0,1]
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

// 色调 H [0,360]
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

// 饱和度 S [0,1]
func Saturation(R, G, B uint32) (S float64) {
	fr, fg, fb := float64(R), float64(G), float64(B)
	max := mathx.Max3Float64(fr, fg, fb)
	min := mathx.Min3Float64(fr, fg, fb)
	return 1 - min/max
}

// 明度 [0,1]
func Value(R, G, B uint32) (V float64) {
	fr, fg, fb := float64(R), float64(G), float64(B)
	max := mathx.Max3Float64(fr, fg, fb)
	return max / math.MaxUint16
}
