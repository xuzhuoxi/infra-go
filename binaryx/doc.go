// Package binaryx provides binary serialization and deserialization helpers for io.Reader and io.Writer.
// Package binaryx 提供基于 io.Reader 与 io.Writer 的二进制序列化与反序列化工具。
//
// Created by xuzhuoxi on 2019-03-18.
// @author xuzhuoxi
//
// Capabilities:
//   - Read and write primitive values (bool, integers, floats, complex, string) with configurable byte order
//   - Length-prefixed strings and slices; int/uint width follows platform (32- or 64-bit) via build tags
//   - ValueKind enum and helpers to classify values, supply defaults, and validate supported types
//   - Generic Read, Write, ReadSlice, WriteSlice dispatch for concrete types and *interface{}
//
// 主要职责：
//   - 按指定字节序读写基本类型（bool、整数、浮点、复数、string）
//   - 长度前缀的 string 与切片读写；int/uint 宽度随平台（32/64 位）由构建标签决定
//   - ValueKind 枚举及类型判断、默认值与合法性校验工具
//   - 通用 Read、Write、ReadSlice、WriteSlice，支持具体类型与 *interface{} 分发
package binaryx
