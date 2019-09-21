package filterx

import (
	"github.com/xuzhuoxi/infra-go/errorsx"
	"github.com/xuzhuoxi/infra-go/slicex"
	"image"
	"image/draw"
)

func CVTWithOTSU(srcImg image.Image, dstImg draw.Image) error {
	switch img := srcImg.(type) {
	case *image.Gray:
		return CVTGrayWithOTSU(img, dstImg)
	case *image.Gray16:
		return CVTGray16WithOTSU(img, dstImg)
	default:
		return CVTGray(img, dstImg, 42767)
	}
	return errorsx.NoCaseCatchError("CVTWithOTSU")
}

// OTSU二值化算法(大津法)
func CVTGrayWithOTSU(graySrcImg *image.Gray, grayDstImg draw.Image) error {
	threshold := GetOtsuThresholdAtGray(graySrcImg)
	return CVTGray(graySrcImg, grayDstImg, uint32(threshold))
}

// OTSU二值化算法(大津法)
func CVTGray16WithOTSU(graySrcImg *image.Gray16, grayDstImg draw.Image) error {
	threshold := GetOtsuThresholdAtGray16(graySrcImg)
	return CVTGray(graySrcImg, grayDstImg, uint32(threshold))
}

// 计算32位灰度图的OTSU算法的阈值(效率很低)
// 64位图像阈值
// threshold [0， 65535]
func GetOtsuThresholdAtGray16(grayImg *image.Gray16) (threshold uint32) {
	min := grayImg.Rect.Min
	max := grayImg.Rect.Max
	//得到图片的以0-65535索引的像素值个数列表
	pixelCount := make([]int, 65536, 65536)
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			pixel := grayImg.Gray16At(x, y).Y
			pixelCount[int(pixel)] += 1
		}
	}
	threshold = uint32(getOtsuThreshold(pixelCount))
	return
}

// 计算32位灰度图的OTSU算法的阈值
// 64位图像阈值
// threshold [0， 65535]
func GetOtsuThresholdAtGray(grayImg *image.Gray) (threshold uint32) {
	min := grayImg.Rect.Min
	max := grayImg.Rect.Max
	//得到图片的以0-255索引的像素值个数列表
	pixelCount := make([]int, 256, 256)
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			pixel := grayImg.GrayAt(x, y).Y
			pixelCount[int(pixel)] += 1
		}
	}
	threshold = uint32(getOtsuThreshold(pixelCount))
	return
}

//--------------------------------------------

func getOtsuThreshold(pixelCount []int) (threshold uint) {
	var w0, w1, u0, u1 int
	var u, g int
	max := -10
	// 遍历所有阈值，根据公式挑选出最好的
	for index, _ := range pixelCount {
		w0, u0 = getOtsuWU(pixelCount[:index])
		w1, u1 = getOtsuWU(pixelCount[index:])
		//总平均灰度
		u = w0*u0 + w1*u1
		//类间方差
		g = w0*(u0-u)*(u0-u) + w1*(u1-u)*(u1-u)
		////类间方差等价公式
		//g = w0 * w1 * (u0 * u1) * (u0 * u1)
		//取最大的
		if g > max {
			threshold, max = uint(index), g
		}
	}
	return
}

func getOtsuWU(slice []int) (w int, u int) {
	w = slicex.SumInt(slice) //得到阈值以下像素个数
	if w == 0 {
		u = 0
	} else {
		sum := 0
		for threshold, count := range slice {
			sum += threshold * count
		}
		u = sum / w
	}
	return
}
