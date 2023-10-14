// Package timex
// Create on 2023/6/16
// @author xuzhuoxi
package timex

import (
	"strconv"
	"strings"
	"time"
)

// ParseDuration
// 解释持续时间字符串获得time.Duration。如果持续时间 Str 不合法，则返回 0。
// parses a duration string. Returns 0 if durationStr is not legal。
func ParseDuration(durationStr string) time.Duration {
	durationStr = strings.ToLower(strings.TrimSpace(durationStr))
	if len(durationStr) == 0 {
		return 0
	}
	d, err := time.ParseDuration(durationStr)
	if nil != err {
		return 0
	}
	return d
}

func parseDuration2(durationStr string) time.Duration {
	durationStr = strings.ToLower(strings.TrimSpace(durationStr))
	if len(durationStr) == 0 {
		return 0
	}
	var d time.Duration
	var err error
	if strings.HasSuffix(durationStr, "h") {
		d, err = parseDuration(durationStr, time.Hour, 1)
	} else if strings.HasSuffix(durationStr, "m") {
		d, err = parseDuration(durationStr, time.Minute, 1)
	} else if strings.HasSuffix(durationStr, "s") {
		d, err = parseDuration(durationStr, time.Second, 1)
	} else if strings.HasSuffix(durationStr, "ms") {
		d, err = parseDuration(durationStr, time.Millisecond, 2)
	} else if strings.HasSuffix(durationStr, "ns") {
		d, err = parseDuration(durationStr, time.Nanosecond, 2)
	} else {
		value, err := strconv.ParseFloat(strings.TrimSpace(durationStr), 64)
		if err != nil {
			return 0
		}
		return time.Duration(value)
	}
	if err != nil {
		return 0
	}
	return d
}

func parseDuration(sizeStr string, unit time.Duration, unitLen int) (d time.Duration, err error) {
	sizeStr = sizeStr[:len(sizeStr)-unitLen]
	sizeStr = strings.TrimSpace(sizeStr)
	value, err1 := strconv.ParseFloat(sizeStr, 64)
	if nil != err1 {
		return 0, err1
	}
	return time.Duration(value * float64(unit)), nil
}
