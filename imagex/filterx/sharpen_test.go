package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"reflect"
	"testing"
)

func TestCheckSharpen(t *testing.T) {
	fmt.Println(Sharpen3Laplace4.CheckValidity())
	fmt.Println(Sharpen3Laplace8.CheckValidity())
	fmt.Println(SharpenStrengthen3All.CheckValidity())
	fmt.Println(Sharpen5All.CheckValidity())
}

func TestSharpen(t *testing.T) {
	sources := SourcePaths
	targets := SharpenPaths
	filter := &SharpenStrengthen3All
	fmt.Println("check:", filter.IsScaleMatrix(), filter.IsPixelUnsafe())
	//filter, _ := CreateMotionBlurFilter(8, imagex.Vertical)
	fmt.Println("Filter: ", filter)
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
