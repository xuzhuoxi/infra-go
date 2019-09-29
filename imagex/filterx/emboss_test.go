package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"reflect"
	"testing"
)

func TestCheckEmboss(t *testing.T) {
	fmt.Println(Emboss3Oblique45.CheckValidity())
	fmt.Println(Emboss5Oblique45.CheckValidity())
	fmt.Println(Emboss3Asymmetrical.CheckValidity())
}

func TestEmbossImage(t *testing.T) {
	sources := SourcePaths
	targets := EmbossPaths
	filter := &Emboss3Asymmetrical
	//filter, _ := CreateEdgeFilter(1, imagex.AllDirection, 2)
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
		dstImg := imagex.CopyImage(img)
		err = FilterImageWithTemplate(img, dstImg, *filter)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(img, osxu.RunningBaseDir()+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
