package graphicx

import "image/color"

// Pixel is a packed 32-bit ARGB pixel value.
// Pixel 表示打包的 32 位 ARGB 像素值。
type Pixel uint32

// RGBA unpacks the pixel as R, G, B, A uint8 channels.
// RGBA 将像素解包为 R、G、B、A 四个 uint8 通道。
func (p Pixel) RGBA() (R, G, B, A uint8) {
	A, R, G, B = Pixel2ARGB(uint32(p))
	return
}

// Pixel64 is a packed 64-bit ARGB pixel value.
// Pixel64 表示打包的 64 位 ARGB 像素值。
type Pixel64 uint64

// RGBA unpacks the pixel as R, G, B, A uint16 channels.
// RGBA 将像素解包为 R、G、B、A 四个 uint16 通道。
func (p Pixel64) RGBA() (R, G, B, A uint16) {
	A, R, G, B = Pixel2ARGB64(uint64(p))
	return
}

//-----------------------------------------

// ARGB2Pixel packs A, R, G, B uint8 channels into a 32-bit ARGB pixel.
// ARGB2Pixel 将 A、R、G、B 四个 uint8 通道打包为 32 位 ARGB 像素。
func ARGB2Pixel(A, R, G, B uint8) (pixel uint32) {
	pA := uint32(A) << 24
	pR := uint32(R) << 16
	pG := uint32(G) << 8
	pB := uint32(B) << 0
	return pA | pR | pG | pB
}

// Pixel2ARGB unpacks a 32-bit ARGB pixel into A, R, G, B uint8 channels.
// Pixel2ARGB 将 32 位 ARGB 像素解包为 A、R、G、B 四个 uint8 通道。
func Pixel2ARGB(pixel uint32) (A, R, G, B uint8) {
	A = uint8((pixel & 0xff000000) >> 24)
	R = uint8((pixel & 0x00ff0000) >> 16)
	G = uint8((pixel & 0x0000ff00) >> 8)
	B = uint8((pixel & 0x000000ff) >> 0)
	return
}

// RGBA2Pixel packs R, G, B, A uint8 channels into a 32-bit RGBA-ordered pixel.
// RGBA2Pixel 将 R、G、B、A 四个 uint8 通道打包为 32 位 RGBA 序像素。
func RGBA2Pixel(R, G, B, A uint8) (pixel uint32) {
	pR := uint32(R) << 24
	pG := uint32(G) << 16
	pB := uint32(B) << 8
	pA := uint32(A) << 0
	return pA | pR | pG | pB
}

// Pixel2RGBA unpacks a 32-bit RGBA-ordered pixel into R, G, B, A uint8 channels.
// Pixel2RGBA 将 32 位 RGBA 序像素解包为 R、G、B、A 四个 uint8 通道。
func Pixel2RGBA(pixel uint32) (R, G, B, A uint8) {
	R = uint8((pixel & 0xff000000) >> 24)
	G = uint8((pixel & 0x00ff0000) >> 16)
	B = uint8((pixel & 0x0000ff00) >> 8)
	A = uint8((pixel & 0x000000ff) >> 0)
	return
}

// ARGB2Pixel64 packs A, R, G, B uint16 channels into a 64-bit ARGB pixel.
// ARGB2Pixel64 将 A、R、G、B 四个 uint16 通道打包为 64 位 ARGB 像素。
func ARGB2Pixel64(A, R, G, B uint16) (pixel uint64) {
	pA := uint64(A) << 48
	pR := uint64(R) << 32
	pG := uint64(G) << 16
	pB := uint64(B) << 0
	return pA | pR | pG | pB
}

// Pixel2ARGB64 unpacks a 64-bit ARGB pixel into A, R, G, B uint16 channels.
// Pixel2ARGB64 将 64 位 ARGB 像素解包为 A、R、G、B 四个 uint16 通道。
func Pixel2ARGB64(pixel uint64) (A, R, G, B uint16) {
	A = uint16((pixel & 0xffff000000000000) >> 48)
	R = uint16((pixel & 0x0000ffff00000000) >> 32)
	G = uint16((pixel & 0x00000000ffff0000) >> 16)
	B = uint16((pixel & 0x000000000000ffff) >> 0)
	return
}

// RGB2APixel64 packs R, G, B, A uint16 channels into a 64-bit RGBA-ordered pixel.
// RGB2APixel64 将 R、G、B、A 四个 uint16 通道打包为 64 位 RGBA 序像素。
func RGB2APixel64(R, G, B, A uint16) (pixel uint64) {
	pR := uint64(R) << 48
	pG := uint64(G) << 32
	pB := uint64(B) << 16
	pA := uint64(A) << 0
	return pA | pR | pG | pB
}

// Pixel2RGBA64 unpacks a 64-bit RGBA-ordered pixel into R, G, B, A uint16 channels.
// Pixel2RGBA64 将 64 位 RGBA 序像素解包为 R、G、B、A 四个 uint16 通道。
func Pixel2RGBA64(pixel uint64) (R, G, B, A uint16) {
	R = uint16((pixel & 0xffff000000000000) >> 48)
	G = uint16((pixel & 0x0000ffff00000000) >> 32)
	B = uint16((pixel & 0x00000000ffff0000) >> 16)
	A = uint16((pixel & 0x000000000000ffff) >> 0)
	return
}

// Color2Pixel converts c to a packed 32-bit pixel using RGBA channel values.
// Color2Pixel 根据 RGBA 通道值将 c 转换为打包的 32 位像素。
func Color2Pixel(c color.Color) uint32 {
	r, g, b, a := c.RGBA()
	return (a << 16) | (r << 8) | g | (b >> 8)
}

// Color2Pixel64 converts c to a packed 64-bit pixel using RGBA channel values.
// Color2Pixel64 根据 RGBA 通道值将 c 转换为打包的 64 位像素。
func Color2Pixel64(c color.Color) uint64 {
	r, g, b, a := c.RGBA()
	return (uint64(a) << 48) | (uint64(r) << 32) | (uint64(g) << 16) | uint64(b)
}
