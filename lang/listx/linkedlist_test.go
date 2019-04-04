//
//Created by xuzhuoxi
//on 2019-04-04.
//@author xuzhuoxi
//
package listx

import (
	"fmt"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList()
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

func TestLinkedList_Get(t *testing.T) {
	list := NewLinkedList()
	list.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(list.First())
	fmt.Println(list.Last())
	fmt.Println(list.FirstMulti(2))
	fmt.Println(list.LastMulti(2))
	fmt.Println(list.GetMulti(2, 3))
	fmt.Println(list.GetMultiLast(2, 3))
}

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList()
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

func TestLinkedList_Index(t *testing.T) {
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

func TestLinkedList_Swap(t *testing.T) {
	list := NewLinkedList()
	list.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	list.Swap(0, 10)
	fmt.Println(list.GetAll())

	list.Swap(0, 10) //回复
	list.Swap(0, 9)
	fmt.Println(list.GetAll())

	list.Swap(0, 9) //回复
	list.Swap(1, 10)
	fmt.Println(list.GetAll())

	list.Swap(1, 10) //回复
	list.Swap(4, 5)
	fmt.Println(list.GetAll())
}
