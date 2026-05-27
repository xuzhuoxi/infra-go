// Package collectionx provides ordered collections keyed by element identifiers.
// Package collectionx 提供按元素唯一标识索引的有序集合工具。
//
// Capabilities:
//   - IOrderHashElement and OrderHashElement: elements with string Id for map and slice storage
//   - IOrderHashGroup and OrderHashGroup: ordered group with O(1) lookup when size exceeds Threshold
//   - Add, remove, update, and batch operations with documented sentinel errors
//
// 主要职责：
//   - IOrderHashElement、OrderHashElement：带 string Id 的元素，供 map 与切片双结构存储
//   - IOrderHashGroup、OrderHashGroup：有序分组，超过 Threshold 时切换为 map 加速查找
//   - 支持增删改及批量操作，并通过预定义错误变量反馈失败原因
package collectionx
