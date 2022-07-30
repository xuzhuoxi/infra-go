package filterx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/graphicx"
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
)

// FastBlur.java
func FastBlur_Java(srcImg image.Image, dstImg draw.Image, radius int) error {
	// Stack Blur v1.0 from
	// http://www.quasimondo.com/StackBlurForCanvas/StackBlurDemo.html
	//
	// Java Author: Mario Klingemann <mario at quasimondo.com>
	// http://incubator.quasimondo.com
	// created Feburary 29, 2004
	// Android port : Yahel Bouaziz <yahel at kayenko.com>
	// http://www.kayenko.com
	// ported april 5th, 2012

	// This is a compromise between Gaussian Blur and Box blur
	// It creates much better looking blurs than Box Blur, but is
	// 7x faster than my Gaussian Blur implementation.
	//
	// I called it Stack Blur because this describes best how this
	// filter works internally: it creates a kind of moving stack
	// of colors whilst scanning through the image. Thereby it
	// just has to add one new block of color to the right side
	// of the stack and remove the leftmost color. The remaining
	// colors on the topmost layer of the stack are either added on
	// or reduced by one, depending on if they are on the right or
	// on the left side of the stack.
	//
	// If you are using this algorithm in your code please add
	// the following line:
	//
	// Stack Blur Algorithm by Mario Klingemann <mario@quasimondo.com>
	if radius < 1 {
		return errors.New("KernelRadius < 1. ")
	}
	size := srcImg.Bounds().Size()
	var w, h = size.X, size.Y
	var wm, hm, wh = w - 1, h - 1, w * h
	var pix = make([]graphicx.Pixel, wh, wh)
	var div = radius + radius + 1
	var r, g, b = make([]int, wh, wh), make([]int, wh, wh), make([]int, wh, wh)
	var rsum, gsum, bsum, x, y, i, yp, yi, yw int
	var p graphicx.Pixel
	var vmin = make([]int, mathx.MaxInt(w, h))
	var divsum = (div + 1) >> 1
	divsum *= divsum
	var dv = make([]int, 256*divsum*512)

	for y = 0; y < h; y++ {
		for x = 0; x < w; x++ {
			pix[y*w+x] = graphicx.Pixel(graphicx.Color2Pixel(srcImg.At(x, y)))
		}
	}
	for i = 0; i < 256*divsum*512; i++ {
		dv[i] = i / divsum
	}

	var stack = make([][4]int, div, div)
	var stackpointer, stackstart, rbs int
	var sir [4]int
	var R, G, B, A uint8
	var r1 = radius + 1
	var routsum, goutsum, boutsum int
	var rinsum, ginsum, binsum int

	for y = 0; y < h; y++ {
		rinsum, ginsum, binsum, routsum, goutsum, boutsum, rsum, gsum, bsum = 0, 0, 0, 0, 0, 0, 0, 0, 0
		for i = -radius; i <= radius; i++ {
			p = pix[yi+mathx.MaxInt(wm, mathx.MaxInt(i, 0))]
			sir = stack[i+radius]
			R, G, B, A = p.RGBA()
			sir[0], sir[1], sir[2] = int(R), int(G), int(B)
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
		fmt.Println(2222, rinsum, ginsum, binsum, routsum, goutsum, boutsum, rsum, gsum, bsum)
		for x = 0; x < w; x++ {
			fmt.Println(1111, rinsum, ginsum, binsum, routsum, goutsum, boutsum, rsum, gsum, bsum)
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
			R, G, B, A = p.RGBA()
			sir[0], sir[1], sir[2], sir[3] = int(R), int(G), int(B), int(A)
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
			pix[yi] = graphicx.Pixel(graphicx.ARGB2Pixel(1, uint8(dv[rsum]), uint8(dv[gsum]), uint8(dv[bsum])))
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
			p = graphicx.Pixel(x + vmin[y])

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
	var setColor = &color.RGBA{}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			setColor.A, setColor.R, setColor.G, setColor.B = pix[y*w+x].RGBA()
			dstImg.Set(x, y, setColor)
		}
	}
	return nil
}

func FastBlur(srcImg image.Image, dstImg draw.Image, radius int) error {
	if nil == srcImg || nil == dstImg {
		return errors.New("SrcImg or dstImg is nil! ")
	}
	if radius < 1 {
		return errors.New("KernelRadius should >=1. ")
	}
	size := srcImg.Bounds().Size()
	w, h := size.X, size.Y
	wm, hm, wh := w-1, h-1, w*h
	div := radius + radius + 1

	var pix = make([]int, wh*3)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := srcImg.At(x, y).RGBA()
			pix[y*w+x+0] = int(r) >> 8
			pix[y*w+x+1] = int(g) >> 8
			pix[y*w+x+2] = int(b) >> 8
		}
	}

	var vMIN, vMAX = make([]int, mathx.MaxInt(w, h)), make([]int, mathx.MaxInt(w, h))
	var r, g, b, dv = make([]int, wh), make([]int, wh), make([]int, wh), make([]int, 256*div)
	var rsum, gsum, bsum, x, y, i, p, p1, p2, yp, yi, yw int
	for i = 0; i < 256*div; i++ {
		dv[i] = i / div
	}

	for y = 0; y < h; y++ {
		rsum, gsum, bsum = 0, 0, 0
		for i = -radius; i <= radius; i++ {
			p = (yi + mathx.MinInt(wm, mathx.MaxInt(i, 0))) * 3
			rsum += pix[p+0]
			gsum += pix[p+1]
			bsum += pix[p+2]
		}
		for x = 0; x < w; x++ {
			r[yi] = dv[rsum]
			g[yi] = dv[gsum]
			b[yi] = dv[bsum]

			if y == 0 {
				vMIN[x] = mathx.MinInt(x+radius+1, wm)
				vMAX[x] = mathx.MaxInt(x-radius, 0)
			}
			p1 = (yw + vMIN[x]) * 3
			p2 = (yw + vMAX[x]) * 3
			rsum += pix[p1+0] - pix[p2+0]
			gsum += pix[p1+1] - pix[p2+1]
			bsum += pix[p1+2] - pix[p2+2]
			yi++
		}
		yw += w
	}

	for x = 0; x < w; x++ {
		rsum, gsum, bsum = 0, 0, 0
		yp = -radius * w
		for i = -radius; i <= radius; i++ {
			yi = mathx.MaxInt(0, yp) + x
			rsum += r[yi]
			gsum += g[yi]
			bsum += b[yi]
			yp += w
		}
		yi = x
		for y = 0; y < h; y++ {
			pix[yi*3+0] = dv[rsum]
			pix[yi*3+1] = dv[gsum]
			pix[yi*3+2] = dv[bsum]
			if x == 0 {
				vMIN[y] = mathx.MinInt(y+radius+1, hm) * w
				vMAX[y] = mathx.MaxInt(y-radius, 0) * w
			}
			p1 = x + vMIN[y]
			p2 = x + vMAX[y]
			rsum += r[p1] - r[p2]
			gsum += g[p1] - g[p2]
			bsum += b[p1] - b[p2]
			yi += w
		}
	}
	dstRect := dstImg.Bounds()
	var c = &color.NRGBA{A: 255}
	var index int
	var add int
	for y := dstRect.Min.Y; y < dstRect.Max.Y; y++ {
		//var str = ""
		for x := dstRect.Min.X; x < dstRect.Max.X; x++ {
			index = y*w + x
			c.R, c.G, c.B = uint8(pix[index+0]), uint8(pix[index+1]), uint8(pix[index+2])
			//str += "," + fmt.Sprint(c)
			dstImg.Set(x, y, c)
			if c.R == c.G && c.G == c.B {
				add++
			}
		}
		//fmt.Println(str)
	}
	fmt.Println("结果：", wh, add)
	return nil
}

// 堆模糊(快速模糊)
// https://www.cnblogs.com/Darksun/p/4681476.html
// FastBlurDemo.zip
func FastBlurARGB8888(srcImg image.Image, dstImg draw.Image, radius int) error {
	if nil == srcImg || nil == dstImg {
		return errors.New("SrcImg or dstImg is nil! ")
	}
	if radius < 1 {
		return errors.New("KernelRadius should >=1. ")
	}
	size := srcImg.Bounds().Size()
	w, h := size.X, size.Y
	wm, hm, wh := w-1, h-1, w*h
	div := radius + radius + 1

	var r, g, b = make([]int, wh), make([]int, wh), make([]int, wh)
	var rsum, gsum, bsum, x, y, i, p1, p2, yp, yi, yw int
	var vmin, vmax = make([]int, mathx.MaxInt(w, h)), make([]int, mathx.MaxInt(w, h))
	var pix = imagex.NewPixelImage(w, h, 0)
	var R1, G1, B1 uint8
	var R2, G2, B2 uint8
	imagex.Copy2PixelImage(srcImg, pix)

	for y = 0; y < h; y++ {
		rsum, gsum, bsum = 0, 0, 0
		for i = -radius; i <= radius; i++ {
			R1, G1, B1, _ = pix.PixelIndexAt(yi + mathx.MinInt(wm, mathx.MaxInt(i, 0))).RGBA()
			rsum += int(R1)
			gsum += int(G1)
			bsum += int(B1)
		}
		for x = 0; x < w; x++ {
			r[yi] = rsum / div
			g[yi] = gsum / div
			b[yi] = bsum / div
			if y == 0 {
				vmin[x] = mathx.MinInt(x+radius+1, wm)
				vmax[x] = mathx.MaxInt(x-radius, 0)
			}
			R1, G1, B1, _ = pix.PixelIndexAt(yw + vmin[x]).RGBA()
			R2, G2, B2, _ = pix.PixelIndexAt(yw + vmax[x]).RGBA()
			rsum += int(R1) - int(R2)
			gsum += int(G1) - int(G2)
			bsum += int(B1) - int(B2)
			yi++
		}
		yw += w
	}

	for x = 0; x < w; x++ {
		rsum, gsum, bsum = 0, 0, 0
		yp = -radius * w
		for i = -radius; i <= radius; i++ {
			yi = mathx.MaxInt(0, yp) + x
			rsum += r[yi]
			gsum += g[yi]
			bsum += b[yi]
			yp += w
		}
		yi = x
		for y = 0; y < h; y++ {
			pix.IndexSet(yi, graphicx.ARGB2Pixel(0xff, uint8(rsum/div), uint8(gsum/div), uint8(bsum/div)))
			if x == 0 {
				vmin[y] = mathx.MinInt(y+radius+1, hm) * w
				vmax[y] = mathx.MaxInt(y-radius, 0) * w
			}
			p1 = x + vmin[y]
			p2 = x + vmax[y]

			rsum += r[p1] - r[p2]
			gsum += g[p1] - g[p2]
			bsum += b[p1] - b[p2]

			yi += w
		}
	}
	imagex.CopyPixel2Image(pix, dstImg)
	return nil
}
