package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(ColorDodge, BlendColorDodgeColor, BlendColorDodgeRGBA)
}

// BlendColorDodgeColor applies color dodge blend.
// BlendColorDodgeColor 应用颜色减淡（Color Dodge）混合。
// R = S + S*D/(1-D)
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func BlendColorDodgeColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendAddRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.NRGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendColorDodgeRGBA applies color dodge on RGBA64 channels.
// BlendColorDodgeRGBA 对 RGBA64 通道应用颜色减淡混合。
// R = S + S*D/(1-D)
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func BlendColorDodgeRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = colorDodge(Sr, Dr)
	G = colorDodge(Sg, Dg)
	B = colorDodge(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = colorDodge(Sa, Da)
	}
	return
}

// R = S + S*D/(1-D)
// R = S + S*D/(255-D) (32位)
// R = S + S*D/(65535-D) (64位)
func colorDodge(S, D uint32) uint32 {
	return S + S*D/(65535-D)
}
