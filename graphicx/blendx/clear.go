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
	RegisterBlendFunc(Clear, BlendClearColor, BlendClearRGBA)
}

// 清除模式
// 同背后模式一样，当在图层上操作时，清除模式才会出现。利用清除模式可将图层中有像素的部分清除掉。当有图层时，利用清除模式，使用喷漆桶工具可以将图层中的颜色相近的区域清除掉。
// 可在喷漆桶工具的选项栏中设定“预值”以确定喷漆桶工具所清除的范围。工具选项栏中的“用于所有图层”选项在清除模式下无效。
// R = 0
func BlendClearColor(foreColor, backColor color.Color, _ float64, _ bool) color.Color {
	return &color.RGBA64{}
}

// 清除模式
// 同背后模式一样，当在图层上操作时，清除模式才会出现。利用清除模式可将图层中有像素的部分清除掉。当有图层时，利用清除模式，使用喷漆桶工具可以将图层中的颜色相近的区域清除掉。
// 可在喷漆桶工具的选项栏中设定“预值”以确定喷漆桶工具所清除的范围。工具选项栏中的“用于所有图层”选项在清除模式下无效。
// R = 0
func BlendClearRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, _ bool) (R, G, B, A uint32) {
	return 0, 0, 0, 0
}
