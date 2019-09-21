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

func TestCVTGray(t *testing.T) {
	sources := GrayPaths
	targets := CVTPaths
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			return
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		err = CVTGray(img, img.(draw.Image), 42767)
		if nil != err {
			fmt.Println(err)
			return
		}
		err = imagex.SaveImage(img, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
