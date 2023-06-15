package filterx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
)

// FilterMatrix
// 滤波器
type FilterMatrix struct {
	// 滤波器卷积核
	Kernel FilterKernel
	// 滤波器半径
	KernelRadius int
	// 滤波器边长
	KernelSize int
	// 滤波器卷积核倍率
	KernelScale int
	// 运算结果偏移量
	ResultOffset int
}

// IsScaleMatrix
// 是否为倍率滤波器
func (fm FilterMatrix) IsScaleMatrix() bool {
	return fm.KernelScale != 0 && fm.KernelScale != 1
}

// IsPixelUnsafe
// 运算结果是否可能超出像素范围(非安全像素值)
func (fm FilterMatrix) IsPixelUnsafe() bool {
	if fm.ResultOffset != 0 {
		return true
	}
	for _, kv := range fm.Kernel {
		if kv.Value < 0 {
			return true
		}
	}
	return false
}

// FlipUD
// 上下翻转
func (fm FilterMatrix) FlipUD() FilterMatrix {
	rs := fm
	rs.Kernel = fm.Kernel.FlipUD()
	rs.Kernel.Sorted()
	return rs
}

// FlipLR
// 左右翻转
func (fm FilterMatrix) FlipLR() FilterMatrix {
	rs := fm
	rs.Kernel = fm.Kernel.FlipLR()
	rs.Kernel.Sorted()
	return rs
}

// Rotate
// 顺时针旋转90度
func (fm FilterMatrix) Rotate(clockwise bool, count90 int) FilterMatrix {
	rs := fm
	rs.Kernel = fm.Kernel.Rotate(clockwise, count90)
	rs.Kernel.Sorted()
	return rs
}

// CheckValidity
// 滤波模板有效性
func (fm FilterMatrix) CheckValidity() error {
	if fm.KernelRadius < 1 {
		return errors.New("KernelRadius < 1. ")
	}
	if fm.KernelScale < 0 {
		return errors.New("KernelScale < 0. ")
	}
	sum := 0
	for _, vector := range fm.Kernel {
		sum += vector.Value
	}
	if sum != fm.KernelScale {
		return errors.New(fmt.Sprintf("KernelScale Error! KernelScale=%d, AddSum=%d.", fm.KernelScale, sum))
	}
	return nil
}

// FilterImageWithMatrix
// 使用滤波器
func FilterImageWithMatrix(srcImg image.Image, dstImg draw.Image, m FilterMatrix) error {
	if nil == srcImg || nil == dstImg {
		return errors.New("SrcImg or dstImg is nil! ")
	}
	if err := m.CheckValidity(); nil != err {
		return err
	}
	sourceImage := srcImg
	targetImage := dstImg
	if srcImg == dstImg {
		sourceImage = imagex.CopyImage(srcImg)
	}
	size := srcImg.Bounds().Size()
	var x, y int
	var ox, oy int
	var sumR, sumG, sumB int
	var R, G, B, A uint32
	var setColor = &color.NRGBA64{A: 0xffff}

	needScale := m.IsScaleMatrix()
	kScale := m.KernelScale
	unSafe := m.IsPixelUnsafe()
	resultOffset := m.ResultOffset
	switch {
	case needScale && unSafe:
		//放大 负结果
		for y = 0; y < size.Y; y++ {
			for x = 0; x < size.X; x++ {
				sumR, sumG, sumB = 0, 0, 0
				_, _, _, A = sourceImage.At(x, y).RGBA()
				setColor.A = uint16(A)
				for _, vector := range m.Kernel {
					ox = mathx.MinInt(size.X, mathx.MaxInt(x+vector.X, 0))
					oy = mathx.MinInt(size.Y, mathx.MaxInt(y+vector.Y, 0))
					R, G, B, _ = sourceImage.At(ox, oy).RGBA()
					sumR += int(R) * vector.Value
					sumG += int(G) * vector.Value
					sumB += int(B) * vector.Value
				}
				sumR, sumG, sumB = sumR/kScale, sumG/kScale, sumB/kScale
				sumR = mathx.MinInt(mathx.MaxInt(sumR+resultOffset, 0), 65535)
				sumG = mathx.MinInt(mathx.MaxInt(sumG+resultOffset, 0), 65535)
				sumB = mathx.MinInt(mathx.MaxInt(sumB+resultOffset, 0), 65535)
				setColor.R, setColor.G, setColor.B = uint16(sumR), uint16(sumG), uint16(sumB)
				targetImage.Set(x, y, setColor)
			}
		}
	case needScale && !unSafe:
		//放大 非负结果
		for y = 0; y < size.Y; y++ {
			for x = 0; x < size.X; x++ {
				sumR, sumG, sumB = 0, 0, 0
				_, _, _, A = sourceImage.At(x, y).RGBA()
				setColor.A = uint16(A)
				for _, vector := range m.Kernel {
					ox = mathx.MinInt(size.X, mathx.MaxInt(x+vector.X, 0))
					oy = mathx.MinInt(size.Y, mathx.MaxInt(y+vector.Y, 0))
					R, G, B, _ = sourceImage.At(ox, oy).RGBA()
					sumR += int(R) * vector.Value
					sumG += int(G) * vector.Value
					sumB += int(B) * vector.Value
				}
				sumR, sumG, sumB = sumR/kScale, sumG/kScale, sumB/kScale
				setColor.R, setColor.G, setColor.B = uint16(sumR), uint16(sumG), uint16(sumB)
				targetImage.Set(x, y, setColor)
			}
		}
	case !needScale && unSafe:
		//非放大 负结果
		for y = 0; y < size.Y; y++ {
			for x = 0; x < size.X; x++ {
				sumR, sumG, sumB = 0, 0, 0
				_, _, _, A = sourceImage.At(x, y).RGBA()
				setColor.A = uint16(A)
				for _, vector := range m.Kernel {
					ox = mathx.MinInt(size.X, mathx.MaxInt(x+vector.X, 0))
					oy = mathx.MinInt(size.Y, mathx.MaxInt(y+vector.Y, 0))
					R, G, B, _ = sourceImage.At(ox, oy).RGBA()
					sumR += int(R) * vector.Value
					sumG += int(G) * vector.Value
					sumB += int(B) * vector.Value
				}
				sumR = mathx.MinInt(mathx.MaxInt(sumR+resultOffset, 0), 65535)
				sumG = mathx.MinInt(mathx.MaxInt(sumG+resultOffset, 0), 65535)
				sumB = mathx.MinInt(mathx.MaxInt(sumB+resultOffset, 0), 65535)
				setColor.R, setColor.G, setColor.B = uint16(sumR), uint16(sumG), uint16(sumB)
				targetImage.Set(x, y, setColor)
			}
		}
	case !needScale && !unSafe:
		//非放大 非负结果
		for y = 0; y < size.Y; y++ {
			for x = 0; x < size.X; x++ {
				sumR, sumG, sumB = 0, 0, 0
				_, _, _, A = sourceImage.At(x, y).RGBA()
				setColor.A = uint16(A)
				for _, vector := range m.Kernel {
					ox = mathx.MinInt(size.X, mathx.MaxInt(x+vector.X, 0))
					oy = mathx.MinInt(size.Y, mathx.MaxInt(y+vector.Y, 0))
					R, G, B, _ = sourceImage.At(ox, oy).RGBA()
					sumR += int(R) * vector.Value
					sumG += int(G) * vector.Value
					sumB += int(B) * vector.Value
				}
				setColor.R, setColor.G, setColor.B = uint16(sumR), uint16(sumG), uint16(sumB)
				targetImage.Set(x, y, setColor)
			}
		}
	}
	return nil
}
