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
	RegisterBlendFunc(HardMix, HardMixBlend)
}

// 实色混合模式
// 根据绘图颜色与底图颜色的颜色数值相加，当相加的颜色数值大于该颜色模式颜色数值的最大值，混合颜色为最大值；
// 当相加的颜色数值小于该颜色模式颜色数值的最大值，混合颜色为0；
// 当相加的颜色数值等于该颜色模式颜色数值的最大值，混合颜色由底图颜色决定，底图颜色值比绘图颜色的颜色值大，则混合颜色为最大值，相反则为0.实色混合能产生颜色较少、边缘较硬的图像效果。
// S+D>=255 : R = 255
// S+D<255 : R = 0
func HardMixBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = HardMixUnit(source.A, target.A, factor)
	}
	source.R = HardMixUnit(source.R, target.R, factor)
	source.G = HardMixUnit(source.G, target.G, factor)
	source.B = HardMixUnit(source.B, target.B, factor)
	return source
}

// S+D>=255 : R = 255
// S+D<255 : R = 0
func HardMixUnit(S uint8, D uint8, _ float64) uint8 {
	if uint16(S)+uint16(D) >= 255 {
		return 255
	} else {
		return 0
	}
}
