//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"image/color"
	"github.com/xuzhuoxi/infra-go/mathx/randx"
)

func init() {
	RegisterBlendFunc(Dissolve, DissolveBlend)
}

// 溶解模式
// 随机选择一个色
func DissolveBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if randx.RandBool() {
		return source
	} else {
		return target
	}
}
