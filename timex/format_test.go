package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatUnixMillisecond(t *testing.T) {
	n := time.Now()
	ms := n.Unix() * 1000
	layout := time.RFC3339
	fmt.Println(FormatTime(n, layout))
	fmt.Println(FormatUnixMilli(ms, layout))
}

func TestFormat(t *testing.T) {
	n := time.Now()
	fmt.Println(n.Format("2006-01-02 15:04:05:999"))
	fmt.Println(n.Format("Jan _2 15:04:05.000000000"))
	fmt.Println(n.Format(".9999"))
	fmt.Println(n)
}
