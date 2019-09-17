package graphicx

import (
	"fmt"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"testing"
)

func TestRGBA_2_RGB(t *testing.T) {
	black := color.Black
	fmt.Println(black.RGBA())
	white := color.White
	fmt.Println(white.RGBA())
}

func TestRGBA(t *testing.T) {
	rgba := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	fmt.Println(rgba)
}

func TestGray(t *testing.T) {
	yellow := colornames.Yellow
	transparent := color.Transparent
	yellowGray := color.GrayModel.Convert(yellow)
	transparentGray := color.GrayModel.Convert(transparent)
	fmt.Println(yellow, yellowGray, transparent, transparentGray)
	fmt.Println(yellow.RGBA())
	fmt.Println(yellowGray.RGBA())
	fmt.Println(transparent.RGBA())
	fmt.Println(transparentGray.RGBA())
}

func TestRGBAFunc(t *testing.T) {
	rgb := &color.RGBA{R: 255, G: 255, B: 255, A: 200}
	fmt.Println(rgb.RGBA())
}

func TestOut(t *testing.T) {
	var m = uint32(math.MaxUint16)
	out := (m*m + m*m) >> 8
	fmt.Println(out)
	fmt.Println(math.MaxUint32, math.MaxUint16, m*m*2, m*m)
}
