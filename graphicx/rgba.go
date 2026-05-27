package graphicx

import (
	"math"
)

// Inverse inverts RGBA channels (65535 - channel).
// Inverse 对 RGBA 各通道取反（65535 - 通道值）。
func Inverse(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB, targetA uint32) {
	max := uint32(math.MaxUint16)
	targetR, targetG, targetB, targetA = max-pixelR, max-pixelG, max-pixelB, max-pixelA
	return
}

// ToNRGBA_White converts premultiplied RGBA to NRGBA using white as the background.
// ToNRGBA_White 以白色为底色将预乘 RGBA 转换为 NRGBA。
func ToNRGBA_White(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB uint32) {
	targetR, targetG, targetB = BlendRGBANormal(pixelR, pixelG, pixelB, pixelA, 65535)
	return
}

// ToNRGBA_Black converts premultiplied RGBA to NRGBA using black as the background.
// ToNRGBA_Black 以黑色为底色将预乘 RGBA 转换为 NRGBA。
func ToNRGBA_Black(pixelR, pixelG, pixelB, pixelA uint32) (targetR, targetG, targetB uint32) {
	targetR, targetG, targetB = BlendRGBANormal(pixelR, pixelG, pixelB, pixelA, 0)
	return
}

// BlendRGBANormal blends foreground (Dr,Dg,Db,Da) over solid background S using normal alpha compositing.
// BlendRGBANormal 以正常模式将前景 (Dr,Dg,Db,Da) 与纯色背景 S 混合。
// With 16-bit image data, R,G,B are in uint16 range. White or black S removes foreground alpha when Da is partial.
// 使用 16 位图像数据时 R、G、B 为 uint16 范围；背景为纯白或纯黑且 Da 非全不透明时可去除前景 Alpha。
// R = D*Da + S*(1-Da)
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
