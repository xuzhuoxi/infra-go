// 膨胀
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
	"image/draw"
)

// 膨胀二值图
// 只要不是纯白，都会发生膨胀
// directions为允许膨胀的方向
func DilateCVT(graySrcImg image.Image, grayDstImg draw.Image, directions imagex.PixelDirection) error {
	return DilateGray(graySrcImg, grayDstImg, directions, 254)
}

// 膨胀灰度图像(用于普通灰度图意义不太，二值图的效果不错)
// 阈值threshold, 只有<=threshold的像素点才会发生膨胀
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
	changeList := imagex.NewPixelImage(size.X, size.Y, whitePix)
	//生成变更记录图
	nextX, nextY := 0, 0
	var current color.Color
	var currentGray, nextGray uint32
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x < size.X-1; x++ {
			current = graySrcImg.At(x, y)
			_, currentGray, _, _ = current.RGBA()
			grayDstImg.Set(x, y, current)
			if currentGray > threshold { //比阈值浅
				continue
			}
			for _, dir := range dirs {
				nextX, nextY = x+dir.X, y+dir.Y
				_, nextGray, _, _ = graySrcImg.At(nextX, nextY).RGBA()
				if nextGray <= threshold || nextGray <= currentGray { //下一个方位:比阈值深 || 比当前色深
					continue
				}
				if currentGray < changeList.At(nextX, nextY) { //下一个方位: 当前比记录色深
					changeList.Set(nextX, nextY, currentGray)
				}
			}
		}
	}
	//fmt.Println("changeList:", changeList)
	setColor := &color.Gray16{}
	var changePixel uint32
	if graySrcImg == grayDstImg { //来源与目标不一致
		for y := 0; y < size.Y; y++ {
			for x := 0; x < size.X; x++ {
				changePixel = changeList.At(x, y)
				if changePixel < whitePix {
					setColor.Y = uint16(changePixel)
					grayDstImg.Set(x, y, setColor)
				}
			}
		}
	} else {
		for y := 0; y < size.Y; y++ {
			for x := 0; x < size.X; x++ {
				changePixel = changeList.At(x, y)
				if changePixel < whitePix {
					setColor.Y = uint16(changePixel)
					grayDstImg.Set(x, y, setColor)
				} else {
					grayDstImg.Set(x, y, graySrcImg.At(x, y))
				}
			}
		}
	}
	return nil
}
