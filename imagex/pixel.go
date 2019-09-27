package imagex

import (
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

func (i *PixelImage) ForEachPixel(eachFunc func(x, y int, pixel uint32)) {
	var x, y int
	for y = 0; y < i.Height; y++ {
		for x = 0; x < i.Width; x++ {
			eachFunc(x, y, i.Pix[i.getIndex(x, y)])
		}
	}
}

func (i *PixelImage) getIndex(x, y int) int {
	return 4 * (y*i.Width + x)
}

func (i *PixelImage) to(x, y int) int {
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

func (i *PixelImage64) ForEachPixel(eachFunc func(x, y int, pixel uint64)) {
	var x, y int
	for y = 0; y < i.Height; y++ {
		for x = 0; x < i.Width; x++ {
			eachFunc(x, y, i.Pix[i.getIndex(x, y)])
		}
	}
}

func (i *PixelImage64) getIndex(x, y int) int {
	return y*i.Width + x
}

//----------------------------------------------

func NewPixelImage(width, height int, pixel uint32) *PixelImage {
	rs := &PixelImage{Width: width, Height: height, Pix: make([]uint32, width*height)}
	if pixel > 0 {
		ln := width * height
		for index := 0; index < ln; index++ {
			rs.Pix[index] = pixel
		}
	}
	return rs
}

func NewPixelImage64(width, height int, pixel uint64) *PixelImage64 {
	rs := &PixelImage64{Width: width, Height: height, Pix: make([]uint64, width*height)}
	if pixel > 0 {
		ln := width * height
		for index := 0; index < ln; index++ {
			rs.Pix[index] = pixel
		}
	}
	return rs
}

func CopyColor2Pixel(src image.Image, dst *PixelImage) {
	size := src.Bounds().Size()
	var a, r, g, b uint32
	var pixel uint32
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			a, r, g, b = src.At(x, y).RGBA()
			pixel = (a << 16) | (r << 8) | g | b>>8
			dst.Set(x, y, pixel)
		}
	}
}

func CopyColor2Pixel64(src image.Image, dst *PixelImage64) {
	size := src.Bounds().Size()
	var a, r, g, b uint32
	var pixel uint64
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			a, r, g, b = src.At(x, y).RGBA()
			pixel = (uint64(a) << 48) | (uint64(r) << 32) | (uint64(g) << 16) | uint64(b)
			dst.Set(x, y, pixel)
		}
	}
}

func CopyPixel2Color(src *PixelImage, dst draw.Image) {
	var a, r, g, b uint32
	var pixel uint32
	var setColor = &color.RGBA{}
	for y := 0; y < src.Height; y++ {
		for x := 0; x < src.Width; x++ {
			pixel = src.At(x, y)
			a, r, g, b = (pixel&0xff000000)>>24, (pixel&0x00ff0000)>>16, (pixel&0x0000ff00)>>8, pixel&0x000000ff
			setColor.A, setColor.R, setColor.G, setColor.B = uint8(a), uint8(r), uint8(g), uint8(b)
			dst.Set(x, y, setColor)
		}
	}
}

func CopyPixel642Color(src *PixelImage64, dst draw.Image) {
	var a, r, g, b uint64
	var pixel uint64
	var setColor = &color.RGBA64{}
	for y := 0; y < src.Height; y++ {
		for x := 0; x < src.Width; x++ {
			pixel = src.At(x, y)
			a, r, g, b = (pixel&0xff000000)>>48, (pixel&0x00ff0000)>>32, (pixel&0x0000ff00)>>16, pixel&0x000000ff
			setColor.A, setColor.R, setColor.G, setColor.B = uint16(a), uint16(r), uint16(g), uint16(b)
			dst.Set(x, y, setColor)
		}
	}
}
