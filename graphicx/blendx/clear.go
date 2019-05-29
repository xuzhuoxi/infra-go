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
	RegisterBlendFunc(Clear, ClearBlend)
}

// 清除模式
// 同背后模式一样，当在图层上操作时，清除模式才会出现。利用清除模式可将图层中有像素的部分清除掉。当有图层时，利用清除模式，使用喷漆桶工具可以将图层中的颜色相近的区域清除掉。
// 可在喷漆桶工具的选项栏中设定“预值”以确定喷漆桶工具所清除的范围。工具选项栏中的“用于所有图层”选项在清除模式下无效。
// R = 0
func ClearBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ClearUnit(source.A, target.A, factor)
	}
	source.R = ClearUnit(source.R, target.R, factor)
	source.G = ClearUnit(source.G, target.G, factor)
	source.B = ClearUnit(source.B, target.B, factor)
	return source
}

// R = 0
func ClearUnit(S uint8, D uint8, _ float64) uint8 {
	return 0
}
