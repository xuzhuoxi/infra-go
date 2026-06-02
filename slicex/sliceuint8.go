package slicex

import (
	"bytes"
)

// MergeUint8 merges multiple slices into one.
// MergeUint8 合并多个切片为一个。
func MergeUint8(slices ...[]uint8) []uint8 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint8, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertUint8 inserts elements at the specified position.
// InsertUint8 按位置插入元素。
func InsertUint8(slice []uint8, pos int, target ...uint8) []uint8 {
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
	rs := make([]uint8, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadUint8 inserts elements at the head.
// InsertHeadUint8 在头部插入元素。
func InsertHeadUint8(slice []uint8, target ...uint8) []uint8 {
	return InsertUint8(slice, 0, target...)
}

// InsertTailUint8 inserts elements at the tail.
// InsertTailUint8 在尾部插入元素。
func InsertTailUint8(slice []uint8, target ...uint8) []uint8 {
	return InsertUint8(slice, len(slice), target...)
}

// RemoveValueUint8 removes the first element equal to target.
// RemoveValueUint8 按值删除第一个匹配元素。
func RemoveValueUint8(slice []uint8, target uint8) ([]uint8, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint8(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint8(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueUint8 removes all elements equal to target.
// RemoveAllValueUint8 按值删除所有匹配元素。
func RemoveAllValueUint8(slice []uint8, target uint8) ([]uint8, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint8, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtUint8 removes the element at the specified index.
// RemoveAtUint8 按索引删除元素。
func RemoveAtUint8(slice []uint8, pos int) ([]uint8, uint8, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint8, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadUint8 removes the first element.
// RemoveHeadUint8 删除头部元素。
func RemoveHeadUint8(slice []uint8) ([]uint8, uint8, bool) {
	return RemoveAtUint8(slice, 0)
}

// RemoveTailUint8 removes the last element.
// RemoveTailUint8 删除尾部元素。
func RemoveTailUint8(slice []uint8) ([]uint8, uint8, bool) {
	return RemoveAtUint8(slice, len(slice)-1)
}

// RemoveFromUint8 removes length elements starting at startPos.
// RemoveFromUint8 从 startPos 起删除 length 个元素。
func RemoveFromUint8(slice []uint8, startPos int, length int) (result []uint8, removed []uint8, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint8(slice, startPos, endPos)
}

// RemoveRangeUint8 removes elements in [startPos, endPos).
// RemoveRangeUint8 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeUint8(slice []uint8, startPos int, endPos int) (result []uint8, removed []uint8, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint8, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsUint8 reports whether slice contains target.
// ContainsUint8 判断切片是否包含目标值。
func ContainsUint8(slice []uint8, target uint8) bool {
	_, ok := IndexUint8(slice, target)
	return ok
}

// IndexUint8 returns the index of the first matching element.
// IndexUint8 从头部查找第一个匹配元素的索引。
func IndexUint8(slice []uint8, target uint8) (int, bool) {
	//if nil == slice || len(slice) == 0 {
	//	return -1, false
	//}
	//for index, value := range slice {
	//	if value == target {
	//		return index, true
	//	}
	//}
	//return -1, false
	index := bytes.IndexByte(slice, target)
	return index, index != -1
}

// LastIndexUint8 returns the index of the last matching element.
// LastIndexUint8 从尾部查找最后一个匹配元素的索引。
func LastIndexUint8(slice []uint8, target uint8) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	for index := len(slice) - 1; index >= 0; index-- {
		if slice[index] == target {
			return index, true
		}
	}
	return -1, false
}

// ReverseUint8 reverses elements in place.
// ReverseUint8 原地倒序元素。
func ReverseUint8(slice []uint8) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// CopyUint8 returns a shallow copy of slice.
// CopyUint8 克隆切片（浅拷贝）。
func CopyUint8(slice []uint8) []uint8 {
	if nil == slice {
		return nil
	}
	ln := len(slice)
	if 0 == ln {
		return []uint8{}
	}
	rs := make([]uint8, ln)
	copy(rs, slice)
	return rs
}

// CopyByte returns a shallow copy of slice as []uint8.
// CopyByte 克隆 []byte 切片（浅拷贝）。
func CopyByte(slice []byte) []uint8 {
	return CopyUint8(slice)
}

// EqualUint8 reports whether two slices are equal element-wise.
// EqualUint8 逐元素比较两个切片是否相等。
func EqualUint8(a, b []uint8) bool {
	return bytes.Equal(a, b)
}

// SumUint8 returns the sum of all elements.
// SumUint8 求和所有元素。
func SumUint8(slice []uint8) uint8 {
	rs := uint8(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
