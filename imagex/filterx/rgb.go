//RGBA转RGB
package filterx

import (
	"errors"
	"github.com/xuzhuoxi/infra-go/graphicx/blendx"
	"github.com/xuzhuoxi/infra-go/imagex"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"image/draw"
	"math"
)

//以绿色(Green)作底转换RGBA图像为RGB图像(去Alpha通道)
func NrgbaAtGreen(rgbaImg image.Image, nrgbaImg draw.Image) error {
	return NrgbaAt(rgbaImg, nrgbaImg, colornames.Green)
}

//以黑色(Black)作底转换RGBA图像为RGB图像(去Alpha通道)
func NrgbaAtBlack(rgbaImg image.Image, nrgbaImg draw.Image) error {
	return NrgbaAt(rgbaImg, nrgbaImg, color.Black)
}

//以白色(White)作底转换RGBA图像为RGB图像(去Alpha通道)
func NrgbaAtWhite(rgbaImg image.Image, nrgbaImg draw.Image) error {
	return NrgbaAt(rgbaImg, nrgbaImg, color.White)
}

//以某个背景色作底转换RGBA图像为NRGBA图像(去Alpha通道)
func NrgbaAt(rgbaImg image.Image, nrgbaImg draw.Image, backgroundColor color.Color) error {
	if nil == rgbaImg {
		return errors.New("SrcImg is nil! ")
	}
	if nil == nrgbaImg {
		return errors.New("DstImg is nil! ")
	}
	if rgbaImg == nrgbaImg {
		imagex.BlendSourceNormal(nrgbaImg, backgroundColor)
	} else {
		rgbaRect := rgbaImg.Bounds()
		Sr, Sg, Sb, Sa := backgroundColor.RGBA()
		setColor := &color.RGBA64{A: 65535}
		var Dr, Dg, Db, Da uint32
		var R, G, B uint32
		var rgbaColor color.Color
		for y := rgbaRect.Min.Y; y < rgbaRect.Max.Y; y++ {
			for x := rgbaRect.Min.X; x < rgbaRect.Max.X; x++ {
				rgbaColor = rgbaImg.At(x, y)
				Dr, Dg, Db, Da = rgbaColor.RGBA()
				if math.MaxUint16 == Da { //前景不透明,直接复制源像素
					nrgbaImg.Set(x, y, rgbaColor)
					continue
				}
				if 0 == Da { //前景全透明，直接复制背景像素
					nrgbaImg.Set(x, y, backgroundColor)
					continue
				}
				R, G, B, _ = blendx.BlendNormalRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, false)
				setColor.R, setColor.G, setColor.B = uint16(R), uint16(G), uint16(B)
				nrgbaImg.Set(x, y, setColor)
			}
		}
	}
	return nil
}
