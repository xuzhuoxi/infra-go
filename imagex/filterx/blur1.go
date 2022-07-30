// 模糊
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
)

type _RGB struct {
	R, G, B int
}

func (rgb _RGB) Pixel() int {
	return (rgb.R&0x0000ff)<<16 | (rgb.R&0x0000ff)<<8 | rgb.R&0x0000ff
}

// 模糊
func FastBlur64(srcImg image.Image, dstImg draw.Image, radius int) error {
	if nil == srcImg {
		return errors.New("SrcImg is nil! ")
	}
	if radius < 1 {
		return errors.New("KernelRadius should >=1. ")
	}
	rect := srcImg.Bounds()

	var w, h = rect.Size().X, rect.Size().Y
	var pix = make([]_RGB, w*h)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, _ := srcImg.At(x, y).RGBA()
			pix[y*w+x] = _RGB{R: int(r), G: int(g), B: int(b)}
		}
	}

	var wm = w - 1
	var hm = h - 1
	var wh = w * h
	var div = radius + radius + 1

	var rgb = make([]_RGB, wh)
	var rgbSum = &_RGB{}

	var x, y, i, p, yp, yi, yw int
	var vMin = make([]int, mathx.MaxInt(w, h))

	var divSum = (div + 1) >> 1
	divSum *= divSum
	var dv = make([]int, 256*divSum)
	for i := 0; i < 256*divSum; i++ {
		dv[i] = i / divSum
	}

	var stack = make([]_RGB, div)
	var stackPointer int
	var stackStart int
	var sir *_RGB
	var rbs int
	var r1 = radius + 1
	var outSum, inSum = &_RGB{}, &_RGB{}
	for y = 0; y < h; y++ {
		inSum.R, inSum.G, inSum.B = 0, 0, 0
		outSum.R, outSum.G, outSum.B = 0, 0, 0
		rgbSum.R, rgbSum.G, rgbSum.B = 0, 0, 0
		for i = -radius; i <= radius; i++ {
			stack[i+radius] = pix[yi+mathx.MinInt(wm, mathx.MaxInt(i, 0))]
			sir = &stack[i+radius]
			rbs = r1 - mathx.AbsInt(i)
			rgbSum.R, rgbSum.G, rgbSum.B = sir.R*rbs, sir.G*rbs, sir.B*rbs
			if i > 0 {
				inSum.R, inSum.G, inSum.B = inSum.R+sir.R, inSum.G+sir.G, inSum.B+sir.B
			} else {
				outSum.R, outSum.G, outSum.B = outSum.R+sir.R, outSum.G+sir.G, outSum.B+sir.B
			}
		}
		stackPointer = radius
		for x = 0; x < w; x++ {
			rgb[yi].R, rgb[yi].G, rgb[yi].B = dv[rgbSum.R], dv[rgbSum.G], dv[rgbSum.B]
			rgbSum.R, rgbSum.G, rgbSum.B = rgbSum.R-outSum.R, rgbSum.G-outSum.G, rgbSum.B-outSum.B
			stackStart = stackPointer - radius + div
			sir = &stack[stackStart%div]
			outSum.R, outSum.G, outSum.B = outSum.R-sir.R, outSum.G-sir.G, outSum.B-sir.B
			if y == 0 {
				vMin[x] = mathx.MinInt(x+radius+1, wm)
			}
			sir.R, sir.G, sir.B = pix[yw+vMin[x]].R, pix[yw+vMin[x]].G, pix[yw+vMin[x]].B
			inSum.R, inSum.G, inSum.B = inSum.R+sir.R, inSum.G+sir.G, inSum.B+sir.B
			rgbSum.R, rgbSum.G, rgbSum.B = rgbSum.R+inSum.R, rgbSum.G+inSum.G, rgbSum.B+inSum.B
			stackPointer = (stackPointer + 1) % div
			sir = &stack[(stackPointer)%div]
			outSum.R, outSum.G, outSum.B = outSum.R+sir.R, outSum.G+sir.G, outSum.B+sir.B
			inSum.R, inSum.G, inSum.B = inSum.R-sir.R, inSum.G-sir.G, inSum.B-sir.B
			yi++
		}
		yw += w
	}
	for x = 0; x < w; x++ {
		inSum.R, inSum.G, inSum.B = 0, 0, 0
		outSum.R, outSum.G, outSum.B = 0, 0, 0
		rgbSum.R, rgbSum.G, rgbSum.B = 0, 0, 0
		yp = -radius * w
		for i = -radius; i <= radius; i++ {
			yi = mathx.MaxInt(0, yp) + x
			sir = &stack[i+radius]
			sir.R, sir.G, sir.B = rgb[yi].R, rgb[yi].G, rgb[yi].B
			rbs = r1 - mathx.AbsInt(i)
			rgbSum.R, rgbSum.G, rgbSum.B = rgb[yi].R*rbs, rgb[yi].G*rbs, rgb[yi].B*rbs

			if i > 0 {
				inSum.R, inSum.G, inSum.B = inSum.R+sir.R, inSum.G+sir.G, inSum.B+sir.B
			} else {
				outSum.R, outSum.G, outSum.B = outSum.R+sir.R, outSum.G+sir.G, outSum.B+sir.B
			}
			if i < hm {
				yp += w
			}
		}
		yi = x
		stackPointer = radius
		for y = 0; y < h; y++ {
			// Preserve alpha channel: ( 0xff000000 & pix[yi] )
			pix[yi].R, pix[yi].G, pix[yi].B = dv[rgbSum.R], dv[rgbSum.G], dv[rgbSum.B]
			rgbSum.R, rgbSum.G, rgbSum.B = rgbSum.R-outSum.R, rgbSum.G-outSum.G, rgbSum.B-outSum.B
			stackStart = stackPointer - radius + div
			sir = &stack[stackStart%div]
			outSum.R, outSum.G, outSum.B = outSum.R-sir.R, outSum.G-sir.G, outSum.B-sir.B
			if x == 0 {
				vMin[y] = mathx.MinInt(y+r1, hm) * w
			}
			p = x + vMin[y]
			sir.R, sir.G, sir.B = rgb[p].R, rgb[p].G, rgb[p].B
			inSum.R, inSum.G, inSum.B = inSum.R+sir.R, inSum.G+sir.G, inSum.B+sir.B
			rgbSum.R, rgbSum.G, rgbSum.B = rgbSum.R+inSum.R, rgbSum.G+inSum.G, rgbSum.B+inSum.B
			stackPointer = (stackPointer + 1) % div
			sir = &stack[stackPointer]
			outSum.R, outSum.G, outSum.B = outSum.R+sir.R, outSum.G+sir.G, outSum.B+sir.B
			inSum.R, inSum.G, inSum.B = inSum.R-sir.R, inSum.G-sir.G, inSum.B-sir.B
			yi += w
		}
	}
	//复制像素
	dstRect := dstImg.Bounds()
	cm := dstImg.ColorModel()
	var c = &color.RGBA{}
	var a uint32
	for y := dstRect.Min.Y; y < dstRect.Max.Y; y++ {
		for x := dstRect.Min.X; x < dstRect.Max.X; x++ {
			_, _, _, a = dstImg.At(x, y).RGBA()
			sir = &pix[w*y+x]
			c.R, c.G, c.B, c.A = uint8(sir.R), uint8(sir.G), uint8(sir.B), uint8(a)
			dstImg.Set(x, y, cm.Convert(*c))
		}
	}
	return nil
}
