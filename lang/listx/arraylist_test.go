// Package listx
// Created by xuzhuoxi
// on 2019-04-04.
// @author xuzhuoxi
//
package listx

import (
	"fmt"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	list := NewArrayList(0, 20)
	list.Add(1)
	list.Add(2, 3)
	list.AddAt(0, 1)
	list.AddAt(list.Len(), 9)
	fmt.Println(list.GetAll())
	list2 := NewLinkedList()
	list2.AddAll(0, list)
	list2.AddAll(list2.Len(), list)
	fmt.Println(list2.GetAll())
}

func TestArrayList_Get(t *testing.T) {
	list := NewArrayList(0, 20)
	list.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(list.First())
	fmt.Println(list.Last())
	fmt.Println(list.FirstMulti(2))
	fmt.Println(list.LastMulti(2))
	fmt.Println(list.GetMulti(2, 3))
	fmt.Println(list.GetMultiLast(2, 3))
	fmt.Println(list.GetMultiLast(9, 3))
}

func TestArrayList_Remove(t *testing.T) {
	list := NewArrayList(0, 20)
	list.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Println(list.RemoveAt(11))
	fmt.Println(list.RemoveMultiAt(8, 2))
	fmt.Println(list.RemoveFirst())
	fmt.Println(list.RemoveLast())
	fmt.Println(list.GetAll())
	fmt.Println(list.RemoveFirstMulti(2))
	fmt.Println(list.RemoveLastMulti(2))
	fmt.Println(list.GetAll())
}

func TestArrayList_Index(t *testing.T) {
	list := NewLinkedList()
	list.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(list.GetAll())
	fmt.Println(list.IndexOf(0))
	fmt.Println(list.IndexOf(10))
	fmt.Println(list.IndexOf(5))
	fmt.Println(list.LastIndexOf(0))
	fmt.Println(list.LastIndexOf(10))
	fmt.Println(list.LastIndexOf(5))
}

func TestArrayList(t *testing.T) {
	list := NewArrayList(0, 20)
	list.Add(1, 2, 3, 4, 5, 6, 7)
	list.AddAt(0, 0)
	list.AddAt(8, 8)
	fmt.Println(list.GetAll())
	list.RemoveFirst()
	list.RemoveLast()
	fmt.Println(list.GetAll())
	fmt.Println(list.First())
	fmt.Println(list.Last())
	fmt.Println("------", list.GetAll())
	fmt.Println(list.FirstMulti(2))
	fmt.Println(list.LastMulti(2))
	fmt.Println(list.IndexOf(3))
	fmt.Println(list.LastIndexOf(3))

	list.RemoveAt(3)
	if e, ok := list.Get(2); ok {
		fmt.Println("2=", e)
	} else {
		panic("error")
	}
}
