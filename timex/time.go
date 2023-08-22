package timex

import "time"

// NowUTC
// returns the current time under UTC
// 返回UTC下的当前时间
func NowUTC() time.Time {
	return time.Now().In(time.UTC)
}

// Now
// returns the current time under loc.
// 返回loc时区下的当前时间
func Now(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

// FromSecondsLocal
// Generate Time from seconds timestamp, using the Local time zone.
// 由秒时间戳生成Time，使用本地时区
func FromSecondsLocal(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// FromSecondsUTC
// Generates Time from a second timestamp, using the UTC time zone.
// 由秒时间戳生成Time，使用UTC时区
func FromSecondsUTC(sec int64) time.Time {
	return FromSecondsLocal(sec).In(time.UTC)
}

// FromSeconds
// Generates Time from a second timestamp, using the specified time zone.
// 由秒时间戳生成Time，使用指定的时区
func FromSeconds(sec int64, loc *time.Location) time.Time {
	if loc == time.Local {
		return FromSecondsLocal(sec)
	}
	return FromSecondsLocal(sec).In(loc)
}

// FromMillisLocal
// Generate Time from milliseconds timestamp, using the Local time zone.
// 由毫秒时间戳生成Time，使用本地时区
func FromMillisLocal(ms int64) time.Time {
	return time.Unix(0, ms*1e6)
}

// FromMillisUTC
// Generates Time from milliseconds timestamp, using the UTC time zone.
// 由毫秒秒时间戳生成Time，使用UTC时区
func FromMillisUTC(ms int64) time.Time {
	return FromMillisLocal(ms).In(time.UTC)
}

// FromMillis
// Generates Time from milliseconds timestamp, using the specified time zone.
// 由毫秒时间戳生成Time，使用指定的时区
func FromMillis(ms int64, loc *time.Location) time.Time {
	if loc == time.Local {
		return FromMillisLocal(ms)
	}
	return FromMillisLocal(ms).In(loc)
}

// FromNanosLocal
// Generate Time from nanoseconds timestamp, using the Local time zone.
// 由纳秒时间戳生成Time，使用本地时区
func FromNanosLocal(ns int64) time.Time {
	return time.Unix(0, ns)
}

// FromNanosUTC
// Generates Time from nanoseconds timestamp, using the UTC time zone.
// 由纳秒秒时间戳生成Time，使用UTC时区
func FromNanosUTC(ns int64) time.Time {
	return FromNanosLocal(ns).In(time.UTC)
}

// FromNanos
// Generates Time from nanoseconds timestamp, using the specified time zone.
// 由纳秒时间戳生成Time，使用指定的时区
func FromNanos(ns int64, loc *time.Location) time.Time {
	if loc == time.Local {
		return FromNanosLocal(ns)
	}
	return FromNanosLocal(ns).In(loc)
}
