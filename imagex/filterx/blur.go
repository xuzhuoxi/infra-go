// 模糊
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/mathx/gaussx"
	"image"
	"image/draw"
)

// 均值模糊滤波器
var (
	//3x3 4邻域均值滤波器
	BoxFourNear3 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 4,
		Kernel: []KernelVector{
			{0, -1, 1}, {-1, 0, 1},
			{1, +0, 1}, {+0, 1, 1}, {0, 0, 1}}}
	//3x3 8邻域均值滤波器
	BoxEightNear3 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 8,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
	//3x3 均值滤波器
	BoxAverage3 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 9,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {0, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
)

//高斯模糊滤波器
var (
	//3x3 高斯滤波器
	Gauss3 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 16,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, -1, 2}, {1, -1, 1},
			{-1, +0, 2}, {0, +0, 4}, {1, +0, 2},
			{-1, +1, 1}, {0, +1, 2}, {1, +1, 1}}}
	//5x5 高斯滤波器
	Gauss5 = FilterMatrix{Radius: 2, Size: 5, KernelScale: 273,
		Kernel: []KernelVector{
			{-2, -2, 1}, {-1, -2, 4}, {0, -2, 7}, {1, -2, 4}, {2, -2, 1},
			{-2, -1, 4}, {-1, -1, 16}, {0, -1, 26}, {1, -1, 16}, {2, -1, 4},
			{-2, +0, 7}, {-1, +0, 26}, {0, +0, 41}, {1, +0, 26}, {2, +0, 7},
			{-2, +1, 4}, {-1, +1, 16}, {0, +1, 26}, {1, +1, 16}, {2, +1, 4},
			{-2, +2, 1}, {-1, +2, 4}, {1, +2, 7}, {1, +2, 4}, {2, +2, 1}}}
)

// 运动模糊滤波器
var (
	//3x3水平运动滤波器
	Motion3Horizontal = FilterMatrix{Radius: 1, Size: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{-1, 0, 1}, {0, 0, 1}, {1, 0, 1}}}
	//3x3垂直运动滤波器
	Motion3Vertical = FilterMatrix{Radius: 1, Size: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{0, -1, 1}, {0, 0, 1}, {0, 1, 1}}}
	//3x3 45度运动滤波器(左上右下)
	Motion3Oblique45 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{1, -1, 1}, {0, +0, 1}, {-1, +1, 1}}}
	//3x3 135度运动滤波器(左下右上)
	Motion3Oblique135 = FilterMatrix{Radius: 1, Size: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, 0, 1}, {1, 1, 1}}}
)

//-------------------------------------------

// 创建高斯模糊滤波器
// radius：	卷积核半径 [1，3]
// sigma:	标准差
func CreateGaussBlurFilter(radius int, sigma float64) (filter *FilterMatrix, err error) {
	if radius < 1 {
		return nil, errors.New("Radius < 1. ")
	}
	kSize := radius + radius + 1
	gaussKernel := gaussx.CreateGaussKernelInt(radius, sigma, 0)
	var kernel = make([]KernelVector, 0, kSize*kSize)
	var scale = 0
	var value int
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			value = gaussKernel[y+radius][x+radius]
			if value != 0 {
				scale += value
				kernel = append(kernel, KernelVector{X: x, Y: y, Value: value})
			}
		}
	}
	return &FilterMatrix{Radius: radius, Size: kSize, KernelScale: scale, Kernel: kernel}, nil
}

// 创建运动模糊滤波器
// radius：		卷积核半径
// direction: 	运动方向
func CreateMotionBlurFilter(radius int, direction imagex.PixelDirection) (filter *FilterMatrix, err error) {
	if radius < 1 {
		return nil, errors.New("Radius < 1. ")
	}
	dirAdds := imagex.GetPixelDirectionAdds(direction)
	if nil == dirAdds {
		return nil, errors.New("Direction Error. ")
	}
	kSize := radius + radius + 1
	scale := len(dirAdds)*radius + 1
	var kernel = make([]KernelVector, 0, scale)
	kernel = append(kernel, KernelVector{X: 0, Y: 0, Value: 1})
	for _, add := range dirAdds {
		for i := 1; i <= radius; i++ {
			kernel = append(kernel, KernelVector{X: add.X * i, Y: add.Y * i, Value: 1})
		}
	}
	return &FilterMatrix{Radius: radius, Size: kSize, KernelScale: scale, Kernel: kernel}, nil
}

//------------------------------------------------------

// 4邻域均值模糊
func BlurWithForeNear3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, BoxFourNear3)
}

// 8邻域均值模糊
func BlurWithEightNear3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, BoxEightNear3)
}

// 均值模糊
func BlurWithAverage3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, BoxAverage3)
}

// 高斯3x3模糊
func BlurWithGauss3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Gauss3)
}

// 高斯5x5模糊
func BlurWithGauss5(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Gauss5)
}

// 水平运动模糊
func BlurWithMotion3Horizontal(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Horizontal)
}

// 垂直运动模糊
func BlurWithMotion3Vertical(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Vertical)
}

// 斜45角运动模糊
func BlurWithMotion3Oblique45(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Oblique45)
}

// 斜135角运动模糊
func BlurWithMotion3Oblique135(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Oblique135)
}

// 自定义运动模糊
func BlurWithMotion(srcImg image.Image, dstImg draw.Image, radius int, direction imagex.PixelDirection) error {
	filter, err := CreateMotionBlurFilter(radius, direction)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, *filter)
}
