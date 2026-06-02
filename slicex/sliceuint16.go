package slicex

// MergeUint16 merges multiple slices into one.
// MergeUint16 合并多个切片为一个。
func MergeUint16(slices ...[]uint16) []uint16 {
	ln := len(slices)
	index := 0
	total := 0
	for index < ln {
		total += len(slices[index])
		index++
	}
	rs := make([]uint16, total)
	rsIndex := 0
	for index = 0; index < ln; index++ {
		copy(rs[rsIndex:], slices[index])
		rsIndex += len(slices[index])
	}
	return rs
}

// InsertUint16 inserts elements at the specified position.
// InsertUint16 按位置插入元素。
func InsertUint16(slice []uint16, pos int, target ...uint16) []uint16 {
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
	rs := make([]uint16, ln+tl)
	copy(rs, slice[:pos])
	copy(rs[pos:], target)
	copy(rs[pos+tl:], slice[pos:])
	return rs
}

// InsertHeadUint16 inserts elements at the head.
// InsertHeadUint16 在头部插入元素。
func InsertHeadUint16(slice []uint16, target ...uint16) []uint16 {
	return InsertUint16(slice, 0, target...)
}

// InsertTailUint16 inserts elements at the tail.
// InsertTailUint16 在尾部插入元素。
func InsertTailUint16(slice []uint16, target ...uint16) []uint16 {
	return InsertUint16(slice, len(slice), target...)
}

// RemoveValueUint16 removes the first element equal to target.
// RemoveValueUint16 按值删除第一个匹配元素。
func RemoveValueUint16(slice []uint16, target uint16) ([]uint16, bool) {
	if len(slice) == 0 {
		return nil, false
	}
	index, ok := IndexUint16(slice, target)
	if ok {
		rs, _, ok2 := RemoveAtUint16(slice, index)
		if ok2 {
			return rs, true
		}
		return nil, false
	}
	return nil, false
}

// RemoveAllValueUint16 removes all elements equal to target.
// RemoveAllValueUint16 按值删除所有匹配元素。
func RemoveAllValueUint16(slice []uint16, target uint16) ([]uint16, bool) {
	sl := len(slice)
	if sl == 0 {
		return nil, false
	}
	cp := make([]uint16, sl)
	index := 0
	for _, value := range slice {
		if value != target {
			cp[index] = value
			index++
		}
	}
	return cp[:index], true
}

// RemoveAtUint16 removes the element at the specified index.
// RemoveAtUint16 按索引删除元素。
func RemoveAtUint16(slice []uint16, pos int) ([]uint16, uint16, bool) {
	ln := len(slice)
	if pos < 0 || pos >= ln {
		return nil, 0, false
	}
	obj := slice[pos]
	rs := make([]uint16, ln-1)
	copy(rs, slice[:pos])
	copy(rs[pos:], slice[pos+1:])
	return rs, obj, true
}

// RemoveHeadUint16 removes the first element.
// RemoveHeadUint16 删除头部元素。
func RemoveHeadUint16(slice []uint16) ([]uint16, uint16, bool) {
	return RemoveAtUint16(slice, 0)
}

// RemoveTailUint16 removes the last element.
// RemoveTailUint16 删除尾部元素。
func RemoveTailUint16(slice []uint16) ([]uint16, uint16, bool) {
	return RemoveAtUint16(slice, len(slice)-1)
}

// RemoveFromUint16 removes length elements starting at startPos.
// RemoveFromUint16 从 startPos 起删除 length 个元素。
func RemoveFromUint16(slice []uint16, startPos int, length int) (result []uint16, removed []uint16, ok bool) {
	endPos := startPos + length
	return RemoveRangeUint16(slice, startPos, endPos)
}

// RemoveRangeUint16 removes elements in [startPos, endPos).
// RemoveRangeUint16 删除 [startPos, endPos) 区间内的元素。
func RemoveRangeUint16(slice []uint16, startPos int, endPos int) (result []uint16, removed []uint16, ok bool) {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	sl := len(slice)
	if startPos < 0 || endPos >= sl || startPos == endPos || sl == 0 {
		return nil, nil, false
	}
	rs := make([]uint16, sl-endPos+startPos)
	copy(rs, slice[:startPos])
	copy(rs[startPos:], slice[endPos:])
	return rs, slice[startPos:endPos], true
}

// ContainsUint16 reports whether slice contains target.
// ContainsUint16 判断切片是否包含目标值。
func ContainsUint16(slice []uint16, target uint16) bool {
	_, ok := IndexUint16(slice, target)
	return ok
}

// IndexUint16 returns the index of the first matching element.
// IndexUint16 从头部查找第一个匹配元素的索引。
func IndexUint16(slice []uint16, target uint16) (int, bool) {
	if nil == slice || len(slice) == 0 {
		return -1, false
	}
	for index, value := range slice {
		if value == target {
			return index, true
		}
	}
	return -1, false
}

// LastIndexUint16 returns the index of the last matching element.
// LastIndexUint16 从尾部查找最后一个匹配元素的索引。
func LastIndexUint16(slice []uint16, target uint16) (int, bool) {
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

// ReverseUint16 reverses elements in place.
// ReverseUint16 原地倒序元素。
func ReverseUint16(slice []uint16) {
	ln := len(slice)
	if 0 == ln {
		return
	}
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// EqualUint16 reports whether two slices are equal element-wise.
// EqualUint16 逐元素比较两个切片是否相等。
func EqualUint16(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}
	for index, val := range a {
		if val != b[index] {
			return false
		}
	}
	return true
}

// SumUint16 returns the sum of all elements.
// SumUint16 求和所有元素。
func SumUint16(slice []uint16) uint16 {
	rs := uint16(0)
	for _, val := range slice {
		rs += val
	}
	return rs
}
