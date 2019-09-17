package filterx

import (
	"github.com/xuzhuoxi/infra-go/graphicx"
	"image"
	"image/color"
)

// 转制图像为灰度图
func GrayImageDefault(src image.Image) (grayImg *image.Gray) {
	return GrayImage(src, graphicx.DefaultAlgMode)
}

// 转制图像为灰度图
func GrayImage(src image.Image, algMode graphicx.GrayAlgMode) (grayImg *image.Gray) {
	grayImg = image.NewGray(src.Bounds())
	size := src.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			gray := graphicx.GrayRGB(uint8(r), uint8(g), uint8(b), algMode)
			grayImg.SetGray(x, y, color.Gray{Y: gray})
		}
	}
	return
}
