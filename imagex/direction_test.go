package imagex

import (
	"fmt"
	"testing"
)

func TestReverseDirection(t *testing.T) {
	for _, dir := range dirs {
		fmt.Println("方向：", dir, dir.ReverseDirection())
	}
}
