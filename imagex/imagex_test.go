package imagex

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"image/draw"
	"math"
	"testing"
)

var rect = image.Rect(0, 0, 64, 64)
var RunningDir = osxu.GetRunningDir()

func TestNewGray(t *testing.T) {
	img := NewGray(rect, math.MaxUint8>>1)
	err := SaveImage(img, RunningDir+"/test/gray.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestNewGray16(t *testing.T) {
	img := NewGray16(rect, math.MaxUint16>>1)
	err := SaveImage(img, RunningDir+"/test/gray16.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestNewRGBA(t *testing.T) {
	img := NewRGBA(rect, color.White)
	err := SaveImage(img, RunningDir+"/test/rgba.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestNewRGBA64(t *testing.T) {
	img := NewRGBA64(rect, color.White)
	err := SaveImage(img, RunningDir+"/test/rgba64.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestFillImage(t *testing.T) {
	img := NewRGBA64(rect, color.White)
	FillImageAt(img, color.Black, image.Rect(16, 16, 48, 48))
	err := SaveImage(img, RunningDir+"/test/fill.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestCopyImageTo(t *testing.T) {
	src := []string{"test/src01.png", "test/src02.png"}
	dst := []string{"test/copy_dst01.jpeg", "test/copy_dst02.jpeg"}
	for index, _ := range src {
		srcImg, err := LoadImage(RunningDir+"/"+src[index], formatx.PNG)
		if nil != err {
			fmt.Println(err)
		}
		dstImg := image.NewNRGBA(srcImg.Bounds())
		CopyImageTo(srcImg, dstImg)
		err = SaveImage(dstImg, RunningDir+"/"+dst[index], formatx.PNG, jpegx.DefaultJPEGOptions)
		if nil != err {
			fmt.Println(err)
		}
	}
}

func TestBlendBackground(t *testing.T) {
	src := []string{"test/src01.png", "test/src02.png"}
	dst := []string{"test/blend_dst01.jpeg", "test/blend_dst02.jpeg"}
	bg := []color.Color{color.White, colornames.Yellow}
	for index, _ := range src {
		img, err := LoadImage(RunningDir+"/"+src[index], formatx.PNG)
		if nil != err {
			fmt.Println(err)
		}
		BlendSourceNormal(img.(draw.Image), bg[index])
		err = SaveImage(img, RunningDir+"/"+dst[index], formatx.PNG, jpegx.DefaultJPEGOptions)
		if nil != err {
			fmt.Println(err)
		}
		fmt.Println()
	}
}
