package blendx

import (
	"image/color"
)

func init() {
	RegisterBlendFunc(LinearDodge, BlendLinearDodgeColor, BlendLinearDodgeRGBA)
}

// BlendLinearDodgeColor applies linear dodge blend.
// BlendLinearDodgeColor 应用线性减淡（Linear Dodge）混合。
// R = S + D
func BlendLinearDodgeColor(S, D color.Color, _ float64, destinationAlpha bool) color.Color {
	Sr, Sg, Sb, Sa := S.RGBA()
	Dr, Dg, Db, Da := D.RGBA()
	R, G, B, A := BlendLinearDodgeRGBA(Sr, Sg, Sb, Sa, Dr, Dg, Db, Da, 0, destinationAlpha)
	return &color.RGBA64{R: uint16(R), G: uint16(G), B: uint16(B), A: uint16(A)}
}

// BlendLinearDodgeRGBA applies linear dodge on RGBA64 channels.
// BlendLinearDodgeRGBA 对 RGBA64 通道应用线性减淡混合。
// R = S + D
func BlendLinearDodgeRGBA(Sr, Sg, Sb, Sa uint32, Dr, Dg, Db, Da uint32, _ float64, destinationAlpha bool) (R, G, B, A uint32) {
	R = linearDodge(Sr, Dr)
	G = linearDodge(Sg, Dg)
	B = linearDodge(Sb, Db)
	if destinationAlpha {
		A = Da
	} else {
		A = linearDodge(Sa, Da)
	}
	return
}

// R = S + D
func linearDodge(S, D uint32) uint32 {
	Add := S + D
	if Add <= 65535 {
		return Add
	} else {
		return 255
	}
}
