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
	RegisterBlendFunc(HardMix, BlendHardMixColor, BlendHardMixRGBA)
}

// 实色混合模式
// 根据绘图颜色与底图颜色的颜色数值相加，当相加的颜色数值大于该颜色模式颜色数值的最大值，混合颜色为最大值；
// 当相加的颜色数值小于该颜色模式颜色数值的最大值，混合颜色为0；
// 当相加的颜色数值等于该颜色模式颜色数值的最大值，混合颜色由底图颜色决定，底图颜色值比绘图颜色的颜色值大，则混合颜色为最大值，相反则为0.实色混合能产生颜色较少、边缘较硬的图像效果。
// R = S+D>=1 ? 1 : 0
// R = S+D>=255 ? 255 : 0
// R = S+D>=65535 ? 65535 : 0
func BlendHardMixColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendHardMixRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 实色混合模式
// 根据绘图颜色与底图颜色的颜色数值相加，当相加的颜色数值大于该颜色模式颜色数值的最大值，混合颜色为最大值；
// 当相加的颜色数值小于该颜色模式颜色数值的最大值，混合颜色为0；
// 当相加的颜色数值等于该颜色模式颜色数值的最大值，混合颜色由底图颜色决定，底图颜色值比绘图颜色的颜色值大，则混合颜色为最大值，相反则为0.实色混合能产生颜色较少、边缘较硬的图像效果。
// R = S+D>=1 ? 1 : 0
// R = S+D>=255 ? 255 : 0
// R = S+D>=65535 ? 65535 : 0
func BlendHardMixRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = hardMix(Sr, Dr)
	G = hardMix(Sg, Dg)
	B = hardMix(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = hardMix(Sa, Da)
	}
	return
}

// R = S+D>=1 ? 1 : 0
// R = S+D>=255 ? 255 : 0
// R = S+D>=65535 ? 65535 : 0
func hardMix(S, D uint32) uint32 {
	if S+D >= 65535 {
		return 65535
	} else {
		return 0
	}
}
