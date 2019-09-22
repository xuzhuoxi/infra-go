package imagex

type PixelImage struct {
	Pix           []uint32
	Width, Height int
}

func (i *PixelImage) Max() (maxX, maxY int) {
	return i.Width, i.Height
}

func (i *PixelImage) At(x, y int) uint32 {
	return i.Pix[i.index(x, y)]
}

func (i *PixelImage) Set(x, y int, gray uint32) {
	i.Pix[i.index(x, y)] = gray
}

func (i *PixelImage) ForEachPixel(eachFunc func(x, y int, pixel uint32)) {
	var x, y int
	for y = 0; y < i.Height; y++ {
		for x = 0; x < i.Width; x++ {
			eachFunc(x, y, i.Pix[i.index(x, y)])
		}
	}
}

func (i *PixelImage) index(x, y int) int {
	return y*i.Width + x
}

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
