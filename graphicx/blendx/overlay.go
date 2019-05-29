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
	RegisterBlendFunc(Overlay, OverlayBlend)
}

// 叠加模式
// 在保留底色明暗变化的基础上使用“正片叠底”或“滤色”模式，绘图的颜色被叠加到底色上，但保留底色的高光和阴影部分。底色的颜色没有被取代，而是和绘图色混合来体现原图的亮部和暗部。
// 使用此模式可使底色的图像的饱和度及对比度得到相应的提高，使图像看起来更加鲜亮。
// 这种模式以一种非艺术逻辑的方式把放置或应用到一个层上的颜色同背景色进行混 合，然而，却能得到有趣的效果。背景图像中的纯黑色或纯白色区域无法在Overlay模式下 显示层上的Overlay着色或图像区域。
// 背景区域上落在黑色和白色之间的亮度值同0verlay 材料的颜色混合在一起，产生最终的合成颜色。为了使背景图像看上去好像是同设计或文本 一起拍摄的，Overlay可用来在背景图像上画上一个设计或文本。
// S <= 128 : R = S * D / 128
// S > 128 : R = 255 - (255 - S) * (255 - D) / 128
func OverlayBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = OverlayUnit(source.A, target.A, factor)
	}
	source.R = OverlayUnit(source.R, target.R, factor)
	source.G = OverlayUnit(source.G, target.G, factor)
	source.B = OverlayUnit(source.B, target.B, factor)
	return source
}

// S <= 128 : R = S * D / 128
// S > 128 : R = 255 - (255 - S) * (255 - D) / 128
func OverlayUnit(S uint8, D uint8, _ float64) uint8 {
	if S <= 128 {
		return uint8((uint16(S) * uint16(D)) >> 7)
	} else {
		return uint8(255 - (uint16(255-S)*uint16(255-D))>>7)
	}
}
