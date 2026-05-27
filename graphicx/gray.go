package graphicx

import "image/color"

// GrayAlgMode selects a grayscale conversion algorithm.
// GrayAlgMode 指定灰度转换算法。
type GrayAlgMode uint8

const (
	Green   GrayAlgMode = iota + 1 // Green channel only. 仅使用绿色通道。
	Average                        // Arithmetic mean of RGB. RGB 算术平均。
	BitMove                        // Fixed-point weighted sum with bit shift. 定点加权求和并移位。
	Integer                        // Integer weighted sum. 整数加权求和。
	Float                          // Floating-point weighted sum. 浮点加权求和。
)

var DefaultAlgMode = BitMove // Default grayscale algorithm. 默认灰度算法。

// GrayWithGreen returns the green channel as grayscale.
// GrayWithGreen 仅取绿色通道作为灰度值。
func GrayWithGreen(R, G, B uint32) uint32 {
	return G
}

// GrayWithAverage returns the arithmetic mean of R, G, B as grayscale.
// GrayWithAverage 以 R、G、B 算术平均作为灰度值。
func GrayWithAverage(R, G, B uint32) uint32 {
	return (R + G + B) / 3
}

// GrayWithBitMove returns grayscale using fixed-point weights and a bit shift.
// GrayWithBitMove 使用定点权重与移位计算灰度值。
func GrayWithBitMove(R, G, B uint32) uint32 {
	return (R*76 + G*151 + B*28) >> 8
}

// GrayWithInteger returns grayscale using integer luminance weights.
// GrayWithInteger 使用整数亮度权重计算灰度值。
func GrayWithInteger(R, G, B uint32) uint32 {
	return (R*299 + G*578 + B*114) / 1000
}

// GrayWithFloat returns grayscale using floating-point luminance weights.
// GrayWithFloat 使用浮点亮度权重计算灰度值。
func GrayWithFloat(R, G, B uint32) uint32 {
	return uint32(float64(R)*0.299 + float64(G)*0.587 + float64(B)*0.114)
}

//-------------------------------------------------

// GrayColorDefault converts c to grayscale using DefaultAlgMode.
// GrayColorDefault 使用 DefaultAlgMode 将 c 转换为灰度颜色。
func GrayColorDefault(c color.Color) color.Color {
	return GrayColor(c, DefaultAlgMode)
}

// GrayColor converts c to grayscale using the given algorithm.
// GrayColor 使用指定算法将 c 转换为灰度颜色。
func GrayColor(c color.Color, algMode GrayAlgMode) color.Color {
	R, G, B, _ := c.RGBA()
	Y := uint16(GrayRGB(R, G, B, algMode))
	return &color.Gray16{Y: Y}
}

// GrayRGB converts R, G, B to a grayscale value using algMode.
// GrayRGB 使用 algMode 将 R、G、B 转换为灰度值。
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
