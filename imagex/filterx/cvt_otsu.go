package filterx

import (
	"github.com/xuzhuoxi/infra-go/graphicx"
	"github.com/xuzhuoxi/infra-go/slicex"
	"image"
	"image/draw"
)

// CVTWithOTSU
// OTSU二值化算法(大津法)
func CVTWithOTSU(srcImg image.Image, dstImg draw.Image) error {
	threshold := GetOtsuThreshold(srcImg)
	return CVTGray(srcImg, dstImg, threshold)
}

// CVTWithOTSU64
// OTSU二值化算法(大津法)
func CVTWithOTSU64(srcImg image.Image, dstImg draw.Image) error {
	threshold := GetOtsuThreshold64(srcImg)
	return CVTGray(srcImg, dstImg, threshold)
}

// GetOtsuThreshold
// 计算灰度图的OTSU算法的阈值
// 64位像素强制转为32位像素参与计算
// 结果返回64位图像级阈值
// threshold [0， 65535]
func GetOtsuThreshold(grayImg image.Image) (threshold uint32) {
	min := grayImg.Bounds().Min
	max := grayImg.Bounds().Max
	//得到图片的以0-255索引的像素值个数列表
	pixelCount := make([]int, 256, 256)
	var pixel uint32
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			pixel = graphicx.GetGreen(grayImg.At(x, y)) >> 8
			pixelCount[int(pixel)] += 1
		}
	}
	threshold = uint32(getOtsuThreshold(pixelCount)) << 8
	return
}

// GetOtsuThreshold64
// 计算灰度图的OTSU算法的阈值
// 结果返回64位图像级阈值
// threshold [0， 65535]
func GetOtsuThreshold64(grayImg image.Image) (threshold uint32) {
	min := grayImg.Bounds().Min
	max := grayImg.Bounds().Max
	//得到图片的以0-65535索引的像素值个数列表
	pixelCount := make([]int, 65536, 65536)
	var pixel uint32
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			pixel = graphicx.GetGreen(grayImg.At(x, y))
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
	for index := range pixelCount {
		w0, u0 = getOtsuWU(pixelCount[:index:index])
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
