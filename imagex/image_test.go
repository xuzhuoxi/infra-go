package imagex

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"image"
	"image/color"
	"math"
	"testing"
)

var rect = image.Rect(0, 0, 64, 64)

func TestNewGray(t *testing.T) {
	img := NewGray(rect, math.MaxUint8>>1)
	err := SaveImage(img, osxu.RunningBaseDir()+"test/gray.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestNewGray16(t *testing.T) {
	img := NewGray16(rect, math.MaxUint16>>1)
	err := SaveImage(img, osxu.RunningBaseDir()+"test/gray16.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestNewRGBA(t *testing.T) {
	img := NewRGBA(rect, math.MaxUint32>>1)
	err := SaveImage(img, osxu.RunningBaseDir()+"test/rgba.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestNewRGBA64(t *testing.T) {
	img := NewRGBA64(rect, math.MaxUint64>>1)
	err := SaveImage(img, osxu.RunningBaseDir()+"test/rgba64.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}

func TestFillImage(t *testing.T) {
	img := NewRGBA64(rect, math.MaxUint64>>1)
	FillImagetAt(img, color.Black, image.Rect(16, 16, 48, 48))
	err := SaveImage(img, osxu.RunningBaseDir()+"test/fill.jpeg", formatx.JPEG, jpegx.DefaultJPEGOptions)
	if nil != err {
		fmt.Println(err)
	}
}
