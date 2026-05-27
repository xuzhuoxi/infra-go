package imagex

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/xuzhuoxi/infra-go/graphicx/blendx"
	"github.com/xuzhuoxi/infra-go/mathx"
)

// SprintImage formats an image as a string of pixel color values, one row per line.
// SprintImage 将图像格式化为像素颜色值字符串，每行对应一行像素。
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

// NewGray creates a Gray image filled with grayY over rect.
// NewGray 创建指定矩形区域、以 grayY 填充的 Gray 图像。
func NewGray(rect image.Rectangle, grayY uint8) *image.Gray {
	rs := image.NewGray(rect)
	FillImage(rs, &color.Gray{Y: grayY})
	return rs
}

// NewGray16 creates a Gray16 image filled with grayY over rect.
// NewGray16 创建指定矩形区域、以 grayY 填充的 Gray16 图像。
func NewGray16(rect image.Rectangle, grayY uint16) *image.Gray16 {
	rs := image.NewGray16(rect)
	FillImage(rs, &color.Gray16{Y: grayY})
	return rs
}

// NewRGBA creates an RGBA image over rect, optionally filled with defaultColor.
// NewRGBA 创建指定矩形区域的 RGBA 图像，可选 defaultColor 填充。
func NewRGBA(rect image.Rectangle, defaultColor color.Color) *image.RGBA {
	rs := image.NewRGBA(rect)
	if nil != defaultColor {
		FillImage(rs, defaultColor)
	}
	return rs
}

// NewNRGBA creates an NRGBA image over rect, optionally filled with defaultColor.
// NewNRGBA 创建指定矩形区域的 NRGBA 图像，可选 defaultColor 填充。
func NewNRGBA(rect image.Rectangle, defaultColor color.Color) *image.NRGBA {
	rs := image.NewNRGBA(rect)
	if nil != defaultColor {
		FillImage(rs, defaultColor)
	}
	return rs
}

// NewRGBA64 creates an RGBA64 image over rect, optionally filled with cdefaultColor.
// NewRGBA64 创建指定矩形区域的 RGBA64 图像，可选 cdefaultColor 填充。
func NewRGBA64(rect image.Rectangle, cdefaultColor color.Color) *image.RGBA64 {
	rs := image.NewRGBA64(rect)
	if nil != cdefaultColor {
		FillImage(rs, cdefaultColor)
	}
	return rs
}

// NewNRGBA64 creates an NRGBA64 image over rect, optionally filled with defaultColor.
// NewNRGBA64 创建指定矩形区域的 NRGBA64 图像，可选 defaultColor 填充。
func NewNRGBA64(rect image.Rectangle, defaultColor color.Color) *image.NRGBA64 {
	rs := image.NewNRGBA64(rect)
	if nil != defaultColor {
		FillImage(rs, defaultColor)
	}
	return rs
}

// FillImage fills every pixel of img with color.
// FillImage 使用 color 填充 img 的全部像素。
func FillImage(img draw.Image, color color.Color) {
	rect := img.Bounds()
	setColor := img.ColorModel().Convert(color)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.Set(x, y, setColor)
		}
	}
}

// FillImageAt fills the intersection of img and rect with color.
// FillImageAt 使用 color 填充 img 与 rect 交集区域。
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

// BlendSourceNormal composites a solid background color onto destination pixels using normal blend.
// BlendSourceNormal 使用正常混合模式将纯色背景合成到目标图像像素上。
// Background alpha is ignored.
// 背景色的透明通道会被忽略。
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

// ForEachLoc invokes eachFunc for every pixel coordinate in img.
// ForEachLoc 遍历 img 的每个像素坐标并调用 eachFunc。
func ForEachLoc(img image.Image, eachFunc func(x, y int)) {
	min, max := img.Bounds().Min, img.Bounds().Max
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			eachFunc(x, y)
		}
	}
}
