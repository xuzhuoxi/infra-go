// 锐化
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/draw"
)

// 锐化滤波器
var (
	//3x3 简单锐化滤波器(拉普拉斯4邻滤波器)
	Sharpen3Laplace4 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 1,
		Kernel: []KernelVector{
			{0, -1, -1}, {-1, +0, -1}, {0, +0, +5}, {1, +0, -1}, {0, +1, -1}}}
	//3x3 锐化滤波器(拉普拉斯8邻滤波器)
	Sharpen3Laplace8 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 1,
		Kernel: []KernelVector{
			{-1, -1, -1}, {0, -1, -1}, {1, -1, -1},
			{-1, +0, -1}, {0, +0, +9}, {1, +0, -1},
			{-1, +1, -1}, {0, +1, -1}, {1, +1, -1}}}
	//3x3 全方向锐化滤波器(强调边缘)
	SharpenStrengthen3All = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 1,
		Kernel: []KernelVector{
			{-1, -1, +1}, {0, -1, +1}, {1, -1, +1},
			{-1, +0, +1}, {0, +0, -7}, {1, +0, +1},
			{-1, +1, +1}, {0, +1, +1}, {1, +1, +1}}}
	//5x5 全方向锐化滤波器
	Sharpen5All = FilterMatrix{KernelRadius: 2, KernelSize: 5, KernelScale: 1,
		Kernel: []KernelVector{ // 中间改为1，原来是8过不了检验
			{-2, -2, -1}, {-1, -2, -1}, {0, -2, -1}, {1, -2, -1}, {2, -2, -1},
			{-2, -1, -1}, {-1, -1, +2}, {0, -1, +2}, {1, -1, +2}, {2, -1, -1},
			{-2, +0, -1}, {-1, +0, +2}, {0, +0, +1}, {1, +0, +2}, {2, +0, -1},
			{-2, +1, -1}, {-1, +1, +2}, {0, +1, +2}, {1, +1, +2}, {2, +1, -1},
			{-2, +2, -1}, {-1, +2, -1}, {1, +2, -1}, {1, +2, -1}, {2, +2, -1}}}
)

//----------------------------------------

// 创建边锐化滤波器
// radius：		卷积核半径
// direction: 	运动方向
// diff:		梯度差
func CreateSharpenFilter(radius int, direction imagex.PixelDirection, diff uint) (filter FilterMatrix, err error) {
	if radius < 1 {
		return filter, errors.New("KernelRadius < 1. ")
	}
	dirAdds := imagex.GetPixelDirectionAdds(direction)
	if nil == dirAdds {
		return filter, errors.New("Direction Error. ")
	}
	kSize := radius + radius + 1
	ln := len(dirAdds)*radius + 1
	var kernel = make([]KernelVector, 0, ln)
	var value int
	var sumValue int
	for _, add := range dirAdds {
		for i := radius; i >= 1; i-- {
			value = (i-radius)*int(diff) - 1
			kernel = append(kernel, KernelVector{X: add.X * i, Y: add.Y * i, Value: value})
			sumValue += value
		}
	}
	kernel = append(kernel, KernelVector{X: 0, Y: 0, Value: -sumValue + 1})
	return FilterMatrix{KernelRadius: radius, KernelSize: kSize, KernelScale: 0, Kernel: kernel}, nil
}

// 创建拉普拉斯4邻锐化滤波器
func CreateSharpenLaplace4Filter(radius int) (filter FilterMatrix, err error) {
	return CreateSharpenFilter(radius, imagex.Horizontal|imagex.Vertical, 0)
}

// 创建拉普拉斯8邻锐化滤波器
func CreateSharpenLaplace8Filter(radius int) (filter FilterMatrix, err error) {
	return CreateSharpenFilter(radius, imagex.AllDirection, 0)
}

//-----------------------------------------------

// 使用拉普拉斯4邻锐化滤波器处理图像
func SharpenWithLaplace4(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Sharpen3Laplace4)
}

// 使用拉普拉斯4邻锐化滤波器处理图像
func SharpenWithLaplace8(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Sharpen3Laplace8)
}

// 使用强调边缘滤波器处理图像
func SharpenWithStrengthen3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, SharpenStrengthen3All)
}

// 使用强调边缘滤波器处理图像
func SharpenWithSharpen5All(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Sharpen5All)
}

//-----------------------------------------------

// 使用定制的拉普拉斯锐化滤波器处理图像
func SharpenWithCustomLaplace(srcImg image.Image, dstImg draw.Image, radius int, direction imagex.PixelDirection) error {
	filter, err := CreateSharpenFilter(radius, direction, 0)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}
