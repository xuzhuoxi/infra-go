package filterx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"image/draw"
	"testing"
)

const (
	RGBTargetPath01 = "test/rgb/01.png"
	RGBTargetPath02 = "test/rgb/02.png"
	RGBTargetPath03 = "test/rgb/03.png"
	RGBTargetPath04 = "test/rgb/04.png"
	RGBTargetPath05 = "test/rgb/05.png"
	RGBTargetPath06 = "test/rgb/06.png"
	RGBTargetPath07 = "test/rgb/07.png"
)

func TestImageRGBA2RGB(t *testing.T) {
	sources := []string{SourcePath01, SourcePath02, SourcePath03, SourcePath04, SourcePath05, SourcePath06, SourcePath07}
	targets := []string{RGBTargetPath01, RGBTargetPath02, RGBTargetPath03, RGBTargetPath04, RGBTargetPath05, RGBTargetPath06, RGBTargetPath07}
	for index, source := range sources {
		img, err := imagex.LoadImage(osxu.RunningBaseDir()+source, formatx.PNG)
		if nil != err {
			fmt.Println(err)
			continue
		}
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
