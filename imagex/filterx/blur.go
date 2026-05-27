package filterx

import (
	"errors"
	"image"
	"image/draw"

	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/mathx/gaussx"
)

// 均值模糊滤波器
var (
	// BoxFourNear3 is a 3×3 four-neighbor box blur kernel.
	// BoxFourNear3 是 3×3 四邻域均值滤波器。
	BoxFourNear3 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 4,
		Kernel: []KernelVector{
			{0, -1, 1}, {-1, 0, 1},
			{1, +0, 1}, {+0, 1, 1}, {0, 0, 1}}}
	// BoxEightNear3 is a 3×3 eight-neighbor box blur kernel.
	// BoxEightNear3 是 3×3 八邻域均值滤波器。
	BoxEightNear3 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 8,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
	// BoxAverage3 is a 3×3 full box average blur kernel.
	// BoxAverage3 是 3×3 全邻域均值滤波器。
	BoxAverage3 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 9,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {0, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
)

// 高斯模糊滤波器
var (
	// Gauss3 is a 3×3 Gaussian blur kernel.
	// Gauss3 是 3×3 高斯滤波器。
	Gauss3 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 16,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, -1, 2}, {1, -1, 1},
			{-1, +0, 2}, {0, +0, 4}, {1, +0, 2},
			{-1, +1, 1}, {0, +1, 2}, {1, +1, 1}}}
	// Gauss5 is a 5×5 Gaussian blur kernel.
	// Gauss5 是 5×5 高斯滤波器。
	Gauss5 = FilterMatrix{KernelRadius: 2, KernelSize: 5, KernelScale: 273,
		Kernel: []KernelVector{
			{-2, -2, 1}, {-1, -2, 4}, {0, -2, 7}, {1, -2, 4}, {2, -2, 1},
			{-2, -1, 4}, {-1, -1, 16}, {0, -1, 26}, {1, -1, 16}, {2, -1, 4},
			{-2, +0, 7}, {-1, +0, 26}, {0, +0, 41}, {1, +0, 26}, {2, +0, 7},
			{-2, +1, 4}, {-1, +1, 16}, {0, +1, 26}, {1, +1, 16}, {2, +1, 4},
			{-2, +2, 1}, {-1, +2, 4}, {1, +2, 7}, {1, +2, 4}, {2, +2, 1}}}
)

// 运动模糊滤波器
var (
	// Motion3Horizontal is a 3×3 horizontal motion blur kernel.
	// Motion3Horizontal 是 3×3 水平运动滤波器。
	Motion3Horizontal = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{-1, 0, 1}, {0, 0, 1}, {1, 0, 1}}}
	// Motion3Vertical is a 3×3 vertical motion blur kernel.
	// Motion3Vertical 是 3×3 垂直运动滤波器。
	Motion3Vertical = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{0, -1, 1}, {0, 0, 1}, {0, 1, 1}}}
	// Motion3Oblique45 is a 3×3 45° motion blur kernel (upper-left to lower-right).
	// Motion3Oblique45 是 3×3 45° 运动滤波器（左上至右下）。
	Motion3Oblique45 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{1, -1, 1}, {0, +0, 1}, {-1, +1, 1}}}
	// Motion3Oblique135 is a 3×3 135° motion blur kernel (lower-left to upper-right).
	// Motion3Oblique135 是 3×3 135° 运动滤波器（左下至右上）。
	Motion3Oblique135 = FilterMatrix{KernelRadius: 1, KernelSize: 3, KernelScale: 3,
		Kernel: []KernelVector{
			{-1, -1, 1}, {0, 0, 1}, {1, 1, 1}}}
)

//-------------------------------------------

// CreateBoxFourNearBlurFilter builds a custom four-neighbor box blur kernel.
// CreateBoxFourNearBlurFilter 创建自定义四邻域均值模糊滤波器。
// radius: kernel radius.
// radius：卷积核半径。
// includeCenter: whether the center pixel is included.
// includeCenter：是否包含中心像素。
func CreateBoxFourNearBlurFilter(radius int, includeCenter bool) (filter FilterMatrix, err error) {
	if radius < 1 {
		return filter, errors.New("KernelRadius < 1. ")
	}
	kSize := radius + radius + 1
	ln := 2*kSize + 3
	if !includeCenter {
		ln -= 1
	}
	var kernel FilterKernel = make([]KernelVector, 0, ln)
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if !includeCenter && 0 == x && 0 == y {
				continue
			}
			if mathx.AbsInt(x)+mathx.AbsInt(y) > radius {
				continue
			}
			kernel = append(kernel, KernelVector{X: x, Y: y, Value: 1})
		}
	}
	return FilterMatrix{KernelRadius: radius, KernelSize: kSize, KernelScale: ln, Kernel: kernel}, nil
}

// CreateBoxEightNearBlurFilter builds a custom eight-neighbor box blur kernel.
// CreateBoxEightNearBlurFilter 创建自定义八邻域均值模糊滤波器。
// radius: kernel radius.
// radius：卷积核半径。
// includeCenter: whether the center pixel is included.
// includeCenter：是否包含中心像素。
func CreateBoxEightNearBlurFilter(radius int, includeCenter bool) (filter FilterMatrix, err error) {
	if radius < 1 {
		return filter, errors.New("KernelRadius < 1. ")
	}
	kSize := radius + radius + 1
	ln := kSize * kSize
	if !includeCenter {
		ln -= 1
	}
	var kernel FilterKernel = make([]KernelVector, 0, ln)
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if !includeCenter && 0 == x && 0 == y {
				continue
			}
			kernel = append(kernel, KernelVector{X: x, Y: y, Value: 1})
		}
	}
	return FilterMatrix{KernelRadius: radius, KernelSize: kSize, KernelScale: ln, Kernel: kernel}, nil
}

// CreateGaussBlurFilter builds a Gaussian blur kernel.
// CreateGaussBlurFilter 创建高斯模糊滤波器。
// radius: kernel radius in [1, 3].
// radius：卷积核半径 [1，3]。
// sigma: standard deviation.
// sigma：标准差。
func CreateGaussBlurFilter(radius int, sigma float64) (filter FilterMatrix, err error) {
	if radius < 1 {
		return filter, errors.New("KernelRadius < 1. ")
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
	return FilterMatrix{KernelRadius: radius, KernelSize: kSize, KernelScale: scale, Kernel: kernel}, nil
}

// CreateMotionBlurFilter builds a directional motion blur kernel.
// CreateMotionBlurFilter 创建运动模糊滤波器。
// radius: kernel radius.
// radius：卷积核半径。
// direction: motion direction.
// direction：运动方向。
func CreateMotionBlurFilter(radius int, direction imagex.PixelDirection) (filter FilterMatrix, err error) {
	if radius < 1 {
		return filter, errors.New("KernelRadius < 1. ")
	}
	dirAdds := imagex.GetPixelDirectionAdds(direction)
	if nil == dirAdds {
		return filter, errors.New("Direction Error. ")
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
	return FilterMatrix{KernelRadius: radius, KernelSize: kSize, KernelScale: scale, Kernel: kernel}, nil
}

//------------------------------------------------------

// BlurWithForeNear3 applies the 3×3 four-neighbor box blur.
// BlurWithForeNear3 使用 3×3 四邻域均值模糊。
func BlurWithForeNear3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, BoxFourNear3)
}

// BlurWithEightNear3 applies the 3×3 eight-neighbor box blur.
// BlurWithEightNear3 使用 3×3 八邻域均值模糊。
func BlurWithEightNear3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, BoxEightNear3)
}

// BlurWithAverage3 applies the 3×3 full box average blur.
// BlurWithAverage3 使用 3×3 全邻域均值模糊。
func BlurWithAverage3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, BoxAverage3)
}

// BlurWithBoxFour applies a custom four-neighbor box blur.
// BlurWithBoxFour 使用自定义四邻域均值模糊。
func BlurWithBoxFour(srcImg image.Image, dstImg draw.Image, radius int, includeCenter bool) error {
	filter, err := CreateBoxFourNearBlurFilter(radius, includeCenter)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}

// BlurWithBoxEight applies a custom eight-neighbor box blur.
// BlurWithBoxEight 使用自定义八邻域均值模糊。
func BlurWithBoxEight(srcImg image.Image, dstImg draw.Image, radius int, includeCenter bool) error {
	filter, err := CreateBoxEightNearBlurFilter(radius, includeCenter)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}

//------------------------------------------------------------------

// BlurWithGauss3 applies the preset 3×3 Gaussian blur.
// BlurWithGauss3 使用预设 3×3 高斯模糊。
func BlurWithGauss3(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Gauss3)
}

// BlurWithGauss5 applies the preset 5×5 Gaussian blur.
// BlurWithGauss5 使用预设 5×5 高斯模糊。
func BlurWithGauss5(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Gauss5)
}

// BlurWithGauss applies a custom Gaussian blur.
// BlurWithGauss 使用自定义高斯模糊。
func BlurWithGauss(srcImg image.Image, dstImg draw.Image, radius int, sigma float64) error {
	filter, err := CreateGaussBlurFilter(radius, sigma)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}

//------------------------------------------------------------------

// BlurWithMotion3Horizontal applies the 3×3 horizontal motion blur.
// BlurWithMotion3Horizontal 使用 3×3 水平运动模糊。
func BlurWithMotion3Horizontal(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Horizontal)
}

// BlurWithMotion3Vertical applies the 3×3 vertical motion blur.
// BlurWithMotion3Vertical 使用 3×3 垂直运动模糊。
func BlurWithMotion3Vertical(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Vertical)
}

// BlurWithMotion3Oblique45 applies the 3×3 45° motion blur.
// BlurWithMotion3Oblique45 使用 3×3 斜 45° 运动模糊。
func BlurWithMotion3Oblique45(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Oblique45)
}

// BlurWithMotion3Oblique135 applies the 3×3 135° motion blur.
// BlurWithMotion3Oblique135 使用 3×3 斜 135° 运动模糊。
func BlurWithMotion3Oblique135(srcImg image.Image, dstImg draw.Image) error {
	return FilterImageWithMatrix(srcImg, dstImg, Motion3Oblique135)
}

// BlurWithMotion applies a custom directional motion blur.
// BlurWithMotion 使用自定义方向运动模糊。
func BlurWithMotion(srcImg image.Image, dstImg draw.Image, radius int, direction imagex.PixelDirection) error {
	filter, err := CreateMotionBlurFilter(radius, direction)
	if nil != err {
		return err
	}
	return FilterImageWithMatrix(srcImg, dstImg, filter)
}
