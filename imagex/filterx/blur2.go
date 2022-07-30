package filterx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/graphicx"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/draw"
)

func FastBlur2(srcImg image.Image, dstImg draw.Image, radius int) error {
	if nil == srcImg {
		return errors.New("SrcImg is nil! ")
	}
	if nil == dstImg {
		return errors.New("DstImg is nil! ")
	}
	if radius < 1 {
		return errors.New("KernelRadius should >=1. ")
	}
	size := srcImg.Bounds().Size()
	w, h := size.X, size.Y
	wm, hm, wh := w-1, h-1, w*h
	var div = radius + radius + 1

	var pix = make([]int, wh) //ARGB
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := srcImg.At(x, y).RGBA()
			pix[y*w+x] = int(graphicx.ARGB2Pixel(uint8(a/257), uint8(r/257), uint8(g/257), uint8(b/257)))
		}
	}

	var r, g, b = make([]int, wh), make([]int, wh), make([]int, wh)
	var rsum, gsum, bsum, x, y, i, p, yp, yi, yw int
	var vmin = make([]int, mathx.MaxInt(w, h))

	var divsum = (div + 1) >> 1
	divsum *= divsum
	var dv = make([]int, 256*divsum)
	for i = 0; i < 256*divsum; i++ {
		dv[i] = i / divsum
	}

	var stack = make([][3]int, div)
	var stackpointer, stackstart int
	var rbs int
	var sir [3]int
	var r1 = radius + 1
	var routsum, goutsum, boutsum int
	var rinsum, ginsum, binsum int

	for y = 0; y < h; y++ {
		rinsum, ginsum, binsum, routsum, goutsum, boutsum, rsum, gsum, bsum = 0, 0, 0, 0, 0, 0, 0, 0, 0
		for i = -radius; i <= radius; i++ {
			p = pix[yi+mathx.MinInt(wm, mathx.MaxInt(i, 0))]
			sir = stack[i+radius]
			sir[0], sir[1], sir[2] = (p&0xff0000)>>16, (p&0x00ff00)>>8, p&0x0000ff
			rbs = r1 - mathx.AbsInt(i)
			rsum += sir[0] * rbs
			gsum += sir[1] * rbs
			bsum += sir[2] * rbs
			if i > 0 {
				rinsum += sir[0]
				ginsum += sir[1]
				binsum += sir[2]
			} else {
				routsum += sir[0]
				goutsum += sir[1]
				boutsum += sir[2]
			}
		}
		stackpointer = radius
		for x = 0; x < w; x++ {
			fmt.Println(111, x, w, len(r), len(g), len(b), len(dv), yi, rsum, gsum, bsum)
			r[yi] = dv[rsum]
			g[yi] = dv[gsum]
			b[yi] = dv[bsum]
			rsum -= routsum
			gsum -= goutsum
			bsum -= boutsum

			stackstart = stackpointer - radius + div
			sir = stack[stackstart%div]

			routsum -= sir[0]
			goutsum -= sir[1]
			boutsum -= sir[2]

			if y == 0 {
				vmin[x] = mathx.MinInt(x+radius+1, wm)
			}
			p = pix[yw+vmin[x]]
			sir[0], sir[1], sir[2] = (p&0xff0000)>>16, (p&0x00ff00)>>8, p&0x0000ff

			rinsum += sir[0]
			ginsum += sir[1]
			binsum += sir[2]
			rsum += rinsum
			gsum += ginsum
			bsum += binsum
			stackpointer = (stackpointer + 1) % div
			sir = stack[(stackpointer)%div]
			routsum += sir[0]
			goutsum += sir[1]
			boutsum += sir[2]
			rinsum -= sir[0]
			ginsum -= sir[1]
			binsum -= sir[2]
			yi++
		}
		yw += w
	}
	for x = 0; x < w; x++ {
		rinsum, ginsum, binsum, routsum, goutsum, boutsum, rsum, gsum, bsum = 0, 0, 0, 0, 0, 0, 0, 0, 0
		yp = -radius * w
		for i = -radius; i <= radius; i++ {
			yi = mathx.MaxInt(0, yp) + x

			sir = stack[i+radius]

			sir[0] = r[yi]
			sir[1] = g[yi]
			sir[2] = b[yi]

			rbs = r1 - mathx.AbsInt(i)

			rsum += r[yi] * rbs
			gsum += g[yi] * rbs
			bsum += b[yi] * rbs

			if i > 0 {
				rinsum += sir[0]
				ginsum += sir[1]
				binsum += sir[2]
			} else {
				routsum += sir[0]
				goutsum += sir[1]
				boutsum += sir[2]
			}

			if i < hm {
				yp += w
			}
		}
		yi = x
		stackpointer = radius
		for y = 0; y < h; y++ {
			// Preserve alpha channel: ( 0xff000000 & pix[yi] )
			pix[yi] = (0xff000000 & pix[yi]) | (dv[rsum] << 16) | (dv[gsum] << 8) | dv[bsum]

			rsum -= routsum
			gsum -= goutsum
			bsum -= boutsum

			stackstart = stackpointer - radius + div
			sir = stack[stackstart%div]

			routsum -= sir[0]
			goutsum -= sir[1]
			boutsum -= sir[2]

			if x == 0 {
				vmin[y] = mathx.MinInt(y+r1, hm) * w
			}
			p = x + vmin[y]

			sir[0] = r[p]
			sir[1] = g[p]
			sir[2] = b[p]

			rinsum += sir[0]
			ginsum += sir[1]
			binsum += sir[2]

			rsum += rinsum
			gsum += ginsum
			bsum += binsum

			stackpointer = (stackpointer + 1) % div
			sir = stack[stackpointer]

			routsum += sir[0]
			goutsum += sir[1]
			boutsum += sir[2]

			rinsum -= sir[0]
			ginsum -= sir[1]
			binsum -= sir[2]

			yi += w
		}
	}
	return nil
}
