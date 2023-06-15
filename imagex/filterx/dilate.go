// Package filterx
// 膨胀
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
	"image/draw"
)

// DilateCVT
// 膨胀二值图
// 只要不是纯白，都会发生膨胀
// directions为允许膨胀的方向
func DilateCVT(cvtSrcImg image.Image, cvtDstImg draw.Image, directions imagex.PixelDirection) error {
	return DilateGray(cvtSrcImg, cvtDstImg, directions, 254)
}

// DilateGray
// 膨胀灰度图像(用于普通灰度图意义不太，二值图的效果不错)
// 阈值threshold, 只有<=threshold的像素点才会向外发生膨胀
// 阈值threshold范围[0,65535]
// directions为允许膨胀的方向
func DilateGray(graySrcImg image.Image, grayDstImg draw.Image, directions imagex.PixelDirection, threshold uint32) error {
	if nil == graySrcImg || nil == grayDstImg {
		return errors.New("SrcImg or grayDstImg is nil. ")
	}
	const whitePix = 65535
	size := graySrcImg.Bounds().Size()
	dirs := imagex.GetPixelDirectionAdds(directions)
	//临时白图
	changeImg := imagex.NewPixelImage(size.X, size.Y, whitePix)
	//生成变更记录图
	var nextX, nextY int
	var current color.Color
	var currentGray, nextGray uint32
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x < size.X-1; x++ {
			current = graySrcImg.At(x, y)
			grayDstImg.Set(x, y, current)
			_, currentGray, _, _ = current.RGBA()
			if currentGray > threshold { //比阈值浅
				continue
			}
			for _, dir := range dirs {
				nextX, nextY = x+dir.X, y+dir.Y
				_, nextGray, _, _ = graySrcImg.At(nextX, nextY).RGBA()
				if nextGray <= threshold || nextGray <= currentGray { //下一个方位:比阈值深 || 比当前色深
					continue
				}
				if currentGray < uint32(changeImg.At(nextX, nextY)) { //下一个方位: 当前比记录色深
					changeImg.Set(nextX, nextY, currentGray)
				}
			}
		}
	}
	setColor := &color.Gray16{}
	theSame := graySrcImg == grayDstImg
	var changePixel uint32
	changeImg.ForEachPixel(func(x, y int) {
		changePixel = changeImg.At(x, y)
		if changePixel < whitePix {
			setColor.Y = uint16(changePixel)
			grayDstImg.Set(x, y, setColor)
			return
		}
		if !theSame {
			grayDstImg.Set(x, y, graySrcImg.At(x, y))
		}
	})
	return nil
}
