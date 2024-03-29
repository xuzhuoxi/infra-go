package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"reflect"
	"testing"
)

func TestFilterKernelRotate(t *testing.T) {
	filter := Gauss5
	fmt.Println(filter)
	filter2 := filter.Rotate(false, 1)
	fmt.Println(filter)
	fmt.Println(filter2)
}

func TestFilterImageWithTemplate(t *testing.T) {
	sources := SourcePaths
	targets := BlurPaths
	filter, _ := CreateMotionBlurFilter(8, imagex.Vertical)
	fmt.Println("Filter: ", filter)
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
		err = FilterImageWithMatrix(img, dstImg, filter)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(img, RunningDir+"/"+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
