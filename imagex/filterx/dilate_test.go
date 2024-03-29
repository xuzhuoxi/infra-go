package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"image/draw"
	"reflect"
	"testing"
)

func TestDilateCVT(t *testing.T) {
	sources := CVTPaths
	targets := DilatePaths
	for index, source := range sources {
		if index >= len(targets) {
			return
		}
		img, _, err := imagex.LoadImage(RunningDir+"/"+source, formatx.Auto)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		dstImg := imagex.CopyImageStruct(img)
		err = DilateCVT(img, dstImg, imagex.AllDirection)
		if nil != err {
			fmt.Println(err)
		}
		err = imagex.SaveImage(dstImg, RunningDir+"/"+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}

func TestDilateGray(t *testing.T) {
	sources := GrayPaths
	targets := DilatePaths
	for index, source := range sources {
		img, _, err := imagex.LoadImage(RunningDir+"/"+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		err = DilateGray(img, img.(draw.Image), imagex.AllDirection, 65535/2)
		if nil != err {
			fmt.Println(err)
		}
		err = imagex.SaveImage(img, RunningDir+"/"+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
