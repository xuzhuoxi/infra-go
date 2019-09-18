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
	RegisterBlendFunc(Screen, BlendScreenColor, BlendScreenRGBA)
}

// 滤色模式
// 作用结果和正片叠底刚好相反，它是将两个颜色的互补色的像素值相乘，然后除以255得到的最终色的像素值。通常执行滤色模式后的颜色都较浅。任何颜色和黑色执行滤色，原色不受影响;任何颜色和白色执行滤色得到的是白色；而与其他颜色执行滤色会产生漂白的效果。
// 此屏幕模式对于在图像中创建霓虹辉光效果是有用的。如果在层上围绕背景对象的边缘涂了白色（或任何淡颜色），然后指定层Screen模式，通过调节层的opacity设置就能 获得饱满或稀薄的辉光效果。
//（附：在Screen和Multipy运算中的重点是----两幅图做Screen运算会加强亮的部分；做Multipy运算则会加强两幅图中暗的部分）
// R = 1 - (1 - B)*(1 - F)
// R = 255 - (255-B)*(255-F) / 255
// R = 65535 - (65535-B)*(65535-F) / 65535
func BlendScreenColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendScreenRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 滤色模式
// 作用结果和正片叠底刚好相反，它是将两个颜色的互补色的像素值相乘，然后除以255得到的最终色的像素值。通常执行滤色模式后的颜色都较浅。任何颜色和黑色执行滤色，原色不受影响;任何颜色和白色执行滤色得到的是白色；而与其他颜色执行滤色会产生漂白的效果。
// 此屏幕模式对于在图像中创建霓虹辉光效果是有用的。如果在层上围绕背景对象的边缘涂了白色（或任何淡颜色），然后指定层Screen模式，通过调节层的opacity设置就能 获得饱满或稀薄的辉光效果。
//（附：在Screen和Multipy运算中的重点是----两幅图做Screen运算会加强亮的部分；做Multipy运算则会加强两幅图中暗的部分）
// R = 1 - (1 - B)*(1 - F)
// R = 255 - (255-B)*(255-F) / 255
// R = 65535 - (65535-B)*(65535-F) / 65535
func BlendScreenRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = screen(foreR, backR)
	G = screen(foreG, backG)
	B = screen(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = screen(foreA, backA)
	}
	return
}

// R = 1 - (1 - B)*(1 - F)
// R = 255 - (255-B)*(255-F) / 255
// R = 65535 - (65535-B)*(65535-F) / 65535
func screen(F, B uint32) uint32 {
	return 65535 - (65535-B)*(65535-F)/65535
}
