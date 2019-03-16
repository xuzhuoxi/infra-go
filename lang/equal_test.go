//
//Created by xuzhuoxi
//on 2019-03-16.
//@author xuzhuoxi
//
package lang

import (
	"fmt"
	"testing"
)

func TestTypeEqual(t *testing.T) {
	a := 1
	b := uint16(1)
	c := 2
	fmt.Println(TypeEqual(a, b))
	fmt.Println(TypeEqual(a, c))
}

func TestEqual(t *testing.T) {
	a := []uint16{1, 2, 3}
	b := []int16{1, 2, 3}
	c := []interface{}{1, 2, 3}
	d := []uint16{1, 2, 2}
	fmt.Println(Equal(a, b))
	fmt.Println(Equal(a, c))
	fmt.Println(Equal(b, c))
	fmt.Println(Equal(a, d))
}

func TestNil(t *testing.T) {
	var a []uint16
	var b []uint16
	fmt.Println(TypeEqual(a, b))
}
