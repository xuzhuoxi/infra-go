package projectionx

import (
	"image"
	"image/color"
)

// ProjectDirection selects horizontal or vertical projection mode.
// ProjectDirection 选择水平或垂直投影模式。
type ProjectDirection int32

const (
	// ModeNone is an invalid projection mode.
	// ModeNone 无效投影模式。
	ModeNone ProjectDirection = iota
	// ModeVertical projects along the vertical axis (per column).
	// ModeVertical 垂直模式（按列统计）。
	ModeVertical
	// ModeHorizontal projects along the horizontal axis (per row).
	// ModeHorizontal 水平模式（按行统计）。
	ModeHorizontal
)

// Result holds per-axis valid-pixel counts from a projection.
// Result 保存投影各轴上的有效像素计数。
type Result struct {
	// 数据
	Data []uint32
	// 投影模式
	Mode ProjectDirection
}

// MaxData returns the maximum count in Data.
// MaxData 返回 Data 中的最大计数值。
func (r *Result) MaxData() uint32 {
	var max uint32 = 0
	for _, d := range r.Data {
		if d > max {
			max = d
		}
	}
	return max
}

// Length returns len(Data).
// Length 返回 len(Data)。
func (r *Result) Length() int {
	return len(r.Data)
}

// ValidFunc decides whether a pixel counts toward the projection.
// ValidFunc 判断像素是否计入投影统计。
type ValidFunc func(color color.Color) bool

// ProjectHorizontal counts valid pixels per row.
// ProjectHorizontal 按行统计有效像素（水平投影）。
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

// ProjectVertical counts valid pixels per column.
// ProjectVertical 按列统计有效像素（垂直投影）。
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

// Project dispatches to horizontal or vertical projection by direction.
// Project 按 direction 分发到水平或垂直投影。
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

// Project2D returns vertical and horizontal projections of src.
// Project2D 同时返回 src 的垂直与水平投影。
func Project2D(src image.Image, validFunc ValidFunc) (v *Result, h *Result) {
	v = ProjectVertical(src, validFunc)
	h = ProjectHorizontal(src, validFunc)
	return
}
