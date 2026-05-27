// Package jpegx registers JPEG encode/decode with formatx and provides quality options.
// Package jpegx 向 formatx 注册 JPEG 编解码，并提供品质选项。
//
// Created by xuzhuoxi on 2019-04-27.
// @author xuzhuoxi
//
// Capabilities:
//   - init registration for jpeg, jpg, and jps formats
//   - EncodeJPEG / DecodeJPEG wrapping image/jpeg
//   - DefaultJPEGOptions and NewJpegOptions for quality control
//
// 主要职责：
//   - init 时注册 jpeg、jpg、jps 格式
//   - EncodeJPEG / DecodeJPEG 封装 image/jpeg
//   - DefaultJPEGOptions 与 NewJpegOptions 控制压缩品质
package jpegx
