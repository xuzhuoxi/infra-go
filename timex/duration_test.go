// Package timex
// Create on 2023/8/21
// @author xuzhuoxi
package timex

import (
	"fmt"
	"testing"
)

func TestParseDuration(t *testing.T) {
	strArr := []string{"300ms", "-1.5h", "2h45m"}
	for _, str := range strArr {
		fmt.Println(ParseDuration(str).Nanoseconds())
	}

	// Output:
	// 300000000
	// -5400000000000
	// 9900000000000
}
