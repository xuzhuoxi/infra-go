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
	RegisterBlendFunc(LinearDodge, LinearDodgeBlend)
}

// 线性减淡模式
// 查看每个通道的颜色信息，通过增加“亮度”使底色的颜色变亮来反映绘图色，和黑色混合没变化。
// R = S + D
func LinearDodgeBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = LinearDodgeUnit(source.A, target.A, factor)
	}
	source.R = LinearDodgeUnit(source.R, target.R, factor)
	source.G = LinearDodgeUnit(source.G, target.G, factor)
	source.B = LinearDodgeUnit(source.B, target.B, factor)
	return source
}

// R = S + D
func LinearDodgeUnit(S uint8, D uint8, _ float64) uint8 {
	Add := uint16(S) + uint16(D)
	if 255 < Add {
		return 255
	} else {
		return uint8(Add)
	}
}
