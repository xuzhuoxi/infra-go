package graphicx

import (
	"math"
)

// YUV converts 16-bit RGB channels to Y, U, V components.
// YUV 将 16 位 RGB 通道转换为 Y、U、V 分量。
func YUV(R, G, B uint32) (Y, U, V float64) {
	Y = Luminance(R, G, B)
	U = Chrominance(R, G, B)
	V = Chroma(R, G, B)
	return
}

// Luminance returns the luma Y in [0, 1]: Y = 0.299*R + 0.587*G + 0.114*B.
// Luminance 返回明亮度（灰阶）Y，范围 [0, 1]：Y = 0.299*R + 0.587*G + 0.114*B。
func Luminance(R, G, B uint32) (Y float64) {
	return (0.299*float64(R) + 0.587*float64(G) + 0.114*float64(B)) / math.MaxUint16
}

// Chrominance returns chroma U in [0, 1]: U = -0.14713*R - 0.28886*G + 0.436*B.
// Chrominance 返回色度 U，范围 [0, 1]：U = -0.14713*R - 0.28886*G + 0.436*B。
func Chrominance(R, G, B uint32) (U float64) {
	return (-0.14713*float64(R) - 0.28886*float64(G) + 0.436*float64(B)) / math.MaxUint16
}

// Chroma returns chroma V in [0, 1]: V = 0.615*R - 0.51499*G - 0.10001*B.
// Chroma 返回浓度 V，范围 [0, 1]：V = 0.615*R - 0.51499*G - 0.10001*B。
func Chroma(R, G, B uint32) (V float64) {
	return (0.615*float64(R) + 0.51499*float64(G) + 0.10001*float64(B)) / math.MaxUint16
}

// 其它：
// Y = 0.299 * r + 0.587 * g + 0.114 * b
// U = 0.1687* r - 0.3313* g + 0.5 * b + 128
// V = 0.5 * r - 0.4187*g - 0.0813*b + 128
