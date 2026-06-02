package timex

import "time"

var (
	ZeroUTC       = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	Zero1970UTC   = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	ZeroLocal     = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	Zero1970Local = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
)

// NowDuration returns the nanosecond duration elapsed from 1/1/1 00:00:00 to now.
// NowDuration 返回从 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的纳秒时间计数。
func NowDuration() time.Duration {
	return time.Now().Sub(ZeroUTC)
}

// NowDuration1970 returns the nanosecond duration elapsed from 1970-01-01 00:00:00 to now.
// NowDuration1970 返回从 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的纳秒时间计数。
func NowDuration1970() time.Duration {
	return time.Now().Sub(Zero1970UTC)
}

// DurationFrom returns the nanosecond duration elapsed from t to now.
// DurationFrom 返回从 time.Time 到现在经过的纳秒时间计数。
func DurationFrom(t time.Time) time.Duration {
	return time.Now().Sub(t)
}

// TimeNowBefore returns the time obtained by subtracting d from the present moment.
// TimeNowBefore 返回从现在时刻向前一段时间长度得到的 time.Time。
func TimeNowBefore(d time.Duration) time.Time {
	return time.Now().Add(-d)
}

// NowNanos returns the nanoseconds elapsed from 1/1/1 00:00:00 to now.
// NowNanos 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的纳秒数。
func NowNanos() int64 {
	return NowDuration().Nanoseconds()
}

// NowNanos1970 returns the nanoseconds elapsed from 1970-01-01 00:00:00 to now.
// NowNanos1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的纳秒数。
func NowNanos1970() int64 {
	return time.Now().UnixNano()
}

// NowTicks returns the ticks elapsed from 1/1/1 00:00:00 to now; one tick equals 100 nanoseconds.
// NowTicks 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的 Tick 刻度数，一个 Tick 刻度等于 100 纳秒。
func NowTicks() int64 {
	return NowDuration().Nanoseconds() / 100
}

// NowTicks1970 returns the ticks elapsed from 1970-01-01 00:00:00 to now; one tick equals 100 nanoseconds.
// NowTicks1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的 Tick 刻度数，一个 Tick 刻度等于 100 纳秒。
func NowTicks1970() int64 {
	return time.Now().UnixNano() / 100
}

// NowMilliseconds returns the milliseconds elapsed from 1/1/1 00:00:00 to now.
// NowMilliseconds 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的毫秒数。
func NowMilliseconds() int64 {
	return NowDuration().Milliseconds()
}

// NowMilliseconds1970 returns the milliseconds elapsed from 1970-01-01 00:00:00 to now.
// NowMilliseconds1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的毫秒数。
func NowMilliseconds1970() int64 {
	return NowDuration1970().Milliseconds()
}

// NowFloatSeconds returns the floating-point seconds elapsed from 1/1/1 00:00:00 to now.
// NowFloatSeconds 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的秒数浮点值。
func NowFloatSeconds() float64 {
	return NowDuration().Seconds()
}

// NowFloatSeconds1970 returns the floating-point seconds elapsed from 1970-01-01 00:00:00 to now.
// NowFloatSeconds1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的秒数浮点值。
func NowFloatSeconds1970() float64 {
	return NowDuration1970().Seconds()
}

// NowSeconds returns the integer seconds elapsed from 1/1/1 00:00:00 to now.
// NowSeconds 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的秒数整数值。
func NowSeconds() int64 {
	return NowDuration().Nanoseconds() / int64(time.Second)
}

// NowSeconds1970 returns the integer seconds elapsed from 1970-01-01 00:00:00 to now.
// NowSeconds1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的秒数整数值。
func NowSeconds1970() int64 {
	return time.Now().Unix()
}

// NowFloatMinutes returns the floating-point minutes elapsed from 1/1/1 00:00:00 to now.
// NowFloatMinutes 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的分钟数浮点值。
func NowFloatMinutes() float64 {
	return NowDuration().Minutes()
}

// NowFloatMinutes1970 returns the floating-point minutes elapsed from 1970-01-01 00:00:00 to now.
// NowFloatMinutes1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的分钟数浮点值。
func NowFloatMinutes1970() float64 {
	return NowDuration1970().Minutes()
}

// NowMinutes returns the integer minutes elapsed from 1/1/1 00:00:00 to now.
// NowMinutes 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的分钟数整数值。
func NowMinutes() int64 {
	return NowDuration().Nanoseconds() / int64(time.Minute)
}

// NowMinutes1970 returns the integer minutes elapsed from 1970-01-01 00:00:00 to now.
// NowMinutes1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的分钟数整数值。
func NowMinutes1970() int64 {
	return NowDuration1970().Nanoseconds() / int64(time.Minute)
}

// NowFloatHours returns the floating-point hours elapsed from 1/1/1 00:00:00 to now.
// NowFloatHours 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的小时数浮点值。
func NowFloatHours() float64 {
	return NowDuration().Hours()
}

// NowFloatHours1970 returns the floating-point hours elapsed from 1970-01-01 00:00:00 to now.
// NowFloatHours1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的小时数浮点值。
func NowFloatHours1970() float64 {
	return NowDuration1970().Hours()
}

// NowHours returns the integer hours elapsed from 1/1/1 00:00:00 to now.
// NowHours 返回自 1 年 1 月 1 日 0 时 0 分 0 秒到现在经过的小时数整数值。
func NowHours() int64 {
	return NowDuration().Nanoseconds() / int64(time.Hour)
}

// NowHours1970 returns the integer hours elapsed from 1970-01-01 00:00:00 to now.
// NowHours1970 返回自 1970 年 1 月 1 日 0 时 0 分 0 秒到现在经过的小时数整数值。
func NowHours1970() int64 {
	return NowDuration1970().Nanoseconds() / int64(time.Hour)
}
