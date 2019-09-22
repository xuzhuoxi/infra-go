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

func TestErodeCVT(t *testing.T) {
	sources := CVTPaths
	targets := ErodePaths
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		fmt.Println("读取的图像内存类型(img)：", reflect.ValueOf(img).Type())
		err = ErodeCVT(img, img.(draw.Image), imagex.AllDirection)
		if nil != err {
			fmt.Println(err)
		}
		err = imagex.SaveImage(img, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
