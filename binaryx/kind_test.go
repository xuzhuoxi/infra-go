//
//Created by xuzhuoxi
//on 2019-03-18.
//@author xuzhuoxi
//
package binaryx

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestKindDefine(t *testing.T) {
	fmt.Println(KindNone, KindBool, KindString)
	fmt.Println(KindSliceNone, KindSliceBool, KindSliceString)
}

func TestKind(t *testing.T) {
	a := []bool(nil)
	fmt.Println(a, reflect.TypeOf(a))
}

func TestLen(t *testing.T) {
	var a complex64
	fmt.Println(unsafe.Sizeof(a))
	var b uint64
	fmt.Println(unsafe.Sizeof(b))
	var s string = "aaaaaa"
	fmt.Println(unsafe.Sizeof(s), len(s))
}

func TestVar(t *testing.T) {
	kind, _ := GetValueKind(false)
	kind2, _ := GetValueKind([]bool(nil))
	fmt.Println(kind, kind2)
}

func TestKindDefault(t *testing.T) {
	a := GetKindValue(KindSliceFloat32, 3)
	a = []float32{1, 2, 3}
	fmt.Println(a, reflect.TypeOf(a))
	b := a.([]float32)
	b = append(b, 555)
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(a, reflect.TypeOf(a))

	a1 := GetKindValue(KindSliceFloat32, 3)
	fmt.Println(a1, reflect.TypeOf(a1))
}
