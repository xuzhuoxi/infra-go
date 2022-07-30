// histogram of Oriented Gradient
// 方向梯度直方图
package featurex

import (
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"math"
)

// 计算红色直方图
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

// 计算红色直方图
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

// 计算绿色直方图
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

// 计算绿色直方图
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

// 计算蓝色直方图
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

// 计算蓝色直方图
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

// 计算透明直方图
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

// 计算透明直方图
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

// 计算透明直方图
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

// 计算透明直方图
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

// 计算透明直方图
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

// 计算透明直方图
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
