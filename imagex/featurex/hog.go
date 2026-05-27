package featurex

import (
	"image"
	"math"

	"github.com/xuzhuoxi/infra-go/mathx"
)

// HogRed builds an 8-bit histogram of the red channel with range mapping k*O+l.
// HogRed 对红色通道构建 8 位直方图，范围映射为 k*O+l。
func HogRed(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(256, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			O, _, _, _ = img.At(x, y).RGBA()
			N = hogRevise(O>>8, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogRed64 builds a 16-bit histogram of the red channel with range mapping k*O+l.
// HogRed64 对红色通道构建 16 位直方图，范围映射为 k*O+l。
func HogRed64(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(65536, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			O, _, _, _ = img.At(x, y).RGBA()
			N = hogRevise(O, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogGreen builds an 8-bit histogram of the green channel with range mapping k*O+l.
// HogGreen 对绿色通道构建 8 位直方图，范围映射为 k*O+l。
func HogGreen(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(256, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			_, O, _, _ = img.At(x, y).RGBA()
			N = hogRevise(O>>8, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogGreen64 builds a 16-bit histogram of the green channel with range mapping k*O+l.
// HogGreen64 对绿色通道构建 16 位直方图，范围映射为 k*O+l。
func HogGreen64(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(65536, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			_, O, _, _ = img.At(x, y).RGBA()
			N = hogRevise(O, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogBlue builds an 8-bit histogram of the blue channel with range mapping k*O+l.
// HogBlue 对蓝色通道构建 8 位直方图，范围映射为 k*O+l。
func HogBlue(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(256, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			_, _, O, _ = img.At(x, y).RGBA()
			N = hogRevise(O>>8, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogBlue64 builds a 16-bit histogram of the blue channel with range mapping k*O+l.
// HogBlue64 对蓝色通道构建 16 位直方图，范围映射为 k*O+l。
func HogBlue64(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(65536, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			_, _, O, _ = img.At(x, y).RGBA()
			N = hogRevise(O, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogAlpha builds an 8-bit histogram of the alpha channel with range mapping k*O+l.
// HogAlpha 对 Alpha 通道构建 8 位直方图，范围映射为 k*O+l。
func HogAlpha(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(256, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			_, _, _, O = img.At(x, y).RGBA()
			N = hogRevise(O>>8, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogAlpha64 builds a 16-bit histogram of the alpha channel with range mapping k*O+l.
// HogAlpha64 对 Alpha 通道构建 16 位直方图，范围映射为 k*O+l。
func HogAlpha64(img image.Image, k, l float64) []int {
	size := img.Bounds().Size()
	ln := hog(65536, k, l)
	min, max := 0, ln
	rs := make([]int, ln, ln)
	var O uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			_, _, _, O = img.At(x, y).RGBA()
			N = hogRevise(O, k, l, min, max)
			rs[N]++
		}
	}
	return rs
}

// HogRGBA builds 8-bit histograms for R, G, B, and A channels.
// HogRGBA 对 R、G、B、A 通道分别构建 8 位直方图。
func HogRGBA(img image.Image, k, l float64) (RHog, GHog, BHog, AHog []int) {
	size := img.Bounds().Size()
	ln := hog(256, k, l)
	min, max := 0, ln
	RHog = make([]int, ln, ln)
	GHog = make([]int, ln, ln)
	BHog = make([]int, ln, ln)
	AHog = make([]int, ln, ln)
	var R, G, B, A uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			R, G, B, A = img.At(x, y).RGBA()
			N = hogRevise(R>>8, k, l, min, max)
			RHog[N]++
			N = hogRevise(G>>8, k, l, min, max)
			GHog[N]++
			N = hogRevise(B>>8, k, l, min, max)
			BHog[N]++
			N = hogRevise(A>>8, k, l, min, max)
			AHog[N]++
		}
	}
	return
}

// HogRGBA64 builds 16-bit histograms for R, G, B, and A channels.
// HogRGBA64 对 R、G、B、A 通道分别构建 16 位直方图。
func HogRGBA64(img image.Image, k, l float64) (RHog, GHog, BHog, AHog []int) {
	size := img.Bounds().Size()
	ln := hog(65536, k, l)
	min, max := 0, ln
	RHog = make([]int, ln, ln)
	GHog = make([]int, ln, ln)
	BHog = make([]int, ln, ln)
	AHog = make([]int, ln, ln)
	var R, G, B, A uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			R, G, B, A = img.At(x, y).RGBA()
			N = hogRevise(R, k, l, min, max)
			RHog[N]++
			N = hogRevise(G, k, l, min, max)
			GHog[N]++
			N = hogRevise(B, k, l, min, max)
			BHog[N]++
			N = hogRevise(A, k, l, min, max)
			AHog[N]++
		}
	}
	return
}

// HogRGB builds 8-bit histograms for R, G, and B channels.
// HogRGB 对 R、G、B 通道分别构建 8 位直方图。
func HogRGB(img image.Image, k, l float64) (RHog, GHog, BHog []int) {
	size := img.Bounds().Size()
	ln := hog(256, k, l)
	min, max := 0, ln
	RHog = make([]int, ln, ln)
	GHog = make([]int, ln, ln)
	BHog = make([]int, ln, ln)
	var R, G, B uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			R, G, B, _ = img.At(x, y).RGBA()
			N = hogRevise(R>>8, k, l, min, max)
			RHog[N]++
			N = hogRevise(G>>8, k, l, min, max)
			GHog[N]++
			N = hogRevise(B>>8, k, l, min, max)
			BHog[N]++
		}
	}
	return
}

// HogRGB64 builds 16-bit histograms for R, G, and B channels.
// HogRGB64 对 R、G、B 通道分别构建 16 位直方图。
func HogRGB64(img image.Image, k, l float64) (RHog, GHog, BHog []int) {
	size := img.Bounds().Size()
	ln := hog(65536, k, l)
	min, max := 0, ln
	RHog = make([]int, ln, ln)
	GHog = make([]int, ln, ln)
	BHog = make([]int, ln, ln)
	var R, G, B uint32
	var N int
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			R, G, B, _ = img.At(x, y).RGBA()
			N = hogRevise(R, k, l, min, max)
			RHog[N]++
			N = hogRevise(G, k, l, min, max)
			GHog[N]++
			N = hogRevise(B, k, l, min, max)
			BHog[N]++
		}
	}
	return
}

//直方图算子(范围修正)
func hogRevise(O uint32, k, l float64, min, max int) int {
	N := hog(O, k, l)
	N = mathx.MinInt(max, N)
	N = mathx.MaxInt(min, N)
	return N
}

//直方图算子(范围修正)
func hog(O uint32, k, l float64) int {
	return int(math.Ceil(k*float64(O) + l))
}
