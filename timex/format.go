package timex

import (
	"time"
)

const (
	StdLongMonth             = "January"
	StdMonth                 = "Jan"
	StdNumMonth              = "1"
	StdZeroMonth             = "01"
	StdLongWeekDay           = "Monday"
	StdWeekDay               = "Mon"
	StdDay                   = "2"
	StdUnderDay              = "_2"
	StdZeroDay               = "02"
	StdHour                  = "15"
	StdHour12                = "3"
	StdZeroHour12            = "03"
	StdMinute                = "4"
	StdZeroMinute            = "04"
	StdSecond                = "5"
	StdZeroSecond            = "05"
	StdLongYear              = "2006"
	StdYear                  = "06"
	StdPM                    = "PM"
	Stdpm                    = "pm"
	StdTZ                    = "MST"
	StdISO8601TZ             = "Z0700"             // Prints Z for UTC. // UTC 时打印 Z。
	StdISO8601SecondsTZ      = "Z070000"
	StdISO8601ShortTZ        = "Z07"
	StdISO8601ColonTZ        = "Z07:00"            // Prints Z for UTC. // UTC 时打印 Z。
	StdISO8601ColonSecondsTZ = "Z07:00:00"
	StdNumTZ                 = "-0700"             // Always numeric. // 始终为数字格式。
	StdNumSecondsTz          = "-070000"
	StdNumShortTZ            = "-07"               // Always numeric. // 始终为数字格式。
	StdNumColonTZ            = "-07:00"            // Always numeric. // 始终为数字格式。
	StdNumColonSecondsTZ     = "-07:00:00"
)

// FormatNow formats the current time with layout in loc.
// FormatNow 格式化当前时间，返回 loc 时区下的字符串表示。
func FormatNow(layout string, loc *time.Location) string {
	return time.Now().In(loc).Format(layout)
}

// FormatNowLocal formats the current time with layout in the local time zone.
// FormatNowLocal 格式化当前时间，返回本地时区的字符串表示。
func FormatNowLocal(layout string) string {
	return time.Now().Format(layout)
}

// FormatNowUTC formats the current time with layout in UTC.
// FormatNowUTC 格式化当前时间，返回 UTC 时区的字符串表示。
func FormatNowUTC(layout string) string {
	return NowUTC().Format(layout)
}

// FormatSeconds formats a second timestamp with layout in loc.
// FormatSeconds 格式化秒数值，返回 loc 时区下的字符串表示。
func FormatSeconds(s int64, loc *time.Location, layout string) string {
	return FromSeconds(s, loc).Format(layout)
}

// FormatSecondsLocal formats a second timestamp with layout in the local time zone.
// FormatSecondsLocal 格式化秒数值，返回本地时区的字符串表示。
func FormatSecondsLocal(s int64, layout string) string {
	return FromSecondsLocal(s).Format(layout)
}

// FormatSecondsUTC formats a second timestamp with layout in UTC.
// FormatSecondsUTC 格式化秒数值，返回 UTC 时区的字符串表示。
func FormatSecondsUTC(s int64, layout string) string {
	return FromSecondsUTC(s).Format(layout)
}

// FormatMillis formats a millisecond timestamp with layout in loc.
// FormatMillis 格式化毫秒值，返回 loc 时区下的字符串表示。
func FormatMillis(ms int64, loc *time.Location, layout string) string {
	return FromMillis(ms, loc).Format(layout)
}

// FormatMillisLocal formats a millisecond timestamp with layout in the local time zone.
// FormatMillisLocal 格式化毫秒值，返回本地时区的字符串表示。
func FormatMillisLocal(ms int64, layout string) string {
	return FromMillisLocal(ms).Format(layout)
}

// FormatMillisUTC formats a millisecond timestamp with layout in UTC.
// FormatMillisUTC 格式化毫秒值，返回 UTC 时区的字符串表示。
func FormatMillisUTC(ms int64, layout string) string {
	return FromMillisUTC(ms).Format(layout)
}

// FormatNanos formats a nanosecond timestamp with layout in loc.
// FormatNanos 格式化纳秒值，返回 loc 时区下的字符串表示。
func FormatNanos(ns int64, loc *time.Location, layout string) string {
	return FromNanos(ns, loc).Format(layout)
}

// FormatNanosLocal formats a nanosecond timestamp with layout in the local time zone.
// FormatNanosLocal 格式化纳秒值，返回本地时区的字符串表示。
func FormatNanosLocal(ns int64, layout string) string {
	return FromNanosLocal(ns).Format(layout)
}

// FormatNanosUTC formats a nanosecond timestamp with layout in UTC.
// FormatNanosUTC 格式化纳秒值，返回 UTC 时区的字符串表示。
func FormatNanosUTC(ns int64, layout string) string {
	return FromNanosUTC(ns).Format(layout)
}
