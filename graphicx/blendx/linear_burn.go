//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(LinearBurn, BlendLinearBurnColor, BlendLinearBurnRGBA)
}

// 线性加深模式
// 查看每个通道的颜色信息，通过降低“亮度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// R = S + D - 1
// R = S + D - 255
// R = S + D - 65535
func BlendLinearBurnColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLinearBurnRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 线性加深模式
// 查看每个通道的颜色信息，通过降低“亮度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// R = S + D - 1
// R = S + D - 255
// R = S + D - 65535
func BlendLinearBurnRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = linearBurn(Sr, Dr)
	G = linearBurn(Sg, Dg)
	B = linearBurn(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = linearBurn(Sa, Da)
	}
	return
}

// R = S + D - 1
// R = S + D - 255
// R = S + D - 65535
func linearBurn(S, D uint32) uint32 {
	Add := S + D
	if Add > 65535 {
		return Add - 65535
	} else {
		return 0
	}
}
