// Package blendx
// Created by xuzhuoxi
// on 2019-05-25.
// @author xuzhuoxi
//
package blendx

import (
	"image/color"
	"math/rand"
)

func init() {
	RegisterBlendFunc(Dissolve, BlendDissolveColor, BlendDissolveRGBA)
}

// BlendDissolveColor
// 溶解模式
// 最终色和绘图色相同，只是根据每个像素点所在的位置的透明度的不同，可随机以绘图色和底色取代。透明度越大，溶解效果就越明显。
// 使用这种模式，像素仿佛是整个的来自一幅图像或是另一幅，看不出有什么混合的迹象，
// 如果你想在两幅图像之间达到看不出混合迹象的效果，而又能比使用溶解模式拥有更多的对图案的控制，那么可以在最上面图层上建一个图层蒙版并用纯白色填充它。
// 这种效果对模拟破损纸的边缘或原图的 “泼溅”类型是重要的。
// 随机选择一个色
func BlendDissolveColor(S, D color.Color, factor float64, destinationAlpha bool) color.Color {
	if rand.Float64() <= factor {
		return S
	} else {
		return D
	}
}

// BlendDissolveRGBA
// 溶解模式
// 最终色和绘图色相同，只是根据每个像素点所在的位置的透明度的不同，可随机以绘图色和底色取代。透明度越大，溶解效果就越明显。
// 使用这种模式，像素仿佛是整个的来自一幅图像或是另一幅，看不出有什么混合的迹象，
// 如果你想在两幅图像之间达到看不出混合迹象的效果，而又能比使用溶解模式拥有更多的对图案的控制，那么可以在最上面图层上建一个图层蒙版并用纯白色填充它。
// 这种效果对模拟破损纸的边缘或原图的 “泼溅”类型是重要的。
// 随机选择一个色
func BlendDissolveRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, factor float64, _ bool) (R, G, B, A uint32) {
	if rand.Float64() <= factor {
		return Sr, Sg, Sb, Sa
	} else {
		return Dr, Dg, Db, Da
	}
}
