// Package gobx provides Gob encoding handlers and buffer codecs built on encodingx and encoding/gob.
// Package gobx 基于 encodingx 与 encoding/gob 提供 Gob 编解码处理器及缓冲编解码器。
//
// Created by xuzhuoxi on 2019-02-16.
// @author xuzhuoxi
//
// Capabilities:
//   - NewGobCodingHandler, NewGobCodingHandlerSync, NewGobCodingHandlerAsync: buffer-reuse vs per-call strategies
//   - NewGobBuffEncoder, NewGobBuffDecoder, NewGobBuffCodecs: dedicated gob buff codecs
//   - Default pools and NewDefault* factories; default handler uses async (full type descriptors each encode)
//
// 主要职责：
//   - 多种 Gob 处理器：复用缓冲区（Sync/默认 Handler）与每次独立编码（Async）
//   - 专用 Gob 缓冲编解码器，与 bytex 数据块配合
//   - 默认对象池及 NewDefault* 工厂；默认处理器为 Async（每次编码附带完整类型描述）
package gobx
