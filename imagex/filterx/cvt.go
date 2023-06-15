// Package filterx
// 二值化
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
)

// CVTGray
// 二值化灰度图
// threshold [0， 65535]
func CVTGray(graySrcImg image.Image, grayDstImg draw.Image, threshold uint32) error {
	if nil == graySrcImg || nil == grayDstImg {
		return errors.New("SrcImg or grayDstImg is nil. ")
	}
	srcRect := graySrcImg.Bounds()
	dstRect := grayDstImg.Bounds()
	minX := mathx.MaxInt(dstRect.Min.X, srcRect.Min.X)
	minY := mathx.MaxInt(dstRect.Min.Y, srcRect.Min.Y)
	maxX := mathx.MinInt(dstRect.Max.X, srcRect.Max.X)
	maxY := mathx.MinInt(dstRect.Max.Y, srcRect.Max.Y)
	grayZero := color.Gray{Y: 0}
	grayOne := color.Gray{Y: 255}
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			_, G, _, _ := graySrcImg.At(x, y).RGBA() //由RGBA表示的灰度图，A必然为65535，取G为灰度值
			if G < threshold {
				grayDstImg.Set(x, y, grayZero)
			} else {
				grayDstImg.Set(x, y, grayOne)
			}
		}
	}
	return nil
}
