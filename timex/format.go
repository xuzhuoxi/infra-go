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
	StdISO8601TZ             = "Z0700" //prints Z for UTC
	StdISO8601SecondsTZ      = "Z070000"
	StdISO8601ShortTZ        = "Z07"
	StdISO8601ColonTZ        = "Z07:00" //prints Z for UTC
	StdISO8601ColonSecondsTZ = "Z07:00:00"
	StdNumTZ                 = "-0700" //always numeric
	StdNumSecondsTz          = "-070000"
	StdNumShortTZ            = "-07"    //always numeric
	StdNumColonTZ            = "-07:00" //always numeric
	StdNumColonSecondsTZ     = "-07:00:00"
)

// FormatNow
// Formats the current time, returning a string representation
// 格式化当前时间，返回字符串表示
func FormatNow(layout string, loc *time.Location) string {
	return time.Now().In(loc).Format(layout)
}

// FormatNowLocal
// Formats the current time, returning a string representation of the local time zone
// 格式化当前时间，返回本地时区的字符串表示
func FormatNowLocal(layout string) string {
	return time.Now().Format(layout)
}

// FormatNowUTC
// Formats the current time, returning a string representation of the utc time zone
// 格式化当前时间，返回UTC时区的字符串表示
func FormatNowUTC(layout string) string {
	return NowUTC().Format(layout)
}

// FormatSeconds
// Formats a second value, returning a string representation
// 格式化秒数值，返回字符串表示
func FormatSeconds(s int64, loc *time.Location, layout string) string {
	return FromSeconds(s, loc).Format(layout)
}

// FormatSecondsLocal
// Formats a second value, returning a string representation of the local time zone
// 格式化秒数值，返回本地时区的字符串表示
func FormatSecondsLocal(s int64, layout string) string {
	return FromSecondsLocal(s).Format(layout)
}

// FormatSecondsUTC
// Formats a second value, returning a string representation of the utc time zone
// 格式化秒数值，返回UTC时区的字符串表示
func FormatSecondsUTC(s int64, layout string) string {
	return FromSecondsUTC(s).Format(layout)
}

// FormatMillis
// Formats a millisecond value, returning a string representation
// 格式化毫秒值，返回字符串表示
func FormatMillis(ms int64, loc *time.Location, layout string) string {
	return FromMillis(ms, loc).Format(layout)
}

// FormatMillisLocal
// Formats a millisecond value, returning a string representation of the local time zone
// 格式化毫秒值，返回本地时区的字符串表示
func FormatMillisLocal(ms int64, layout string) string {
	return FromMillisLocal(ms).Format(layout)
}

// FormatMillisUTC
// Formats a millisecond value, returning a string representation of the utc time zone
// 格式化毫秒值，返回UTC时区的字符串表示
func FormatMillisUTC(ms int64, layout string) string {
	return FromMillisUTC(ms).Format(layout)
}

// FormatNanos
// Formats a nanosecond value, returning a string representation
// 格式化纳秒值，返回字符串表示
func FormatNanos(ns int64, loc *time.Location, layout string) string {
	return FromNanos(ns, loc).Format(layout)
}

// FormatNanosLocal
// Formats a nanosecond value, returning a string representation of the local time zone
// 格式化纳秒值，返回本地时区的字符串表示
func FormatNanosLocal(ns int64, layout string) string {
	return FromNanosLocal(ns).Format(layout)
}

// FormatNanosUTC
// Formats a nanosecond value, returning a string representation of the utc time zone
// 格式化纳秒值，返回UTC时区的字符串表示
func FormatNanosUTC(ns int64, layout string) string {
	return FromNanosUTC(ns).Format(layout)
}
