package filterx

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
)

//填充颜色到整图
func FillColor(srcImg draw.Image, c color.Color) error {
	if nil == srcImg {
		return errors.New("SrcImg or dstImg is nil. ")
	}
	fillRect(srcImg, srcImg.Bounds(), c)
	return nil
}

//填充颜色到区域
func FillColorWithRect(srcImg draw.Image, rect image.Rectangle, c color.Color) error {
	if nil == srcImg {
		return errors.New("SrcImg or dstImg is nil. ")
	}
	fillRect(srcImg, rect, c)
	return nil
}

//--------------------------------------------------------

func fillRect(srcImg draw.Image, rect image.Rectangle, c color.Color) {
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			srcImg.Set(x, y, c)
		}
	}
}
