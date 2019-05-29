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
	RegisterBlendFunc(ColorBurn, ColorBurnBlend)
}

// 颜色加深模式
// 查看每个通道的颜色信息，通过增加“对比度”使底色的颜色变暗来反映绘图色，和白色混合没变化。
// 除了背景上的较淡区域消失，且图像区域呈现尖锐的边缘特性之外，这种Color Burn模式创建的效果类似于由MuItiply模式创建的效果。
// R = S - ((1 - S) * (1 - D)) / D
// R = S - (255-S)*(255-D) / D
func ColorBurnBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ColorBurnUnit(source.A, target.A, factor)
	}
	source.R = ColorBurnUnit(source.R, target.R, factor)
	source.G = ColorBurnUnit(source.G, target.G, factor)
	source.B = ColorBurnUnit(source.B, target.B, factor)
	return source
}

// R = S - ((1 - S) * (1 - D)) / D
// R = S - (255-S)*(255-D) / D
func ColorBurnUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	return uint8(S16 - (255-S16)*(255-D16)/D16)
}
