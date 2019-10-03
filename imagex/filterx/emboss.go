//浮雕
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/errorsx"
)

type EmbossDirection int

//以下顺序不能乱
const (
	Lu2Rd EmbossDirection = iota
	Ru2Ld
	Rd2Lu
	Ld2Ru

	U2D
	R2L
	D2U
	L2R
)

// 浮雕滤波器
var (
	//3x3 (左上->右下)浮雕滤波器
	Emboss3Lu2Rd = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, -1}, {-0, -1, -1},
			{-1, -0, -1}, {+1, +0, +1},
			{+0, +1, +1}, {+1, +1, +1}}}
	//3x3 (上->下)浮雕滤波器
	Emboss3U2D = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, -1}, {0, -1, -1}, {+1, -1, -1},
			{-1, +1, +1}, {0, +1, +1}, {+1, +1, +1}}}
	//5x5 (左上->右下)浮雕滤波器
	Emboss5Lu2Rd = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -2, -1}, {+0, -2, -1}, {+1, -2, -1},
			{-2, -1, -1}, {-1, -1, -1}, {+0, -1, -1}, {+2, -1, +1},
			{-2, -0, -1}, {-1, -0, -1}, {+1, -0, +1}, {+2, -0, +1},
			{-2, +1, -1}, {+0, +1, +1}, {+1, +1, +1}, {+2, +1, +1},
			{-1, +2, +1}, {+0, +2, +1}, {+1, +2, +1}, {+2, +2, +1}}}
	//5x5 (上->下)浮雕滤波器
	Emboss5U2D = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -2, -1}, {0, -2, -1}, {+1, -2, -1}, {+2, -2, -1},
			{-2, -1, -1}, {-1, -1, -1}, {0, -1, -1}, {+1, -0, -1}, {+2, -1, -1},
			{-2, +1, +1}, {-1, +1, +1}, {0, +1, +1}, {+1, +1, +1}, {+2, +1, +1},
			{-2, +2, +1}, {-1, +2, +1}, {0, +2, +1}, {+1, +2, +1}, {+2, +2, +1}}}
	//3x3 非对称浮雕滤波器
	Emboss3Asymmetrical = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0, ResultOffset: 32768,
		Kernel: []KernelVector{
			{-1, -1, 2}, {+0, +0, -1}, {+1, +1, -1}}}
)

//---------------------------------------

// 创建对称浮雕滤波器
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

// 创建对称浮雕滤波器
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
