package imagex

import "image/color"

type PixelChanImage struct {
	C             []uint8
	Width, Height int
}

func (i *PixelChanImage) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

func (i *PixelChanImage) ChanAt(x, y int) (C uint8) {
	index := i.getIndex(x, y)
	return i.C[index]
}

func (i *PixelChanImage) SetChan(x, y int, C uint8) {
	index := i.getIndex(x, y)
	i.C[index] = C
}

func (i *PixelChanImage) getIndex(x, y int) int {
	return 4 * (y*i.Width + x)
}

//---------------------------

type PixelChanImage16 struct {
	C             []uint16
	Width, Height int
}

func (i *PixelChanImage16) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

func (i *PixelChanImage16) ChanAt(x, y int) (C uint16) {
	index := i.getIndex(x, y)
	return i.C[index]
}

func (i *PixelChanImage16) SetChan(x, y int, C uint16) {
	index := i.getIndex(x, y)
	i.C[index] = C
}

func (i *PixelChanImage16) getIndex(x, y int) int {
	return 4 * (y*i.Width + x)
}

//---------------------------

type PixelChan4Image struct {
	A, R, G, B    []uint8
	Width, Height int
}

func (i *PixelChan4Image) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

func (i *PixelChan4Image) RGBAt(x, y int) (A, R, G, B uint8) {
	index := i.getIndex(x, y)
	return i.A[index], i.R[index], i.G[index], i.B[index]
}

func (i *PixelChan4Image) SetPixel(x, y int, pixel uint32) {
	A := (pixel & 0xff000000) >> 24
	R := (pixel & 0x00ff0000) >> 16
	G := (pixel & 0x0000ff00) >> 8
	B := (pixel & 0x000000ff) >> 0
	index := i.getIndex(x, y)
	i.A[index], i.R[index], i.G[index], i.B[index] = uint8(A), uint8(R), uint8(G), uint8(B)
}

func (i *PixelChan4Image) SetRGB(x, y int, A, R, G, B uint8) {
	index := i.getIndex(x, y)
	i.A[index], i.R[index], i.G[index], i.B[index] = A, R, G, B
}

func (i *PixelChan4Image) SetColor(x, y int, c color.Color) {
	R, G, B, A := c.RGBA()
	index := i.getIndex(x, y)
	i.A[index], i.R[index], i.G[index], i.B[index] = uint8(A>>8), uint8(R>>8), uint8(G>>8), uint8(B>>8)
}

func (i *PixelChan4Image) getIndex(x, y int) int {
	return 4 * (y*i.Width + x)
}

//--------------------------------------------------------------

type PixelChan4Image64 struct {
	A, R, G, B    []uint16
	Width, Height int
}

func (i *PixelChan4Image64) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

func (i *PixelChan4Image64) RGB64At(x, y int) (A, R, G, B uint16) {
	index := i.getIndex(x, y)
	return i.A[index], i.R[index], i.G[index], i.B[index]
}

func (i *PixelChan4Image64) SetPixel(x, y int, pixel uint64) {
	A := (pixel & 0xffff00000000000000) >> 48
	R := (pixel & 0x0000ffff0000000000) >> 32
	G := (pixel & 0x00000000ffff000000) >> 16
	B := (pixel & 0x00000000000000ffff) >> 0
	index := i.getIndex(x, y)
	i.A[index], i.R[index], i.G[index], i.B[index] = uint16(A), uint16(R), uint16(G), uint16(B)
}

func (i *PixelChan4Image64) SetRGB64(x, y int, A, R, G, B uint16) {
	index := i.getIndex(x, y)
	i.A[index], i.R[index], i.G[index], i.B[index] = A, R, G, B
}

func (i *PixelChan4Image64) SetColor(x, y int, c color.Color) {
	R, G, B, A := c.RGBA()
	index := i.getIndex(x, y)
	i.A[index], i.R[index], i.G[index], i.B[index] = uint16(A), uint16(R), uint16(G), uint16(B)
}

func (i *PixelChan4Image64) getIndex(x, y int) int {
	return 4 * (y*i.Width + x)
}
