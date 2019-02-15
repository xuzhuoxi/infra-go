package slicex

import (
	"github.com/xuzhuoxi/infra-go/lang"
)

//一些说明:
//对于slice，一但发生扩容(cap增大)则会重新分配内存，新的slice就与源slice脱离关系了，也就是对于下标操作不会再影响源slice了
//cap不变的,append操作会直接加在len索引后
//
//这里的方法性能都不高，性能消耗在以下几个方面：
//1.interface{}在分配内存的效率相对于具体类型较差.
//2.返回切片都是重新分配的内存

//合并
func MergeT(slices ...[]interface{}) []interface{} {
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

//按位置插入
func InsertT(slice []interface{}, pos int, target ...interface{}) []interface{} {
	ln := len(slice)
	if pos < 0 {
		pos = 0
		goto Start
	}
	if pos > ln {
		pos = ln
	}
Start:
	tl := len(target)
	rs := make([]interface{}, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

//头插入
func InsertHeadT(slice []interface{}, target ...interface{}) []interface{} {
	return InsertT(slice, 0, target...)
}

//尾插入
func InsertTailT(slice []interface{}, target ...interface{}) []interface{} {
	return InsertT(slice, len(slice), target...)
}

//按值删除
func RemoveValueT(slice []interface{}, target interface{}) ([]interface{}, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexT(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtT(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

//按值删除
func RemoveAllValueT(slice []interface{}, target interface{}) ([]interface{}, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]interface{}, sl)
	index := 0
	for _, value := range slice {
		if !lang.Equal(value, target) {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

//按点删除
func RemoveAtT(slice []interface{}, pos int) ([]interface{}, interface{}, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, nil, false
	}
	obj := slice[pos]
	rs := make([]interface{}, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

//删除尾
func RemoveHeadT(slice []interface{}, pos int) ([]interface{}, interface{}, bool) {
	return RemoveAtT(slice, 0)
}

//删除头
func RemoveTailT(slice []interface{}, pos int) ([]interface{}, interface{}, bool) {
	return RemoveAtT(slice, len(slice)-1)
}

//删除区间
func RemoveFromT(slice []interface{}, startPos int, length int) (result []interface{}, removed []interface{}, ok bool) {
	endPos := startPos + length
	return RemoveRangeT(slice, startPos, endPos)
}

//删除区间
func RemoveRangeT(slice []interface{}, startPos int, endPos int) (result []interface{}, removed []interface{}, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]interface{}, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

//包含
func ContainsT(slice []interface{}, target interface{}) bool {
	_, ok := IndexT(slice, target)
	return ok
}

//从头部查找
func IndexT(slice []interface{}, target interface{}) (int, bool) {
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

//从尾部查找
func LastIndexT(slice []interface{}, target interface{}) (int, bool) {
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

//倒序
func ReverseT(slice []interface{}) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func CopyT(slice []interface{}) []interface{} {
	if nil == slice {
		return nil
	}
	ln := len(slice)
	if 0 == ln {
		return []interface{}{}
	}
	rs := make([]interface{}, ln)
	copy(rs, slice)
	return rs
}
