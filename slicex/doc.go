// Package slicex provides typed slice manipulation helpers for merge, insert, remove, search, reverse, copy, equal, and sum operations.
// Package slicex 提供针对多种元素类型的切片操作辅助：合并、插入、删除、查找、倒序、拷贝、比较与求和等。
//
// Notes on slice behavior:
//   - When a slice grows beyond its capacity, Go reallocates backing memory and the new slice no longer shares storage with the source.
//   - Appends that stay within capacity write after len without reallocating.
//
// Performance considerations:
//   - Methods operating on []interface{} incur allocation overhead compared to concrete types.
//   - Most operations return newly allocated slices rather than modifying in place.
//
// 切片行为说明：
//   - 切片扩容（cap 增大）时会重新分配内存，新切片与源切片不再共享底层存储。
//   - cap 不变时，append 直接在 len 索引后追加。
//
// 性能说明：
//   - interface{} 相关方法的内存分配效率低于具体类型。
//   - 多数操作返回新分配的切片，而非原地修改。
//
// Capabilities:
//   - Generic []interface{} helpers (MergeT, InsertT, Remove*, IndexT, etc.)
//   - Typed helpers for int/uint/float/complex/string and byte slices
//   - String-specific case-insensitive search and duplicate removal
//   - Numeric Sum* aggregations and uint/uint64 slice conversion
//
// 主要职责：
//   - 面向 []interface{} 的通用切片操作（MergeT、InsertT、Remove*、IndexT 等）
//   - 针对 int/uint/float/complex/string 及 byte 切片的类型化辅助函数
//   - 字符串大小写不敏感查找与去重
//   - 数值求和及 uint/uint64 切片互转
package slicex
