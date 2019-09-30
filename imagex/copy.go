package imagex

import (
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/draw"
)

//复制图像
func CopyImage(srcImg image.Image) draw.Image {
	copyImg := CopyImageStruct(srcImg)
	if nil == copyImg {
		return nil
	}
	CopyImageTo(srcImg, copyImg)
	return copyImg
}

//复制图像
func XCopyImage(srcImg image.Image) draw.Image {
	copyImg := CopyImageStruct(srcImg)
	if nil == copyImg {
		return nil
	}
	XCopyImageTo(srcImg, copyImg)
	return copyImg
}

//把源图像复制到目标图像
func CopyImageTo(srcImg image.Image, dstImg draw.Image) {
	srcRect := srcImg.Bounds()
	dstRect := dstImg.Bounds()
	minX := mathx.MaxInt(dstRect.Min.X, srcRect.Min.X)
	minY := mathx.MaxInt(dstRect.Min.Y, srcRect.Min.Y)
	maxX := mathx.MinInt(dstRect.Max.X, srcRect.Max.X)
	maxY := mathx.MinInt(dstRect.Max.Y, srcRect.Max.Y)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			dstImg.Set(x, y, srcImg.At(x, y))
		}
	}
}

//把源图像复制到目标图像
func XCopyImageTo(srcImg image.Image, dstImg draw.Image) {
	draw.Draw(dstImg, srcImg.Bounds(), srcImg, image.Pt(0, 0), draw.Over)
}

//复制图像
func CopyImageStruct(srcImg image.Image) draw.Image {
	rect := srcImg.Bounds()
	var clone draw.Image
	switch t := srcImg.(type) {
	case *image.Alpha:
		clone = image.NewAlpha(rect)
	case *image.CMYK:
		clone = image.NewCMYK(rect)
	case *image.Gray:
		clone = image.NewGray(rect)
	case *image.Gray16:
		clone = image.NewGray16(rect)
	case *image.NRGBA:
		clone = image.NewNRGBA(rect)
	case *image.NRGBA64:
		clone = image.NewNRGBA64(rect)
	case *image.Paletted:
		clone = image.NewPaletted(rect, t.Palette)
	case *image.RGBA:
		clone = image.NewRGBA(rect)
	case *image.RGBA64:
		clone = image.NewRGBA64(rect)
	case *image.YCbCr:
		clone = image.NewRGBA(rect)
	case *image.NYCbCrA:
		clone = image.NewNRGBA(rect)
	default:
		return nil
	}
	return clone
}
