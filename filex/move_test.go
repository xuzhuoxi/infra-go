package filex

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMove(t *testing.T) {
	var current = filepath.Dir(os.Args[0])

	fmt.Println(current)

	var sf = current + "/source/aa.txt"
	var sd = current + "/source/empty"
	var tf = current + "/target/aa.move"
	var td = current + "/target/temp/empty.move"

	Move(sf, tf)
	Move(sd, td)
}

func TestMoveTo(t *testing.T) {
	var current = filepath.Dir(os.Args[0])

	fmt.Println(filepath.Split("/source/empty"))
	fmt.Println(filepath.Split("/source/empty/"))
	fmt.Println(filepath.Split("/source/empty.txt"))

	fmt.Println(current)

	var sf = current + "/source/aa.txt"
	var sd = current + "/source/empty"
	var to = current + "/target"

	MoveTo(sf, to)
	MoveTo(sd, to)
}
