//
//Created by xuzhuoxi
//on 2019-05-29.
//@author xuzhuoxi
//
package graphicx

import "image/color"

// 64位像素单元转32位像素单元
func Color64To32(pixelUnit uint32) uint32 {
	return pixelUnit >> 8
}

// 64位像素单元转32位像素单元
func Color64ToFloat(pixelUnit uint32) float64 {
	return float64(pixelUnit) / 65536
}

// 32位像素单元转64位像素单元
func Color32To64(pixelUnit uint32) uint32 {
	return pixelUnit << 8
}

// 32位像素单元转32位像素单元
func Color32ToFloat(pixelUnit uint32) float64 {
	return float64(pixelUnit) / 256
}

// 浮点像素单元转32位像素单元
func ColorFloatTo32(pixelUnit float64) uint32 {
	return uint32(pixelUnit * 256)
}

func ColorFloatTo64(pixelUnit float64) uint32 {
	return uint32(pixelUnit * 65536)
}

//取像素红色部分
func GetRed(c color.Color) (R uint32) {
	R, _, _, _ = c.RGBA()
	return
}

//取像素绿色通道部分
func GetGreen(c color.Color) (G uint32) {
	_, G, _, _ = c.RGBA()
	return
}

//取像素蓝色通道部分
func GetBlue(c color.Color) (B uint32) {
	_, _, B, _ = c.RGBA()
	return
}

//取像素透明通道部分
func GetAlpha(c color.Color) (A uint32) {
	_, _, _, A = c.RGBA()
	return
}
