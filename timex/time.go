package timex

import "time"

// FromSecond
// 由秒时间戳生成Time
func FromSecond(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// FromMilli
// 由毫秒时间戳生成Time
func FromMilli(ms int64) time.Time {
	return time.Unix(0, ms*1e6)
}

// FromNano
// 由纳秒时间戳生成Time
func FromNano(ns int64) time.Time {
	return time.Unix(0, ns)
}
