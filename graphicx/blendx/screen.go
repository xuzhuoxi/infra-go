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
	RegisterBlendFunc(Screen, ScreenBlend)
}

// 滤色模式
// 作用结果和正片叠底刚好相反，它是将两个颜色的互补色的像素值相乘，然后除以255得到的最终色的像素值。通常执行滤色模式后的颜色都较浅。任何颜色和黑色执行滤色，原色不受影响;任何颜色和白色执行滤色得到的是白色；而与其他颜色执行滤色会产生漂白的效果。
// 此屏幕模式对于在图像中创建霓虹辉光效果是有用的。如果在层上围绕背景对象的边缘涂了白色（或任何淡颜色），然后指定层Screen模式，通过调节层的opacity设置就能 获得饱满或稀薄的辉光效果。
//（附：在Screen和Multipy运算中的重点是----两幅图做Screen运算会加强亮的部分；做Multipy运算则会加强两幅图中暗的部分）
// R = 1 - (1 - S)*(1 - D)
// R = 255 - (255-S)*(255-D) / 255
func ScreenBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = ScreenUnit(source.A, target.A, factor)
	}
	source.R = ScreenUnit(source.R, target.R, factor)
	source.G = ScreenUnit(source.G, target.G, factor)
	source.B = ScreenUnit(source.B, target.B, factor)
	return source
}

// R = 1 - (1 - S)*(1 - D)
// R = 255 - (255-S)*(255-D) / 255
func ScreenUnit(S uint8, D uint8, _ float64) uint8 {
	S16 := uint16(S)
	D16 := uint16(D)
	return uint8(255 - (255-S16)*(255-D16)/255)
}
