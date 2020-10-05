package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"reflect"
	"testing"
)

func TestCVTGray(t *testing.T) {
	sources := GrayPaths
	targets := CVTPaths
	for index, source := range sources {
		if index >= len(targets) {
			return
		}
		img, err := imagex.LoadImage(RunningDir+"/"+source, formatx.Auto)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		dstImg := imagex.CopyImageStruct(img)
		err = CVTGray(img, dstImg, 42767)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(dstImg, RunningDir+"/"+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
