package imagex

import (
	"bytes"
	"fmt"
	"github.com/xuzhuoxi/infra-go/graphicx/blendx"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
	"math"
)

//图像字符串化
func SprintImage(img image.Image) string {
	bs := bytes.NewBufferString("")
	rect := img.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; y < rect.Max.X; x++ {
			bs.WriteString(fmt.Sprint(img.At(x, y)))
		}
		bs.WriteString("\n")
	}
	return bs.String()
}

// 新建灰度图像
func NewGray(rect image.Rectangle, grayY uint8) *image.Gray {
	rs := image.NewGray(rect)
	FillImage(rs, &color.Gray{Y: grayY})
	return rs
}

// 新建灰度图像
func NewGray16(rect image.Rectangle, grayY uint16) *image.Gray16 {
	rs := image.NewGray16(rect)
	FillImage(rs, &color.Gray16{Y: grayY})
	return rs
}

// 新建RGBA图像
func NewRGBA(rect image.Rectangle, defaultColor color.Color) *image.RGBA {
	rs := image.NewRGBA(rect)
	if nil != defaultColor {
		FillImage(rs, defaultColor)
	}
	return rs
}

// 新建NRGBA图像
func NewNRGBA(rect image.Rectangle, defaultColor color.Color) *image.NRGBA {
	rs := image.NewNRGBA(rect)
	if nil != defaultColor {
		FillImage(rs, defaultColor)
	}
	return rs
}

// 新建RGBA64图像
func NewRGBA64(rect image.Rectangle, cdefaultColor color.Color) *image.RGBA64 {
	rs := image.NewRGBA64(rect)
	if nil != cdefaultColor {
		FillImage(rs, cdefaultColor)
	}
	return rs
}

// 新建RGBA64图像
func NewNRGBA64(rect image.Rectangle, defaultColor color.Color) *image.NRGBA64 {
	rs := image.NewNRGBA64(rect)
	if nil != defaultColor {
		FillImage(rs, defaultColor)
	}
	return rs
}

//使用颜色填充图像
func FillImage(img draw.Image, color color.Color) {
	rect := img.Bounds()
	setColor := img.ColorModel().Convert(color)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.Set(x, y, setColor)
		}
	}
}

//使用颜色填充图像部分区域
func FillImageAt(img draw.Image, color color.Color, rect image.Rectangle) {
	rect2 := img.Bounds()
	minX := mathx.MaxInt(rect.Min.X, rect2.Min.X)
	minY := mathx.MaxInt(rect.Min.Y, rect2.Min.Y)
	maxX := mathx.MinInt(rect.Max.X, rect2.Max.X)
	maxY := mathx.MinInt(rect.Max.Y, rect2.Max.Y)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			img.Set(x, y, color)
		}
	}
}

// 增加背景色
// 背景色的透明通道会被忽略
// R = S*(1-Da) + D*Da [0,1]
func BlendSourceNormal(destinationImg draw.Image, sourceColor color.Color) {
	rect := destinationImg.Bounds()
	Sr, Sg, Sb, _ := sourceColor.RGBA()
	setColor := &color.RGBA64{A: 65535}
	var Dr, Dg, Db, Da uint32
	var R, G, B uint32
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			Dr, Dg, Db, Da = destinationImg.At(x, y).RGBA()
			if math.MaxUint16 == Da { //前景不透明
				continue
			}
			if 0 == Da { //前景全透明
				destinationImg.Set(x, y, sourceColor)
				continue
			}
			R, G, B, _ = blendx.BlendNormalRGBA(Sr, Sg, Sb, 0, Dr, Dg, Db, Da, 0, false)
			setColor.R, setColor.G, setColor.B = uint16(R), uint16(G), uint16(B)
			destinationImg.Set(x, y, setColor)
		}
	}
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

func ForEachLoc(img image.Image, eachFunc func(x, y int)) {
	min, max := img.Bounds().Min, img.Bounds().Max
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			eachFunc(x, y)
		}
	}
}
