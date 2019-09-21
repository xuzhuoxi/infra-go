package filterx

import (
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/draw"
)

// Kittle二值化算法(灰度拉伸)
func CVTGrayWithKittle(srcImg *image.Gray, dstImg draw.Image) error {
	threshold := GetKittleThresholdAtGray(srcImg)
	return CVTGray(srcImg, dstImg, uint32(threshold)+1)
}

// Kittle二值化算法(灰度拉伸)
func CVTGray16WithKittle(srcImg *image.Gray16, dstImg draw.Image) error {
	threshold := GetKittleThresholdAtGray16(srcImg)
	return CVTGray(srcImg, dstImg, uint32(threshold))
}

//---------------------------------------------

// 计算Kittler算法的阈值
// 64位图像阈值
// threshold [0， 65535]
func GetKittleThresholdAtGray(grayImg *image.Gray) (threshold uint32) {
	min := grayImg.Rect.Min
	max := grayImg.Rect.Max
	sumGrayGrads, sumGrads, grads := 0, 0, 0
	for y := min.Y + 1; y < max.Y-1; y++ {
		for x := min.X + 1; x < max.X-1; x++ {
			current := grayImg.GrayAt(x, y).Y
			left, right, up, down := grayImg.GrayAt(x-1, y).Y, grayImg.GrayAt(x+1, y).Y, grayImg.GrayAt(x, y-1).Y, grayImg.GrayAt(x, y+1).Y
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
	return uint32(sumGrayGrads/sumGrads) * 257
}

// 计算Kittler算法的阈值
// 64位图像阈值
// threshold [0， 65535]
func GetKittleThresholdAtGray16(grayImg *image.Gray16) (threshold uint32) {
	min := grayImg.Rect.Min
	max := grayImg.Rect.Max
	var sumGrayGrads, sumGrads, grads int
	var left, right, up, down, current int
	for y := min.Y + 1; y < max.Y-1; y++ {
		for x := min.X + 1; x < max.X-1; x++ {
			current = int(grayImg.Gray16At(x, y).Y)
			left, right, up, down = int(grayImg.Gray16At(x-1, y).Y), int(grayImg.Gray16At(x+1, y).Y), int(grayImg.Gray16At(x, y-1).Y), int(grayImg.Gray16At(x, y+1).Y)
			hGrads, vGrads := mathx.AbsInt(left-right), mathx.AbsInt(up-down)
			grads = mathx.MaxInt(hGrads, vGrads)
			sumGrads += grads
			sumGrayGrads += grads * current
		}
	}
	return uint32(sumGrayGrads / sumGrads)
}
