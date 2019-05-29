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
	RegisterBlendFunc(Multiply, MultiplyBlend)
}

// 正片叠底
// 此模式就象是将两副透明的图像重叠夹在一起放在一张发光的桌子上。
// 将两个颜色的像素值相乘，然后除以255得到的结果就是最终色的像素值。通常执行正片叠底模式后的颜色比原来两种颜色都深。
// 任何颜色和黑色正片叠底得到的仍然是黑色，任何颜色和白色执行正片叠底则保持原来的颜色不变，而与其他颜色执行此模式会产生暗室中以此种颜色照明的效果。
// 在MuItiply模式 中应用较淡的颜色对图像的最终像素颜色没有影响。 MuItiply模式模拟阴影是很捧的。现实 中的阴影从来也不会描绘出比源材料（阴影）或背景（获得阴影的区域）更淡的颜色或色调的 特征。用户将在本章中使用MuItiply模式在恢复的图像中对Lee加入一个下拉阴影。
// 在RGB模式下，每一个像素点的色阶范围是0-255，纯黑色的色阶值是0，纯白色的色阶值是255。
// R = S*D 		[0,1]
// R = S*D/255	[0,255]
func MultiplyBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = MultiplyUnit(source.A, target.A, factor)
	}
	source.R = MultiplyUnit(source.R, target.R, factor)
	source.G = MultiplyUnit(source.G, target.G, factor)
	source.B = MultiplyUnit(source.B, target.B, factor)
	return source
}

// R = S*D 		[0,1]
// R = S*D/255	[0,255]
func MultiplyUnit(S uint8, D uint8, _ float64) uint8 {
	return uint8(uint16(S) * uint16(D) / 255)
}
