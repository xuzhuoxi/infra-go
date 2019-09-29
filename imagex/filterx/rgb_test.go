package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"reflect"
	"testing"
)

func TestImage2NRGBA(t *testing.T) {
	sources := SourcePaths
	targets := RGBPaths
	for index, source := range sources {
		if index >= len(targets) {
			return
		}
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.Auto)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型：", reflect.ValueOf(img).Type())
		dstImg := imagex.CopyImage(img)
		err = NrgbaAtWhite(img, dstImg)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(img, osxu.RunningBaseDir()+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
			continue
		}
	}
}
