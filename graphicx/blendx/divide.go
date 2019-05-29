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
	RegisterBlendFunc(Divide, DivideBlend)
}

// 划分模式
// 假设上面图层选择划分，那么所看到的图像是，下面的可见图层根据上面这个图层颜色的纯度，相应减去了同等纯度的该颜色，同时上面颜色的明暗度不同，被减去区域图像明度也不同，上面图层颜色的亮，图像亮度变化就会越小，上面图层越暗，被减区域图像就会越亮。
// 也就是说，如果上面图层是白色，那么也不会减去颜色也不会提高明度，如果上面图层是黑色，那么所有不纯的颜色都会被减去，只留着最纯的光的三原色，及其混合色，青品黄与白色。
// C = 255 * A / B
func DivideBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = DivideUnit(source.A, target.A, factor)
	}
	source.R = DivideUnit(source.R, target.R, factor)
	source.G = DivideUnit(source.G, target.G, factor)
	source.B = DivideUnit(source.B, target.B, factor)
	return source
}

// C = 255 * A / B
func DivideUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(255 * uint16(S) / uint16(D))
}
