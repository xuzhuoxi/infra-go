package timex

import "time"

// NowUTC returns the current time in UTC.
// NowUTC 返回 UTC 时区下的当前时间。
func NowUTC() time.Time {
	return time.Now().In(time.UTC)
}

// Now returns the current time in loc.
// Now 返回 loc 时区下的当前时间。
func Now(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

// FromSecondsLocal builds time.Time from a second timestamp using the local time zone.
// FromSecondsLocal 由秒时间戳生成 Time，使用本地时区。
func FromSecondsLocal(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// FromSecondsUTC builds time.Time from a second timestamp using the UTC time zone.
// FromSecondsUTC 由秒时间戳生成 Time，使用 UTC 时区。
func FromSecondsUTC(sec int64) time.Time {
	return FromSecondsLocal(sec).In(time.UTC)
}

// FromSeconds builds time.Time from a second timestamp using the specified time zone.
// FromSeconds 由秒时间戳生成 Time，使用指定的时区。
func FromSeconds(sec int64, loc *time.Location) time.Time {
	if loc == time.Local {
		return FromSecondsLocal(sec)
	}
	return FromSecondsLocal(sec).In(loc)
}

// FromMillisLocal builds time.Time from a millisecond timestamp using the local time zone.
// FromMillisLocal 由毫秒时间戳生成 Time，使用本地时区。
func FromMillisLocal(ms int64) time.Time {
	return time.Unix(0, ms*1e6)
}

// FromMillisUTC builds time.Time from a millisecond timestamp using the UTC time zone.
// FromMillisUTC 由毫秒时间戳生成 Time，使用 UTC 时区。
func FromMillisUTC(ms int64) time.Time {
	return FromMillisLocal(ms).In(time.UTC)
}

// FromMillis builds time.Time from a millisecond timestamp using the specified time zone.
// FromMillis 由毫秒时间戳生成 Time，使用指定的时区。
func FromMillis(ms int64, loc *time.Location) time.Time {
	if loc == time.Local {
		return FromMillisLocal(ms)
	}
	return FromMillisLocal(ms).In(loc)
}

// FromNanosLocal builds time.Time from a nanosecond timestamp using the local time zone.
// FromNanosLocal 由纳秒时间戳生成 Time，使用本地时区。
func FromNanosLocal(ns int64) time.Time {
	return time.Unix(0, ns)
}

// FromNanosUTC builds time.Time from a nanosecond timestamp using the UTC time zone.
// FromNanosUTC 由纳秒时间戳生成 Time，使用 UTC 时区。
func FromNanosUTC(ns int64) time.Time {
	return FromNanosLocal(ns).In(time.UTC)
}

// FromNanos builds time.Time from a nanosecond timestamp using the specified time zone.
// FromNanos 由纳秒时间戳生成 Time，使用指定的时区。
func FromNanos(ns int64, loc *time.Location) time.Time {
	if loc == time.Local {
		return FromNanosLocal(ns)
	}
	return FromNanosLocal(ns).In(loc)
}
