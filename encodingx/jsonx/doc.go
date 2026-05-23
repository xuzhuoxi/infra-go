// Package jsonx provides JSON encoding handlers and buffer codecs built on encodingx and json-iterator.
// Package jsonx 基于 encodingx 与 json-iterator 提供 JSON 编解码处理器及缓冲编解码器。
//
// Created by xuzhuoxi on 2019-02-24.
// @author xuzhuoxi
//
// Capabilities:
//   - NewJsonCodingHandlerAsync/Sync: ICodingHandler backed by jsoniter
//   - NewJsonBuffEncoder, NewJsonBuffDecoder, NewJsonBuffCodecs: bytex buffer integration
//   - Default pools and NewDefault* factories for handler and buff codec reuse
//
// 主要职责：
//   - JSON 编解码处理器（jsoniter 实现 ICodingHandler）
//   - 与 bytex 集成的 JSON 缓冲编码器、解码器与双向编解码器
//   - 默认对象池及 NewDefault* 工厂，便于复用处理器与缓冲实例
package jsonx
