package filex

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/osxu"
	"testing"
)

func TestCopyToLittle(t *testing.T) {
	current := osxu.GetRunningDir()
	src := current + "/source/little.txt"
	dst := current + "/target/"
	c, err := CopyTo(src, dst)
	fmt.Println(c, err)
}

func TestCopyToMiddle(t *testing.T) {
	current := osxu.GetRunningDir()
	src := current + "/source/middle.png"
	dst := current + "/target"
	c, err := CopyTo(src, dst)
	fmt.Println(c, err)
}

func TestCopyToLarge(t *testing.T) {
	current := osxu.GetRunningDir()
	src := current + "/source/large.zip"
	dst := current + "\\target"
	c, err := CopyTo(src, dst)
	fmt.Println(c, err)
}
