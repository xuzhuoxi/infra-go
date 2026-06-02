// Package stringx provides Unicode-aware string helpers and password validation utilities.
// Package stringx 提供 Unicode 感知的字符串辅助函数与密码校验工具。
//
// Capabilities:
//   - Concatenate heterogeneous values into a string (Print, PrintSafe)
//   - Rune-based counting, indexing, substring, prefix/suffix, and split operations
//   - Token splitting with optional trimming and empty-token filtering
//   - Password validation via bitmask character-class flags (PasswdFlag)
//
// 主要职责：
//   - 将多种类型参数拼接为字符串（Print、PrintSafe）
//   - 基于 rune 的计数、查找、截取、前后缀拆分等操作
//   - 按分隔符拆分 token，支持去空白与忽略空串
//   - 通过位掩码字符类别标志校验密码（PasswdFlag）
package stringx
