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

// resultColor		: 结果色
// sourceColor		: 源颜色/背景色/混合色(ps概念)
// destinationColor	：目标色/前景色/基色(ps概念)
type FuncBlendColor func(source, destination color.Color, factor float64, destinationAlpha bool) (resultColor color.Color)
type FuncBlendRGBA func(sourceR, sourceG, sourceB, sourceA uint32, destinationR, destinationG, destinationB, destinationA uint32,
	factor float64, destinationAlpha bool) (R, G, B, A uint32)

var (
	funcBlendColorArr = make([]FuncBlendColor, 128, 128)
	funcBlendRGBAArr  = make([]FuncBlendRGBA, 128, 128)
)

func (m BlendMode) BlendColor(source, destination color.Color, factor float64, destinationAlpha bool) (c color.Color, err error) {
	if funcBlendColor := funcBlendColorArr[m]; nil != funcBlendColor {
		c = funcBlendColor(source, destination, factor, destinationAlpha)
		return
	}
	err = errors.New(fmt.Sprint("BlendMode undefinde: ", m))
	return
}

func (m BlendMode) BlendRGBA(sourceR, sourceG, sourceB, sourceA uint32, destinationR, destinationG, destinationB, destinationA uint32,
	factor float64, destinationAlpha bool) (R, G, B, A uint32, err error) {
	if funcBlendRGB := funcBlendRGBAArr[m]; nil != funcBlendRGB {
		R, G, B, A = funcBlendRGB(sourceR, sourceG, sourceB, destinationR, destinationG, destinationB, destinationA, sourceA, factor, destinationAlpha)
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

//kCGBlendModeNormal,
//kCGBlendModeMultiply,
//kCGBlendModeScreen,
//kCGBlendModeOverlay,
//kCGBlendModeDarken,
//kCGBlendModeLighten,
//kCGBlendModeColorDodge,
//kCGBlendModeColorBurn,
//kCGBlendModeSoftLight,
//kCGBlendModeHardLight,
//kCGBlendModeDifference,
//kCGBlendModeExclusion,
//kCGBlendModeHue,
//kCGBlendModeSaturation,
//kCGBlendModeColor,
//kCGBlendModeLuminosity,
//kCGBlendModeClear,                  /* R = 0 */
//kCGBlendModeCopy,                   /* R = S */
//kCGBlendModeSourceIn,               /* R = S*Da */
//kCGBlendModeSourceOut,              /* R = S*(1 - Da) */
//kCGBlendModeSourceAtop,             /* R = S*Da + D*(1 - Sa) */
//kCGBlendModeDestinationOver,        /* R = S*(1 - Da) + D */
//kCGBlendModeDestinationIn,          /* R = D*Sa */
//kCGBlendModeDestinationOut,         /* R = D*(1 - Sa) */
//kCGBlendModeDestinationAtop,        /* R = S*(1 - Da) + D*Sa */
//kCGBlendModeXOR,                    /* R = S*(1 - Da) + D*(1 - Sa) */
//kCGBlendModePlusDarker,             /* R = MAX(0, (1 - D) + (1 - S)) */
//kCGBlendModePlusLighter             /* R = MIN(1, S + D) */

//Apple额外定义的枚举
//R: premultiplied result, 表示混合结果
//S: Source, 表示源颜色(Sa对应透明度值: 0.0-1.0)
//D: destination colors with alpha, 表示带透明度的目标颜色(Da对应透明度值: 0.0-1.0)
// 	 R表示结果，S表示包含alpha的原色，D表示包含alpha的目标色，Ra，Sa和Da分别是三个的alpha。
//   明白了这些以后，就可以开始寻找我们所需要的blend模式了。相信你可以和我一样，很快找到这个模式
