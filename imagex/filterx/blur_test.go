package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"reflect"
	"testing"
)

func TestPoint(t *testing.T) {
	var pix = make([]_RGB, 2)
	var pix2 = make([]_RGB, 2)
	var p0 = pix[0]
	var p1 = &pix[1]
	p0.R = 1
	p1.R = 1
	fmt.Println(pix)
	fmt.Println(p0)
	fmt.Println(p1)
	pix2[0].R = 2
	pix[0] = pix2[0]
	fmt.Println(pix)
	pix2[0].G = 1
	fmt.Println(pix)
}

func TestBlurMatrix(t *testing.T) {
	sources := SourcePaths
	targets := BlurPaths
	filter := &Gauss5
	//filter, _ := CreateMotionBlurFilter(8, imagex.Vertical)
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.Auto)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		dstImg := imagex.CopyImageStruct(img)
		err = FilterImageWithMatrix(img, dstImg, *filter)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(dstImg, osxu.RunningBaseDir()+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}

func TestSimple(t *testing.T) {
	var rgb = make([]_RGB, 20)
	fmt.Println(rgb[2])
	rgb[0].B = 111
	fmt.Println(rgb[0])
}

func TestCreateBoxEightNearBlurFilter(t *testing.T) {
	var temp FilterMatrix
	temp, _ = CreateBoxFourNearBlurFilter(2, true)
	fmt.Println(temp)
	temp, _ = CreateBoxFourNearBlurFilter(2, false)
	fmt.Println(temp)
}

func TestCreateBoxFourNearBlurFilter(t *testing.T) {
	var temp FilterMatrix
	temp, _ = CreateBoxEightNearBlurFilter(2, true)
	fmt.Println(temp)
	temp, _ = CreateBoxEightNearBlurFilter(2, false)
	fmt.Println(temp)
}

func TestCreateGaussBlurFilter(t *testing.T) {
	var temp FilterMatrix
	temp, _ = CreateGaussBlurFilter(2, 1.4)
	fmt.Println(temp)
	temp, _ = CreateGaussBlurFilter(2, 1.0)
	fmt.Println(temp)
	temp, _ = CreateGaussBlurFilter(2, 0.8)
	fmt.Println(temp)
}

func TestCreateMotionBlurFilter(t *testing.T) {
	var temp FilterMatrix
	temp, _ = CreateMotionBlurFilter(1, imagex.AllDirection)
	fmt.Println(temp)
	temp, _ = CreateMotionBlurFilter(2, imagex.Vertical)
	fmt.Println(temp)
	temp, _ = CreateMotionBlurFilter(2, imagex.Horizontal)
	fmt.Println(temp)
	temp, _ = CreateMotionBlurFilter(2, imagex.Oblique)
	fmt.Println(temp)
}
