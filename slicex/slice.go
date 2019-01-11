package slicex

import (
	"github.com/xuzhuoxi/util-go/lang"
)

//一些说明:
//对于slice，一但发生扩容(cap增大)则会重新分配内存，新的slice就与源slice脱离关系了，也就是对于下标操作不会再影响源slice了
//cap不变的,append操作会直接加在len索引后
//
//这里的方法性能都不高，性能消耗在以下几个方面：
//1.interface{}在分配内存的效率相对于具体类型较差.
//2.返回切片都是重新分配的内存

func MergeSlice(slices ...[]interface{}) []interface{} {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]interface{}, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

func InsertAt(slice []interface{}, target interface{}, pos int) []interface{} {
	ln := len(slice)
	if pos < 0 {
		pos = 0
		goto Start
	}
	if pos > ln {
		pos = ln
	}
Start:
	rs := make([]interface{}, ln+1)
	copy(rs, slice[:pos])
	rs[pos] = target
	copy(rs[pos+1:], slice[pos:])
	return rs
}

func InsertHead(slice []interface{}, target interface{}) []interface{} {
	return InsertAt(slice, target, 0)
}

func InsertTail(slice []interface{}, target interface{}) []interface{} {
	return InsertAt(slice, target, len(slice))
}

func RemoveAt(slice []interface{}, index int) ([]interface{}, interface{}, bool) {
	ln := len(slice)
	if index < 0 || index >= ln {
		return slice, nil, false
	}
	obj := slice[index]
	rs := make([]interface{}, ln-1)
	copy(rs, slice[:index])
	copy(rs[index:], slice[index+1:])
	return rs, obj, true
}

func RemoveHead(slice []interface{}, index int) ([]interface{}, interface{}, bool) {
	return RemoveAt(slice, 0)
}

func RemoveTail(slice []interface{}, index int) ([]interface{}, interface{}, bool) {
	return RemoveAt(slice, len(slice)-1)
}

func RemoveObject(slice []interface{}, target interface{}) ([]interface{}, bool) {
	index, ok := IndexOf(slice, target)
	if ok {
		rs, _, ok2 := RemoveAt(slice, index)
		if ok2 {
			return rs, true
		}
		return slice, false
	}
	return slice, false
}

func Contains(slice []interface{}, target interface{}) bool {
	_, ok := IndexOf(slice, target)
	return ok
}

func IndexOf(slice []interface{}, target interface{}) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	for index, value := range slice {
		if lang.Equal(value, target) {
			return index, true
		}
	}
	return -1, false
}

func LastIndexOf(slice []interface{}, target interface{}) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	for index := len(slice) - 1; index >= 0; index-- {
		if lang.Equal(slice[index], target) {
			return index, true
		}
	}
	return -1, false
}

func Reverse(slice []interface{}) []interface{} {
	if nil == slice {
		return nil
	}
	ln := len(slice)
	if 0 == ln {
		return []interface{}{}
	}
	rs := make([]interface{}, ln)
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = slice[j], slice[i]
	}
	return rs
}
