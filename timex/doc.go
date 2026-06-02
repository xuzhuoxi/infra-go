// Package timex provides time, timestamp, duration, and formatting helpers built on Go's standard time package.
// Package timex 基于标准库 time 包，提供时间、时间戳、持续时间解析与格式化等辅助能力。
//
// Create on 2023/6/16
// @author xuzhuoxi
//
// Capabilities:
//   - Current time in local, UTC, or arbitrary location (Now, NowUTC)
//   - Build time.Time from second, millisecond, or nanosecond timestamps (FromSeconds*, FromMillis*, FromNanos*)
//   - Format current time or numeric timestamps with Go reference layout tokens (FormatNow*, FormatSeconds*, etc.)
//   - Lenient duration string parsing (ParseDuration)
//   - Elapsed time since 1/1/1 or Unix epoch (1970) in nanoseconds, ticks, milliseconds, seconds, minutes, and hours
//   - Reference zero times for UTC and local zones (ZeroUTC, Zero1970UTC, ZeroLocal, Zero1970Local)
//
// 主要职责：
//   - 获取本地、UTC 或指定时区的当前时间（Now、NowUTC）
//   - 由秒、毫秒、纳秒时间戳构造 time.Time（FromSeconds*、FromMillis*、FromNanos*）
//   - 使用 Go 参考时间布局常量格式化当前时间或数值时间戳（FormatNow*、FormatSeconds* 等）
//   - 宽松解析持续时间字符串（ParseDuration）
//   - 计算自 1/1/1 或 Unix 纪元（1970）起经过的纳秒、Tick、毫秒、秒、分钟、小时
//   - 提供 UTC 与本地时区的零时刻参考值（ZeroUTC、Zero1970UTC、ZeroLocal、Zero1970Local）
package timex
