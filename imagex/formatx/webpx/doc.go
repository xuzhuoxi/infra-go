// Package webpx registers WebP format with formatx (encode/decode stubs).
// Package webpx 向 formatx 注册 WebP 格式（编解码为占位实现）。
//
// Created by xuzhuoxi on 2023-05-03.
// @author xuzhuoxi
//
// Capabilities:
//   - init registration for webp format
//   - EncodeWEBP / DecodeWEBP placeholders returning not-implemented errors
//
// 主要职责：
//   - init 时注册 webp 格式
//   - EncodeWEBP / DecodeWEBP 占位实现，返回未实现错误
package webpx
