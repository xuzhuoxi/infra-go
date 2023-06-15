// Package filterx
// 腐蚀
package filterx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
)

// ErodeCVT 腐蚀二值图
func ErodeCVT(cvtSrcImg image.Image, cvtDstImg draw.Image, directions imagex.PixelDirection) error {
	return ErodeGray(cvtSrcImg, cvtDstImg, directions, 65535, 0)
}

// ErodeGray
// 腐蚀灰度图像
// 腐蚀阈值erodeThreshold (0,65535]
// 		>=erodeThreshold像素的点才具有腐蚀性
// 		erodeThreshold=1时只要不是纯黑都具有腐蚀性，erodeThreshold=65535时只有纯白具有腐蚀性
// 抗腐蚀阈值antiErodeThreshold [0, 65535)
// 		<antiErodeThreshold的像素点不会被腐蚀,
// 		antiErodeThreshold=0时不抗腐蚀
// 要求 erodeThreshold > antiErodeThreshold
func ErodeGray(graySrcImg image.Image, grayDstImg draw.Image, directions imagex.PixelDirection, erodeThreshold uint32, antiErodeThreshold uint32) error {
	if nil == graySrcImg || nil == grayDstImg {
		return errors.New("SrcImg or grayDstImg is nil. ")
	}
	if erodeThreshold == 0 {
		return errors.New("erodeThreshold = 0. ")
	}
	if erodeThreshold <= antiErodeThreshold {
		return errors.New(fmt.Sprint("antiErodeThreshold <= erodeThreshold: ", erodeThreshold, antiErodeThreshold))
	}
	const blackPix = 0
	size := graySrcImg.Bounds().Size()
	dirs := imagex.GetPixelDirectionAdds(directions.ReverseDirection()) //腐蚀字面意思以浅色作为参考对象，因此要反向
	//临时黑图
	changeImg := imagex.NewPixelImage(size.X, size.Y, blackPix)
	//生成变更记录图
	var current color.Color
	var currentGray, nextGray uint32
	for y := 1; y < size.Y-1; y++ {
		for x := 1; x < size.X-1; x++ {
			current = graySrcImg.At(x, y)
			grayDstImg.Set(x, y, current)
			_, currentGray, _, _ = current.RGBA()
			if currentGray < antiErodeThreshold { //比抗腐蚀阈值深，保留
				continue
			}
			if currentGray < uint32(changeImg.At(x, y)) { //比变更集深，代表已经腐蚀。
				continue
			}
			//以下为查找周围最强腐蚀点
			var erodeGray uint32
			for _, dir := range dirs {
				_, nextGray, _, _ = graySrcImg.At(x+dir.X, y+dir.Y).RGBA()
				if nextGray >= erodeThreshold { //找到腐蚀点
					erodeGray = mathx.MaxUint32(erodeGray, nextGray)
				}
			}
			//fmt.Println("查找...", x, y, currentGray, erodeGray)
			if erodeGray > 0 && erodeGray > currentGray { //比当前更浅
				changeImg.Set(x, y, erodeGray)
			}
		}
	}
	setColor := &color.Gray16{}
	theSame := graySrcImg == grayDstImg
	var changePixel uint32
	changeImg.ForEachPixel(func(x, y int) {
		changePixel = changeImg.At(x, y)
		if changePixel > blackPix {
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
