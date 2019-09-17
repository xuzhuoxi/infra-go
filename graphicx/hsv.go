package graphicx

import (
	"github.com/xuzhuoxi/infra-go/mathx"
	"math"
)

// 取HSV
// 色调 H [0,360]
// 饱和度 S [0,1]
// 明度 [0,1]
func HSV(R, G, B uint32) (H, S, V float64) {
	fr, fg, fb := float64(R), float64(G), float64(B)
	max := mathx.Max3Float64(fr, fg, fb)
	min := mathx.Min3Float64(fr, fg, fb)
	S = 1 - min/max
	V = max / (math.MaxUint16 + 1)

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
	return max / (math.MaxUint16 + 1)
}
