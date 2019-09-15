package filterx

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {
	var pix = make([]_RGB, 2)
	var pix2 = make([]_RGB, 2)
	var p0 = pix[0]
	var p1 = &pix[1]
	p0.R = 1
	p1.R = 1
	fmt.Println(pix)
	fmt.Println(p0)
	fmt.Println(p1)
	pix2[0].R = 2
	pix[0] = pix2[0]
	fmt.Println(pix)
	pix2[0].G = 1
	fmt.Println(pix)
}
