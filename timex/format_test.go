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
