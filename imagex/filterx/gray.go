package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/graphicx"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
)

// 转制NRGBA图像为Gray16灰度图
// 使用默认灰度计算算法
func GrayDefaultFromNRGBA(src image.Image) (resultImage image.Image, err error) {
	return GrayFromNRGBA(src, graphicx.DefaultAlgMode)
}

// 转制RGBA图像为Gray16灰度图
// 使用默认灰度计算算法
func GrayDefaultFromRGBA(src image.Image, backgroundColor color.Color) (resultImage image.Image, err error) {
	return GrayFromRGBA(src, backgroundColor, graphicx.DefaultAlgMode)
}

// 转制NRGBA图像为Gray16灰度图
func GrayFromNRGBA(rgbImg image.Image, algMode graphicx.GrayAlgMode) (resultImage image.Image, err error) {
	if nil == rgbImg {
		return nil, errors.New("RgbImg is nil! ")
	}
	grayImg := image.NewGray16(rgbImg.Bounds())
	size := rgbImg.Bounds().Size()
	setColor := &color.Gray16{}
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			r, g, b, _ := rgbImg.At(x, y).RGBA()
			gray := graphicx.GrayRGB(r, g, b, algMode)
			setColor.Y = uint16(gray)
			grayImg.SetGray16(x, y, *setColor)
		}
	}
	return grayImg, nil
}

// 转制RGBA图像为Gray16灰度图
func GrayFromRGBA(rgbaImg image.Image, background color.Color, algMode graphicx.GrayAlgMode) (resultImage image.Image, err error) {
	if nil == rgbaImg {
		return nil, errors.New("RgbaImg is nil! ")
	}
	rgbaRect := rgbaImg.Bounds()
	copyImg := imagex.NewNRGBA64(rgbaRect, nil)
	NrgbaAt(rgbaImg, copyImg, background)
	return GrayFromNRGBA(copyImg, algMode)
}
