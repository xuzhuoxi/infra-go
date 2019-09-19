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
	RegisterBlendFunc(VividLight, BlendVividLightColor, BlendVividLightRGBA)
}

// 亮光模式
// 根据绘图色通过增加或降低“对比度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过降低对比度被照亮，如果绘图色比50%的灰暗，图像通过增加对比度变暗。
// (D<=128): R = S - (255-S)*(255-2*D) / (2*D)
// (D>128): R = S + S*(2*D-255)/(2*(255-D))
//
// (D<=32768): R = S - (65535-S)*(65535-2*D) / (2*D)
// (D>32768): R = S + S*(2*D-65535)/(2*(6553-D))
func BlendVividLightColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendVividLightRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 亮光模式
// 根据绘图色通过增加或降低“对比度”，加深或减淡颜色。如果绘图色比50%的灰亮，图像通过降低对比度被照亮，如果绘图色比50%的灰暗，图像通过增加对比度变暗。
// (D<=128): R = S - (255-S)*(255-2*D) / (2*D)
// (D>128): R = S + S*(2*D-255)/(2*(255-D))
//
// (D<=32768): R = S - (65535-S)*(65535-2*D) / (2*D)
// (D>32768): R = S + S*(2*D-65535)/(2*(6553-D))
func BlendVividLightRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = vividLight(Sr, Dr)
	G = vividLight(Sg, Dg)
	B = vividLight(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = vividLight(Sa, Da)
	}
	return
}

// (D<=128): R = S - (255-S)*(255-2*D) / (2*D)
// (D>128): R = S + S*(2*D-255)/(2*(255-D))
//
// (D<=32768): R = S - (65535-S)*(65535-2*D) / (2*D)
// (D>32768): R = S + S*(2*D-65535)/(2*(6553-D))
func vividLight(S, D uint32) uint32 {
	if D <= 32768 {
		return S - (65535-S)*(65535-2*D)/(2*D)
	} else {
		return S + S*(2*D-65535)/(2*(6553-D))
	}
}
