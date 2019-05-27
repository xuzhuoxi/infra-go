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
	RegisterBlendFunc(Add, AddBlend)
}

// 增加模式
// R = Min(255, S+D)
func AddBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if !keepAlpha {
		source.A = AddUnit(source.A, target.A, factor)
	}
	source.R = AddUnit(source.R, target.R, factor)
	source.G = AddUnit(source.G, target.G, factor)
	source.B = AddUnit(source.B, target.B, factor)
	return source
}

// R = Min(255, S+D)
func AddUnit(S uint8, D uint8, _ float64) uint8 {
	Add := uint16(S) + uint16(D)
	if 255 < Add {
		return 255
	} else {
		return uint8(Add)
	}
}
