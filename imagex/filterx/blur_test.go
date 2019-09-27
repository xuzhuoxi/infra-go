package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"image/draw"
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

func TestFastBlur(t *testing.T) {
	sources := SourcePaths
	targets := BlurPaths
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		//err = BlurWithAverage(img, img.(draw.Image), 2)
		err = BlurWithTemplate(img, img.(draw.Image), FourNear3)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(img, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
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

func TestCreateGaussTemplate(t *testing.T) {
	var temp *BlurTemplate
	temp = CreateGaussBlurTemplate(2, 1.4)
	fmt.Println(temp)
	temp = CreateGaussBlurTemplate(2, 1.0)
	fmt.Println(temp)
}

func TestBlurWithTemplate(t *testing.T) {
	sources := SourcePaths
	targets := BlurPaths
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		err = BlurWithTemplate(img, img.(draw.Image), EightNear3)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(img, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
