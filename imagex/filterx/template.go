package filterx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"image"
	"image/color"
	"image/draw"
)

// 滤波器向量单元
type FilterOffset struct {
	//向量X
	X int
	//向量Y
	Y int
	//向量值
	Value int
}

// 滤波器
type FilterTemplate struct {
	// 滤波器半径
	Radius int
	// 滤波器边长
	Size int
	// 滤波器核
	Offsets []FilterOffset
	// 滤波器卷积核倍率
	Scale int
}

// 滤波模板有效性
func (t FilterTemplate) CheckValidity() error {
	if t.Radius < 1 {
		return errors.New("Radius < 1. ")
	}
	if t.Scale <= 0 {
		return errors.New("Scale <= 0. ")
	}
	sum := 0
	for _, offset := range t.Offsets {
		sum += offset.Value
	}
	if sum != t.Scale {
		return errors.New(fmt.Sprintf("Scale Error! Scale=%d, AddSum=%d.", t.Scale, sum))
	}
	return nil
}

//使用滤波器
func FilterImageWithTemplate(srcImg image.Image, dstImg draw.Image, template FilterTemplate) error {
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
			if template.Scale != 0 && template.Scale != 1 {
				sumR, sumG, sumB, sumA = sumR/template.Scale, sumG/template.Scale, sumB/template.Scale, sumA/template.Scale
			}
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
		if sumValue != 0 && sumValue != 1 {
			sumR, sumG, sumB, sumA = sumR/sumValue, sumG/sumValue, sumB/sumValue, sumA/sumValue
		}
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
