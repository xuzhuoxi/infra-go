// Package pngx registers PNG encode/decode with formatx.
// Package pngx 向 formatx 注册 PNG 编解码。
//
// Created by xuzhuoxi on 2019-04-27.
// @author xuzhuoxi
//
// Capabilities:
//   - init registration for png format
//   - EncodePNG with optional png.CompressionLevel; DecodePNG wrapping image/png
//
// 主要职责：
//   - init 时注册 png 格式
//   - EncodePNG 支持可选 png.CompressionLevel；DecodePNG 封装 image/png
package pngx
