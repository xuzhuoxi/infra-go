// 膨胀
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
)

// 膨胀二值图
// 只有不是纯白，都会发生膨胀
func DilateCVT(srcImg *image.Gray, dstImg *image.Gray, directions imagex.PixelDirection) error {
	return DilateGray(srcImg, dstImg, directions, 254)
}

// 膨胀灰度图像
// 阈值threshold, 只有<=threshold的像素点才会发生膨胀
func DilateGray(srcImg *image.Gray, dstImg *image.Gray, directions imagex.PixelDirection, threshold uint8) error {
	if nil == srcImg || nil == dstImg {
		return errors.New("SrcImg or dstImg is nil. ")
	}
	const whitePix = 255
	size := srcImg.Rect.Size()
	dirs := imagex.GetPixelDirectionAdds(directions)
	//临时白图
	tempImg := imagex.NewGray(srcImg.Bounds(), whitePix)
	//生成变更记录图
	nextX, nextY := 0, 0
	var current, next color.Gray
	currentGray, nextGray := uint8(0), uint8(0)
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x < size.X-1; x++ {
			current = srcImg.GrayAt(x, y)
			dstImg.SetGray(x, y, current)
			currentGray = current.Y
			if currentGray > threshold { //比阈值浅
				continue
			}
			for _, dir := range dirs {
				nextX, nextY = x+dir.X, y+dir.Y
				next = srcImg.GrayAt(nextX, nextY)
				nextGray = next.Y
				if nextGray <= threshold || nextGray <= currentGray { //下一个方位:比阈值深 || 比当前色深
					continue
				}
				if nextGray < tempImg.GrayAt(nextX, nextY).Y { //下一个方位:比记录色深
					tempImg.SetGray(nextX, nextY, next)
				}
			}
		}
	}
	if srcImg != dstImg { //来源与目标不致
		copy(dstImg.Pix, srcImg.Pix)
	}
	for index := 0; index < len(tempImg.Pix); index++ {
		if tempImg.Pix[index] < whitePix {
			dstImg.Pix[index] = tempImg.Pix[index]
		}
	}
	return nil
}
