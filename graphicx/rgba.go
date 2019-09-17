//
//Created by xuzhuoxi
//on 2019-05-29.
//@author xuzhuoxi
//
package graphicx

import (
	"math"
)

// 反相
func Inverse(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB, targetA uint32) {
	max := uint32(math.MaxUint16)
	targetR, targetG, targetB, targetA = max-pixelR, max-pixelG, max-pixelB, max-pixelA
	return
}

// RGBA颜色转RGB，使用白色作底
func RGBA2RGB_White(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB uint32) {
	targetR, targetG, targetB, _ = BlendPixelNormal(pixelR, pixelG, pixelB, pixelA, math.MaxUint16, math.MaxUint16, math.MaxUint16, math.MaxUint16)
	return
}

// RGBA颜色转RGB，使用黑色作底
func RGBA2RGB_Black(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB uint32) {
	targetR, targetG, targetB, _ = BlendPixelNormal(pixelR, pixelG, pixelB, pixelA, 0, 0, 0, math.MaxUint16)
	return
}

// 混全两个像素，使用normal模式
// 使用纯黑或纯白作背景色，可实现去除前景Alpha通道功能
// 使用64位图像数据，R,G,B的值范围为uint16
//	Target.R = BGColorR *(1-Source.A ) + Source.R*Source.A ;
//	Target.G = BGColorG *(1-Source.A ) + Source.G*Source.A ;
//	Target.B = BGColorB *(1-Source.A ) + Source.B*Source.A ;
//	Target.A = BGColorA *(1-Source.A ) + Source.A*Source.A ;
func BlendPixelNormal(foreR, foreG, foreB, foreA uint32, bgR, bgG, bgB, bgA uint32) (targetR, targetG, targetB, targetA uint32) {
	if math.MaxUint16 == foreA { //前景不透明
		return foreR, foreG, foreB, foreA
	}
	if 0 == foreA { //前景全透明
		return bgR, bgG, bgB, bgA
	}
	rfa := math.MaxUint16 - foreA
	targetR = (bgR*rfa + foreR*foreA) >> 16
	targetG = (bgG*rfa + foreG*foreA) >> 16
	targetB = (bgB*rfa + foreB*foreA) >> 16
	targetA = (bgA*rfa + foreA*foreA) >> 16
	return
}
