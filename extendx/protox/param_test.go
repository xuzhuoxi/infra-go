// Create on 2023/7/19
// @author xuzhuoxi
package protox

import (
	"fmt"
	"testing"
)

type CloneObject struct {
	A int
	B string
}

func TestInterfaceClone(t *testing.T) {
	c := CloneObject{A: 1, B: "HelloWorld"}
	var cc interface{} = c
	dd := cc
	ddo := dd.(CloneObject)
	ddo.A = 0
	fmt.Println(c)
	fmt.Println(ddo)
}
