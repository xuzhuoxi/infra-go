// 投影： 按水平或垂直方向统计色素有效性
package projectionx

import (
	"image"
	"image/color"
)

// 投影模式
type ProjectDirection int32

const (
	// 无效
	ModeNone ProjectDirection = iota
	// 垂直模式
	ModeVertical
	// 水平模式
	ModeHorizontal
)

// 投影结果
type Result struct {
	// 数据
	Data []uint32
	// 投影模式
	Mode ProjectDirection
}

func (r *Result) MaxData() uint32 {
	var max uint32 = 0
	for _, d := range r.Data {
		if d > max {
			max = d
		}
	}
	return max
}

func (r *Result) Length() int {
	return len(r.Data)
}

// 投影有效性函数
type ValidFunc func(color color.Color) bool

// 水平投影
func ProjectHorizontal(src image.Image, validFunc ValidFunc) *Result {
	size := src.Bounds().Size()
	w, h := size.X, size.Y
	rs := make([]uint32, h, h)
	for y := 0; y < h; y++ {
		rs[y] = 0
		for x := 0; x < w; x++ {
			if validFunc(src.At(x, y)) {
				rs[y]++
			}
		}
	}
	return &Result{Data: rs, Mode: ModeHorizontal}
}

// 垂直投影
func ProjectVertical(src image.Image, validFunc ValidFunc) *Result {
	size := src.Bounds().Size()
	w, h := size.X, size.Y
	rs := make([]uint32, w, w)
	for x := 0; x < w; x++ {
		rs[x] = 0
		for y := 0; y < h; y++ {
			if validFunc(src.At(x, y)) {
				rs[x]++
			}
		}
	}
	return &Result{Data: rs, Mode: ModeVertical}
}

// 投影
func Project(src image.Image, validFunc ValidFunc, direction ProjectDirection) *Result {
	switch direction {
	case ModeHorizontal:
		return ProjectHorizontal(src, validFunc)
	case ModeVertical:
		return ProjectVertical(src, validFunc)
	default:
		return nil
	}
}

// 二维投影
func Project2D(src image.Image, validFunc ValidFunc) (v *Result, h *Result) {
	v = ProjectVertical(src, validFunc)
	h = ProjectHorizontal(src, validFunc)
	return
}
