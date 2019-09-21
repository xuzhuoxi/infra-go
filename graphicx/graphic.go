//
//Created by xuzhuoxi
//on 2019-05-29.
//@author xuzhuoxi
//
package graphicx

// 64位像素单元转32位像素单元
func Color64To32(pixelUnit uint32) uint32 {
	return pixelUnit / 257
}

// 64位像素单元转32位像素单元
func Color64ToFloat(pixelUnit uint32) float64 {
	return float64(pixelUnit) / 65535
}

// 32位像素单元转64位像素单元
func Color32To64(pixelUnit uint32) uint32 {
	return pixelUnit * 257
}

// 32位像素单元转32位像素单元
func Color32ToFloat(pixelUnit uint32) float64 {
	return float64(pixelUnit) / 255
}

// 浮点像素单元转32位像素单元
func ColorFloatTo32(pixelUnit float64) uint32 {
	return uint32(pixelUnit * 255)
}

func ColorFloatTo64(pixelUnit float64) uint32 {
	return uint32(pixelUnit * 65535)
}
