// Package encodingx provides serialization abstractions, buffer codecs, key-value maps, and object pools.
// Package encodingx 提供序列化抽象、缓冲编解码、键值映射及对象池等能力。
//
// Created by xuzhuoxi on 2019-02-11.
// @author xuzhuoxi
//
// Capabilities:
//   - IEncodingData, IDecodingData, ICodingData: encode and decode to byte slices
//   - IEncodeHandler, IDecodeHandler, ICodingHandler: pluggable encode/decode handlers
//   - IBuffEncoder, IBuffDecoder, IBuffCodecs: write/read multiple values via bytex buffers
//   - IKeyValue, CodingMap: string-keyed maps with binaryx-backed binary wire format
//   - IPoolKeyValue and related pools: reuse handlers, key-value maps, and buff codecs
//   - NewUtf8StringCodingHandle, StringToByteArray, ByteArrayToString: UTF-8 and zero-copy helpers
//
// Subpackages jsonx, yamlx, and gobx supply format-specific ICodingHandler and buff codec factories.
//
// 主要职责：
//   - IEncodingData、IDecodingData、ICodingData：与字节切片互转的序列化契约
//   - IEncodeHandler、IDecodeHandler、ICodingHandler：可插拔的编解码处理器
//   - IBuffEncoder、IBuffDecoder、IBuffCodecs：基于 bytex 缓冲区的多值读写
//   - IKeyValue、CodingMap：基于 binaryx 的字符串键值二进制格式
//   - 各类对象池：复用处理器、键值映射与缓冲编解码器
//   - UTF-8 字符串处理器及 unsafe 零拷贝 string/[]byte 转换
//
// 子包 jsonx、yamlx、gobx 提供 JSON、YAML、Gob 等格式的处理器与缓冲编解码工厂。
package encodingx
