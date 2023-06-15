package graphicx

import "image/color"

type GrayAlgMode uint8

const (
	Green GrayAlgMode = iota + 1
	Average
	BitMove
	Integer
	Float
)

var DefaultAlgMode = BitMove

// GrayWithGreen
// 仅取绿色求灰度图
func GrayWithGreen(R, G, B uint32) uint32 {
	return G
}

// GrayWithAverage
// 平均值法求灰度图
func GrayWithAverage(R, G, B uint32) uint32 {
	return (R + G + B) / 3
}

// GrayWithBitMove
// 移位方法求灰度图
func GrayWithBitMove(R, G, B uint32) uint32 {
	return (R*76 + G*151 + B*28) >> 8
}

// GrayWithInteger
// 整数方法求灰度图
func GrayWithInteger(R, G, B uint32) uint32 {
	return (R*299 + G*578 + B*114) / 1000
}

// GrayWithFloat
// 浮点算法求灰度图
func GrayWithFloat(R, G, B uint32) uint32 {
	return uint32(float64(R)*0.299 + float64(G)*0.587 + float64(B)*0.114)
}

//-------------------------------------------------

// GrayColorDefault
// 返回灰度颜色
func GrayColorDefault(c color.Color) color.Color {
	return GrayColor(c, DefaultAlgMode)
}

// GrayColor
// 返回灰度颜色
func GrayColor(c color.Color, algMode GrayAlgMode) color.Color {
	R, G, B, _ := c.RGBA()
	Y := uint16(GrayRGB(R, G, B, algMode))
	return &color.Gray16{Y: Y}
}

// GrayRGB
// 转换为灰度值
func GrayRGB(R, G, B uint32, algMode GrayAlgMode) uint32 {
	switch algMode {
	case Green:
		return GrayWithGreen(R, G, B)
	case Average:
		return GrayWithAverage(R, G, B)
	case BitMove:
		return GrayWithBitMove(R, G, B)
	case Integer:
		return GrayWithInteger(R, G, B)
	case Float:
		return GrayWithFloat(R, G, B)
	default:
		return 0
	}
}
