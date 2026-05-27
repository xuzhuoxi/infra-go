package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(VividLight, BlendVividLightColor, BlendVividLightRGBA)
}

// BlendVividLightColor applies vivid light blend.
// BlendVividLightColor 应用亮光（Vivid Light）混合。
func BlendVividLightColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendVividLightRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendVividLightRGBA applies vivid light on RGBA64 channels.
// BlendVividLightRGBA 对 RGBA64 通道应用亮光混合。
func BlendVividLightRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = vividLight(Sr, Dr)
	G = vividLight(Sg, Dg)
	B = vividLight(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = vividLight(Sa, Da)
	}
	return
}

// (D<=128): R = S - (255-S)*(255-2*D) / (2*D)
// (D>128): R = S + S*(2*D-255)/(2*(255-D))
//
// (D<=32768): R = S - (65535-S)*(65535-2*D) / (2*D)
// (D>32768): R = S + S*(2*D-65535)/(2*(6553-D))
func vividLight(S, D uint32) uint32 {
	if D <= 32768 {
		return S - (65535-S)*(65535-2*D)/(2*D)
	} else {
		return S + S*(2*D-65535)/(2*(6553-D))
	}
}
