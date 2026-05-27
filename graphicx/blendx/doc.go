// Package blendx provides Photoshop- and Core Graphics–style image blend modes for color and RGBA64 channels.
// Package blendx 提供类似 Photoshop / Core Graphics 的图像混合模式，支持 color.Color 与 RGBA64 通道运算。
//
// Created by xuzhuoxi on 2019-05-25.
// @author xuzhuoxi
//
// Capabilities:
//   - BlendMode enum covering normal, darken/lighten, overlay, HSL, Porter-Duff, and advanced modes
//   - BlendColor and BlendRGBA dispatch on BlendMode via registered FuncBlendColor / FuncBlendRGBA callbacks
//   - Per-mode Blend*Color and Blend*RGBA functions (source S, destination D; 16-bit channel range)
//   - RegisterBlendFunc to plug custom implementations into the mode registry at init time
//
// Naming convention (aligned with Photoshop / CGBlendMode):
//   - source (S): background / blend color
//   - destination (D): foreground / base color
//   - result (R): blended output; formulas in function docs use S, D, Da, Sa
//
// 主要职责：
//   - BlendMode 枚举：正常、变暗/变亮、叠加、HSL、Porter-Duff 及高级混合等模式
//   - 通过注册的 FuncBlendColor / FuncBlendRGBA 按 BlendMode 分发 BlendColor、BlendRGBA
//   - 各模式独立的 Blend*Color、Blend*RGBA 实现（源色 S、目标色 D；16 位通道范围）
//   - RegisterBlendFunc 在 init 时向模式注册表挂载实现
//
// 命名约定（与 Photoshop / CGBlendMode 一致）：
//   - source（S）：背景色 / 混合色
//   - destination（D）：前景色 / 基色
//   - result（R）：混合结果；函数注释中的公式使用 S、D、Da、Sa
package blendx
