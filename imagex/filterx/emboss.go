package filterx

import (
	"errors"

	"github.com/xuzhuoxi/infra-go/errorsx"
)

// EmbossDirection specifies emboss light direction.
// EmbossDirection 表示浮雕光照方向。
type EmbossDirection int

// 以下顺序不能乱
const (
	// Lu2Rd is upper-left to lower-right emboss direction.
	// Lu2Rd 是左上至右下浮雕方向。
	Lu2Rd EmbossDirection = iota
	// Ru2Ld is upper-right to lower-left emboss direction.
	// Ru2Ld 是右上至左下浮雕方向。
	Ru2Ld
	// Rd2Lu is lower-right to upper-left emboss direction.
	// Rd2Lu 是右下至上左浮雕方向。
	Rd2Lu
	// Ld2Ru is lower-left to upper-right emboss direction.
	// Ld2Ru 是左下至右上浮雕方向。
	Ld2Ru

	// U2D is top-to-bottom emboss direction.
	// U2D 是上至下浮雕方向。
	U2D
	// R2L is right-to-left emboss direction.
	// R2L 是右至左浮雕方向。
	R2L
	// D2U is bottom-to-top emboss direction.
	// D2U 是下至上浮雕方向。
	D2U
	// L2R is left-to-right emboss direction.
	// L2R 是左至右浮雕方向。
	L2R
)

// 浮雕滤波器
var (
	// Emboss3Lu2Rd is a 3×3 upper-left to lower-right emboss kernel.
	// Emboss3Lu2Rd 是 3×3（左上→右下）浮雕滤波器。
	Emboss3Lu2Rd = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, -1}, {-0, -1, -1},
			{-1, -0, -1}, {+1, +0, +1},
			{+0, +1, +1}, {+1, +1, +1}}}
	// Emboss3U2D is a 3×3 top-to-bottom emboss kernel.
	// Emboss3U2D 是 3×3（上→下）浮雕滤波器。
	Emboss3U2D = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, -1}, {0, -1, -1}, {+1, -1, -1},
			{-1, +1, +1}, {0, +1, +1}, {+1, +1, +1}}}
	// Emboss5Lu2Rd is a 5×5 upper-left to lower-right emboss kernel.
	// Emboss5Lu2Rd 是 5×5（左上→右下）浮雕滤波器。
	Emboss5Lu2Rd = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -2, -1}, {+0, -2, -1}, {+1, -2, -1},
			{-2, -1, -1}, {-1, -1, -1}, {+0, -1, -1}, {+2, -1, +1},
			{-2, -0, -1}, {-1, -0, -1}, {+1, -0, +1}, {+2, -0, +1},
			{-2, +1, -1}, {+0, +1, +1}, {+1, +1, +1}, {+2, +1, +1},
			{-1, +2, +1}, {+0, +2, +1}, {+1, +2, +1}, {+2, +2, +1}}}
	// Emboss5U2D is a 5×5 top-to-bottom emboss kernel.
	// Emboss5U2D 是 5×5（上→下）浮雕滤波器。
	Emboss5U2D = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -2, -1}, {0, -2, -1}, {+1, -2, -1}, {+2, -2, -1},
			{-2, -1, -1}, {-1, -1, -1}, {0, -1, -1}, {+1, -0, -1}, {+2, -1, -1},
			{-2, +1, +1}, {-1, +1, +1}, {0, +1, +1}, {+1, +1, +1}, {+2, +1, +1},
			{-2, +2, +1}, {-1, +2, +1}, {0, +2, +1}, {+1, +2, +1}, {+2, +2, +1}}}
	// Emboss3Asymmetrical is a 3×3 asymmetrical emboss kernel.
	// Emboss3Asymmetrical 是 3×3 非对称浮雕滤波器。
	Emboss3Asymmetrical = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, 2}, {+0, +0, -1}, {+1, +1, -1}}}
)

//---------------------------------------

// CreateEmbossFilterSymmetry builds a symmetric emboss kernel.
// CreateEmbossFilterSymmetry 创建对称浮雕滤波器。
func CreateEmbossFilterSymmetry(radius int, direction EmbossDirection, offset int) (filter FilterMatrix, err error) {
	if radius < 1 {
		return filter, errors.New("KernelRadius < 1. ")
	}
	kSize := radius + radius + 1
	var kernel FilterKernel
	switch direction {
	case Lu2Rd, Ru2Ld, Rd2Lu, Ld2Ru:
		kernel = getEmbossLu2RdKernel(radius)
	case U2D, R2L, D2U, L2R:
		kernel = getEmbossU2DKernel(radius)
	default:
		return filter, errorsx.NoCaseCatchError("CreateEmbossFilterSymmetry")
	}
	kernel.RotateSelf(true, int(direction))
	return FilterMatrix{KernelRadius: radius, KernelSize: kSize, KernelScale: 0, ResultOffset: offset, Kernel: kernel}, nil
}

// CreateAsymmetricalEmbossFilter builds an asymmetrical emboss kernel.
// CreateAsymmetricalEmbossFilter 创建非对称浮雕滤波器。
func CreateAsymmetricalEmbossFilter(radius int, value int, offset int, angle int) (filter FilterMatrix, err error) {
	return
}

// 生成(左上->右下)浮雕卷积核心
func getEmbossLu2RdKernel(radius int) FilterKernel {
	kSize := radius + radius + 1
	ln := kSize*kSize - kSize
	var lu2rdKernel FilterKernel = make([]KernelVector, 0, ln)
	var add int
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			add = x + y
			if 0 == add {
				continue
			}
			if add < 0 {
				lu2rdKernel = append(lu2rdKernel, KernelVector{X: x, Y: y, Value: -1})
			} else {
				lu2rdKernel = append(lu2rdKernel, KernelVector{X: x, Y: y, Value: +1})
			}
		}
	}
	return lu2rdKernel
}

// 生成(上->下)浮雕卷积核心
func getEmbossU2DKernel(radius int) FilterKernel {
	kSize := radius + radius + 1
	ln := kSize*kSize - kSize
	var u2dKernel FilterKernel = make([]KernelVector, 0, ln)
	for y := -radius; y <= radius && y != 0; y++ {
		for x := -radius; x <= radius; x++ {
			if y < 0 {
				u2dKernel = append(u2dKernel, KernelVector{X: x, Y: y, Value: -1})
			} else {
				u2dKernel = append(u2dKernel, KernelVector{X: x, Y: y, Value: +1})
			}
		}
	}
	return u2dKernel
}
