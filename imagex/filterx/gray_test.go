package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"image/color"
	"testing"
)

const (
	SourcePath01     = "test/01.png"
	SourcePath02     = "test/02.png"
	SourcePath03     = "test/03.png"
	SourcePath04     = "test/04.png"
	SourcePath05     = "test/05.png"
	SourcePath06     = "test/06.png"
	SourcePath07     = "test/07.png"
	GrayTargetPath01 = "test/gray/01.png"
	GrayTargetPath02 = "test/gray/02.png"
	GrayTargetPath03 = "test/gray/03.png"
	GrayTargetPath04 = "test/gray/04.png"
	GrayTargetPath05 = "test/gray/05.png"
	GrayTargetPath06 = "test/gray/06.png"
	GrayTargetPath07 = "test/gray/07.png"
)

func TestGrayRGBAImage(t *testing.T) {
	sources := []string{SourcePath01, SourcePath02, SourcePath03, SourcePath04, SourcePath05, SourcePath06, SourcePath07}
	targets := []string{GrayTargetPath01, GrayTargetPath02, GrayTargetPath03, GrayTargetPath04, GrayTargetPath05, GrayTargetPath06, GrayTargetPath07}
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		grayImg, err := GrayDefaultFromRGBA(img, color.Black)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(grayImg, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}

func TestGrayRGBImage(t *testing.T) {
	sources := []string{RGBTargetPath01, RGBTargetPath02, RGBTargetPath03, RGBTargetPath04, RGBTargetPath05, RGBTargetPath06, RGBTargetPath07}
	targets := []string{GrayTargetPath01, GrayTargetPath02, GrayTargetPath03, GrayTargetPath04, GrayTargetPath05, GrayTargetPath06, GrayTargetPath07}
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
		grayImg, err := GrayDefaultFromNRGBA(img)
		if nil != err {
			fmt.Println(err)
			continue
		}
		err = imagex.SaveImage(grayImg, osxu.RunningBaseDir()+targets[index], formatx.PNG, nil)
		if nil != err {
			fmt.Println(err)
		}
	}
}
