package slicex

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(s) //[1 2 3]
	copy(s, []int{4, 5, 6, 7, 8, 9})
	fmt.Println(s) //[4 5 6]
}

func TestAppend(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	slice := src[:4]
	fmt.Println("cap:", cap(src), cap(slice))
	fmt.Println(src, slice)
	//newSlice := append(slice, 6)
	newSlice := append(slice, 6, 7)
	fmt.Println(src, slice, newSlice)
	newSlice[0] = 8
	fmt.Println(src, slice, newSlice)
	fmt.Println(cap(src), cap(slice), cap(newSlice))
}

func TestInsertAt(t *testing.T) {
	s := []interface{}{2, 3, 4, 5, 6}
	slice := InsertT(s, 99, 0)
	fmt.Println(slice)
	slice = InsertT(s, 88, 99)
	fmt.Println(slice)
}

func TestRemoveAt(t *testing.T) {
	s := []interface{}{2, 3, 4, 5, 6}
	fmt.Println(s)
	slice, r, ok := RemoveAtT(s, 0)
	fmt.Println(s, slice, r, ok)
	slice2, r2, ok2 := RemoveAtT(s, 4)
	fmt.Println(s, slice2, r2, ok2)
}

func TestMergeSlice(t *testing.T) {
	s1 := []interface{}{2, 3, 4, 5, 6}
	s2 := []interface{}{12, 13, 14, 15, 16}
	s3 := []interface{}{22, 23, 24, 25, 26}
	m := MergeT(s1, s2, s3)
	fmt.Println(s1, s2, s3, m)
	ss1 := []byte{2, 3, 4, 5, 6}
	ss2 := []byte{12, 13, 14, 15, 16}
	ss3 := []byte{22, 23, 24, 25, 26}
	sm := MergeUint8(ss1, ss2, ss3)
	fmt.Println(ss1, ss2, ss3, sm)
}

func TestRemoveBetween(t *testing.T) {
	s1 := []interface{}{2, 3, 4, 5, 6}
	fmt.Println(RemoveRangeT(s1, 0, 2))
	fmt.Println(RemoveRangeT(s1, 1, 4))
	fmt.Println(RemoveRangeT(s1, 3, 4))
}

func TestRemoveValueAll(t *testing.T) {
	s1 := []interface{}{2, 3, 4, 5, 6, 6, 6, 7}
	fmt.Println(RemoveAllValueT(s1, 4))
	fmt.Println(RemoveAllValueT(s1, 6))
}
