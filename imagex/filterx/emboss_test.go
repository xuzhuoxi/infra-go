package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"reflect"
	"testing"
)

func TestCheckEmboss(t *testing.T) {
	fmt.Println(Emboss3Lu2Rd.CheckValidity())
	fmt.Println(Emboss5Lu2Rd.CheckValidity())
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
		img, _, err := imagex.LoadImage(RunningDir+"/"+source, formatx.Auto)
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
		err = imagex.SaveImage(dstImg, RunningDir+"/"+targets[index], formatx.Auto, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
