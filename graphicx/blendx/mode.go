package blendx

import (
	"errors"
	"fmt"
	"image/color"
)

// BlendMode identifies a registered image blend mode.
// BlendMode 标识已注册的图像混合模式。
type BlendMode int

// FuncBlendColor blends two color.Color values.
// FuncBlendColor 混合两个 color.Color 值。
// resultColor is the blended result.
// resultColor 为混合结果色。
// source is the source/background/blend color (Photoshop terminology).
// source 为源颜色/背景色/混合色（Photoshop 概念）。
// destination is the destination/foreground/base color (Photoshop terminology).
// destination 为目标色/前景色/基色（Photoshop 概念）。
type FuncBlendColor func(source, destination color.Color, factor float64, destinationAlpha bool) (resultColor color.Color)

// FuncBlendRGBA blends two RGBA64 channel tuples (16-bit per channel).
// FuncBlendRGBA 混合两组 RGBA64 通道（每通道 16 位）。
type FuncBlendRGBA func(sourceR, sourceG, sourceB, sourceA uint32, destinationR, destinationG, destinationB, destinationA uint32,
	factor float64, destinationAlpha bool) (R, G, B, A uint32)

var (
	funcBlendColorArr = make([]FuncBlendColor, 128, 128)
	funcBlendRGBAArr  = make([]FuncBlendRGBA, 128, 128)
)

// BlendColor applies blend mode m to source and destination colors.
// BlendColor 对源色与目标色应用混合模式 m。
func (m BlendMode) BlendColor(source, destination color.Color, factor float64, destinationAlpha bool) (c color.Color, err error) {
	if funcBlendColor := funcBlendColorArr[m]; nil != funcBlendColor {
		c = funcBlendColor(source, destination, factor, destinationAlpha)
		return
	}
	err = errors.New(fmt.Sprint("BlendMode undefinde: ", m))
	return
}

// BlendRGBA applies blend mode m to source and destination RGBA64 channels.
// BlendRGBA 对源与目标 RGBA64 通道应用混合模式 m。
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
	// None indicates no blend mode.
	// None 表示无混合模式。
	None BlendMode = iota

	//-----------------------------------

	// Normal is normal blend mode (implemented).
	// Normal 正常模式（已实现）。
	Normal
	// NormalThreshold is threshold blend mode (implemented).
	// NormalThreshold 阈值模式（已实现）。
	NormalThreshold
	// Dissolve is dissolve blend mode (implemented).
	// Dissolve 溶解模式（已实现）。
	Dissolve
	// Behind is behind blend mode (implemented).
	// Behind 背后模式（已实现）。
	Behind
	// Clear is clear blend mode (implemented).
	// Clear 清除模式（已实现）。
	Clear
	// Copy is copy blend mode (implemented).
	// Copy 覆盖模式（已实现）。
	Copy

	//-----------------------------------

	// Darken is darken blend mode (implemented).
	// Darken 变暗模式（已实现）。
	Darken
	// Multiply is multiply blend mode (implemented).
	// Multiply 正片叠底（已实现）。
	Multiply
	// ColorBurn is color burn blend mode (implemented).
	// ColorBurn 颜色加深模式（已实现）。
	ColorBurn
	// LinearBurn is linear burn blend mode (implemented).
	// LinearBurn 线性加深模式（已实现）。
	LinearBurn
	// DarkerColor is darker color blend mode.
	// DarkerColor 深色模式。
	DarkerColor

	//-----------------------------------

	// Add is add blend mode (implemented).
	// Add 增加模式（已实现）。
	Add
	// Lighten is lighten blend mode (implemented).
	// Lighten 变亮模式（已实现）。
	Lighten
	// Screen is screen blend mode (implemented).
	// Screen 滤色模式（已实现）。
	Screen
	// ColorDodge is color dodge blend mode (implemented).
	// ColorDodge 颜色减淡模式（已实现）。
	ColorDodge
	// LinearDodge is linear dodge blend mode (implemented).
	// LinearDodge 线性减淡模式（已实现）。
	LinearDodge
	// LighterColor is lighter color blend mode.
	// LighterColor 浅色模式。
	LighterColor

	//-----------------------------------

	// Overlay is overlay blend mode (implemented).
	// Overlay 叠加模式（已实现）。
	Overlay
	// SoftLight is soft light blend mode (implemented).
	// SoftLight 柔光模式（已实现）。
	SoftLight
	// HardLight is hard light blend mode (implemented).
	// HardLight 强光模式（已实现）。
	HardLight
	// VividLight is vivid light blend mode (implemented).
	// VividLight 亮光模式（已实现）。
	VividLight
	// LinearLight is linear light blend mode (implemented).
	// LinearLight 线性光模式（已实现）。
	LinearLight
	// PinLight is pin light blend mode (implemented).
	// PinLight 点光模式（已实现）。
	PinLight
	// HardMix is hard mix blend mode (implemented).
	// HardMix 实色混合模式（已实现）。
	HardMix

	//-----------------------------------

	// Difference is difference blend mode (implemented).
	// Difference 差值模式（已实现）。
	Difference
	// Exclusion is exclusion blend mode (implemented).
	// Exclusion 排除模式（已实现）。
	Exclusion
	// Subtract is subtract blend mode (implemented).
	// Subtract 减去模式（已实现）。
	Subtract
	// Divide is divide blend mode (implemented).
	// Divide 划分模式（已实现）。
	Divide

	//-----------------------------------

	// Hue is hue blend mode.
	// Hue 色相模式。
	Hue
	// Saturation is saturation blend mode.
	// Saturation 饱和度模式。
	Saturation
	// Color is color blend mode.
	// Color 颜色模式。
	Color
	// Luminosity is luminosity blend mode.
	// Luminosity 亮度模式。
	Luminosity

	//-----------------------------------

	// DestinationAtop is destination atop Porter-Duff mode (implemented).
	// DestinationAtop 目标色在上混合（已实现）。
	DestinationAtop
	// DestinationIn is destination in Porter-Duff mode (implemented).
	// DestinationIn 目标色在内混合（已实现）。
	DestinationIn
	// DestinationOut is destination out Porter-Duff mode (implemented).
	// DestinationOut 目标色在外混合（已实现）。
	DestinationOut
	// DestinationOver is destination over Porter-Duff mode (implemented).
	// DestinationOver 目标色覆盖混合（已实现）。
	DestinationOver
	// PlusDarker is plus darker blend mode (implemented).
	// PlusDarker 高级深色（已实现）。
	PlusDarker
	// PlusLighter is plus lighter blend mode (implemented).
	// PlusLighter 高级浅色（已实现）。
	PlusLighter
	// SourceAtop is source atop Porter-Duff mode (implemented).
	// SourceAtop 源色在上混合（已实现）。
	SourceAtop
	// SourceIn is source in Porter-Duff mode (implemented).
	// SourceIn 源色在内混合（已实现）。
	SourceIn
	// SourceOut is source out Porter-Duff mode (implemented).
	// SourceOut 源色在外混合（已实现）。
	SourceOut
	// SourceOver is source over Porter-Duff mode (implemented).
	// SourceOver 源色覆盖混合（已实现）。
	SourceOver
	// Xor is XOR Porter-Duff mode (implemented).
	// Xor 异或模式（已实现）。
	Xor
)

// RegisterBlendFunc registers color and RGBA blend callbacks for blendMode.
// RegisterBlendFunc 为 blendMode 注册 color 与 RGBA 混合回调。
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
