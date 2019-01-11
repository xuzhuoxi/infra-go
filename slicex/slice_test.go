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
	slice := InsertAt(s, 99, 0)
	fmt.Println(slice)
	slice = InsertAt(s, 88, 99)
	fmt.Println(slice)
}

func TestRemoveAt(t *testing.T) {
	s := []interface{}{2, 3, 4, 5, 6}
	fmt.Println(s)
	slice, r, ok := RemoveAt(s, 0)
	fmt.Println(s, slice, r, ok)
	slice2, r2, ok2 := RemoveAt(s, 4)
	fmt.Println(s, slice2, r2, ok2)
}

func TestMergeSlice(t *testing.T) {
	s1 := []interface{}{2, 3, 4, 5, 6}
	s2 := []interface{}{12, 13, 14, 15, 16}
	s3 := []interface{}{22, 23, 24, 25, 26}
	m := MergeSlice(s1, s2, s3)
	fmt.Println(s1, s2, s3, m)
}
