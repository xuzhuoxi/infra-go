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

func FormatNow(layout string) string {
	return time.Now().Format(layout)
}

func FormatTime(time time.Time, layout string) string {
	return time.Format(layout)
}

func FormatUnixSecond(second int64, layout string) string {
	return FromSecond(second).Format(layout)
}

func FormatUnixMilli(millisecond int64, layout string) string {
	return FromMilli(millisecond).Format(layout)
}

func FormatUnixNano(nanosecond int64, layout string) string {
	return FromNano(nanosecond).Format(layout)
}
