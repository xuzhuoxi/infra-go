package timex

import (
	"fmt"
	"testing"
	"time"
)

var (
	layoutFormat     = "15:04:05 -0700"
	testFormatLoc, _ = time.LoadLocation("Australia/Darwin")
)

func TestFormatSecondsLocal(t *testing.T) {
	s := int64(3666)
	fmt.Println(FormatSecondsLocal(s, layoutFormat))

	// Output:
	// 09:01:06 +0800
}

func TestFormatSecondsUTC(t *testing.T) {
	s := int64(3666)
	fmt.Println(FormatSecondsUTC(s, layoutFormat))

	// Output:
	// 01:01:06 +0000
}

func TestFormatSeconds(t *testing.T) {
	s := int64(3666)
	fmt.Println(FormatSeconds(s, testFormatLoc, layoutFormat))

	// Output:
	// 10:31:06 +0930
}

func TestFormatMillisLocal(t *testing.T) {
	s := int64(3666000)
	fmt.Println(FormatMillisLocal(s, layoutFormat))

	// Output:
	// 09:01:06 +0800
}

func TestFormatMillisUTC(t *testing.T) {
	s := int64(3666000)
	fmt.Println(FormatMillisUTC(s, layoutFormat))

	// Output:
	// 01:01:06 +0000
}

func TestFormatMillis(t *testing.T) {
	s := int64(3666000)
	fmt.Println(FormatMillis(s, testFormatLoc, layoutFormat))

	// Output:
	// 10:31:06 +0930
}

func TestFormatNanosLocal(t *testing.T) {
	s := int64(3666000 * time.Millisecond)
	fmt.Println(FormatNanosLocal(s, layoutFormat))

	// Output:
	// 09:01:06 +0800
}

func TestFormatNanosUTC(t *testing.T) {
	s := int64(3666000 * time.Millisecond)
	fmt.Println(FormatNanosUTC(s, layoutFormat))

	// Output:
	// 01:01:06 +0000
}

func TestFormatNanos(t *testing.T) {
	s := int64(3666000 * time.Millisecond)
	fmt.Println(FormatNanos(s, testFormatLoc, layoutFormat))

	// Output:
	// 10:31:06 +0930
}
