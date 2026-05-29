// Package mathx provides numeric helpers: abs/min/max/clamp, bit ops, random ranges, size units, and radix conversion.
// Package mathx 提供数值辅助能力：绝对值、最小/最大/钳制、位运算、随机范围、存储单位解析及进制转换。
//
// Created by xuzhuoxi on 2019-04-03.
// @author xuzhuoxi
//
// Capabilities:
//   - Abs*, Min*, Max*, Clamp*, and PowInt for common integer and float operations
//   - BitString, BitValid, BitFit, BitSet, and combined BitValidAnd/Or helpers
//   - RandBool and Rand* functions for bounded pseudo-random values
//   - SizeUnit constants (Byte through EB) and ParseSize for human-readable size strings
//   - SystemTo and base-10/26/36 conversion helpers (System10To26, System26To10, etc.)
//   - Subpackage gaussx: 2D Gaussian function and normalized convolution kernel builders
//
// 主要职责：
//   - Abs*、Min*、Max*、Clamp*、PowInt 等常用整数与浮点运算
//   - BitString、BitValid、BitFit、BitSet 及 BitValidAnd/Or 位操作辅助
//   - RandBool 与 Rand* 系列有界伪随机数生成
//   - SizeUnit 常量（Byte 至 EB）与 ParseSize 解析可读大小字符串
//   - SystemTo 及十/二十六/三十六进制互转（System10To26、System26To10 等）
//   - 子包 gaussx：二维高斯函数与归一化卷积核构造
package mathx
