// 模糊
package filterx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
	"image/draw"
)

type BlurOffset struct {
	X, Y  int
	Value int
}

type BlurTemplate struct {
	Radius, Size int
	Offsets      []BlurOffset
	Sum          int
}

func (t BlurTemplate) CheckValidity() error {
	if t.Radius < 1 {
		return errors.New("Radius < 1. ")
	}
	if t.Sum <= 0 {
		return errors.New("Sum <= 0. ")
	}
	sum := 0
	for _, offset := range t.Offsets {
		sum += offset.Value
	}
	if sum != t.Sum {
		return errors.New(fmt.Sprintf("Sum Error! Sum=%d, AddSum=%d.", t.Sum, sum))
	}
	return nil
}

var (
	FourNear3 = BlurTemplate{Radius: 1, Size: 3, Sum: 4,
		Offsets: []BlurOffset{
			{0, -1, 1}, {-1, 0, 1},
			{1, +0, 1}, {+0, 1, 1}}}
	EightNear3 = BlurTemplate{Radius: 1, Size: 3, Sum: 8,
		Offsets: []BlurOffset{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
	Average3 = BlurTemplate{Radius: 1, Size: 3, Sum: 9,
		Offsets: []BlurOffset{
			{-1, -1, 1}, {0, -1, 1}, {1, -1, 1},
			{-1, +0, 1}, {0, +0, 1}, {1, +0, 1},
			{-1, +1, 1}, {0, +1, 1}, {1, +1, 1}}}
	Gauss3 = BlurTemplate{Radius: 1, Size: 3, Sum: 16,
		Offsets: []BlurOffset{
			{-1, -1, 1}, {0, -1, 2}, {1, -1, 1},
			{-1, +0, 2}, {0, +0, 4}, {1, +0, 2},
			{-1, +1, 1}, {0, +1, 2}, {1, +1, 1}}}
	Gauss5 = BlurTemplate{Radius: 2, Size: 5, Sum: 273,
		Offsets: []BlurOffset{
			{-2, -2, 1}, {-1, -2, 4}, {0, -2, 7}, {1, -2, 4}, {2, -2, 1},
			{-2, -1, 4}, {-1, -1, 16}, {0, -1, 26}, {1, -1, 16}, {2, -1, 4},
			{-2, +0, 7}, {-1, +0, 26}, {0, +0, 41}, {1, +0, 26}, {2, +0, 7},
			{-2, +1, 4}, {-1, +1, 16}, {0, +1, 26}, {1, +1, 16}, {2, +1, 4},
			{-2, +2, 1}, {-1, +2, 4}, {1, +2, 7}, {1, +2, 4}, {2, +2, 1}}}
)

// radius：	卷积核半径 [1，3]
// sigma:	标准差
func CreateGaussBlurTemplate(radius int, sigma float64) *BlurTemplate {
	kSize := radius*2 + 1
	return &BlurTemplate{Radius: radius, Size: kSize, Sum: 0, Offsets: nil}
}

//使用模板进行滤波模糊
func BlurWithTemplate(srcImg image.Image, dstImg draw.Image, template BlurTemplate) error {
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
	radius := template.Radius
	var x, y int
	var sumR, sumG, sumB, sumA int
	var R, G, B, A uint32
	var setColor = &color.NRGBA64{}
	//内部处理
	for y = radius; y < size.Y-radius; y++ {
		for x = radius; x < size.X-radius; x++ {
			sumR, sumG, sumB, sumA = 0, 0, 0, 0
			for _, offset := range template.Offsets {
				R, G, B, A = sourceImage.At(x+offset.X, y+offset.Y).RGBA()
				sumR += int(R) * offset.Value
				sumG += int(G) * offset.Value
				sumB += int(B) * offset.Value
				sumA += int(A) * offset.Value
			}
			sumR, sumG, sumB, sumA = sumR/template.Sum, sumG/template.Sum, sumB/template.Sum, sumA/template.Sum
			setColor.R, setColor.G, setColor.B, setColor.A = uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)
			targetImage.Set(x, y, setColor)
		}
	}
	//边缘处理
	handleTemp := func(x, y int) {
		sumR, sumG, sumB, sumA = 0, 0, 0, 0
		sumValue := 0
		var cx, cy int
		for _, offset := range template.Offsets {
			cx = x + offset.X
			cy = y + offset.Y
			if cx < 0 || cx >= size.X || cy < 0 || cy >= size.Y {
				continue
			}
			R, G, B, A = sourceImage.At(x+offset.X, y+offset.Y).RGBA()
			sumR += int(R) * offset.Value
			sumG += int(G) * offset.Value
			sumB += int(B) * offset.Value
			sumA += int(A) * offset.Value
			sumValue += offset.Value
		}
		sumR, sumG, sumB, sumA = sumR/sumValue, sumG/sumValue, sumB/sumValue, sumA/sumValue
		setColor.R, setColor.G, setColor.B, setColor.A = uint16(sumR), uint16(sumG), uint16(sumB), uint16(sumA)
		dstImg.Set(x, y, setColor)
	}
	for y = 0; y < radius; y++ {
		for x = 0; x < size.X; x++ {
			handleTemp(x, y)
		}
	}
	for y = size.Y - 1 - radius; y < size.Y-1; y++ {
		for x = 0; x < size.X; x++ {
			handleTemp(x, y)
		}
	}
	for x := 0; x < radius; x++ {
		for y = 0; y < size.Y; y++ {
			handleTemp(x, y)
		}
	}
	for x := size.X - 1 - radius; x < size.X-1; x++ {
		for y = 0; y < size.Y; y++ {
			handleTemp(x, y)
		}
	}
	return nil
}

func BlurWithForeNear3(srcImg image.Image, dstImg draw.Image) error {
	return BlurWithTemplate(srcImg, dstImg, FourNear3)
}

func BlurWithEightNear3(srcImg image.Image, dstImg draw.Image) error {
	return BlurWithTemplate(srcImg, dstImg, EightNear3)
}

func BlurWithAverage3(srcImg image.Image, dstImg draw.Image) error {
	return BlurWithTemplate(srcImg, dstImg, Average3)
}

func BlurWithGauss3(srcImg image.Image, dstImg draw.Image) error {
	return BlurWithTemplate(srcImg, dstImg, Gauss3)
}

func BlurWithGauss5(srcImg image.Image, dstImg draw.Image) error {
	return BlurWithTemplate(srcImg, dstImg, Gauss5)
}
