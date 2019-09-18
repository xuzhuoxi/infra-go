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
	RegisterBlendFunc(Multiply, BlendMultiplyColor, BlendMultiplyRGBA)
}

// 正片叠底
// 此模式就象是将两副透明的图像重叠夹在一起放在一张发光的桌子上。
// 将两个颜色的像素值相乘，然后除以255得到的结果就是最终色的像素值。通常执行正片叠底模式后的颜色比原来两种颜色都深。
// 任何颜色和黑色正片叠底得到的仍然是黑色，任何颜色和白色执行正片叠底则保持原来的颜色不变，而与其他颜色执行此模式会产生暗室中以此种颜色照明的效果。
// 在Multiply模式 中应用较淡的颜色对图像的最终像素颜色没有影响。 MuItiply模式模拟阴影是很捧的。现实 中的阴影从来也不会描绘出比源材料（阴影）或背景（获得阴影的区域）更淡的颜色或色调的 特征。用户将在本章中使用MuItiply模式在恢复的图像中对Lee加入一个下拉阴影。
// 在RGB模式下，每一个像素点的色阶范围是0-255(或0-65535)，纯黑色的色阶值是0，纯白色的色阶值是255。
// R = B*F 		[0,1]
// R = B*F/255	[0,255]
// R = B*F/65535	[0,65535]
func BlendMultiplyColor(foreColor, backColor color.Color, _ float64, keepForegroundAlpha bool) color.Color {
	fR, fG, fB, fA := foreColor.RGBA()
	bR, bG, bB, bA := backColor.RGBA()
	R, G, B, A := BlendMultiplyRGBA(fR, fG, fB, fA, bR, bG, bB, bA, 0, keepForegroundAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// 正片叠底
// 此模式就象是将两副透明的图像重叠夹在一起放在一张发光的桌子上。
// 将两个颜色的像素值相乘，然后除以255得到的结果就是最终色的像素值。通常执行正片叠底模式后的颜色比原来两种颜色都深。
// 任何颜色和黑色正片叠底得到的仍然是黑色，任何颜色和白色执行正片叠底则保持原来的颜色不变，而与其他颜色执行此模式会产生暗室中以此种颜色照明的效果。
// 在Multiply模式 中应用较淡的颜色对图像的最终像素颜色没有影响。 MuItiply模式模拟阴影是很捧的。现实 中的阴影从来也不会描绘出比源材料（阴影）或背景（获得阴影的区域）更淡的颜色或色调的 特征。用户将在本章中使用MuItiply模式在恢复的图像中对Lee加入一个下拉阴影。
// 在RGB模式下，每一个像素点的色阶范围是0-255(或0-65535)，纯黑色的色阶值是0，纯白色的色阶值是255。
// R = B*F 		[0,1]
// R = B*F/255	[0,255]
// R = B*F/65535	[0,65535]
func BlendMultiplyRGBA(foreR, foreG, foreB, foreA uint32, backR, backG, backB, backA uint32, _ float64, keepForegroundAlpha bool) (R, G, B, A uint32) {
	R = multiply(foreR, backR)
	G = multiply(foreG, backG)
	B = multiply(foreB, backB)
	if keepForegroundAlpha {
		A = foreA
	} else {
		A = multiply(foreA, backA)
	}
	return
}

// R = B*F 		[0,1]
// R = B*F/255	[0,255]
// R = B*F/65535	[0,65535]
func multiply(F, B uint32) uint32 {
	return B * F / 65535
}
