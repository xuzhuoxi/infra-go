package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/graphicx"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
)

// GrayDefaultFromNRGBA
// 转制NRGBA图像为Gray16灰度图
// 使用默认灰度计算算法
func GrayDefaultFromNRGBA(src image.Image) (resultImage image.Image, err error) {
	return GrayFromNRGBA(src, graphicx.DefaultAlgMode)
}

// GrayDefaultFromRGBA
// 转制RGBA图像为Gray16灰度图
// 使用默认灰度计算算法
func GrayDefaultFromRGBA(src image.Image, backgroundColor color.Color) (resultImage image.Image, err error) {
	return GrayFromRGBA(src, backgroundColor, graphicx.DefaultAlgMode)
}

// GrayFromNRGBA
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

// GrayFromRGBA
// 转制RGBA图像为Gray16灰度图
func GrayFromRGBA(rgbaImg image.Image, background color.Color, algMode graphicx.GrayAlgMode) (resultImage image.Image, err error) {
	if nil == rgbaImg {
		return nil, errors.New("RgbaImg is nil! ")
	}
	rgbaRect := rgbaImg.Bounds()
	copyImg := imagex.NewNRGBA64(rgbaRect, nil)
	NRGBAAt(rgbaImg, copyImg, background)
	return GrayFromNRGBA(copyImg, algMode)
}

func GetGrayPixel(grayImg *image.Gray) (pixel [][]uint8) {
	size := grayImg.Rect.Size()
	pixel = make([][]uint8, size.Y, size.Y)
	for y := 0; y < size.Y; y++ {
		pixel[y] = grayImg.Pix[y : y+size.X]
	}
	return
}

func GetGray16Pixel(grayImg *image.Gray16) (pixel [][]uint16) {
	size := grayImg.Rect.Size()
	pixel = make([][]uint16, size.Y, size.Y)
	for y := 0; y < size.Y; y++ {
		pixel[y] = make([]uint16, size.X, size.X)
		for x := 0; x < size.X; x++ {
			pixel[y][x] = grayImg.Gray16At(x, y).Y
		}
	}
	return
}
