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

func TestImage2NRGBA(t *testing.T) {
	sources := SourcePaths
	targets := RGBPaths
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型：", reflect.ValueOf(img).Type())
		err = NrgbaAtWhite(img, img.(draw.Image))
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(img, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
			continue
		}
	}
}
