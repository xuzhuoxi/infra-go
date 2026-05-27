// Package projectionx counts valid pixels along horizontal or vertical axes for layout analysis.
// Package projectionx 沿水平或垂直方向统计有效像素，用于版面与区域分析。
//
// Capabilities:
//   - ProjectHorizontal / ProjectVertical with custom ValidFunc predicates
//   - Project dispatcher and Project2D for simultaneous H/V projection
//   - Result with per-axis counts, MaxData, and Length helpers
//
// 主要职责：
//   - ProjectHorizontal / ProjectVertical：按自定义 ValidFunc 统计有效像素
//   - Project 分发器与 Project2D 同时获取水平/垂直投影
//   - Result 保存各轴计数，提供 MaxData、Length 辅助方法
package projectionx
