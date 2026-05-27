// Package graphicx provides pixel packing, color space conversion, grayscale, and RGBA compositing utilities.
// Package graphicx 提供像素打包、色彩空间转换、灰度化与 RGBA 合成等图像处理工具。
//
// Created by xuzhuoxi on 2019-05-29.
// @author xuzhuoxi
//
// Capabilities:
//   - Pack and unpack ARGB/RGBA pixels (8-bit and 16-bit) and convert to/from image/color
//   - Convert between 16-bit, 8-bit, and normalized float color component representations
//   - Extract RGBA channel values from image/color types
//   - HSV ↔ RGB conversion and per-channel hue, saturation, and value extraction
//   - Grayscale conversion with multiple luminance algorithms (GrayAlgMode)
//   - YUV luminance and chrominance component computation
//   - Premultiplied RGBA blending, color inversion, and NRGBA conversion on white/black backgrounds
//
// 主要职责：
//   - ARGB/RGBA 像素（8/16 位）打包解包及与 image/color 互转
//   - 16 位、8 位与归一化浮点色阶分量之间的转换
//   - 从 image/color 提取 RGBA 通道值
//   - HSV 与 RGB 互转，以及色相、饱和度、明度分量计算
//   - 多种亮度算法的灰度化（GrayAlgMode）
//   - YUV 明亮度与色度分量计算
//   - 预乘 RGBA 混合、反相及基于白/黑底的 NRGBA 转换
package graphicx
