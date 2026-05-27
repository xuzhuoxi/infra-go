package filterx

import (
	"errors"
	"image"
	"image/draw"

	"github.com/xuzhuoxi/infra-go/imagex"
)

// 边缘滤波器
var (
	// Edge5LR is a left-right edge detection kernel.
	// Edge5LR 是左右边缘检测滤波器。
	Edge5LR = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-2, 0, -1}, {-1, 0, -1}, {0, 0, 4}, {1, 0, -1}, {2, 0, -1}}}
	// Edge5UD is an up-down edge detection kernel.
	// Edge5UD 是上下边缘检测滤波器。
	Edge5UD = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{0, -2, -1}, {0, -1, -1}, {0, 0, 4}, {0, 1, -1}, {0, 2, -1}}}
	// Edge5LuRd is an upper-left to lower-right edge detection kernel.
	// Edge5LuRd 是左上至右下边缘检测滤波器。
	Edge5LuRd = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-2, -2, -1}, {-1, -1, -1}, {0, 0, 4}, {1, 1, -1}, {1, 2, -1}}}
	// Edge5LdRu is a lower-left to upper-right edge detection kernel.
	// Edge5LdRu 是左下至右上边缘检测滤波器。
	Edge5LdRu = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-2, 2, -1}, {-1, 1, -1}, {0, 0, 4}, {1, -1, -1}, {2, -2, -1}}}

	// Edge3Laplace4 is a 3×3 four-neighbor Laplacian edge kernel.
	// Edge3Laplace4 是 3×3 四邻域拉普拉斯边缘检测滤波器。
	Edge3Laplace4 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{0, -1, -1}, {-1, +0, -1}, {0, +0, +4}, {1, +0, -1}, {0, +1, -1}}}
	// Edge3Laplace8 is a 3×3 eight-neighbor Laplacian edge kernel.
	// Edge3Laplace8 是 3×3 八邻域拉普拉斯边缘检测滤波器。
	Edge3Laplace8 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 0,
		Kernel: []KernelVector{
			{-1, -1, -1}, {0, -1, -1}, {1, -1, -1},
			{-1, +0, -1}, {0, +0, +8}, {1, +0, -1},
			{-1, +1, -1}, {0, +1, -1}, {1, +1, -1}}}
)

// CreateEdgeFilter builds a custom edge detection kernel.
// CreateEdgeFilter 创建边缘检测滤波器。
// radius: kernel radius.
// radius：卷积核半径。
// direction: pixel direction.
// direction：像素方向。
// diff: gradient difference.
// diff：梯度差。
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

// CreateEdgeLaplace4Filter builds a four-neighbor Laplacian edge kernel.
// CreateEdgeLaplace4Filter 创建拉普拉斯四邻域边缘检测滤波器。
func CreateEdgeLaplace4Filter(radius int) (filter FilterMatrix, err error) {
	return CreateEdgeFilter(radius, imagex.Horizontal|imagex.Vertical, 0)
}

// CreateEdgeLaplace8Filter builds an eight-neighbor Laplacian edge kernel.
// CreateEdgeLaplace8Filter 创建拉普拉斯八邻域边缘检测滤波器。
func CreateEdgeLaplace8Filter(radius int) (filter FilterMatrix, err error) {
	return CreateEdgeFilter(radius, imagex.AllDirection, 0)
}

//-----------------------------------------------

// EdgeWithLR applies the left-right edge kernel to an image.
// EdgeWithLR 使用左右边缘滤波器处理图像。
func EdgeWithLR(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5LR)
}

// EdgeWithUD applies the up-down edge kernel to an image.
// EdgeWithUD 使用上下边缘滤波器处理图像。
func EdgeWithUD(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5UD)
}

// EdgeWithLuRd applies the upper-left to lower-right edge kernel to an image.
// EdgeWithLuRd 使用左上至右下边缘滤波器处理图像。
func EdgeWithLuRd(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5LuRd)
}

// EdgeWithLdRu applies the lower-left to upper-right edge kernel to an image.
// EdgeWithLdRu 使用左下至右上边缘滤波器处理图像。
func EdgeWithLdRu(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge5LdRu)
}

// EdgeWithLaplace4 applies the four-neighbor Laplacian edge kernel to an image.
// EdgeWithLaplace4 使用拉普拉斯四邻域边缘滤波器处理图像。
func EdgeWithLaplace4(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge3Laplace4)
}

// EdgeWithLaplace8 applies the eight-neighbor Laplacian edge kernel to an image.
// EdgeWithLaplace8 使用拉普拉斯八邻域边缘滤波器处理图像。
func EdgeWithLaplace8(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Edge3Laplace8)
}

//------------------------------------------------

// EdgeCustom applies a custom edge kernel to an image.
// EdgeCustom 使用自定义边缘滤波器处理图像。
func EdgeCustom(srcImg image.Image, dstImg draw.Image, radius int, direction imagex.PixelDirection, diff uint) error {
	filter, err := CreateEdgeFilter(radius, direction, diff)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}

// EdgeWithCustomLaplace4 applies a custom four-neighbor Laplacian edge kernel to an image.
// EdgeWithCustomLaplace4 使用自定义拉普拉斯四邻域边缘滤波器处理图像。
func EdgeWithCustomLaplace4(srcImg image.Image, dstImg draw.Image, radius int) error {
	filter, err := CreateEdgeLaplace4Filter(radius)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}

// EdgeWithCustomLaplace8 applies a custom eight-neighbor Laplacian edge kernel to an image.
// EdgeWithCustomLaplace8 使用自定义拉普拉斯八邻域边缘滤波器处理图像。
func EdgeWithCustomLaplace8(srcImg image.Image, dstImg draw.Image, radius int) error {
	filter, err := CreateEdgeLaplace8Filter(radius)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}
