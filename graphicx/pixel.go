package graphicx

type Pixel uint32

type Pixel64 uint64

func ARGB2Pixel(A, R, G, B uint8) (pixel uint32) {
	pA := uint32(A) << 24
	pR := uint32(R) << 16
	pG := uint32(G) << 8
	pB := uint32(B) << 0
	return pA | pR | pG | pB
}

func Pixel2ARGB(pixel uint32) (A, R, G, B uint8) {
	A = uint8((pixel & 0xff000000) >> 24)
	R = uint8((pixel & 0x00ff0000) >> 16)
	G = uint8((pixel & 0x0000ff00) >> 8)
	B = uint8((pixel & 0x000000ff) >> 0)
	return
}

func RGB2APixel(R, G, B, A uint8) (pixel uint32) {
	pR := uint32(R) << 24
	pG := uint32(G) << 16
	pB := uint32(B) << 8
	pA := uint32(A) << 0
	return pA | pR | pG | pB
}

func Pixel2RGBA(pixel uint32) (R, G, B, A uint8) {
	R = uint8((pixel & 0xff000000) >> 24)
	G = uint8((pixel & 0x00ff0000) >> 16)
	B = uint8((pixel & 0x0000ff00) >> 8)
	A = uint8((pixel & 0x000000ff) >> 0)
	return
}

func ARGB2Pixel64(A, R, G, B uint16) (pixel uint64) {
	pA := uint64(A) << 48
	pR := uint64(R) << 32
	pG := uint64(G) << 16
	pB := uint64(B) << 0
	return pA | pR | pG | pB
}

func Pixel2ARGB64(pixel uint64) (A, R, G, B uint16) {
	A = uint16((pixel & 0xffff000000000000) >> 48)
	R = uint16((pixel & 0x0000ffff00000000) >> 32)
	G = uint16((pixel & 0x00000000ffff0000) >> 16)
	B = uint16((pixel & 0x000000000000ffff) >> 0)
	return
}

func RGB2APixel64(R, G, B, A uint16) (pixel uint64) {
	pR := uint64(R) << 48
	pG := uint64(G) << 32
	pB := uint64(B) << 16
	pA := uint64(A) << 0
	return pA | pR | pG | pB
}

func Pixel2RGBA64(pixel uint64) (R, G, B, A uint16) {
	R = uint16((pixel & 0xffff000000000000) >> 48)
	G = uint16((pixel & 0x0000ffff00000000) >> 32)
	B = uint16((pixel & 0x00000000ffff0000) >> 16)
	A = uint16((pixel & 0x000000000000ffff) >> 0)
	return
}
