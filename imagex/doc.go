// Package imagex provides foundational image I/O, pixel buffers, copying, filling, blending, and eight-direction utilities.
// Package imagex 提供图像读写、像素缓冲区、复制填充、背景混合及八邻域方向等基础图像处理能力。
//
// Created by xuzhuoxi on 2019-05-31.
// @author xuzhuoxi
//
// Capabilities:
//   - LoadImage / SaveImage with formatx format registry (auto-detect or explicit format)
//   - Construct Gray, Gray16, RGBA, NRGBA, RGBA64, and NRGBA64 images with optional fill
//   - CopyImage, CopyImageTo, CopyImageStruct for standard image types
//   - PixelImage / PixelImage64 flat buffers with graphicx.Pixel access and image conversion
//   - PixelChanImage and channel-specialized views (Gray, RGB, ARGB, 8- and 16-bit)
//   - PixelDirection bitmask for eight-neighbor traversal used by filterx morphological ops
//
// 主要职责：
//   - LoadImage / SaveImage：基于 formatx 格式注册表读写图像（自动识别或指定格式）
//   - 创建 Gray、Gray16、RGBA、NRGBA、RGBA64、NRGBA64 图像并可选填充
//   - CopyImage、CopyImageTo、CopyImageStruct 复制标准 image 类型
//   - PixelImage / PixelImage64 扁平像素缓冲，支持 graphicx.Pixel 访问及与 image 互转
//   - PixelChanImage 及 Gray、RGB、ARGB 等通道视图（8/16 位）
//   - PixelDirection 八邻域方向位掩码，供 filterx 形态学运算使用
package imagex
