package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"reflect"
	"testing"
)

func TestCreateEdgeFilter(t *testing.T) {
	var temp FilterMatrix
	temp, _ = CreateEdgeFilter(1, imagex.AllDirection, 0)
	fmt.Println(temp)
	temp, _ = CreateEdgeFilter(2, imagex.Vertical, 1)
	fmt.Println(temp)
	temp, _ = CreateEdgeFilter(3, imagex.Oblique, 2)
	fmt.Println(temp)
}

func TestCheckEdge(t *testing.T) {
	fmt.Println(Edge5LR.CheckValidity())
	fmt.Println(Edge5UD.CheckValidity())
	fmt.Println(Edge5LuRd.CheckValidity())
	fmt.Println(Edge5LdRu.CheckValidity())
	fmt.Println(Edge3Laplace8.CheckValidity())
}

func Test000(t *testing.T) {
	var v = -8
	var v2 = uint32(v)
	fmt.Println(v, v2)
}

func TestEdgeImage(t *testing.T) {
	sources := SourcePaths
	targets := EdgePaths
	filter := &Edge3Laplace8
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
