package filterx

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/mathx"
	"image"
	"image/color"
	"image/draw"
)

func FastBlur3(srcImg image.Image, dstImg draw.Image, radius int) error {
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
	div := radius + radius + 1

	var pix = make([]int, wh*3)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := srcImg.At(x, y).RGBA()
			pix[y*w+x+0] = int(r) / 256
			pix[y*w+x+1] = int(g) / 256
			pix[y*w+x+2] = int(b) / 256
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
