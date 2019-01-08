package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestFromMilli(t *testing.T) {
	n := time.Now()
	nn := FromMilli(n.Unix() * 1000)
	fmt.Println(n.UnixNano())
	fmt.Println(nn.UnixNano())
}

func TestFromNano(t *testing.T) {
	n := time.Now()
	nn := FromNano(n.UnixNano())
	fmt.Println(n.UnixNano())
	fmt.Println(nn.UnixNano())
}
