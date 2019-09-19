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
	RegisterBlendFunc(LinearLight, BlendLinearLightColor, BlendLinearLightRGBA)
}

// 线性光模式
// 根据绘图色通过增加或降低“亮度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过增加亮度被照亮，如果绘图色比50%的灰暗，图像通过降低亮度变暗。
// R = S + 2*D - 1
// R = S + 2*D - 255
// R = S + 2*D - 65535
func BlendLinearLightColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLinearLightRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 线性光模式
// 根据绘图色通过增加或降低“亮度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过增加亮度被照亮，如果绘图色比50%的灰暗，图像通过降低亮度变暗。
// R = S + 2*D - 1
// R = S + 2*D - 255
// R = S + 2*D - 65535
func BlendLinearLightRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = linearLight(Sr, Dr)
	G = linearLight(Sg, Dg)
	B = linearLight(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = linearLight(Sa, Da)
	}
	return
}

// R = S + 2*D - 1
// R = S + 2*D - 255
// R = S + 2*D - 65535
func linearLight(S, D uint32) uint32 {
	Add := S + 2*D
	if Add > 65535 {
		return Add - 65535
	} else {
		return 0
	}
}
