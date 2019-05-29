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
// 最终色和绘图色相同。当在有透明区域的图层上操作时背后模式才会出现，可将绘制的线条放在图层中图像的后面。
// 这模式被用来在一个图层内透明的部分进行涂画；但当图层里的“保持透明区域”选中时就不可用了。 它只可以在你用涂画工具（画笔，喷枪，图章，历史记录画笔，油漆桶）或是填充命令在图层内的一个对象之后画上阴影或色彩。
// 当在有透明区域的图层上操作时背后模式才会出现
func BehindBlend(source color.RGBA, target color.RGBA, factor float64, keepAlpha bool) color.RGBA {
	if 255 == source.A {
		return color.RGBA{255,255,255,0}
	} else {
		return target
	}
}
