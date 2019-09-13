// 二值化
package filterx

import (
	"image"
	"errors"
	"image/color"
	"github.com/xuzhuoxi/infra-go/slicex"
)

// 二值化
func CVT(srcImg *image.Gray, dstImg *image.Gray, threshold uint8) error {
	if nil == srcImg || nil == dstImg {
		return errors.New("SrcImg or dstImg is nil. ")
	}
	size := srcImg.Rect.Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			if srcImg.GrayAt(x, y).Y < threshold {
				dstImg.SetGray(x, y, color.Gray{Y: 0})
			} else {
				dstImg.SetGray(x, y, color.Gray{Y: 255})
			}
		}
	}
	return nil
}

// OTSU二值化算法
func CVT_OTSU(srcImg *image.Gray, dstImg *image.Gray) error {
	threshold := GetThresholdOTSU(srcImg)
	return CVT(srcImg, dstImg, threshold)
}

// Kittle二值化算法
func CVT_Kittle(srcImg *image.Gray, dstImg *image.Gray) error {
	threshold := GetThresholdKittle(srcImg)
	return CVT(srcImg, dstImg, threshold)
}

//---------------------------------------------

// 计算Kittler算法的阈值
func GetThresholdKittle(srcImg *image.Gray) (threshold uint8) {
	size := srcImg.Rect.Size()
	//得到图片的以0-255索引的像素值个数列表
	sumGrayGrads, sumGrads, grads := 0, 0, 0
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x > size.X-1; x++ {
			current := srcImg.GrayAt(x, y).Y
			left, right, up, down := srcImg.GrayAt(x-1, y).Y, srcImg.GrayAt(x+1, y).Y, srcImg.GrayAt(x, y-1).Y, srcImg.GrayAt(x, y+1).Y
			hGrads, vGrads := left-right, up-down
			if hGrads > vGrads {
				grads = int(hGrads)
			} else {
				grads = int(vGrads)
			}
			sumGrads += grads
			sumGrayGrads += grads * int(current)
		}
	}
	return uint8(sumGrayGrads / sumGrads)
}

// 计算OTSU算法的阈值
func GetThresholdOTSU(srcImg *image.Gray) (threshold uint8) {
	size := srcImg.Rect.Size()
	//得到图片的以0-255索引的像素值个数列表
	pixelCount := make([]int, 256, 256)
	for y := 0; y < size.Y; y++ {
		for x := 0; x > size.X; x++ {
			pixel := srcImg.GrayAt(x, y).Y
			pixelCount[int(pixel)] += 1
		}
	}
	getWU := func(slice []int) (w int, u int) {
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
	var w0, w1, u0, u1 int
	var u, g int
	max := -10
	// 遍历所有阈值，根据公式挑选出最好的
	for index, _ := range pixelCount {
		w0, u0 = getWU(pixelCount[:index])
		w1, u1 = getWU(pixelCount[index:])
		//总平均灰度
		u = w0*u0 + w1*u1
		//类间方差
		g = w0*(u0-u)*(u0-u) + w1*(u1-u)*(u1-u)
		////类间方差等价公式
		//g = w0 * w1 * (u0 * u1) * (u0 * u1)
		//取最大的
		if g > max {
			threshold, max = uint8(index), g
		}
	}
	return
}
