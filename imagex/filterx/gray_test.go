package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"image/color"
	"reflect"
	"testing"
)

func TestGrayRGBAImage(t *testing.T) {
	sources := SourcePaths
	targets := GrayPaths
	for index, source := range sources {
		if index >= len(targets) {
			return
		}
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.Auto)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		grayImg, err := GrayDefaultFromRGBA(img, color.White)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("灰度图像内存类型(grayImg)：", reflect.ValueOf(grayImg).Type())
		err = imagex.SaveImage(grayImg, osxu.RunningBaseDir()+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}

func TestGrayRGBImage(t *testing.T) {
	sources := RGBPaths
	targets := GrayPaths
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		grayImg, err := GrayDefaultFromNRGBA(img)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(grayImg, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
