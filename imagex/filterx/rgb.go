package filterx

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/xuzhuoxi/infra-go/graphicx/blendx"
	"github.com/xuzhuoxi/infra-go/imagex"
	"golang.org/x/image/colornames"
)

// NRGBAAtGreen converts an RGBA image to NRGBA using green as the background.
// NRGBAAtGreen 以绿色（Green）作底将 RGBA 图像转换为 NRGBA 图像（去除 Alpha 通道）。
func NRGBAAtGreen(rgbaImg image.Image, nrgbaImg draw.Image) error {
	return NRGBAAt(rgbaImg, nrgbaImg, colornames.Green)
}

// NRGBAAtBlack converts an RGBA image to NRGBA using black as the background.
// NRGBAAtBlack 以黑色（Black）作底将 RGBA 图像转换为 NRGBA 图像（去除 Alpha 通道）。
func NRGBAAtBlack(rgbaImg image.Image, nrgbaImg draw.Image) error {
	return NRGBAAt(rgbaImg, nrgbaImg, color.Black)
}

// NRGBAAtWhite converts an RGBA image to NRGBA using white as the background.
// NRGBAAtWhite 以白色（White）作底将 RGBA 图像转换为 NRGBA 图像（去除 Alpha 通道）。
func NRGBAAtWhite(rgbaImg image.Image, nrgbaImg draw.Image) error {
	return NRGBAAt(rgbaImg, nrgbaImg, color.White)
}

// NRGBAAt converts an RGBA image to NRGBA composited over a background color.
// NRGBAAt 以指定背景色作底将 RGBA 图像转换为 NRGBA 图像（去除 Alpha 通道）。
func NRGBAAt(rgbaImg image.Image, nrgbaImg draw.Image, backgroundColor color.Color) error {
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
