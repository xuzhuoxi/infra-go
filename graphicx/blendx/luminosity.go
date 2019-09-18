//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"github.com/xuzhuoxi/infra-go/graphicx"
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(Luminosity, BlendLuminosityColor, BlendLuminosityRGBA)
}

// 亮度模式
// 是采用底色的色相和饱和度以及绘图色的亮度来创建最终色。此模式创建于颜色模式相反效果。
// 注： “差值”、“排除”、“色相”、“饱和度”、“颜色”和“明度”模式都不能与专色相混合，而且对于多数混合模式而言，指定为 100% K 的黑色会挖空下方图层中的颜色。请不要使用 100% 黑色，应改为使用 CMYK 值来指定复色黑。
// HSV = bH, bS, fV
func BlendLuminosityColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, _ := backColor.RGBA()
	R, G, B, A := BlendLuminosityRGBA(fR, fG, fB, fA, bR, bG, bB, 0, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 亮度模式
// 是采用底色的色相和饱和度以及绘图色的亮度来创建最终色。此模式创建于颜色模式相反效果。
// 注： “差值”、“排除”、“色相”、“饱和度”、“颜色”和“明度”模式都不能与专色相混合，而且对于多数混合模式而言，指定为 100% K 的黑色会挖空下方图层中的颜色。请不要使用 100% 黑色，应改为使用 CMYK 值来指定复色黑。
// HSV = bH, bS, fV
func BlendLuminosityRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, _ uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R, G, B = luminosity(foreR, foreG, foreB, backR, backG, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = bH, bS, fV
func luminosity(foreR, foreG, foreB uint32, backR, backG, backB uint32) (R, G, B uint32) {
	backH, backS, _ := graphicx.RGB2HSV(backR, backG, backB)
	_, _, foreV := graphicx.RGB2HSV(foreR, foreG, foreB)
	R, G, B = graphicx.HSV2RGB(backH, backS, foreV)
	return
}
