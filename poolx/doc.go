// Package poolx provides reusable object pools to reduce allocations, starting with byte slice pools.
// Package poolx 提供可复用对象池以降低分配开销，当前包含字节切片池实现。
//
// Created by xuzhuoxi on 2025-04-19.
// @author xuzhuoxi
//
// Capabilities:
//   - NewByteSlicePool: create or reuse a shared IByteSlicePool keyed by capacity (default 2048)
//   - IByteSlicePool: Get/GetL acquire *[]byte; Put returns slices; Cap reports configured capacity
//
// 主要职责：
//   - NewByteSlicePool：按容量创建或复用共享的 IByteSlicePool（默认容量 2048）
//   - IByteSlicePool：通过 Get/GetL 获取 *[]byte，通过 Put 归还，Cap 返回配置的容量上限
package poolx
