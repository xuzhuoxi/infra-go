package imagex

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/xuzhuoxi/infra-go/graphicx"
)

// PixelImage is a flat uint32 pixel buffer (A,R,G,B packed) with width and height.
// PixelImage 是扁平 uint32 像素缓冲（A,R,G,B 打包），含宽高。
type PixelImage struct {
	//A,R,G,B
	Pix           []uint32
	Width, Height int
}

// Max returns Width and Height as maxX and maxY.
// Max 返回 Width、Height 作为 maxX、maxY。
func (i *PixelImage) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

// At returns the packed pixel at (x, y).
// At 返回 (x, y) 处的打包像素值。
func (i *PixelImage) At(x, y int) uint32 {
	return i.Pix[i.getIndex(x, y)]
}

// Set stores a packed pixel at (x, y).
// Set 在 (x, y) 处写入打包像素值。
func (i *PixelImage) Set(x, y int, pixel uint32) {
	i.Pix[i.getIndex(x, y)] = pixel
}

// PixelAt returns the graphicx.Pixel at (x, y).
// PixelAt 返回 (x, y) 处的 graphicx.Pixel。
func (i *PixelImage) PixelAt(x, y int) graphicx.Pixel {
	return graphicx.Pixel(i.Pix[i.getIndex(x, y)])
}

// SetPixel stores a graphicx.Pixel at (x, y).
// SetPixel 在 (x, y) 处写入 graphicx.Pixel。
func (i *PixelImage) SetPixel(x, y int, pixel graphicx.Pixel) {
	i.Pix[i.getIndex(x, y)] = uint32(pixel)
}

// IndexAt returns the packed pixel at flat index.
// IndexAt 返回扁平索引处的打包像素值。
func (i *PixelImage) IndexAt(index int) uint32 {
	return i.Pix[index]
}

// IndexSet stores a packed pixel at flat index.
// IndexSet 在扁平索引处写入打包像素值。
func (i *PixelImage) IndexSet(index int, pixel uint32) {
	i.Pix[index] = pixel
}

// PixelIndexAt returns the graphicx.Pixel at flat index.
// PixelIndexAt 返回扁平索引处的 graphicx.Pixel。
func (i *PixelImage) PixelIndexAt(index int) graphicx.Pixel {
	return graphicx.Pixel(i.Pix[index])
}

// IndexSetPixel stores a graphicx.Pixel at flat index.
// IndexSetPixel 在扁平索引处写入 graphicx.Pixel。
func (i *PixelImage) IndexSetPixel(index int, pixel graphicx.Pixel) {
	i.Pix[index] = uint32(pixel)
}

// ForEachPixel invokes eachFunc for every (x, y) coordinate.
// ForEachPixel 遍历每个 (x, y) 坐标并调用 eachFunc。
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

// PixelImage64 is a flat uint64 pixel buffer (A,R,G,B packed) with width and height.
// PixelImage64 是扁平 uint64 像素缓冲（A,R,G,B 打包），含宽高。
type PixelImage64 struct {
	//A,R,G,B
	Pix           []uint64
	Width, Height int
}

// Max returns Width and Height as maxX and maxY.
// Max 返回 Width、Height 作为 maxX、maxY。
func (i *PixelImage64) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

// At returns the packed pixel at (x, y).
// At 返回 (x, y) 处的打包像素值。
func (i *PixelImage64) At(x, y int) uint64 {
	return i.Pix[i.getIndex(x, y)]
}

// Set stores a packed pixel at (x, y).
// Set 在 (x, y) 处写入打包像素值。
func (i *PixelImage64) Set(x, y int, pixel uint64) {
	i.Pix[i.getIndex(x, y)] = pixel
}

// PixelAt returns the graphicx.Pixel64 at (x, y).
// PixelAt 返回 (x, y) 处的 graphicx.Pixel64。
func (i *PixelImage64) PixelAt(x, y int) graphicx.Pixel64 {
	return graphicx.Pixel64(i.Pix[i.getIndex(x, y)])
}

// SetPixel stores a graphicx.Pixel64 at (x, y).
// SetPixel 在 (x, y) 处写入 graphicx.Pixel64。
func (i *PixelImage64) SetPixel(x, y int, pixel graphicx.Pixel64) {
	i.Pix[i.getIndex(x, y)] = uint64(pixel)
}

// IndexAt returns the packed pixel at flat index.
// IndexAt 返回扁平索引处的打包像素值。
func (i *PixelImage64) IndexAt(index int) uint64 {
	return i.Pix[index]
}

// IndexSet stores a packed pixel at flat index.
// IndexSet 在扁平索引处写入打包像素值。
func (i *PixelImage64) IndexSet(index int, pixel uint64) {
	i.Pix[index] = pixel
}

// PixelIndexAt returns the graphicx.Pixel64 at flat index.
// PixelIndexAt 返回扁平索引处的 graphicx.Pixel64。
func (i *PixelImage64) PixelIndexAt(index int) graphicx.Pixel64 {
	return graphicx.Pixel64(i.Pix[index])
}

// IndexSetPixel stores a graphicx.Pixel64 at flat index.
// IndexSetPixel 在扁平索引处写入 graphicx.Pixel64。
func (i *PixelImage64) IndexSetPixel(index int, pixel graphicx.Pixel64) {
	i.Pix[index] = uint64(pixel)
}

// ForEachPixel invokes eachFunc for every (x, y) coordinate.
// ForEachPixel 遍历每个 (x, y) 坐标并调用 eachFunc。
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

// NewPixelImage allocates a width×height PixelImage, optionally filled with pixel.
// NewPixelImage 分配 width×height 的 PixelImage，可选 pixel 填充。
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

// NewPixelImage64 allocates a width×height PixelImage64, optionally filled with pixel.
// NewPixelImage64 分配 width×height 的 PixelImage64，可选 pixel 填充。
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

// Copy2PixelImage copies src into dst as packed uint32 pixels.
// Copy2PixelImage 将 src 复制为打包 uint32 像素写入 dst。
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

// Copy2PixelImage64 copies src into dst as packed uint64 pixels.
// Copy2PixelImage64 将 src 复制为打包 uint64 像素写入 dst。
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

// CopyPixel2Image writes src pixels into dst as color.Image values.
// CopyPixel2Image 将 src 像素写入 dst 为标准 color.Image 值。
func CopyPixel2Image(src *PixelImage, dst draw.Image) {
	var setColor = &color.RGBA{}
	for y := 0; y < src.Height; y++ {
		for x := 0; x < src.Width; x++ {
			setColor.R, setColor.G, setColor.B, setColor.A = src.PixelAt(x, y).RGBA()
			dst.Set(x, y, setColor)
		}
	}
}

// CopyPixel642Image writes src pixels into dst as color.Image values.
// CopyPixel642Image 将 src 像素写入 dst 为标准 color.Image 值。
func CopyPixel642Image(src *PixelImage64, dst draw.Image) {
	var setColor = &color.RGBA64{}
	for y := 0; y < src.Height; y++ {
		for x := 0; x < src.Width; x++ {
			setColor.R, setColor.G, setColor.B, setColor.A = src.PixelAt(x, y).RGBA()
			dst.Set(x, y, setColor)
		}
	}
}
