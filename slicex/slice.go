package slicex

import (
	"github.com/xuzhuoxi/infra-go/lang"
)

// MergeT merges multiple slices into one.
// MergeT 合并多个切片为一个。
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

// InsertT inserts elements at the specified position.
// InsertT 按位置插入元素。
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

// InsertHeadT inserts elements at the head.
// InsertHeadT 在头部插入元素。
func InsertHeadT(slice []interface{}, target ...interface{}) []interface{} {
	return InsertT(slice, 0, target...)
}

// InsertTailT inserts elements at the tail.
// InsertTailT 在尾部插入元素。
func InsertTailT(slice []interface{}, target ...interface{}) []interface{} {
	return InsertT(slice, len(slice), target...)
}

// RemoveValueT removes the first element equal to target.
// RemoveValueT 按值删除第一个匹配元素。
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

// RemoveAllValueT removes all elements equal to target.
// RemoveAllValueT 按值删除所有匹配元素。
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

// RemoveAtT removes the element at the specified index.
// RemoveAtT 按索引删除元素。
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

// RemoveHeadT removes the first element.
// RemoveHeadT 删除头部元素。
func RemoveHeadT(slice []interface{}) ([]interface{}, interface{}, bool) {
	return RemoveAtT(slice, 0)
}

// RemoveTailT removes the last element.
// RemoveTailT 删除尾部元素。
func RemoveTailT(slice []interface{}) ([]interface{}, interface{}, bool) {
	return RemoveAtT(slice, len(slice)-1)
}

// RemoveFromT removes length elements starting at startPos.
// RemoveFromT 从 startPos 起删除 length 个元素。
func RemoveFromT(slice []interface{}, startPos int, length int) (result []interface{}, removed []interface{}, ok bool) {
	endPos := startPos + length
	return RemoveRangeT(slice, startPos, endPos)
}

// RemoveRangeT removes elements in [startPos, endPos).
// RemoveRangeT 删除 [startPos, endPos) 区间内的元素。
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

// ContainsT reports whether slice contains target.
// ContainsT 判断切片是否包含目标值。
func ContainsT(slice []interface{}, target interface{}) bool {
	_, ok := IndexT(slice, target)
	return ok
}

// IndexT returns the index of the first matching element.
// IndexT 从头部查找第一个匹配元素的索引。
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

// LastIndexT returns the index of the last matching element.
// LastIndexT 从尾部查找最后一个匹配元素的索引。
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

// ReverseT reverses elements in place.
// ReverseT 原地倒序元素。
func ReverseT(slice []interface{}) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// CopyT returns a shallow copy of slice.
// CopyT 克隆切片（浅拷贝）。
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

// EqualT reports whether two slices are equal element-wise.
// EqualT 逐元素比较两个切片是否相等。
func EqualT(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for index, val := range a {
		if !lang.Equal(val, b[index]) {
			return false
		}
	}
	return true
}
