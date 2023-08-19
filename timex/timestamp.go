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

func NowDuration() time.Duration {
	return time.Now().Sub(ZeroUTC)
}

func NowDuration1970() time.Duration {
	return time.Now().Sub(Zero1970UTC)
}

func DurationFrom(t time.Time) time.Duration {
	return time.Now().Sub(t)
}

func TimeFrom(d time.Duration) time.Time {
	return time.Now().Add(-d)
}

func NowNanos() int64 {
	return NowDuration().Nanoseconds()
}

func NowNanos1970() int64 {
	return time.Now().UnixNano()
}

func NowTicks() int64 {
	return NowDuration().Nanoseconds() / 100
}

func NowTicks1970() int64 {
	return time.Now().UnixNano() / 100
}

func NowMilliseconds() int64 {
	return NowDuration().Milliseconds()
}

func NowMilliseconds1970() int64 {
	return NowDuration1970().Milliseconds()
}

func NowFloatSeconds() float64 {
	return NowDuration().Seconds()
}

func NowFloatSeconds1970() float64 {
	return NowDuration1970().Seconds()
}

func NowSeconds() int64 {
	return NowDuration().Nanoseconds() / int64(time.Second)
}

func NowSeconds1970() int64 {
	return time.Now().Unix()
}

func NowFloatMinutes() float64 {
	return NowDuration().Minutes()
}

func NowFloatMinutes1970() float64 {
	return NowDuration1970().Minutes()
}

func NowMinutes() int64 {
	return NowDuration().Nanoseconds() / int64(time.Minute)
}

func NowMinutes1970() int64 {
	return NowDuration1970().Nanoseconds() / int64(time.Minute)
}

func NowFloatHours() float64 {
	return NowDuration().Hours()
}

func NowFloatHours1970() float64 {
	return NowDuration1970().Hours()
}

func NowHours() int64 {
	return NowDuration().Nanoseconds() / int64(time.Hour)
}

func NowHours1970() int64 {
	return NowDuration1970().Nanoseconds() / int64(time.Hour)
}
