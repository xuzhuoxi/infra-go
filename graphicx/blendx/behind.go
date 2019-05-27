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
	RegisterBlendFunc(Behind, BehindBlend)
}

// 背后模式
// 当在有透明区域的图层上操作时背后模式才会出现
func BehindBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if 255 == source.A {
		return color.RGBA{255,255,255,0}
	} else {
		return target
	}
}
