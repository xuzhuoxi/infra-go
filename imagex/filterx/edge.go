// 边缘化
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/draw"
)

// 边缘滤波器
var (
	//5x5 左右边缘滤波器
	Edge5LR = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-2, 0, -1}, {-1, 0, -1}, {0, 0, 4}, {1, 0, -1}, {2, 0, -1}}}
	//5x5 上下边缘滤波器
	Edge5UD = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{0, -2, -1}, {0, -1, -1}, {0, 0, 4}, {0, 1, -1}, {0, 2, -1}}}
	//5x5 左上右下边缘滤波器
	Edge5LuRd = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -1, -1}, {0, 0, 4}, {1, 1, -1}, {1, 2, -1}}}
	//5x5 左下右上边缘滤波器
	Edge5LdRu = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-2, 2, -1}, {-1, 1, -1}, {0, 0, 4}, {1, -1, -1}, {2, -2, -1}}}

	//3x3 Laplace拉普拉斯4邻滤波器
	Edge3Laplace4 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{0, -1, -1}, {-1, +0, -1}, {0, +0, +4}, {1, +0, -1}, {0, +1, -1}}}
	//5x5 Laplace拉普拉斯8邻滤波器
	Edge3Laplace8 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-1, -1, -1}, {0, -1, -1}, {1, -1, -1},
			{-1, +0, -1}, {0, +0, +8}, {1, +0, -1},
			{-1, +1, -1}, {0, +1, -1}, {1, +1, -1}}}
)

// 创建边缘检测滤波器
// radius：		卷积核半径
// direction: 	运动方向
// diff:		梯度差
func CreateEdgeFilter(radius int, direction imagex.PixelDirection, diff uint) (filter FilterMatrix, err error) {
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
	kernel = append(kernel, KernelVector{X: 0, Y: 0, Value: -sumValue})
	return FilterMatrix{KernelRadius: radius, KernelSize: kSize, KernelScale: 0, Kernel: kernel}, nil
}

// 创建拉普拉斯4邻边缘检测滤波器
func CreateEdgeLaplace4Filter(radius int) (filter FilterMatrix, err error) {
	return CreateEdgeFilter(radius, imagex.Horizontal|imagex.Vertical, 0)
}

// 创建拉普拉斯8邻边缘检测滤波器
func CreateEdgeLaplace8Filter(radius int) (filter FilterMatrix, err error) {
	return CreateEdgeFilter(radius, imagex.AllDirection, 0)
}

//-----------------------------------------------

// 使用左右边缘滤波器处理图像
func EdgeWithLR(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5LR)
}

// 使用上下边缘滤波器处理图像
func EdgeWithUD(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5UD)
}

// 使用左上右下边缘滤波器处理图像
func EdgeWithLuRd(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5LuRd)
}

// 使用左下右上边缘滤波器处理图像
func EdgeWithLdRu(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5LdRu)
}

// 使用全方向边缘滤波器处理图像
func EdgeWithLaplace4(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge3Laplace4)
}

// 使用全方向边缘滤波器处理图像
func EdgeWithLaplace8(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge3Laplace8)
}

//------------------------------------------------

// 使用自定义边缘滤波器处理图像
func EdgeCustom(srcImg image.Image, dstImg draw.Image, radius int, direction imagex.PixelDirection, diff uint) error {
	filter, err := CreateEdgeFilter(radius, direction, diff)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}

// 使用自定义拉普拉斯4邻边缘滤波器处理图像
func EdgeWithCustomLaplace4(srcImg image.Image, dstImg draw.Image, radius int) error {
	filter, err := CreateEdgeLaplace4Filter(radius)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}

// 使用自定义拉普拉斯8邻边缘滤波器处理图像
func EdgeWithCustomLaplace8(srcImg image.Image, dstImg draw.Image, radius int) error {
	filter, err := CreateEdgeLaplace8Filter(radius)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}
