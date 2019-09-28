package imagex

import (
	"github.com/xuzhuoxi/infra-go/graphicx"
	"image"
	"image/color"
	"image/draw"
)

type PixelImage struct {
	//A,R,G,B
	Pix           []uint32
	Width, Height int
}

func (i *PixelImage) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

func (i *PixelImage) At(x, y int) uint32 {
	return i.Pix[i.getIndex(x, y)]
}

func (i *PixelImage) Set(x, y int, pixel uint32) {
	i.Pix[i.getIndex(x, y)] = pixel
}

func (i *PixelImage) PixelAt(x, y int) graphicx.Pixel {
	return graphicx.Pixel(i.Pix[i.getIndex(x, y)])
}

func (i *PixelImage) SetPixel(x, y int, pixel graphicx.Pixel) {
	i.Pix[i.getIndex(x, y)] = uint32(pixel)
}

func (i *PixelImage) IndexAt(index int) uint32 {
	return i.Pix[index]
}

func (i *PixelImage) IndexSet(index int, pixel uint32) {
	i.Pix[index] = pixel
}

func (i *PixelImage) PixelIndexAt(index int) graphicx.Pixel {
	return graphicx.Pixel(i.Pix[index])
}

func (i *PixelImage) IndexSetPixel(index int, pixel graphicx.Pixel) {
	i.Pix[index] = uint32(pixel)
}

func (i *PixelImage) ForEachPixel(eachFunc func(x, y int)) {
	var x, y int
	for y = 0; y < i.Height; y++ {
		for x = 0; x < i.Width; x++ {
			eachFunc(x, y)
		}
	}
}

func (i *PixelImage) getIndex(x, y int) int {
	return y*i.Width + x
}

//----------------------------------------------

type PixelImage64 struct {
	//A,R,G,B
	Pix           []uint64
	Width, Height int
}

func (i *PixelImage64) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

func (i *PixelImage64) At(x, y int) uint64 {
	return i.Pix[i.getIndex(x, y)]
}

func (i *PixelImage64) Set(x, y int, pixel uint64) {
	i.Pix[i.getIndex(x, y)] = pixel
}

func (i *PixelImage64) PixelAt(x, y int) graphicx.Pixel64 {
	return graphicx.Pixel64(i.Pix[i.getIndex(x, y)])
}

func (i *PixelImage64) SetPixel(x, y int, pixel graphicx.Pixel64) {
	i.Pix[i.getIndex(x, y)] = uint64(pixel)
}

func (i *PixelImage64) IndexAt(index int) uint64 {
	return i.Pix[index]
}

func (i *PixelImage64) IndexSet(index int, pixel uint64) {
	i.Pix[index] = pixel
}

func (i *PixelImage64) PixelIndexAt(index int) graphicx.Pixel64 {
	return graphicx.Pixel64(i.Pix[index])
}

func (i *PixelImage64) IndexSetPixel(index int, pixel graphicx.Pixel64) {
	i.Pix[index] = uint64(pixel)
}

func (i *PixelImage64) ForEachPixel(eachFunc func(x, y int)) {
	var x, y int
	for y = 0; y < i.Height; y++ {
		for x = 0; x < i.Width; x++ {
			eachFunc(x, y)
		}
	}
}

func (i *PixelImage64) getIndex(x, y int) int {
	return y*i.Width + x
}

//----------------------------------------------

func NewPixelImage(width, height int, pixel uint32) *PixelImage {
	ln := width * height
	rs := &PixelImage{Width: width, Height: height, Pix: make([]uint32, ln)}
	if pixel > 0 {
		for index := 0; index < ln; index++ {
			rs.Pix[index] = uint32(pixel)
		}
	}
	return rs
}

func NewPixelImage64(width, height int, pixel uint64) *PixelImage64 {
	ln := width * height
	rs := &PixelImage64{Width: width, Height: height, Pix: make([]uint64, ln)}
	if pixel > 0 {
		for index := 0; index < ln; index++ {
			rs.Pix[index] = uint64(pixel)
		}
	}
	return rs
}

func Copy2PixelImage(src image.Image, dst *PixelImage) {
	size := src.Bounds().Size()
	var r, g, b, a uint32
	var pixel uint32
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			r, g, b, a = src.At(x, y).RGBA()
			pixel = (a << 16) | (r << 8) | g | (b >> 8)
			dst.Set(x, y, pixel)
		}
	}
}

func Copy2PixelImage64(src image.Image, dst *PixelImage64) {
	size := src.Bounds().Size()
	var r, g, b, a uint32
	var pixel uint64
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			r, g, b, a = src.At(x, y).RGBA()
			pixel = (uint64(a) << 48) | (uint64(r) << 32) | (uint64(g) << 16) | uint64(b)
			dst.Set(x, y, pixel)
		}
	}
}

func CopyPixel2Image(src *PixelImage, dst draw.Image) {
	var setColor = &color.RGBA{}
	for y := 0; y < src.Height; y++ {
		for x := 0; x < src.Width; x++ {
			setColor.R, setColor.G, setColor.B, setColor.A = src.PixelAt(x, y).RGBA()
			dst.Set(x, y, setColor)
		}
	}
}

func CopyPixel642Image(src *PixelImage64, dst draw.Image) {
	var setColor = &color.RGBA64{}
	for y := 0; y < src.Height; y++ {
		for x := 0; x < src.Width; x++ {
			setColor.R, setColor.G, setColor.B, setColor.A = src.PixelAt(x, y).RGBA()
			dst.Set(x, y, setColor)
		}
	}
}
