// Package graphicx
// Created by xuzhuoxi
// on 2019-05-29.
// @author xuzhuoxi
//
package graphicx

import (
	"math"
)

// Inverse 反相
func Inverse(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB, targetA uint32) {
	max := uint32(math.MaxUint16)
	targetR, targetG, targetB, targetA = max-pixelR, max-pixelG, max-pixelB, max-pixelA
	return
}

// ToNRGBA_White
// RGBA颜色转NRGBA，使用白色作底
func ToNRGBA_White(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB uint32) {
	targetR, targetG, targetB = BlendRGBANormal(pixelR, pixelG, pixelB, pixelA, 65535)
	return
}

// ToNRGBA_Black
// RGBA颜色转NRGBA，使用黑色作底
func ToNRGBA_Black(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB uint32) {
	targetR, targetG, targetB = BlendRGBANormal(pixelR, pixelG, pixelB, pixelA, 0)
	return
}

// BlendRGBANormal
// 混全两个像素，使用normal模式
// 使用纯黑或纯白作背景色，可实现去除前景Alpha通道功能
// 使用64位图像数据，R,G,B的值范围为uint16
// R ＝ D*Da + S*(1-Da)
func BlendRGBANormal(Dr, Dg, Db, Da uint32, S uint32) (targetR, targetG, targetB uint32) {
	if math.MaxUint16 == Da { //前景不透明
		return Dr, Dg, Db
	}
	if 0 == Da { //前景全透明
		return S, S, S
	}
	rfa := math.MaxUint16 - Da
	targetR = (S*rfa + Dr*Da) / 65535
	targetG = (S*rfa + Dg*Da) / 65535
	targetB = (S*rfa + Db*Da) / 65535
	return
}
