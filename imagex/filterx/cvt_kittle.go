package filterx

import (
	"github.com/xuzhuoxi/infra-go/graphicx"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/draw"
)

// CVTGrayImageWithKittle
// Kittle二值化算法(灰度拉伸)
func CVTGrayImageWithKittle(srcImg image.Image, dstImg draw.Image) error {
	threshold := GetKittleThresholdAtGreen(srcImg)
	return CVTGray(srcImg, dstImg, uint32(threshold))
}

//---------------------------------------------

// GetKittleThresholdAtGreen
// 计算Kittler算法的阈值
// 64位图像阈值
// threshold [0， 65535]
func GetKittleThresholdAtGreen(grayImg image.Image) (threshold uint32) {
	min := grayImg.Bounds().Min
	max := grayImg.Bounds().Max
	var sumGrayGrads, sumGrads, grads int
	var left, right, up, down, current int
	for y := min.Y + 1; y < max.Y-1; y++ {
		for x := min.X + 1; x < max.X-1; x++ {
			current = int(graphicx.GetGreen(grayImg.At(x, y)))
			left, right, up, down = int(graphicx.GetGreen(grayImg.At(x-1, y))), int(graphicx.GetGreen(grayImg.At(x+1, y))),
				int(graphicx.GetGreen(grayImg.At(x, y-1))), int(graphicx.GetGreen(grayImg.At(x, y+1)))
			hGrads, vGrads := mathx.AbsInt(left-right), mathx.AbsInt(up-down)
			grads = mathx.MaxInt(hGrads, vGrads)
			sumGrads += grads
			sumGrayGrads += grads * current
		}
	}
	return uint32(sumGrayGrads / sumGrads)
}
