// Package cmdx
// Create on 2023/8/28
// @author xuzhuoxi
package cmdx

import (
	"fmt"
	"strings"
	"testing"
)

func TestArgSet_ParseArgs(t *testing.T) {
	argStr := "-b=true -i=12 -d=12m -ui=24 -f=12.34 -s=hello"
	args := strings.Split(argStr, " ")
	set := &ArgSet{}
	err := set.ParseArgs(args)
	if nil != err {
		t.Fatal(err)
	}

	fmt.Println(set.Len())

	fmt.Println("---")

	keys := []string{"b", "i", "d", "ui", "f", "s"}
	for _, k := range keys {
		v := set.Get(k)
		fmt.Println(k, v)
	}

	fmt.Println("---")

	b, okb := set.GetBool("b")
	fmt.Println("b", b, okb)
	i, oki := set.GetInt("i")
	fmt.Println("i", i, oki)
	d, okd := set.GetDuration("d")
	fmt.Println("d", d, okd)
	ui, okui := set.GetUint("ui")
	fmt.Println("ui", ui, okui)
	f, okf := set.GetFloat64("f")
	fmt.Println("f", f, okf)
	s, oks := set.GetString("s")
	fmt.Println("s", s, oks)

	// Output:
	// 6
	// ---
	// b true
	// i 12
	// d 12m
	// ui 24
	// f 12.34
	// s hello
	// ---
	// b true true
	// i 12 true
	// d 12m0s true
	// ui 24 true
	// f 12.34 true
	// s hello true
}
