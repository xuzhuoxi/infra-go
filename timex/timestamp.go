// Package timex
// Create on 2023/8/16
// @author xuzhuoxi
package timex

import "time"

var (
	ZeroUTC       = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	Zero1970UTC   = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	ZeroLocal     = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	Zero1970Local = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
)

// NowDuration
// Returns the count of nanosecond time that has elapsed from 0:00:00 on 1/1/1 to the present
// 返回从 1年1月1日 0时0分0秒 到现在经过的纳秒时间计数
func NowDuration() time.Duration {
	return time.Now().Sub(ZeroUTC)
}

// NowDuration1970
// Returns the count of nanosecond time that has elapsed from January 1, 1970 0:00:00 to the present
// 返回从 1970年1月1日 0时0分0秒 到现在经过的纳秒时间计数
func NowDuration1970() time.Duration {
	return time.Now().Sub(Zero1970UTC)
}

// DurationFrom
// Returns the count of nanosecond time that has elapsed from time.Time to the present
// 返回从 time.Time 到现在经过的纳秒时间计数
func DurationFrom(t time.Time) time.Duration {
	return time.Now().Sub(t)
}

// TimeNowBefore
// Returns the time obtained by the length of some time forward from the present moment. Time
// 返回从现在时刻向前一段时间长度得到的 time.Time
func TimeNowBefore(d time.Duration) time.Time {
	return time.Now().Add(-d)
}

// NowNanos
// Returns the number of nanoseconds that have elapsed since 1/1/1/ 00:00：00 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的纳秒数
func NowNanos() int64 {
	return NowDuration().Nanoseconds()
}

// NowNanos1970
// Returns the number of nanoseconds that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的纳秒数
func NowNanos1970() int64 {
	return time.Now().UnixNano()
}

// NowTicks 1 ticks = 100 ns
// Returning the number of ticks that have elapsed since 00:00:00 on 1/1/1, one tick is equal to 100 nanoseconds
// 返回自 1年1月1日 0时0分0秒 到现在经过的Tick刻度数, 一个Tick刻度等于100纳秒
func NowTicks() int64 {
	return NowDuration().Nanoseconds() / 100
}

// NowTicks1970 1 ticks = 100 ns
// Returning the number of ticks that have elapsed since January 1, 1970 00:00:00, one tick is equal to 100 nanoseconds
// 返回自 1970年1月1日 0时0分0秒 到现在经过的Tick刻度数, 一个Tick刻度等于100纳秒
func NowTicks1970() int64 {
	return time.Now().UnixNano() / 100
}

// NowMilliseconds
// Returning the number of milliseconds that have elapsed since 00:00:00 on 1/1/1 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的毫秒数
func NowMilliseconds() int64 {
	return NowDuration().Milliseconds()
}

// NowMilliseconds1970
// Returning the number of milliseconds that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的毫秒数
func NowMilliseconds1970() int64 {
	return NowDuration1970().Milliseconds()
}

// NowFloatSeconds
// Returns the floating-point value of seconds that have elapsed since 1/1/1 00:00:00 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的秒数浮点值
func NowFloatSeconds() float64 {
	return NowDuration().Seconds()
}

// NowFloatSeconds1970
// Returns the floating-point value of seconds that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的秒数浮点值
func NowFloatSeconds1970() float64 {
	return NowDuration1970().Seconds()
}

// NowSeconds
// Returns an integer value of seconds that have elapsed since 1/1/1 00:00:00 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的秒数整数值
func NowSeconds() int64 {
	return NowDuration().Nanoseconds() / int64(time.Second)
}

// NowSeconds1970
// Returns the integer value of seconds that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的秒数整数值
func NowSeconds1970() int64 {
	return time.Now().Unix()
}

// NowFloatMinutes
// Returns the floating-point value of minutes that have elapsed since 1/1/1 00:00:00 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的分钟数浮点值
func NowFloatMinutes() float64 {
	return NowDuration().Minutes()
}

// NowFloatMinutes1970
// Returns the floating-point value of minutes that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的分钟数浮点值
func NowFloatMinutes1970() float64 {
	return NowDuration1970().Minutes()
}

// NowMinutes
// Returns an integer value of minutes that have elapsed since 1/1/1 00:00:00 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的分钟数整数值
func NowMinutes() int64 {
	return NowDuration().Nanoseconds() / int64(time.Minute)
}

// NowMinutes1970
// Returns the integer value of minutes that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的分钟数整数值
func NowMinutes1970() int64 {
	return NowDuration1970().Nanoseconds() / int64(time.Minute)
}

// NowFloatHours
// Returns the floating-point value of hours that have elapsed since 1/1/1 00:00:00 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的小时数浮点值
func NowFloatHours() float64 {
	return NowDuration().Hours()
}

// NowFloatHours1970
// Returns the floating-point value of hours that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的小时数浮点值
func NowFloatHours1970() float64 {
	return NowDuration1970().Hours()
}

// NowHours
// Returns an integer value of hours that have elapsed since 1/1/1 00:00:00 to the present
// 返回自 1年1月1日 0时0分0秒 到现在经过的小时数整数值
func NowHours() int64 {
	return NowDuration().Nanoseconds() / int64(time.Hour)
}

// NowHours1970
// Returns the integer value of hours that have elapsed since January 1, 1970 00:00:00 to the present
// 返回自 1970年1月1日 0时0分0秒 到现在经过的小时数整数值
func NowHours1970() int64 {
	return NowDuration1970().Nanoseconds() / int64(time.Hour)
}
