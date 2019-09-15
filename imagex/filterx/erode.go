// 腐蚀
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
)

// 腐蚀二值图
func ErodeCVT(srcImg *image.Gray, dstImg *image.Gray, directions imagex.PixelDirection) error {
	return ErodeGray(srcImg, dstImg, directions, 0, 255)
}

// 腐蚀灰度图像
// 保留阈值keepThreshold,
// 		<keepThreshold的像素点不会被腐蚀,keepThreshold=0时腐蚀全部
// 腐蚀阈值erodeThreshold,
// 		像素包围有>erodeThreshold像素的点才会被腐蚀,erodeThreshold=255时只腐蚀边界
func ErodeGray(srcImg *image.Gray, dstImg *image.Gray, directions imagex.PixelDirection, keepThreshold uint8, erodeThreshold uint8) error {
	if nil == srcImg || nil == dstImg {
		return errors.New("SrcImg or dstImg is nil. ")
	}
	const blackPix = 0
	size := srcImg.Rect.Size()
	dirs := imagex.GetPixelDirectionAdds(imagex.ReverseDirection(directions)) //腐蚀字面意思以浅色作为参考对象，因此要反向
	//临时白图
	tempImg := imagex.NewGray(srcImg.Bounds(), blackPix)
	//生成变更记录图
	nextX, nextY := 0, 0
	var current, next color.Gray
	currentGray, nextGray := uint8(0), uint8(0)
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x < size.X-1; x++ {
			current = srcImg.GrayAt(x, y)
			dstImg.SetGray(x, y, current)
			currentGray = current.Y
			if currentGray < keepThreshold { //比阈值深
				continue
			}
			for _, dir := range dirs {
				nextX, nextY = x+dir.X, y+dir.Y
				next = srcImg.GrayAt(nextX, nextY)
				nextGray = next.Y
				if nextGray > erodeThreshold || nextGray < currentGray { //下一个方位:比腐蚀阈值浅 || 比当前色深
					continue
				}
				if nextGray > tempImg.GrayAt(nextX, nextY).Y { //下一个方位:比记录色浅
					tempImg.SetGray(nextX, nextY, next)
				}
			}
		}
	}
	if srcImg != dstImg { //来源与目标不致
		copy(dstImg.Pix, srcImg.Pix)
	}
	for index := 0; index < len(tempImg.Pix); index++ {
		if tempImg.Pix[index] > blackPix {
			dstImg.Pix[index] = tempImg.Pix[index]
		}
	}
	return nil
}
