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

func TestFilterImageWithTemplate(t *testing.T) {
	sources := SourcePaths
	targets := BlurPaths
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		//filter := Gauss5
		filter, _ := CreateMotionBlurFilter(5, imagex.Vertical)
		err = FilterImageWithTemplate(img, img.(draw.Image), *filter)
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
