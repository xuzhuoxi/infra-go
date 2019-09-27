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
	//4邻域均值滤波器
	FourNear3 = FilterTemplate{Radius: 1, Size: 3, Scale: 4,
		Offsets: []FilterOffset{
			{0, -1, 1}, {-1, 0, 1},
			{1, +0, 1}, {+0, 1, 1}, {0, 0, 1}}}
	//8邻域均值滤波器
	EightNear3 = FilterTemplate{Radius: 1, Size: 3, Scale: 8,
		Offsets: []FilterOffset{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
	//均值滤波器
	Average3 = FilterTemplate{Radius: 1, Size: 3, Scale: 9,
		Offsets: []FilterOffset{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {0, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
)

//高斯模糊滤波器
var (
	//3x3高斯滤波器
	Gauss3 = FilterTemplate{Radius: 1, Size: 3, Scale: 16,
		Offsets: []FilterOffset{
			{-1, -1, 1}, {0, -1, 2}, {1, -1, 1},
			{-1, +0, 2}, {0, +0, 4}, {1, +0, 2},
			{-1, +1, 1}, {0, +1, 2}, {1, +1, 1}}}
	//5x5高斯滤波器
	Gauss5 = FilterTemplate{Radius: 2, Size: 5, Scale: 273,
		Offsets: []FilterOffset{
			{-2, -2, 1}, {-1, -2, 4}, {0, -2, 7}, {1, -2, 4}, {2, -2, 1},
			{-2, -1, 4}, {-1, -1, 16}, {0, -1, 26}, {1, -1, 16}, {2, -1, 4},
			{-2, +0, 7}, {-1, +0, 26}, {0, +0, 41}, {1, +0, 26}, {2, +0, 7},
			{-2, +1, 4}, {-1, +1, 16}, {0, +1, 26}, {1, +1, 16}, {2, +1, 4},
			{-2, +2, 1}, {-1, +2, 4}, {1, +2, 7}, {1, +2, 4}, {2, +2, 1}}}
)

// 运动滤波器
var (
	//3x3水平运动滤波器
	Motion3Horizontal = FilterTemplate{Radius: 1, Size: 3, Scale: 3,
		Offsets: []FilterOffset{
			{-1, 0, 1}, {0, 0, 1}, {1, 0, 1}}}
	//3x3垂直运动滤波器
	Motion3Vertical = FilterTemplate{Radius: 1, Size: 3, Scale: 3,
		Offsets: []FilterOffset{
			{0, -1, 1}, {0, 0, 1}, {0, 1, 1}}}
	//3x3 45度运动滤波器(左上右下)
	Motion3Oblique45 = FilterTemplate{Radius: 1, Size: 3, Scale: 3,
		Offsets: []FilterOffset{
			{1, -1, 1}, {0, +0, 1}, {-1, +1, 1}}}
	//3x3 135度运动滤波器(左下右上)
	Motion3Oblique135 = FilterTemplate{Radius: 1, Size: 3, Scale: 3,
		Offsets: []FilterOffset{
			{-1, -1, 1}, {0, 0, 1}, {1, 1, 1}}}
)

// 创建高斯模糊滤波器
// radius：	卷积核半径 [1，3]
// sigma:	标准差
func CreateGaussBlurFilter(radius int, sigma float64) (filter *FilterTemplate, err error) {
	if radius < 1 {
		return nil, errors.New("Radius < 1. ")
	}
	kSize := radius*2 + 1
	kernel := gaussx.CreateGaussKernelInt(radius, sigma, 0)
	var offsets = make([]FilterOffset, 0, kSize*kSize)
	var scale = 0
	var value int
	for y := -kSize; y <= kSize; y++ {
		for x := -kSize; x <= kSize; x++ {
			value = kernel[y+kSize][x+kSize]
			if value != 0 {
				scale += value
				offsets = append(offsets, FilterOffset{X: x, Y: y, Value: value})
			}
		}
	}
	return &FilterTemplate{Radius: radius, Size: kSize, Scale: 0, Offsets: offsets}, nil
}

// 创建运动模糊滤波器
// radius：		卷积核半径
// direction: 	运动方向
func CreateMotionBlurFilter(radius int, direction imagex.PixelDirection) (filter *FilterTemplate, err error) {
	if radius < 1 {
		return nil, errors.New("Radius < 1. ")
	}
	dirAdds := imagex.GetPixelDirectionAdds(direction)
	if nil == dirAdds {
		return nil, errors.New("Direction Error. ")
	}
	kSize := radius*2 + 1
	scale := len(dirAdds)*radius + 1
	var offsets = make([]FilterOffset, 0, scale)
	offsets = append(offsets, FilterOffset{X: 0, Y: 0, Value: 1})
	for _, add := range dirAdds {
		for i := 1; i <= radius; i++ {
			offsets = append(offsets, FilterOffset{X: add.X * i, Y: add.Y * i, Value: 1})
		}
	}
	return &FilterTemplate{Radius: radius, Size: kSize, Scale: scale, Offsets: offsets}, nil
}

//------------------------------------------------------

// 4邻域均值模糊
func BlurWithForeNear3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, FourNear3)
}

// 8邻域均值模糊
func BlurWithEightNear3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, EightNear3)
}

// 均值模糊
func BlurWithAverage3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, Average3)
}

// 高斯3x3模糊
func BlurWithGauss3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, Gauss3)
}

// 高斯5x5模糊
func BlurWithGauss5(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, Gauss5)
}

// 水平运动模糊
func BlurWithMotion3Horizontal(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, Motion3Horizontal)
}

// 垂直运动模糊
func BlurWithMotion3Vertical(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, Motion3Vertical)
}

// 斜45角运动模糊
func BlurWithMotion3Oblique45(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, Motion3Oblique45)
}

// 斜135角运动模糊
func BlurWithMotion3Oblique135(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithTemplate(srcImg, dstImg, Motion3Oblique135)
}

// 自定义运动模糊
func BlurWithMotion(srcImg image.Image, dstImg draw.Image, radius int, direction imagex.PixelDirection) error {
	filter, err := CreateMotionBlurFilter(radius, direction)
	if nil != err {
		return err
	}
	return FilterImageWithTemplate(srcImg, dstImg, *filter)
}
