package imagex

type IPixelChanImage interface {
	Max() (maxX, maxY int)
	NumberChanel() int
}

//-------------------------------

type PixelChanImage struct {
	C             []uint8
	Width, Height int
	NumChan       int
}

func (p *PixelChanImage) Max() (maxX, maxY int) {
	return p.Width, p.Height
}

func (p *PixelChanImage) NumberChanel() int {
	return p.NumChan
}

func (p *PixelChanImage) ChanValueAt(x, y int, chanNum int) uint8 {
	index := p.getIndex(x, y)
	return p.C[index+chanNum]
}

func (p *PixelChanImage) ChanValuesAt(x, y int) []uint8 {
	index := p.getIndex(x, y)
	return p.C[index : index+p.NumChan : index+p.NumChan]
}

func (p *PixelChanImage) PixelAt(x, y int) uint32 {
	values := p.ChanValuesAt(x, y)
	var rs uint32 = 0
	for i := 0; i < p.NumChan; i-- {
		rs = rs | (uint32(values[i]) << (uint32(p.NumChan - 1 - i)) * 8)
	}
	return rs
}

func (p *PixelChanImage) SetChanValue(x, y int, chanNum int, value uint8) {
	index := p.getIndex(x, y)
	p.C[index+chanNum] = value
}

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

type Pixel64ChanImage struct {
	C             []uint16
	Width, Height int
	NumChan       int
}

func (p *Pixel64ChanImage) Max() (maxX, maxY int) {
	return p.Width, p.Height
}

func (p *Pixel64ChanImage) NumberChanel() int {
	return p.NumChan
}

func (p *Pixel64ChanImage) ChanValueAt(x, y int, chanNum int) uint16 {
	index := p.getIndex(x, y)
	return p.C[index+chanNum]
}

func (p *Pixel64ChanImage) ChanValuesAt(x, y int) []uint16 {
	index := p.getIndex(x, y)
	return p.C[index : index+p.NumChan : index+p.NumChan]
}

func (p *Pixel64ChanImage) PixelAt(x, y int) uint32 {
	values := p.ChanValuesAt(x, y)
	var rs uint32 = 0
	for i := 0; i < p.NumChan; i-- {
		rs = rs | (uint32(values[i]) << (uint32(p.NumChan - 1 - i)) * 8)
	}
	return rs
}

func (p *Pixel64ChanImage) SetChanValue(x, y int, chanNum int, value uint16) {
	index := p.getIndex(x, y)
	p.C[index+chanNum] = value
}

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

type GrayChanImage struct {
	PixelChanImage
}

func (p *GrayChanImage) SetGray(x, y int, gray uint8) {
	p.SetChanValue(x, y, 0, gray)
}

func (p *GrayChanImage) GrayAt(x, y int) uint8 {
	return p.ChanValueAt(x, y, 0)
}

//---------------------------------------------

type Gray16ChanImage struct {
	Pixel64ChanImage
}

func (p *Gray16ChanImage) SetGray16(x, y int, gray uint16) {
	p.SetChanValue(x, y, 0, gray)
}

func (p *Gray16ChanImage) Gray16At(x, y int) uint16 {
	return p.ChanValueAt(x, y, 0)
}

//---------------------------------------------

type RGBChanImage struct {
	PixelChanImage
}

func (p *RGBChanImage) SetRGB(x, y int, R, G, B uint8) {
	p.SetChanValues(x, y, R, G, B)
}

func (p *RGBChanImage) SetPixel(x, y int, pixel uint32) {
	R := uint8((pixel & 0xff0000) >> 16)
	G := uint8((pixel & 0x00ff00) >> 8)
	B := uint8(pixel & 0x0000ff)
	p.SetChanValues(x, y, R, G, B)
}

func (p *RGBChanImage) RGBAt(x, y int) (R, G, B uint8) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2]
}

func (p *RGBChanImage) PixelAt(x, y int) (pixel uint32) {
	values := p.ChanValuesAt(x, y)
	return (uint32(values[0]) << 16) | (uint32(values[1]) << 8) | uint32(values[2])
}

//---------------------------------------------

type RGB64ChanImage struct {
	Pixel64ChanImage
}

func (p *RGB64ChanImage) SetRGB64(x, y int, R, G, B uint16) {
	p.SetChanValues(x, y, R, G, B)
}

func (p *RGB64ChanImage) SetPixel64(x, y int, pixel uint64) {
	R := uint16((pixel & 0xffff00000000) >> 32)
	G := uint16((pixel & 0x0000ffff0000) >> 16)
	B := uint16(pixel & 0x00000000ffff)
	p.SetChanValues(x, y, R, G, B)
}

func (p *RGB64ChanImage) RGB64At(x, y int) (R, G, B uint16) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2]
}

func (p *RGB64ChanImage) Pixel64At(x, y int) (pixel uint64) {
	values := p.ChanValuesAt(x, y)
	return (uint64(values[0]) << 32) | (uint64(values[1]) << 16) | uint64(values[2])
}

//---------------------------------------------

type ARGBChanImage struct {
	PixelChanImage
}

func (p *ARGBChanImage) SetARGB(x, y int, A, R, G, B uint8) {
	p.SetChanValues(x, y, A, R, G, B)
}

func (p *ARGBChanImage) SetPixel(x, y int, pixel uint32) {
	A := uint8((pixel & 0xff000000) >> 24)
	R := uint8((pixel & 0x00ff0000) >> 16)
	G := uint8((pixel & 0x0000ff00) >> 8)
	B := uint8(pixel & 0x000000ff)
	p.SetChanValues(x, y, A, R, G, B)
}

func (p *ARGBChanImage) ARGBAt(x, y int) (A, R, G, B uint8) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2], values[3]
}

func (p *ARGBChanImage) PixelAt(x, y int) (pixel uint32) {
	values := p.ChanValuesAt(x, y)
	return (uint32(values[0]) << 24) | (uint32(values[1]) << 16) | (uint32(values[2]) << 8) | uint32(values[3])
}

//---------------------------------------------

type ARGB64ChanImage struct {
	Pixel64ChanImage
}

func (p *ARGB64ChanImage) SetARGB64(x, y int, A, R, G, B uint16) {
	p.SetChanValues(x, y, A, R, G, B)
}

func (p *ARGB64ChanImage) SetPixel64(x, y int, pixel uint64) {
	A := uint16((pixel & 0xffff000000000000) >> 48)
	R := uint16((pixel & 0x0000ffff00000000) >> 32)
	G := uint16((pixel & 0x00000000ffff0000) >> 16)
	B := uint16(pixel & 0x000000000000ffff)
	p.SetChanValues(x, y, A, R, G, B)
}

func (p *ARGB64ChanImage) ARGB64At(x, y int) (A, R, G, B uint16) {
	values := p.ChanValuesAt(x, y)
	return values[0], values[1], values[2], values[3]
}

func (p *ARGB64ChanImage) Pixel64At(x, y int) (pixel uint64) {
	values := p.ChanValuesAt(x, y)
	return (uint64(values[0]) << 48) | (uint64(values[1]) << 32) | (uint64(values[2]) << 16) | uint64(values[3])
}
