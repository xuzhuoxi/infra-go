package filterx

import (
	"image"
	"image/draw"

	"github.com/xuzhuoxi/infra-go/graphicx"
	"github.com/xuzhuoxi/infra-go/mathx"
)

// CVTGrayImageWithKittle binarizes an image using Kittler's method (gray stretch).
// CVTGrayImageWithKittle 使用 Kittler 算法（灰度拉伸）对图像进行二值化。
func CVTGrayImageWithKittle(srcImg image.Image, dstImg draw.Image) error {
	threshold := GetKittleThresholdAtGreen(srcImg)
	return CVTGray(srcImg, dstImg, uint32(threshold))
}

//---------------------------------------------

// GetKittleThresholdAtGreen computes the Kittler threshold from the green channel.
// GetKittleThresholdAtGreen 计算 Kittler 算法的阈值。
// Returns a 16-bit image-level threshold in [0, 65535].
// 结果返回 64 位图像级阈值 [0，65535]。
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
