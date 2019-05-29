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
	RegisterBlendFunc(LinearBurn, LinearBurnBlend)
}

// 线性加深模式
// 查看每个通道的颜色信息，通过降低“亮度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// R = S + D - 255
func LinearBurnBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = LinearBurnUnit(source.A, target.A, factor)
	}
	source.R = LinearBurnUnit(source.R, target.R, factor)
	source.G = LinearBurnUnit(source.G, target.G, factor)
	source.B = LinearBurnUnit(source.B, target.B, factor)
	return source
}

// R = S + D - 255
func LinearBurnUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(uint16(S) + uint16(D) - 255)
}
