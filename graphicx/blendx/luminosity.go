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
// HSV = Sh, Ss, Dv
func BlendLuminosityColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, _ := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLuminosityRGBA(Sr, Sg, Sb, 0, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 亮度模式
// 是采用底色的色相和饱和度以及绘图色的亮度来创建最终色。此模式创建于颜色模式相反效果。
// 注： “差值”、“排除”、“色相”、“饱和度”、“颜色”和“明度”模式都不能与专色相混合，而且对于多数混合模式而言，指定为 100% K 的黑色会挖空下方图层中的颜色。请不要使用 100% 黑色，应改为使用 CMYK 值来指定复色黑。
// HSV = Sh, Ss, Dv
func BlendLuminosityRGBA(Sr, Sg, Sb, _ uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R, G, B = luminosity(Sr, Sg, Sb, Dr, Dg, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = math.MaxUint16
	}
	return
}

// HSV = Sh, Ss, Dv
func luminosity(Sr, Sg, Sb uint32, Dr, Dg, Db uint32) (R, G, B uint32) {
	Sh, Ss, _ := graphicx.RGB2HSV(Sr, Sg, Sb)
	_, _, Dv := graphicx.RGB2HSV(Dr, Dg, Db)
	R, G, B = graphicx.HSV2RGB(Sh, Ss, Dv)
	return
}
