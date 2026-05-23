// Package bytex provides length-prefixed block framing, thread-safe buffers, and byte processing pipelines.
// Package bytex 提供长度前缀块封装、线程安全缓冲区及字节处理流水线。
//
// Created by xuzhuoxi on 2019-02-11.
// @author xuzhuoxi
//
// Capabilities:
//   - Encode and decode uint16-length-prefixed blocks with configurable byte order and handlers
//   - IBuffDataBlock / IBuffToBlock / IBuffToData: concurrent buffers for raw bytes or framed data, integrated with binaryx
//   - sync.Pool wrappers (IPoolBuff*) to reuse buffer instances
//   - BytesProcessor pipeline with built-in little/big-endian block and unblock handlers
//
// 主要职责：
//   - 可配置字节序与 handler 的 uint16 长度前缀块编解码
//   - IBuffDataBlock / IBuffToBlock / IBuffToData：支持原始字节或分帧数据的并发安全缓冲，并与 binaryx 集成
//   - sync.Pool 封装（IPoolBuff*）用于缓冲实例复用
//   - BytesProcessor 流水线，内置小端/大端封包与解包 handler
package bytex
