package graphicx

import "image/color"

type GrayAlgMode uint8

const (
	Green   GrayAlgMode = iota + 1
	Average
	BitMove
	Integer
	Float
)

var DefaultAlgMode = BitMove

// 仅取绿色求灰度图
func GrayWithGreen(r, g, b uint8) uint8 {
	return g
}

// 平均值法求灰度图
func GrayWithAverage(r, g, b uint8) uint8 {
	tmp := uint16(r) + uint16(g) + uint16(b)
	return uint8(tmp / 3)
}

// 移位方法求灰度图
func GrayWithBitMove(r, g, b uint8) uint8 {
	return uint8((uint16(r)*76 + uint16(g)*151 + uint16(b)*28) >> 8)
}

// 整数方法求灰度图
func GrayWithInteger(r, g, b uint8) uint8 {
	return uint8((uint16(r)*30 + uint16(g)*59 + uint16(b)*11) / 100)
}

// 浮点算法求灰度图
func GrayWithFloat(r, g, b uint8) uint8 {
	return uint8(float64(r)*0.3 + float64(g)*0.59 + float64(b)*0.11)
}

//-------------------------------------------------

// 转换自身为灰度
func GraySelfDefault(rgba *color.RGBA) {
	gray := GrayValue(rgba.R, rgba.G, rgba.B, DefaultAlgMode)
	rgba.R, rgba.G, rgba.B = gray, gray, gray
}

// 转换自身为灰度
func GraySelf(rgba *color.RGBA, algMode GrayAlgMode) {
	gray := GrayValue(rgba.R, rgba.G, rgba.B, algMode)
	rgba.R, rgba.G, rgba.B = gray, gray, gray
}

// 返回灰度
func GrayDefault(rgba *color.RGBA) color.Gray {
	gray := GrayValue(rgba.R, rgba.G, rgba.B, DefaultAlgMode)
	return color.Gray{gray}
}

// 返回灰度
func Gray(rgba *color.RGBA, algMode GrayAlgMode) color.Gray {
	gray := GrayValue(rgba.R, rgba.G, rgba.B, algMode)
	return color.Gray{gray}
}

// 灰度
func GrayValue(r, g, b uint8, algMode GrayAlgMode) uint8 {
	switch algMode {
	case Green:
		return GrayWithGreen(r, g, b)
	case Average:
		return GrayWithAverage(r, g, b)
	case BitMove:
		return GrayWithBitMove(r, g, b)
	case Integer:
		return GrayWithInteger(r, g, b)
	case Float:
		return GrayWithFloat(r, g, b)
	default:
		return 0
	}
}
