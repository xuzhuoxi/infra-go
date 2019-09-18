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
	RegisterBlendFunc(Overlay, BlendOverlayColor, BlendOverlayRGBA)
}

// 叠加模式
// 在保留底色明暗变化的基础上使用“正片叠底”或“滤色”模式，绘图的颜色被叠加到底色上，但保留底色的高光和阴影部分。底色的颜色没有被取代，而是和绘图色混合来体现原图的亮部和暗部。
// 使用此模式可使底色的图像的饱和度及对比度得到相应的提高，使图像看起来更加鲜亮。
// 这种模式以一种非艺术逻辑的方式把放置或应用到一个层上的颜色同背景色进行混 合，然而，却能得到有趣的效果。背景图像中的纯黑色或纯白色区域无法在Overlay模式下 显示层上的Overlay着色或图像区域。
// 背景区域上落在黑色和白色之间的亮度值同0verlay 材料的颜色混合在一起，产生最终的合成颜色。为了使背景图像看上去好像是同设计或文本 一起拍摄的，Overlay可用来在背景图像上画上一个设计或文本。
// R = B<=0.5 ? 2*B*F : 1-2*(1-B)*(1-F)
// R = B<=128 ? B*F/128 : 255-(255-B)*(255-F)/128
// R = B<=32768 ? B*F/32768 : 65535-(65535-B)*(65534-F)/32768
func BlendOverlayColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendOverlayRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 叠加模式
// 在保留底色明暗变化的基础上使用“正片叠底”或“滤色”模式，绘图的颜色被叠加到底色上，但保留底色的高光和阴影部分。底色的颜色没有被取代，而是和绘图色混合来体现原图的亮部和暗部。
// 使用此模式可使底色的图像的饱和度及对比度得到相应的提高，使图像看起来更加鲜亮。
// 这种模式以一种非艺术逻辑的方式把放置或应用到一个层上的颜色同背景色进行混 合，然而，却能得到有趣的效果。背景图像中的纯黑色或纯白色区域无法在Overlay模式下 显示层上的Overlay着色或图像区域。
// 背景区域上落在黑色和白色之间的亮度值同0verlay 材料的颜色混合在一起，产生最终的合成颜色。为了使背景图像看上去好像是同设计或文本 一起拍摄的，Overlay可用来在背景图像上画上一个设计或文本。
// R = B<=0.5 ? 2*B*F : 1-2*(1-B)*(1-F)
// R = B<=128 ? B*F/128 : 255-(255-B)*(255-F)/128
// R = B<=32768 ? B*F/32768 : 65535-(65535-B)*(65534-F)/32768
func BlendOverlayRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = overlay(foreR, backR)
	G = overlay(foreG, backG)
	B = overlay(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = overlay(foreA, backA)
	}
	return
}

// R = B<=0.5 ? 2*B*F : 1-2*(1-B)*(1-F)
// R = B<=128 ? B*F/128 : 255-(255-B)*(255-F)/128
// R = B<=32768 ? B*F/32768 : 65535-(65535-B)*(65534-F)/32768
func overlay(F, B uint32) uint32 {
	if B <= 32768 {
		return B * F / 32768
	} else {
		return 65535 - (65535-B)*(65534-F)/32768
	}
}
