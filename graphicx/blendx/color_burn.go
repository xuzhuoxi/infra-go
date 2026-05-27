package blendx

import (
	"image/color"
	"math"
)

func init() {
	RegisterBlendFunc(ColorBurn, BlendColorBurnColor, BlendColorBurnRGBA)
}

// BlendColorBurnColor applies color burn blend.
// BlendColorBurnColor 应用颜色加深（Color Burn）混合。
// R = S - (1-S) * (1-D) / D
// R = S - (255-S) * (255-D) / D
// R = S - (65535-S) * (65535-D) / D (64位图像)
func BlendColorBurnColor(S, D color.Color, factor float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendColorBurnRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.NRGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendColorBurnRGBA applies color burn on RGBA64 channels.
// BlendColorBurnRGBA 对 RGBA64 通道应用颜色加深混合。
// R = S - (1-S) * (1-D) / D
// R = S - (255-S) * (255-D) / D
// R = S - (65535-S) * (65535-D) / D (64位图像)
func BlendColorBurnRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = colorBurn(Sr, Dr)
	G = colorBurn(Sg, Dg)
	B = colorBurn(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = colorBurn(Sa, Da)
	}
	return
}

// R = S - (1-S) * (1-D) / D
// R = S - (255-S) * (255-D) / D
// R = S - (65535-S) * (65535-D) / D (64位图像)
func colorBurn(S, D uint32) uint32 {
	return S - (math.MaxUint16-S)*(math.MaxUint16-D)/D
}
