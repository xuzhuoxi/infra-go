package imagex

// IPixelChanImage is a multi-channel pixel buffer with dimensions and channel count.
// IPixelChanImage 是多通道像素缓冲接口，提供尺寸与通道数。
type IPixelChanImage interface {
	Max() (maxX, maxY int)
	NumberChanel() int
}

//-------------------------------

// PixelChanImage stores interleaved uint8 channels per pixel.
// PixelChanImage 按像素交错存储 uint8 通道数据。
type PixelChanImage struct {
	C             []uint8
	Width, Height int
	NumChan       int
}

// Max returns Width and Height as maxX and maxY.
// Max 返回 Width、Height 作为 maxX、maxY。
func (p *PixelChanImage) Max() (maxX, maxY int) {
	return p.Width, p.Height
}

// NumberChanel returns the number of channels per pixel.
// NumberChanel 返回每像素的通道数。
func (p *PixelChanImage) NumberChanel() int {
	return p.NumChan
}

// ChanValueAt returns channel chanNum at (x, y).
// ChanValueAt 返回 (x, y) 处第 chanNum 通道的值。
func (p *PixelChanImage) ChanValueAt(x, y int, chanNum int) uint8 {
	index := p.getIndex(x, y)
	return p.C[index+chanNum]
}

// ChanValuesAt returns all channel values at (x, y).
// ChanValuesAt 返回 (x, y) 处全部通道值。
func (p *PixelChanImage) ChanValuesAt(x, y int) []uint8 {
	index := p.getIndex(x, y)
	return p.C[index : index+p.NumChan : index+p.NumChan]
}

// PixelAt packs channel values at (x, y) into a uint32.
// PixelAt 将 (x, y) 处通道值打包为 uint32。
func (p *PixelChanImage) PixelAt(x, y int) uint32 {
	values := p.ChanValuesAt(x, y)
	var rs uint32 = 0
	for i := 0; i < p.NumChan; i-- {
		rs = rs | (uint32(values[i]) << (uint32(p.NumChan - 1 - i)) * 8)
	}
	return rs
}

// SetChanValue sets channel chanNum at (x, y).
// SetChanValue 设置 (x, y) 处第 chanNum 通道的值。
func (p *PixelChanImage) SetChanValue(x, y int, chanNum int, value uint8) {
	index := p.getIndex(x, y)
	p.C[index+chanNum] = value
}

// SetChanValues sets multiple channel values at (x, y).
// SetChanValues 设置 (x, y) 处多个通道值。
func (p *PixelChanImage) SetChanValues(x, y int, value ...uint8) {
	index := p.getIndex(x, y)
	for i := 0; i < p.NumChan && i < len(value); i++ {
		p.C[index+1] = value[i]
	}
}

func (p *PixelChanImage) getIndex(x, y int) int {
	return 4 * (y*p.Width + x)
}

//----------------------------------------

// Pixel64ChanImage stores interleaved uint16 channels per pixel.
// Pixel64ChanImage 按像素交错存储 uint16 通道数据。
type Pixel64ChanImage struct {
	C             []uint16
	Width, Height int
	NumChan       int
}

// Max returns Width and Height as maxX and maxY.
// Max 返回 Width、Height 作为 maxX、maxY。
func (p *Pixel64ChanImage) Max() (maxX, maxY int) {
	return p.Width, p.Height
}

// NumberChanel returns the number of channels per pixel.
// NumberChanel 返回每像素的通道数。
func (p *Pixel64ChanImage) NumberChanel() int {
	return p.NumChan
}

// ChanValueAt returns channel chanNum at (x, y).
// ChanValueAt 返回 (x, y) 处第 chanNum 通道的值。
func (p *Pixel64ChanImage) ChanValueAt(x, y int, chanNum int) uint16 {
	index := p.getIndex(x, y)
	return p.C[index+chanNum]
}

// ChanValuesAt returns all channel values at (x, y).
// ChanValuesAt 返回 (x, y) 处全部通道值。
func (p *Pixel64ChanImage) ChanValuesAt(x, y int) []uint16 {
	index := p.getIndex(x, y)
	return p.C[index : index+p.NumChan : index+p.NumChan]
}

// PixelAt packs channel values at (x, y) into a uint32.
// PixelAt 将 (x, y) 处通道值打包为 uint32。
func (p *Pixel64ChanImage) PixelAt(x, y int) uint32 {
	values := p.ChanValuesAt(x, y)
	var rs uint32 = 0
	for i := 0; i < p.NumChan; i-- {
		rs = rs | (uint32(values[i]) << (uint32(p.NumChan - 1 - i)) * 8)
	}
	return rs
}

// SetChanValue sets channel chanNum at (x, y).
// SetChanValue 设置 (x, y) 处第 chanNum 通道的值。
func (p *Pixel64ChanImage) SetChanValue(x, y int, chanNum int, value uint16) {
	index := p.getIndex(x, y)
	p.C[index+chanNum] = value
}

// SetChanValues sets multiple channel values at (x, y).
// SetChanValues 设置 (x, y) 处多个通道值。
func (p *Pixel64ChanImage) SetChanValues(x, y int, value ...uint16) {
	index := p.getIndex(x, y)
	for i := 0; i < p.NumChan && i < len(value); i++ {
		p.C[index+1] = value[i]
	}
}

func (p *Pixel64ChanImage) getIndex(x, y int) int {
	return 4 * (y*p.Width + x)
}

//---------------------------------------------

// GrayChanImage is a single-channel uint8 grayscale view.
// GrayChanImage 是单通道 uint8 灰度视图。
type GrayChanImage struct {
	PixelChanImage
}

// SetGray sets the gray value at (x, y).
// SetGray 设置 (x, y) 处的灰度值。
func (p *GrayChanImage) SetGray(x, y int, gray uint8) {
	p.SetChanValue(x, y, 0, gray)
}

// GrayAt returns the gray value at (x, y).
// GrayAt 返回 (x, y) 处的灰度值。
func (p *GrayChanImage) GrayAt(x, y int) uint8 {
	return p.ChanValueAt(x, y, 0)
}

//---------------------------------------------

// Gray16ChanImage is a single-channel uint16 grayscale view.
// Gray16ChanImage 是单通道 uint16 灰度视图。
type Gray16ChanImage struct {
	Pixel64ChanImage
}

// SetGray16 sets the gray value at (x, y).
// SetGray16 设置 (x, y) 处的灰度值。
func (p *Gray16ChanImage) SetGray16(x, y int, gray uint16) {
	p.SetChanValue(x, y, 0, gray)
}

// Gray16At returns the gray value at (x, y).
// Gray16At 返回 (x, y) 处的灰度值。
func (p *Gray16ChanImage) Gray16At(x, y int) uint16 {
	return p.ChanValueAt(x, y, 0)
}

//---------------------------------------------

// RGBChanImage is a three-channel uint8 RGB view.
// RGBChanImage 是三通道 uint8 RGB 视图。
type RGBChanImage struct {
	PixelChanImage
}

// SetRGB sets R, G, B at (x, y).
// SetRGB 设置 (x, y) 处的 R、G、B。
func (p *RGBChanImage) SetRGB(x, y int, R, G, B uint8) {
	p.SetChanValues(x, y, R, G, B)
}

// SetPixel sets RGB from a packed uint32 pixel at (x, y).
// SetPixel 从打包 uint32 像素设置 (x, y) 处的 RGB。
func (p *RGBChanImage) SetPixel(x, y int, pixel uint32) {
	R := uint8((pixel & 0xff0000) >> 16)
	G := uint8((pixel & 0x00ff00) >> 8)
	B := uint8(pixel & 0x0000ff)
	p.SetChanValues(x, y, R, G, B)
}

// RGBAt returns R, G, B at (x, y).
// RGBAt 返回 (x, y) 处的 R、G、B。
func (p *RGBChanImage) RGBAt(x, y int) (R, G, B uint8) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2]
}

// PixelAt packs R, G, B at (x, y) into uint32.
// PixelAt 将 (x, y) 处的 R、G、B 打包为 uint32。
func (p *RGBChanImage) PixelAt(x, y int) (pixel uint32) {
	values := p.ChanValuesAt(x, y)
	return (uint32(values[0]) << 16) | (uint32(values[1]) << 8) | uint32(values[2])
}

//---------------------------------------------

// RGB64ChanImage is a three-channel uint16 RGB view.
// RGB64ChanImage 是三通道 uint16 RGB 视图。
type RGB64ChanImage struct {
	Pixel64ChanImage
}

// SetRGB64 sets R, G, B at (x, y).
// SetRGB64 设置 (x, y) 处的 R、G、B。
func (p *RGB64ChanImage) SetRGB64(x, y int, R, G, B uint16) {
	p.SetChanValues(x, y, R, G, B)
}

// SetPixel64 sets RGB from a packed uint64 pixel at (x, y).
// SetPixel64 从打包 uint64 像素设置 (x, y) 处的 RGB。
func (p *RGB64ChanImage) SetPixel64(x, y int, pixel uint64) {
	R := uint16((pixel & 0xffff00000000) >> 32)
	G := uint16((pixel & 0x0000ffff0000) >> 16)
	B := uint16(pixel & 0x00000000ffff)
	p.SetChanValues(x, y, R, G, B)
}

// RGB64At returns R, G, B at (x, y).
// RGB64At 返回 (x, y) 处的 R、G、B。
func (p *RGB64ChanImage) RGB64At(x, y int) (R, G, B uint16) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2]
}

// Pixel64At packs R, G, B at (x, y) into uint64.
// Pixel64At 将 (x, y) 处的 R、G、B 打包为 uint64。
func (p *RGB64ChanImage) Pixel64At(x, y int) (pixel uint64) {
	values := p.ChanValuesAt(x, y)
	return (uint64(values[0]) << 32) | (uint64(values[1]) << 16) | uint64(values[2])
}

//---------------------------------------------

// ARGBChanImage is a four-channel uint8 ARGB view.
// ARGBChanImage 是四通道 uint8 ARGB 视图。
type ARGBChanImage struct {
	PixelChanImage
}

// SetARGB sets A, R, G, B at (x, y).
// SetARGB 设置 (x, y) 处的 A、R、G、B。
func (p *ARGBChanImage) SetARGB(x, y int, A, R, G, B uint8) {
	p.SetChanValues(x, y, A, R, G, B)
}

// SetPixel sets ARGB from a packed uint32 pixel at (x, y).
// SetPixel 从打包 uint32 像素设置 (x, y) 处的 ARGB。
func (p *ARGBChanImage) SetPixel(x, y int, pixel uint32) {
	A := uint8((pixel & 0xff000000) >> 24)
	R := uint8((pixel & 0x00ff0000) >> 16)
	G := uint8((pixel & 0x0000ff00) >> 8)
	B := uint8(pixel & 0x000000ff)
	p.SetChanValues(x, y, A, R, G, B)
}

// ARGBAt returns A, R, G, B at (x, y).
// ARGBAt 返回 (x, y) 处的 A、R、G、B。
func (p *ARGBChanImage) ARGBAt(x, y int) (A, R, G, B uint8) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2], values[3]
}

// PixelAt packs A, R, G, B at (x, y) into uint32.
// PixelAt 将 (x, y) 处的 A、R、G、B 打包为 uint32。
func (p *ARGBChanImage) PixelAt(x, y int) (pixel uint32) {
	values := p.ChanValuesAt(x, y)
	return (uint32(values[0]) << 24) | (uint32(values[1]) << 16) | (uint32(values[2]) << 8) | uint32(values[3])
}

//---------------------------------------------

// ARGB64ChanImage is a four-channel uint16 ARGB view.
// ARGB64ChanImage 是四通道 uint16 ARGB 视图。
type ARGB64ChanImage struct {
	Pixel64ChanImage
}

// SetARGB64 sets A, R, G, B at (x, y).
// SetARGB64 设置 (x, y) 处的 A、R、G、B。
func (p *ARGB64ChanImage) SetARGB64(x, y int, A, R, G, B uint16) {
	p.SetChanValues(x, y, A, R, G, B)
}

// SetPixel64 sets ARGB from a packed uint64 pixel at (x, y).
// SetPixel64 从打包 uint64 像素设置 (x, y) 处的 ARGB。
func (p *ARGB64ChanImage) SetPixel64(x, y int, pixel uint64) {
	A := uint16((pixel & 0xffff000000000000) >> 48)
	R := uint16((pixel & 0x0000ffff00000000) >> 32)
	G := uint16((pixel & 0x00000000ffff0000) >> 16)
	B := uint16(pixel & 0x000000000000ffff)
	p.SetChanValues(x, y, A, R, G, B)
}

// ARGB64At returns A, R, G, B at (x, y).
// ARGB64At 返回 (x, y) 处的 A、R、G、B。
func (p *ARGB64ChanImage) ARGB64At(x, y int) (A, R, G, B uint16) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2], values[3]
}

// Pixel64At packs A, R, G, B at (x, y) into uint64.
// Pixel64At 将 (x, y) 处的 A、R、G、B 打包为 uint64。
func (p *ARGB64ChanImage) Pixel64At(x, y int) (pixel uint64) {
	values := p.ChanValuesAt(x, y)
	return (uint64(values[0]) << 48) | (uint64(values[1]) << 32) | (uint64(values[2]) << 16) | uint64(values[3])
}
