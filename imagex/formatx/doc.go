// Package formatx provides a pluggable image format registry with encode/decode dispatch by format name.
// Package formatx 提供可插拔的图像格式注册表，按格式名分发编解码。
//
// Created by xuzhuoxi on 2019-04-26.
// @author xuzhuoxi
//
// Capabilities:
//   - ImageFormat type with Encode / Decode methods
//   - RegisterFormat, RegisterFormatExt, CheckFormatRegistered, GetExtName
//   - Built-in format constants: Auto, PNG, JPEG, JPG, JPS, BMP, WEBP
//   - Subpackages jpegx, pngx, webpx register handlers at init time
//
// 主要职责：
//   - ImageFormat 类型及 Encode / Decode 方法
//   - RegisterFormat、RegisterFormatExt、CheckFormatRegistered、GetExtName
//   - 内置格式常量：Auto、PNG、JPEG、JPG、JPS、BMP、WEBP
//   - 子包 jpegx、pngx、webpx 在 init 时注册编解码实现
package formatx
