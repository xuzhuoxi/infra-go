//
//Created by xuzhuoxi
//on 2019-05-25.
//@author xuzhuoxi
//
package blendx

import (
	"errors"
	"fmt"
	"image/color"
)

type BlendMode int

type FuncBlendColor func(foregroundColor, backgroundColor color.Color, factor float64, keepForegroundAlpha bool) color.Color
type FuncBlendRGBA func(foregroundR, foregroundG, foregroundB, foregroundA uint32,
	backgroundR, backgroundG, backgroundB, backgroundA uint32,
	factor float64, keepForegroundAlpha bool) (R, G, B, A uint32)

var (
	funcBlendColorArr = make([]FuncBlendColor, 128, 128)
	funcBlendRGBAArr  = make([]FuncBlendRGBA, 128, 128)
)

func (m BlendMode) BlendColor(foregroundColor, backgroundColor color.Color, factor float64, keepForegroundAlpha bool) (c color.Color, err error) {
	if funcBlendColor := funcBlendColorArr[m]; nil != funcBlendColor {
		c = funcBlendColor(foregroundColor, backgroundColor, factor, keepForegroundAlpha)
		return
	}
	err = errors.New(fmt.Sprint("BlendMode undefinde: ", m))
	return
}

func (m BlendMode) BlendRGBA(foregroundR, foregroundG, foregroundB, foregroundA uint32,
	backgroundR, backgroundG, backgroundB, backgroundA uint32,
	factor float64, keepForegroundAlpha bool) (R, G, B, A uint32, err error) {
	if funcBlendRGB := funcBlendRGBAArr[m]; nil != funcBlendRGB {
		R, G, B, A = funcBlendRGB(foregroundR, foregroundG, foregroundB, foregroundA, backgroundR, backgroundG, backgroundB, backgroundA, factor, keepForegroundAlpha)
		return
	}
	err = errors.New(fmt.Sprint("BlendMode undefinde: ", m))
	return
}

const (
	// 无
	None BlendMode = iota

	//-----------------------------------

	// 正常模式(已实现)
	Normal
	// 阈值模式(已实现)
	NormalThreshold
	// 溶解模式(已实现)
	Dissolve
	// 背后模式(已实现)
	Behind
	// 清除模式(已实现)
	Clear
	// 覆盖模式(已实现)
	Copy

	//-----------------------------------

	// 变暗模式(已实现)
	Darken
	// 正片叠底(已实现)
	Multiply
	// 颜色加深模式(已实现)
	ColorBurn
	// 线性加深模式(已实现)
	LinearBurn
	// 深色模式----------------------------------------
	DarkerColor

	//-----------------------------------

	// 增加模式(已实现)
	Add
	// 变亮模式(已实现)
	Lighten
	// 滤色模式(已实现)
	Screen
	// 颜色减淡模式(已实现)
	ColorDodge
	// 线性减淡模式(已实现)
	LinearDodge
	// 浅色模式----------------------------------
	LighterColor

	//-----------------------------------

	// 叠加模式(已实现)
	Overlay
	// 柔光模式(已实现)
	SoftLight
	// 强光模式(已实现)
	HardLight
	// 亮光模式(已实现)
	VividLight
	// 线性光模式(已实现)
	LinearLight
	// 点光模式(已实现)
	PinLight
	// 实色混合模式(已实现)
	HardMix

	//-----------------------------------

	// 差值模式(已实现)
	Difference
	// 排除模式(已实现)
	Exclusion
	// 减去模式(已实现)
	Subtract
	// 划分模式(已实现)
	Divide

	//-----------------------------------

	// 色相模式
	Hue
	// 饱和度模式
	Saturation
	// 颜色模式
	Color
	// 亮度模式
	Luminosity

	//-----------------------------------

	// (已实现)
	DestinationAtop
	// (已实现)
	DestinationIn
	// (已实现)
	DestinationOut
	// (已实现)
	DestinationOver
	// 高级深色(已实现)
	PlusDarker
	// 高级浅色(已实现)
	PlusLighter
	// (已实现)
	SourceAtop
	// (已实现)
	SourceIn
	// (已实现)
	SourceOut
	// (已实现)
	SourceOver
	// 异或模式(已实现)
	Xor
)

func RegisterBlendFunc(blendMode BlendMode, funcBlendColor FuncBlendColor, funcBlendRGBA FuncBlendRGBA) {
	funcBlendColorArr[blendMode] = funcBlendColor
	funcBlendRGBAArr[blendMode] = funcBlendRGBA
}
