package graphicx

import "image/color"

// Color64To32 converts a 16-bit color component to 8-bit by shifting right 8 bits.
// Color64To32 将 16 位色阶分量右移 8 位转换为 8 位。
func Color64To32(pixelUnit uint32) uint32 {
	return pixelUnit >> 8
}

// Color64ToFloat converts a 16-bit color component to a normalized float in [0, 1].
// Color64ToFloat 将 16 位色阶分量转换为 [0, 1] 区间的浮点值。
func Color64ToFloat(pixelUnit uint32) float64 {
	return float64(pixelUnit) / 65536
}

// Color32To64 converts an 8-bit color component to 16-bit by shifting left 8 bits.
// Color32To64 将 8 位色阶分量左移 8 位转换为 16 位。
func Color32To64(pixelUnit uint32) uint32 {
	return pixelUnit << 8
}

// Color32ToFloat converts an 8-bit color component to a normalized float in [0, 1].
// Color32ToFloat 将 8 位色阶分量转换为 [0, 1] 区间的浮点值。
func Color32ToFloat(pixelUnit uint32) float64 {
	return float64(pixelUnit) / 256
}

// ColorFloatTo32 converts a normalized float in [0, 1] to an 8-bit color component.
// ColorFloatTo32 将 [0, 1] 区间的浮点值转换为 8 位色阶分量。
func ColorFloatTo32(pixelUnit float64) uint32 {
	return uint32(pixelUnit * 256)
}

// ColorFloatTo64 converts a normalized float in [0, 1] to a 16-bit color component.
// ColorFloatTo64 将 [0, 1] 区间的浮点值转换为 16 位色阶分量。
func ColorFloatTo64(pixelUnit float64) uint32 {
	return uint32(pixelUnit * 65536)
}

// GetRed returns the red channel of c from RGBA().
// GetRed 从 RGBA() 返回颜色 c 的红色通道。
func GetRed(c color.Color) (R uint32) {
	R, _, _, _ = c.RGBA()
	return
}

// GetGreen returns the green channel of c from RGBA().
// GetGreen 从 RGBA() 返回颜色 c 的绿色通道。
func GetGreen(c color.Color) (G uint32) {
	_, G, _, _ = c.RGBA()
	return
}

// GetBlue returns the blue channel of c from RGBA().
// GetBlue 从 RGBA() 返回颜色 c 的蓝色通道。
func GetBlue(c color.Color) (B uint32) {
	_, _, B, _ = c.RGBA()
	return
}

// GetAlpha returns the alpha channel of c from RGBA().
// GetAlpha 从 RGBA() 返回颜色 c 的透明通道。
func GetAlpha(c color.Color) (A uint32) {
	_, _, _, A = c.RGBA()
	return
}
