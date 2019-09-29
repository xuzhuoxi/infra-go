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

// 滤波器向量单元
type KernelVector struct {
	//向量X
	X int
	//向量Y
	Y int
	//向量值
	Value int
}

// 滤波器
type FilterMatrix struct {
	// 滤波器半径
	Radius int
	// 滤波器边长
	Size int
	// 滤波器卷积核
	Kernel []KernelVector
	// 滤波器卷积核倍率
	KernelScale int
}

// 滤波模板有效性
func (t FilterMatrix) CheckValidity() error {
	if t.Radius < 1 {
		return errors.New("Radius < 1. ")
	}
	if t.KernelScale < 0 {
		return errors.New("KernelScale < 0. ")
	}
	sum := 0
	for _, vector := range t.Kernel {
		sum += vector.Value
	}
	if sum != t.KernelScale {
		return errors.New(fmt.Sprintf("KernelScale Error! KernelScale=%d, AddSum=%d.", t.KernelScale, sum))
	}
	return nil
}

//使用滤波器
func FilterImageWithTemplate(srcImg image.Image, dstImg draw.Image, template FilterMatrix) error {
	if nil == srcImg || nil == dstImg {
		return errors.New("SrcImg or dstImg is nil! ")
	}
	if err := template.CheckValidity(); nil != err {
		return err
	}
	sourceImage := srcImg
	targetImage := dstImg
	if srcImg == dstImg {
		sourceImage = imagex.CopyImage(srcImg)
	}
	size := srcImg.Bounds().Size()
	//radius := template.Radius
	var x, y int
	var ox, oy int
	var sumR, sumG, sumB, sumA int
	var R, G, B, A uint32
	var setColor = &color.NRGBA64{}
	//内部处理
	for y = 0; y < size.Y; y++ {
		for x = 0; x < size.X; x++ {
			sumR, sumG, sumB, sumA = 0, 0, 0, 0
			for _, vector := range template.Kernel {
				ox = x + mathx.MinInt(size.X, mathx.MaxInt(vector.X, 0))
				oy = y + mathx.MinInt(size.Y, mathx.MaxInt(vector.Y, 0))
				R, G, B, A = sourceImage.At(ox, oy).RGBA()
				sumR += int(R) * vector.Value
				sumG += int(G) * vector.Value
				sumB += int(B) * vector.Value
				sumA += int(A) * vector.Value
			}
			if template.KernelScale != 0 && template.KernelScale != 1 {
				sumR, sumG, sumB, sumA = sumR/template.KernelScale, sumG/template.KernelScale, sumB/template.KernelScale, sumA/template.KernelScale
			}
			setColor.R, setColor.G, setColor.B, setColor.A = uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)
			targetImage.Set(x, y, setColor)
		}
	}
	////边缘处理
	//handleEdge := func(x, y int) {
	//	sumR, sumG, sumB, sumA = 0, 0, 0, 0
	//	for _, offset := range template.Kernel {
	//		ox = x + mathx.MinInt(size.X, mathx.MaxInt(offset.X, 0))
	//		oy = y + mathx.MinInt(size.Y, mathx.MaxInt(offset.Y, 0))
	//		R, G, B, A = sourceImage.At(ox, oy).RGBA()
	//		sumR += int(R) * offset.Value
	//		sumG += int(G) * offset.Value
	//		sumB += int(B) * offset.Value
	//		sumA += int(A) * offset.Value
	//	}
	//	if template.KernelScale != 0 && template.KernelScale != 1 {
	//		sumR, sumG, sumB, sumA = sumR/template.KernelScale, sumG/template.KernelScale, sumB/template.KernelScale, sumA/template.KernelScale
	//	}
	//	setColor.R, setColor.G, setColor.B, setColor.A = uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)
	//	targetImage.Set(x, y, setColor)
	//}
	//for y = 0; y < radius; y++ {
	//	for x = 0; x < size.X; x++ {
	//		handleEdge(x, y)
	//	}
	//}
	//for y = size.Y - radius; y < size.Y; y++ {
	//	for x = 0; x < size.X; x++ {
	//		handleEdge(x, y)
	//	}
	//}
	//for x := 0; x < radius; x++ {
	//	for y = 0; y < size.Y; y++ {
	//		handleEdge(x, y)
	//	}
	//}
	//for x := size.X - radius; x < size.X; x++ {
	//	for y = 0; y < size.Y; y++ {
	//		handleEdge(x, y)
	//	}
	//}
	return nil
}
