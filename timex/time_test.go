package timex

import (
	"fmt"
	"testing"
	"time"
)

var (
	layoutFrom     = "15:04:05 -0700"
	testFromLoc, _ = time.LoadLocation("US/Hawaii")
)

func TestFromSecondsLocal(t *testing.T) {
	s := int64(1)
	nt := FromSecondsLocal(s)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 08:00:01 +0800
}

func TestFromSecondsUTC(t *testing.T) {
	s := int64(1)
	nt := FromSecondsUTC(s)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 00:00:01 +0000
}

func TestFromSeconds(t *testing.T) {
	s := int64(1)
	nt := FromSeconds(s, testFromLoc)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 14:00:01 -1000
}

func TestFromMillisLocal(t *testing.T) {
	ms := int64(1000)
	nt := FromMillisLocal(ms)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 08:00:01 +0800
}

func TestFromMillisUTC(t *testing.T) {
	ms := int64(1000)
	nt := FromSecondsUTC(ms)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 00:00:01 +0000
}

func TestFromMillis(t *testing.T) {
	ms := int64(1000)
	nt := FromMillis(ms, testFromLoc)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 14:00:01 -1000
}

func TestFromNanosLocal(t *testing.T) {
	ns := int64(time.Second)
	nt := FromNanosLocal(ns)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 08:00:01 +0800
}

func TestFromNanosUTC(t *testing.T) {
	ns := int64(time.Second)
	nt := FromNanosUTC(ns)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 00:00:01 +0000
}

func TestFromNanos(t *testing.T) {
	ns := int64(time.Second)
	nt := FromNanos(ns, testFromLoc)
	fmt.Println(nt.UnixNano())
	fmt.Println(nt.Format(layoutFrom))

	// Output:
	// 1000000000
	// 14:00:01 -1000
}
